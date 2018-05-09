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
	t := time.Now()
	cc, err := grpc.Dial(config.Viper.Membership.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return err
	}
	defer cc.Close()
	log.Debug(time.Since(t))

	c := pb.NewMembershipClient(cc)

	t = time.Now()
	_, err = c.Register(context.Background(), candidate)
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
	log.Debug(time.Since(t))

	return nil
}

// todo through pb struct or only string and convert in this func
func FetchAccount(key *pb.PublicKey) (*pb.Account, error) {
	cc, err := grpc.Dial(config.Viper.Membership.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cc.Close()

	c := pb.NewMembershipClient(cc)

	r, err := c.Fetch(context.Background(), key)
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

// todo rename and rename at .proto
func Network() (*pb.NetOpts, error) {
	t := time.Now()
	cc, err := grpc.Dial(config.Viper.Discovery.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cc.Close()
	log.Debug(time.Since(t))

	t = time.Now()
	c := pb.NewDiscoveryClient(cc)
	log.Debug(time.Since(t))

	t = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	r, err := c.Network(ctx, &pb.Empty{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debug(time.Since(t))

	return r, nil
}
