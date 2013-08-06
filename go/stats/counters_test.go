// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stats

import (
	"expvar"
	"testing"
)

func TestCounters(t *testing.T) {
	Register(nil)
	c := NewCounters("counter1")
	c.Add("c1", 1)
	c.Add("c2", 1)
	c.Add("c2", 1)
	want := `{"c1": 1, "c2": 2}`
	if c.String() != want {
		t.Errorf("want %s, got %s", want, c.String())
	}
	counts := c.Counts()
	if counts["c1"] != 1 {
		t.Errorf("want 1, got %d", counts["c1"])
	}
	if counts["c2"] != 2 {
		t.Errorf("want 2, got %d", counts["c2"])
	}
}

func TestCountersHook(t *testing.T) {
	var gotname string
	var gotv *Counters
	Register(func(name string, v expvar.Var) {
		gotname = name
		gotv = v.(*Counters)
	})

	v := NewCounters("counter2")
	if gotname != "counter2" {
		t.Errorf("want counter2, got %s", gotname)
	}
	if gotv != v {
		t.Errorf("want %#v, got %#v", v, gotv)
	}
}