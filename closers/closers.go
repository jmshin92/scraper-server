package closers

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type Closer interface {
	Close()
}

var (
	closers []Closer
	lock sync.Mutex
)

func AddClosers(c Closer) {
	lock.Lock()
	defer lock.Unlock()
	
	closers = append(closers, c)
}


func Close() {
	lock.Lock()
	defer lock.Unlock()
	
	logrus.Info("start to close all the closers")
	defer logrus.Info("closed all the closers")
	for _, c := range closers{
		c.Close()
	}
	closers = closers[len(closers):]
}