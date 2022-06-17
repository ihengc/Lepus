package Lepus

import (
	"sync"
	"testing"
	"time"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 14:22
* @version: 1.0
* @description:
*********************************************************/

var app *Application

func setup() {
	app = NewApplication()
}

func TestApplication_Run(t *testing.T) {
	setup()
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	go func() {
		app.Run()
		waitGroup.Done()
	}()
	go func() {
		time.Sleep(3 * time.Second)
		t.Log("Stop 1")
		app.Stop()
		time.Sleep(3 * time.Second)
		t.Log("Stop 2")
		app.Stop()
		waitGroup.Done()
	}()
	waitGroup.Wait()
}
