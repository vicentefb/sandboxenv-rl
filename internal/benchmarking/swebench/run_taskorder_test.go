// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swebench

import "testing"

func TestRoundOrderInterleavesReplays(t *testing.T) {
	const instances, fanout = 3, 2
	var envs []*environment
	for i := 0; i < instances; i++ {
		for k := 0; k < fanout; k++ {
			envs = append(envs, &environment{inst: Instance{InstanceID: string(rune('a' + i))}, task: k})
		}
	}
	got := roundOrder(envs, instances, fanout)
	want := []string{"a0", "b0", "c0", "a1", "b1", "c1"}
	if len(got) != len(want) {
		t.Fatalf("roundOrder returned %d environments, want %d", len(got), len(want))
	}
	for i, e := range got {
		if id := e.inst.InstanceID + string(rune('0'+e.task)); id != want[i] {
			t.Errorf("position %d: got %s, want %s", i, id, want[i])
		}
	}
	if same := roundOrder(envs, instances, 1); len(same) != len(envs) {
		t.Errorf("fanout 1 must return the input unchanged")
	}
}

func TestValidateTaskOrder(t *testing.T) {
	c := Config{Mode: ModeEval, TasksPerImage: 1, TaskOrder: "sideways"}
	if err := c.validate(); err == nil {
		t.Fatal("validate accepted an unknown task order")
	}
}
