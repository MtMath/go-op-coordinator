package clients

import (
	"context"
	"log"
	"sync"
	"time"

	addpb "notask/op-coordinator/api/addpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AddClient struct {
	client addpb.AddServiceClient
}

var addOnce sync.Once
var addInstance *AddClient

func NewAddClient(addr string) *AddClient {
	addOnce.Do(func() {
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Error connecting to AddService: %v", err)
		}

		addInstance = &AddClient{
			client: addpb.NewAddServiceClient(conn),
		}
	})

	return addInstance
}

func (c *AddClient) Compute(a, b float64) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := c.client.Compute(ctx, &addpb.OperationRequest{
		A: a,
		B: b,
	})
	if err != nil {
		return 0, err
	}

	return resp.Result, nil
}
