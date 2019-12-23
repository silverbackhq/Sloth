// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/nbio/st"
	"testing"
)

// TestHttpGet test cases
func TestHttpGet(t *testing.T) {
	encryptionKey, _ := GenerateRandomString(32)
	text := []byte(`{"Id":1,"hostname":"localhost","message":"Hello World","type":"agent"}`)
	ciphertext, err := Encrypt(text, []byte(encryptionKey))
	st.Expect(t, nil, err)
	plaintext, err := Decrypt(ciphertext, []byte(encryptionKey))
	st.Expect(t, nil, err)
	st.Expect(t, text, plaintext)
	st.Expect(t, string(text), string(plaintext))
}
