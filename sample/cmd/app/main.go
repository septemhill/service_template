package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sample/internal/service"
	"sample/internal/transport/http"
	"syscall"
	"time"
)

const workerDoneTimeout = 2 * time.Second

type stopFunc func(context.Context)

type worker interface {
	Start(context.Context) error
	Stop(context.Context)
}

func serviceInit(ctx context.Context, workers ...worker) stopFunc {
	for i := 0; i < len(workers); i++ {
		if err := workers[i].Start(ctx); err != nil {
			panic(fmt.Sprintf("service start failed: %v", err))
		}
	}

	return func(ctx context.Context) {
		for i := 0; i < len(workers); i++ {
			workers[i].Stop(ctx)
		}
	}
}

func systemSignalInit(ctx context.Context, stop stopFunc, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-ch

		// send cancel signal
		cancel()

		// Shut all of workers down
		stop(ctx)

		time.Sleep(workerDoneTimeout)
	}()
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	service := service.NewService()
	httpTransport := http.NewHTTPTransport(service, 8888)

	stopper := serviceInit(ctx, httpTransport)
	systemSignalInit(ctx, stopper, cancel)

	<-ctx.Done()
}
