package main

import (
	"context"
	v1 "github.com/golang-learning/inventory/api/storage/v1"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	service, err := InitStorageService()
	if err != nil {
		log.Panicf("service init fail: %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterStorageServiceServer(s, service)

	// signal control
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		<-ctx.Done()
		log.Println("shutting down server...")
		s.GracefulStop()
		return nil
	})

	g.Go(func() error {
		l, err := net.Listen("tcp", ":8080")
		if err != nil {
			errors.Wrap(err, "start server port :8080")
		}

		log.Println("grpc server will listen on :8080")
		return s.Serve(l)
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("get os signal: %v", sig)
		}
	})

	log.Printf("errgroup existing: %+v\n", g.Wait())
}
