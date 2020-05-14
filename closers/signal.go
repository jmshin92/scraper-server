package closers

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/sirupsen/logrus"
)

func init() {
	go signalHandler()
}

func signalHandler() {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	for {
		select {
		case s := <-sigCh:
			// Signal handling
			logrus.Info("received signal[%v]. start to closing closers...", s)
			shutdown()
		}
	}
}

func shutdown() {
	timeout := 60 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	go func() {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					logrus.Info("shutdown timeout[%v] exceeded. force shutting down...", timeout)
					os.Exit(1)
				}
			}
		}
	}()
	
	Close()
	os.Exit(0)
}

func AddCloser(c Closer) {
	closers = append(closers, c)
}
