package main

import (
	"context"
	"log"

	"github.com/nahaktarun/grpc-module2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// client := proto.NewHelloServiceClient(conn)

	// res, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: ""})
	// if err != nil {
	// 	status, ok := status.FromError(err)
	// 	if ok {
	// 		log.Fatalf("status code: %s, error: %s", status.Code().String(), status.Message())
	// 	}
	// 	log.Fatal(err)
	// }

	client := proto.NewTodoServiceClient(conn)

	task1, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "Wake up"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Existing tasks: %s", task1.GetId())

	task2, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "Study"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Existing tasks: %s", task2.GetId())

	tasks, err := client.ListTask(ctx, &proto.ListTasksRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Existing tasks: %s", tasks.GetTasks())

	_, err = client.CompleteTask(ctx, &proto.CompleteTaskRequest{Id: task1.Id})
	if err != nil {
		log.Fatal(err)
	}

	task3, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "Sleep"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Existing tasks: %s", task3.GetId())

	tasks, err = client.ListTask(ctx, &proto.ListTasksRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Existing tasks: %s", tasks.GetTasks())
}
