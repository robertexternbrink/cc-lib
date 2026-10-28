// Copyright (C) NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved. This file is part of cc-lib.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package ccunits

import (
	"testing"
)

func TestInvalidUnit_Sentinel(t *testing.T) {
	if INVALID_UNIT.Valid() || INVALID_UNIT.GetPrefix() != InvalidPrefix || INVALID_UNIT.GetMeasure() != InvalidMeasure || INVALID_UNIT.GetUnitDenominator() != InvalidMeasure {
		t.Errorf("INVALID_UNIT: valid=%v prefix=%g measure=%v denominator=%v", INVALID_UNIT.Valid(), INVALID_UNIT.GetPrefix(), INVALID_UNIT.GetMeasure(), INVALID_UNIT.GetUnitDenominator())
	}
	convert, out := GetUnitPrefixFactor(NewUnit("unknown"), Base)
	if convert != nil || out.Valid() || out.GetMeasure() != InvalidMeasure {
		t.Errorf("invalid input returned converter=%v, valid output=%v (%q)", convert != nil, out.Valid(), out.Short())
	}
	// Sentinel initialization must not depend on tightening the existing measure parser.
	if u := NewUnit("foobar"); !u.Valid() || u.GetMeasure() != Flops {
		t.Errorf("unrelated permissive parser behavior changed")
	}
}
