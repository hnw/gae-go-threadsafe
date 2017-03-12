// How to deploy:
//   $ appcfg.py update . -A [application_id]

// +build appengine

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"google.golang.org/appengine"
	aelog "google.golang.org/appengine/log"
)

var (
	m   sync.Mutex
	cnt int
)

func init() {
	http.HandleFunc("/sleep", sleepHandler)
}

func sleepHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	startTime := time.Now()
	aelog.Infof(ctx, "Now: %v", startTime)
	aelog.Infof(ctx, "InstanceID: %v", appengine.InstanceID())

	m.Lock()
	cnt++
	curCnt := cnt
	time.Sleep(100 * time.Millisecond)
	m.Unlock()

	aelog.Infof(ctx, "Cnt: %v", curCnt)
	time.Sleep(400 * time.Millisecond)
	aelog.Infof(ctx, "500ms after: %v", time.Now())
	time.Sleep(500 * time.Millisecond)
	aelog.Infof(ctx, "1sec after: %v", time.Now())

	fmt.Fprintln(w, "<html><pre>")
	fmt.Fprintf(w, "duration: %v\n", time.Now().Sub(startTime))
	fmt.Fprintln(w, "</pre></html>")
}
