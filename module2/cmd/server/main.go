package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/nahaktarun/grpc-module2/internal/streaming"
	"github.com/nahaktarun/grpc-module2/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	defer cancel()

	if err := run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("error running application", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// log graceful shutdown
	slog.Info("closing server gracefully")
}

func run(ctx context.Context) error {
	grpcServer := grpc.NewServer()

	// helloService := hello.Service{}
	// todoService := todo.NewService()

	streamService := &streaming.Service{}

	// proto.RegisterHelloServiceServer(grpcServer, helloService)

	proto.RegisterStreamingServiceServer(grpcServer, streamService)

	const addr = ":50051"

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			return fmt.Errorf("failed to listen on address %q: %w", addr, err)
		}

		slog.Info("Starting server on the port", slog.String("address", addr))
		if err := grpcServer.Serve(lis); err != nil {
			return fmt.Errorf("failed to serve grpc service: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		grpcServer.GracefulStop()
		return nil
	})

	return g.Wait()
}
