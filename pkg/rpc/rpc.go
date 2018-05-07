package rpc

import (
	"context"

	pb "github.com/spacelavr/pandora/pkg/proto"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"google.golang.org/grpc"
)

func GetBrokerOpts(endpoint string) (*pb.BrokerOpts, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer conn.Close()

	c := pb.NewTrackerClient(conn)

	r, err := c.GetBrokerOpts(context.Background(), &pb.Empty{})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return r, nil
}
