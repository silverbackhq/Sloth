// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"github.com/nbio/st"
	"testing"
)

// TestInArray test cases
func TestInArray(t *testing.T) {
	st.Expect(t, true, InArray("A", []string{"A", "B", "C", "D"}))
	st.Expect(t, true, InArray("B", []string{"A", "B", "C", "D"}))
	st.Expect(t, false, InArray("H", []string{"A", "B", "C", "D"}))
	st.Expect(t, true, InArray(1, []int{2, 3, 1}))
	st.Expect(t, false, InArray(9, []int{2, 3, 1}))
}