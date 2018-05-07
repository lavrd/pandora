package rpc

import (
	"context"

	"github.com/spacelavr/pandora/pkg/config"
	pb "github.com/spacelavr/pandora/pkg/proto"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/utils/network"
	"google.golang.org/grpc"
)

func NotifyTracker() error {
	conn, err := grpc.Dial(config.Viper.Validator.Tracker, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return err
	}
	defer conn.Close()

	c := pb.NewTrackerClient(conn)

	myIP := network.IP()

	_, err = c.NewValidator(context.Background(), &pb.NVRequest{Ip: myIP})
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
