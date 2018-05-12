package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/membership/distribution"
	"github.com/spacelavr/pandora/pkg/membership/env"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/utils/network"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (s *server) ConfirmMember(ctx context.Context, in *pb.Candidate) (*pb.PublicKey, error) {
	dist := &distribution.Distribution{
		Storage: env.GetStorage(),
		Runtime: env.GetRuntime(),
	}

	key, err := dist.AcceptCandidate(in)
	if err != nil {
		if err == errors.AlreadyExists {
			log.Debug(key)
			return key, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
		}
		return &pb.PublicKey{}, err
	}

	return key, nil
}

func (s *server) Node(ctx context.Context, in *pb.Candidate) (*pb.PublicKey, error) {
	dist := &distribution.Distribution{
		Storage: env.GetStorage(),
		Runtime: env.GetRuntime(),
	}

	log.Debug(99999, in)

	key, err := dist.AcceptCandidate(in)
	if err != nil && err != errors.AlreadyExists {
		log.Debug(1)
		return &pb.PublicKey{}, err
	}

	log.Debug(8777)
	log.Debug(key)

	return key, nil
}

func (s *server) SignCert(ctx context.Context, in *pb.Cert) (*pb.Empty, error) {
	cert, err := distribution.New().Issue(in)
	if err != nil {
		return &pb.Empty{}, err
	}

	log.Debug(10)
	if err := Issue(cert); err != nil {
		return &pb.Empty{}, err
	}
	log.Debug(11)

	return &pb.Empty{}, nil
}

func Issue(cert *pb.Cert) error {
	cc, err := grpc.Dial(config.Viper.Master.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
	}
	defer cc.Close()

	c := pb.NewMasterClient(cc)

	_, err = c.ConfirmCert(context.Background(), cert)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *server) FetchMember(ctx context.Context, in *pb.PublicKey) (*pb.Member, error) {
	dist := &distribution.Distribution{
		Storage: env.GetStorage(),
		Runtime: env.GetRuntime(),
	}

	acc, err := dist.AccountFetch(in)
	if err != nil {
		if err == errors.NotFound {
			// todo may be nil? and in all place
			return &pb.Member{}, status.Error(codes.NotFound, codes.NotFound.String())
		}
		return &pb.Member{}, err
	}
	return acc, nil
}

func Listen() error {
	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(config.Viper.Membership.Endpoint))
	if err != nil {
		log.Error(err)
		return err
	}
	defer listen.Close()

	s := grpc.NewServer()
	defer s.GracefulStop()

	pb.RegisterMembershipServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
