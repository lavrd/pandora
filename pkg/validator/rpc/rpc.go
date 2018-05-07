package rpc

import (
	"context"

	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/config"
	pb "github.com/spacelavr/pandora/pkg/proto"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"google.golang.org/grpc"
)

func NotifyTracker(brkOpts *broker.Opts) error {
	conn, err := grpc.Dial(config.Viper.Validator.Tracker, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return err
	}
	defer conn.Close()

	c := pb.NewTrackerClient(conn)

	_, err = c.NewValidator(context.Background(), &pb.BrokerOpts{
		Endpoint: brkOpts.Endpoint,
		User:     brkOpts.User,
		Password: brkOpts.Password,
	})
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
