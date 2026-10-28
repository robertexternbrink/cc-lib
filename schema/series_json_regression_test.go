// Copyright (C) NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved. This file is part of cc-lib.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package schema

import (
	"encoding/json"
	"testing"
)

func TestSeriesJSON_PreservesFloat64Data(t *testing.T) {
	for _, value := range []Float{16777217, -16777217, 1e40, -1e40} {
		series := Series{Hostname: "node001", Data: []Float{value}}
		encoded, err := json.Marshal(&series)
		if err != nil {
			t.Errorf("finite value %v rejected: %v", value, err)
			continue
		}
		var restored Series
		if err := json.Unmarshal(encoded, &restored); err != nil {
			t.Fatal(err)
		}
		if len(restored.Data) != 1 || restored.Data[0] != value {
			t.Errorf("float64 changed: %v -> %v", value, restored.Data)
		}
	}
}
