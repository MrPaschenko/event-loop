package engine

import (
	"sync"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type Queue struct {
	sync.Mutex
	data             []Command
	receiveSignal    chan struct{}
	receiveRequested bool
}

type Loop struct {
	mq          *Queue
	stopSignal  chan struct{}
	stopRequest bool
}

func (mq *Queue) push(command Command) {
	mq.Lock()
	defer mq.Unlock()

	mq.data = append(mq.data, command)
	if mq.receiveRequested {
		mq.receiveRequested = false
		mq.receiveSignal <- struct{}{}
	}

}

func (mq *Queue) pull() Command {
	mq.Lock()
	defer mq.Unlock()

	if mq.empty() {
		mq.receiveRequested = true
		mq.Unlock()
		<-mq.receiveSignal
		mq.Lock()
	}

	res := mq.data[0]
	mq.data[0] = nil
	mq.data = mq.data[1:]
	return res
}

func (mq *Queue) empty() bool {
	return len(mq.data) == 0
}

func (l *Loop) Start() {
	l.mq = &Queue{receiveSignal: make(chan struct{})}
	l.stopSignal = make(chan struct{})
	go func() {
		for !l.stopRequest || !l.mq.empty() {
			cmd := l.mq.pull()
			cmd.Execute(l)
		}
		l.stopSignal <- struct{}{}
	}()
}

func (l *Loop) Post(cmd Command) {
	l.mq.push(cmd)
}

type CommandFunc func(h Handler)

func (cf CommandFunc) Execute(h Handler) {
	cf(h)
}

func (l *Loop) AwaitFinish() {
	l.Post(CommandFunc(func(h Handler) {
		l.stopRequest = true
	}))
	<-l.stopSignal
}
