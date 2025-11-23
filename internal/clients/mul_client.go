package clients

import (
	"context"
	"log"
	"sync"
	"time"

	mulpb "notask/op-coordinator/api/mulpb"

	"google.golang.org/grpc"
)

type MulClient struct {
	client mulpb.MulServiceClient
}

var mulOnce sync.Once
var mulInstance *MulClient

func NewMulClient(addr string) *MulClient {
	mulOnce.Do(func() {
		conn, err := grpc.NewClient(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Error connecting to MulService: %v", err)
		}

		mulInstance = &MulClient{
			client: mulpb.NewMulServiceClient(conn),
		}
	})

	return mulInstance
}

func (c *MulClient) Compute(a, b float64) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := c.client.Compute(ctx, &mulpb.OperationRequest{
		A: a,
		B: b,
	})
	if err != nil {
		return 0, err
	}

	return resp.Result, nil
}
