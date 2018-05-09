package rpc

import (
	"context"

	"github.com/spacelavr/pandora/pkg/config"
	dpb "github.com/spacelavr/pandora/pkg/discovery/pb"
	mpb "github.com/spacelavr/pandora/pkg/membership/pb"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Register(candidate *mpb.Candidate) error {
	conn, err := grpc.Dial(config.Viper.Membership.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return err
	}
	defer conn.Close()

	c := mpb.NewMembershipClient(conn)

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

	return nil
}

func Network() (*dpb.Net, error) {
	conn, err := grpc.Dial(config.Viper.Discovery.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer conn.Close()

	c := dpb.NewDiscoveryClient(conn)

	r, err := c.Network(context.Background(), &dpb.Empty{})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return r, nil
}
