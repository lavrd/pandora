package rpc

import (
	"context"
	"time"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Register(candidate *pb.Candidate) error {
	cc, err := grpc.Dial(config.Viper.Membership.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return err
	}
	defer cc.Close()

	c := pb.NewMembershipClient(cc)

	_, err = c.ConfirmMember(context.Background(), candidate)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return err
		}
		if st.Code() == codes.AlreadyExists {
			return errors.AlreadyExists
		}

		log.Error(err)
		return err
	}

	return nil
}

func Issue(cert *pb.Cert) error {
	cc, err := grpc.Dial(config.Viper.Membership.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
	}
	defer cc.Close()

	c := pb.NewMembershipClient(cc)

	_, err = c.SignCert(context.Background(), cert)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func NodeReg(candidate *pb.Candidate) (*pb.PublicKey, error) {
	cc, err := grpc.Dial(config.Viper.Membership.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cc.Close()

	c := pb.NewMembershipClient(cc)

	r, err := c.ConfirmMember(context.Background(), candidate)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, err
		}
		if st.Code() == codes.AlreadyExists {
			log.Debug(r)
			return r, errors.AlreadyExists
		}

		log.Error(err)
		return nil, err
	}

	return r, nil
}

// todo through pb struct or only string and convert in this func
func FetchAccount(key *pb.PublicKey) (*pb.Member, error) {
	cc, err := grpc.Dial(config.Viper.Membership.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cc.Close()

	c := pb.NewMembershipClient(cc)

	r, err := c.FetchMember(context.Background(), key)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, err
		}
		if st.Code() == codes.NotFound {
			return nil, errors.NotFound
		}

		log.Error(err)
		return nil, err
	}

	return r, nil
}

func Node(key *pb.PublicKey) (*pb.MasterChain, error) {
	cc, err := grpc.Dial(config.Viper.Master.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cc.Close()

	c := pb.NewMasterClient(cc)

	mc, err := c.ConfirmNode(context.Background(), key)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return mc, nil
}

// todo rename and rename at .proto
func Network() (*pb.NetworkOpts, error) {
	cc, err := grpc.Dial(config.Viper.Discovery.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cc.Close()

	c := pb.NewDiscoveryClient(cc)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	r, err := c.Network(ctx, &pb.Empty{})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return r, nil
}
