package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/nahaktarun/grpc-module2/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// init our grpc connection

	// create a client
	// initialize our stream
	//
	//
	// Create a separate go routine to list the server response
	// 	loop for each message from server
	//
	// log the message
	//
	// check if stream is closed
	//
	// close the client stream
	//
	// send some message from the client
	// close the client stream
	// wait for the server go routine to finish

	ctx := context.Background()
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewStreamingServiceClient(conn)
	stream, err := client.Echo(ctx)

	if err != nil {
		log.Fatal(err)
	}
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {

		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}

			log.Printf("message received from server: %s", res.GetMessage())

		}
		return nil
	})

	for i := range 5 {
		req := &proto.EchoRequest{
			Message: fmt.Sprintf("Hello %d", i),
		}
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 2)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatal(err)
	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("bi-directional stream closed")
}
