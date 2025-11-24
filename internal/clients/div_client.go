package clients

import (
	"context"
	"log"
	"sync"
	"time"

	divpb "notask/op-coordinator/api/divpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DivClient struct {
	client divpb.DivServiceClient
}

var divOnce sync.Once
var divInstance *DivClient

func NewDivClient(addr string) *DivClient {
	divOnce.Do(func() {
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Error connecting to DivService: %v", err)
		}

		divInstance = &DivClient{
			client: divpb.NewDivServiceClient(conn),
		}
	})

	return divInstance
}

func (c *DivClient) Compute(a, b float64) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := c.client.Compute(ctx, &divpb.OperationRequest{
		A: a,
		B: b,
	})
	if err != nil {
		return 0, err
	}

	return resp.Result, nil
}
