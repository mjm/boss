package server

import (
	"log"

	"github.com/mjm/boss/cfg"
)

type startTaskMsg struct{}

type taskWorker struct {
	Task *cfg.Task
	Ch   chan interface{}
	Done chan struct{}
}

func (w *taskWorker) Run() {
	for {
		select {
		case msg := <-w.Ch:
			log.Printf("[%s] got msg %#v", w.Task.Name, msg)
		case <-w.Done:
			log.Printf("[%s] stopping", w.Task.Name)
			return
		}
	}
}

func (w *taskWorker) Stop() {
	close(w.Done)
}
