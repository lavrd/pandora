package rpc

import (
	"context"
	"fmt"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	pb "github.com/spacelavr/pandora/pkg/proto"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"google.golang.org/grpc"
)

type server struct {
	BrokerOpts *pb.BrokerOpts
}

func init() {

}

func (s *server) GetValidator(ctx context.Context, in *pb.Empty) (*pb.BrokerOpts, error) {
	return &pb.BrokerOpts{
		Endpoint: s.BrokerOpts.Endpoint,
		User:     s.BrokerOpts.User,
		Password: s.BrokerOpts.Password,
	}, nil
}

func (s *server) NewValidator(ctx context.Context, in *pb.BrokerOpts) (*pb.Empty, error) {

	log.Debug(s.BrokerOpts)

	s.BrokerOpts.Endpoint = in.Endpoint
	s.BrokerOpts.User = in.User
	s.BrokerOpts.Password = in.Password

	return &pb.Empty{}, nil
}

func Listen() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Viper.Tracker.Port))
	if err != nil {
		log.Error(err)
		return err
	}
	defer listen.Close()

	s := grpc.NewServer()
	defer s.GracefulStop()

	pb.RegisterTrackerServer(s, &server{BrokerOpts: &pb.BrokerOpts{}})

	if err := s.Serve(listen); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
