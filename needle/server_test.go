// Copyright 2010 GoNeedle Authors. All rights reserved.
// Use of this source code is governed by a 
// license that can be found in the LICENSE file.

package needle

import (
	//"fmt"
	"testing"
	"time"
)

func startServer(t *testing.T) {
	_, err := MakeServer(":62077", ":62070")
	if err != nil {
		t.Fatalf("Starting server %s\n", err)
	}
}

func startListener(id int64, t *testing.T) {
	_, err := MakeListener(id, ":34000", "127.0.0.1:62077")
	if err != nil {
		t.Fatalf("Starting listener %s\n", err)
	}
}

func TestServer(t *testing.T) {
	startServer(t)
	time.Sleep(1e9)
	startListener(1, t)
	<-make(chan int)
}