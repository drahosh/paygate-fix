// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package admin

import (
	"net/http"
	"testing"

	"github.com/moov-io/base/admin"

	"github.com/go-kit/kit/log"
)

func TestFlushIncomingFiles(t *testing.T) {
	svc := admin.NewServer(":0")
	go svc.Listen()
	defer svc.Shutdown()

	flushIncoming := make(FlushChan, 1)
	RegisterAdminRoutes(log.NewNopLogger(), svc, flushIncoming, nil, getZeroFiles)

	// invalid request, wrong HTTP verb
	req, err := http.NewRequest("GET", "http://"+svc.BindAddr()+"/files/flush/incoming", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d", resp.StatusCode)
	}

	// valid request
	req, err = http.NewRequest("POST", "http://"+svc.BindAddr()+"/files/flush/incoming", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", resp.StatusCode)
	}

	<-flushIncoming
}

func TestFlushOutgoingFiles(t *testing.T) {
	svc := admin.NewServer(":0")
	go svc.Listen()
	defer svc.Shutdown()

	flushOutgoing := make(FlushChan, 1)
	RegisterAdminRoutes(log.NewNopLogger(), svc, nil, flushOutgoing, getZeroFiles)

	// invalid request, wrong HTTP verb
	req, err := http.NewRequest("GET", "http://"+svc.BindAddr()+"/files/flush/outgoing", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d", resp.StatusCode)
	}

	// valid request
	req, err = http.NewRequest("POST", "http://"+svc.BindAddr()+"/files/flush/outgoing", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", resp.StatusCode)
	}

	<-flushOutgoing
}

func TestFlushFilesUpload(t *testing.T) {
	svc := admin.NewServer(":0")
	go svc.Listen()
	defer svc.Shutdown()

	flushIncoming, flushOutgoing := make(FlushChan, 1), make(FlushChan, 1) // buffered channel
	RegisterAdminRoutes(log.NewNopLogger(), svc, flushIncoming, flushOutgoing, getZeroFiles)

	req, err := http.NewRequest("POST", "http://"+svc.BindAddr()+"/files/flush", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", resp.StatusCode)
	}

	// we need to read from this channel to ensure a message was sent
	// if there's no message the test will timeout
	<-flushIncoming
	<-flushOutgoing

	// use the wrong HTTP verb and get an error
	req, err = http.NewRequest("GET", "http://"+svc.BindAddr()+"/files/flush", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d", resp.StatusCode)
	}
}
