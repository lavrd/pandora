package rpc

import (
	"context"

	"github.com/spacelavr/pandora/pkg/config"
	pb "github.com/spacelavr/pandora/pkg/proto"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"google.golang.org/grpc"
)

func GetValidators() (*pb.BrokerOpts, error) {
	conn, err := grpc.Dial(config.Viper.Validator.Tracker, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer conn.Close()

	c := pb.NewTrackerClient(conn)

	r, err := c.GetValidator(context.Background(), &pb.Empty{})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return r, nil
}
