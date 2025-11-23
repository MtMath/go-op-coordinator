package clients

import (
	"context"
	"log"
	"sync"
	"time"

	subpb "notask/op-coordinator/api/subpb"

	"google.golang.org/grpc"
)

type SubClient struct {
	client subpb.SubServiceClient
}

var subOnce sync.Once
var subInstance *SubClient

func NewSubClient(addr string) *SubClient {
	subOnce.Do(func() {
		conn, err := grpc.NewClient(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Error connecting to SubService: %v", err)
		}

		subInstance = &SubClient{
			client: subpb.NewSubServiceClient(conn),
		}
	})

	return subInstance
}

func (c *SubClient) Compute(a, b float64) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := c.client.Compute(ctx, &subpb.OperationRequest{
		A: a,
		B: b,
	})
	if err != nil {
		return 0, err
	}

	return resp.Result, nil
}
