package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	pb "github.com/spacelavr/pandora/pkg/proto"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetBrokerOpts(ctx context.Context, in *pb.Empty) (*pb.BrokerOpts, error) {
	return &pb.BrokerOpts{
		Endpoint: config.Viper.Broker.Endpoint,
		User:     config.Viper.Broker.User,
		Password: config.Viper.Broker.Password,
	}, nil
}

func Listen() error {
	listen, err := net.Listen("tcp", config.Viper.Tracker.Endpoint)
	if err != nil {
		log.Error(err)
		return err
	}
	defer listen.Close()

	s := grpc.NewServer()
	defer s.GracefulStop()

	pb.RegisterTrackerServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
