package rpc

import (
	"context"

	"github.com/spacelavr/pandora/pkg/config"
	pb "github.com/spacelavr/pandora/pkg/proto"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"google.golang.org/grpc"
)

func GetValidators() error {
	conn, err := grpc.Dial(config.Viper.Validator.Tracker, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return err
	}
	defer conn.Close()

	c := pb.NewTrackerClient(conn)

	r, err := c.GetValidators(context.Background(), &pb.GVRequest{})
	if err != nil {
		log.Error(err)
		return err
	}

	log.Error(r.Ips)

	return nil
}
