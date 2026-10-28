// Copyright (C) NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved. This file is part of cc-lib.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package ccunits

import (
	"testing"
)

func TestNewUnit_MicroNanoPrefixes(t *testing.T) {
	cases := []struct {
		input       string
		prefix      Prefix
		measure     Measure
		denominator Measure
	}{
		{"us", Micro, Time, InvalidMeasure},
		{"usec", Micro, Time, InvalidMeasure},
		{"useconds", Micro, Time, InvalidMeasure},
		{"ns", Nano, Time, InvalidMeasure},
		{"nsec", Nano, Time, InvalidMeasure},
		{"nseconds", Nano, Time, InvalidMeasure},
		{"us/s", Micro, Time, Time},
		{"ns/s", Nano, Time, Time},
		{"uW", Micro, Watt, InvalidMeasure},
		{"nJ", Nano, Joule, InvalidMeasure},
		{"uHz", Micro, Frequency, InvalidMeasure},
		{"nHz", Nano, Frequency, InvalidMeasure},
	}
	for _, tc := range cases {
		u := NewUnit(tc.input)
		if !u.Valid() || u.GetPrefix() != tc.prefix || u.GetMeasure() != tc.measure || u.GetUnitDenominator() != tc.denominator {
			t.Errorf("NewUnit(%q): valid=%v prefix=%g measure=%v denominator=%v, want %g/%v/%v", tc.input, u.Valid(), u.GetPrefix(), u.GetMeasure(), u.GetUnitDenominator(), tc.prefix, tc.measure, tc.denominator)
		}
	}
	// Micro and nano remain unsupported for the documented non-dividable measures.
	for _, p := range []string{"u", "n"} {
		for _, m := range []string{"B", "Flops", "packets", "events", "cyc", "requests"} {
			for _, d := range []string{"", "/s"} {
				if u := NewUnit(p + m + d); u.Valid() {
					t.Errorf("non-dividable unit %q accepted as %q", p+m+d, u.Short())
				}
			}
		}
	}
	// Characterize every previously supported canonical prefix, measure and denominator.
	prefixes := []struct {
		short string
		value Prefix
	}{
		{"", Base}, {"K", Kilo}, {"M", Mega}, {"G", Giga}, {"T", Tera}, {"P", Peta}, {"E", Exa}, {"Z", Zetta}, {"Y", Yotta},
		{"m", Milli}, {"Ki", Kibi}, {"Mi", Mebi}, {"Gi", Gibi}, {"Ti", Tebi}, {"Pi", Pebi}, {"Ei", Exbi}, {"Zi", Zebi}, {"Yi", Yobi},
	}
	measures := []struct {
		short string
		value Measure
	}{
		{"B", Bytes}, {"Flops", Flops}, {"%", Percentage}, {"degC", TemperatureC}, {"degF", TemperatureF}, {"RPM", Rotation}, {"Hz", Frequency},
		{"s", Time}, {"W", Watt}, {"J", Joule}, {"cyc", Cycles}, {"requests", Requests}, {"packets", Packets}, {"events", Events},
	}
	denominators := append([]struct {
		short string
		value Measure
	}{{"", InvalidMeasure}}, measures...)
	for _, p := range prefixes {
		for _, m := range measures {
			wantPrefix := p.value
			if m.value == Percentage {
				wantPrefix = Base
			}
			if p.value == Milli && (m.value == Bytes || m.value == Flops || m.value == Packets || m.value == Events || m.value == Cycles || m.value == Requests) {
				wantPrefix = Mega
			}
			for _, d := range denominators {
				input := p.short + m.short
				if d.value != InvalidMeasure {
					input += "/" + d.short
				}
				u := NewUnit(input)
				if !u.Valid() || u.GetPrefix() != wantPrefix || u.GetMeasure() != m.value || u.GetUnitDenominator() != d.value {
					t.Errorf("old unit %q changed: valid=%v prefix=%g measure=%v denominator=%v", input, u.Valid(), u.GetPrefix(), u.GetMeasure(), u.GetUnitDenominator())
				}
			}
		}
	}
	for _, s := range []string{"ks", "gs", "ts", "ps", "es", "zs", "ys", "kis", "gis", "tis", "pis", "eis", "zis", "yis"} {
		if !NewUnit(s).Valid() {
			t.Errorf("old prefix alias %q became invalid", s)
		}
	}
}
