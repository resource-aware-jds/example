package main

import (
	"context"
	"fmt"
	"github.com/resource-aware-jds/common-go/proto"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

func main() {

	grpcPath := "host.docker.internal:" + os.Getenv("HOST_PORT")
	fmt.Println("Sending request to: ", grpcPath)
	grpcConn, err := grpc.Dial(grpcPath, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic("Can't connect gRPC Server: ", err)
	}
	defer grpcConn.Close()

	computeNodeClient := proto.NewComputeNodeClient(grpcConn)
	ctx := context.Background()

	strJobID := os.Getenv("JOB_ID")
	jobID := int32(0)
	i, err := strconv.ParseInt(strJobID, 10, 32)
	if err == nil {
		jobID = int32(i)
	}

	totalJob := 100

	for i := 1; i < totalJob; i++ {
		fmt.Println("Current: ", i)
		computeNodeClient.ReportJob(ctx, &proto.ReportJobRequest{JobID: jobID, TotalJob: int64(totalJob), CurrentJob: int64(i)})
		time.Sleep(1 * time.Second)
	}
}
