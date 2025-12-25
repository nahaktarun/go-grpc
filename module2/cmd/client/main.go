package main

import (
	"context"
	"io"
	"log"

	"github.com/nahaktarun/grpc-module2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// client := proto.NewHelloServiceClient(conn)

	// res, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: ""})
	// if err != nil {
	// 	status, ok := status.FromError(err)
	// 	if ok {
	// 		log.Fatalf("status code: %s, error: %s", status.Code().String(), status.Message())
	// 	}
	// 	log.Fatal(err)
	// }

	// client := proto.NewTodoServiceClient(conn)

	// task1, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "Wake up"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Existing tasks: %s", task1.GetId())

	// task2, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "Study"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Existing tasks: %s", task2.GetId())

	// tasks, err := client.ListTask(ctx, &proto.ListTasksRequest{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Existing tasks: %s", tasks.GetTasks())

	// _, err = client.CompleteTask(ctx, &proto.CompleteTaskRequest{Id: task1.Id})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// task3, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "Sleep"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Existing tasks: %s", task3.GetId())

	// tasks, err = client.ListTask(ctx, &proto.ListTasksRequest{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Existing tasks: %s", tasks.GetTasks())

	// 2. server side streaming

	// 2.a first initalize the grpc connection
	// 2.b create the client
	// 2.c initalize the stream
	// 2.d loop through all the responses we get back from the server
	// 	log it
	// 2.e once the server closes the stream exit gracefully
	//
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := proto.NewStreamingServiceClient(conn)

	stream, err := client.StreamServerTime(ctx, &proto.StreamServerTimeRequest{
		IntervalSeconds: 2,
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		log.Printf("received time from server: %s", res.CurrentTime.AsTime())
	}

	log.Println("server stream closed")
}
