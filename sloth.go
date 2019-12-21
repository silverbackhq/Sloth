// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/silverbackhq/sloth/internal/app/module/common"

	"github.com/drone/envsubst"
	"github.com/spf13/viper"
)

const (
	// AgentRole var
	AgentRole = "agent"
	// WorkerRole var
	WorkerRole = "worker"
	// OrchestratorRole var
	OrchestratorRole = "orchestrator"
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "config", "config.prod.yml", "config")
	flag.Parse()

	configUnparsed, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(fmt.Sprintf(
			"Error while reading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	configParsed, err := envsubst.EvalEnv(string(configUnparsed))

	if err != nil {
		panic(fmt.Sprintf(
			"Error while parsing config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer([]byte(configParsed)))

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading configs [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	if viper.GetString("log.output") != "stdout" {
		fs := module.FileSystem{}
		dir, _ := filepath.Split(viper.GetString("log.output"))

		if !fs.DirExists(dir) {
			if _, err := fs.EnsureDir(dir, 777); err != nil {
				panic(fmt.Sprintf(
					"Directory [%s] creation failed with error: %s",
					dir,
					err.Error(),
				))
			}
		}

		if !fs.FileExists(viper.GetString("log.output")) {
			f, err := os.Create(viper.GetString("log.output"))
			if err != nil {
				panic(fmt.Sprintf(
					"Error while creating log file [%s]: %s",
					viper.GetString("log.output"),
					err.Error(),
				))
			}
			defer f.Close()
		}
	}

	if strings.Contains(
		strings.ToLower(viper.GetString("roles")),
		strings.ToLower(AgentRole),
	) {
		InitializeNewAgent().Run()
	}

	if strings.Contains(
		strings.ToLower(viper.GetString("roles")),
		strings.ToLower(WorkerRole),
	) {
		InitializeNewWorker().Run()
	}

	if strings.Contains(
		strings.ToLower(viper.GetString("roles")),
		strings.ToLower(OrchestratorRole),
	) {
		InitializeNewOrchestrator().Run()
	}
}
