// Copyright (C) NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved. This file is part of cc-lib.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package resampler

import (
	"github.com/ClusterCockpit/cc-lib/v2/schema"
	"testing"
)

func TestLTTB_SingleTargetUsesConsistentResolution(t *testing.T) {
	data := make([]schema.Float, 1000)
	for i := range data {
		data[i] = schema.Float(i)
	}
	got, step, err := LargestTriangleThreeBucket(data, 1, 1000)
	if err != nil {
		t.Fatal(err)
	}
	if !((len(got) == 1 && step == 1000) || (len(got) == len(data) && step == 1)) {
		t.Fatalf("inconsistent output: len=%d step=%d", len(got), step)
	}
}
