package server

import (
	"log"
	"os"
	"os/exec"

	"github.com/mjm/boss/cfg"
)

type startTaskMsg struct{}
type taskExitedMsg struct {
	Err error
}

type taskWorker struct {
	Task *cfg.Task
	Ch   chan interface{}
	Done chan struct{}

	cmd *exec.Cmd
}

func (w *taskWorker) Run() {
	for {
		select {
		case msg := <-w.Ch:
			if err := w.handleMessage(msg); err != nil {
				log.Printf("[%s] error handling %T message: %v", w.Task.Name, msg, err)
				return
			}
		case <-w.Done:
			log.Printf("[%s] stopping", w.Task.Name)
			return
		}
	}
}

func (w *taskWorker) Stop() {
	close(w.Done)
}

func (w *taskWorker) handleMessage(msg interface{}) error {
	switch m := msg.(type) {
	case startTaskMsg:
		return w.start()
	case taskExitedMsg:
		if m.Err == nil {
			log.Printf("[%s] exited normally", w.Task.Name)
		} else {
			log.Printf("[%s] exited with error: %v", w.Task.Name, m.Err)
		}
	}

	log.Printf("[%s] unexpected message type %T", w.Task.Name, msg)
	return nil
}

func (w *taskWorker) start() error {
	w.cmd = exec.Command("/bin/sh", "-c", w.Task.Command)
	// TODO connect these to the logger
	w.cmd.Stdout = os.Stdout
	w.cmd.Stderr = os.Stderr

	if err := w.cmd.Start(); err != nil {
		return err
	}

	go func() {
		err := w.cmd.Wait()
		w.Ch <- taskExitedMsg{Err: err}
	}()
	return nil
}
