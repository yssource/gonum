// Code generated by "go generate gonum.org/v1/gonum/unit; DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"fmt"
	"testing"
)

func TestMole(t *testing.T) {
	for _, value := range []float64{-1, 0, 1} {
		var got Mole
		err := got.From(Mole(value).Unit())
		if err != nil {
			t.Errorf("unexpected error for %T conversion: %v", got, err)
		}
		if got != Mole(value) {
			t.Errorf("unexpected result from round trip of %T(%v): got: %v want: %v", got, float64(value), got, value)
		}
	}
}

func TestMoleFormat(t *testing.T) {
	for _, test := range []struct {
		value  Mole
		format string
		want   string
	}{
		{1.23456789, "%v", "1.23456789 mol"},
		{1.23456789, "%.1v", "1 mol"},
		{1.23456789, "%20.1v", "               1 mol"},
		{1.23456789, "%20v", "      1.23456789 mol"},
		{1.23456789, "%1v", "1.23456789 mol"},
		{1.23456789, "%#v", "unit.Mole(1.23456789)"},
		{1.23456789, "%s", "%!s(unit.Mole=1.23456789 mol)"},
	} {
		got := fmt.Sprintf(test.format, test.value)
		if got != test.want {
			t.Errorf("Format %q %v: got: %q want: %q", test.format, float64(test.value), got, test.want)
		}
	}
}
