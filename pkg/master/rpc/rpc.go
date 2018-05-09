package rpc

// type server struct{}
//
// func (s *server) Candidate(ctx context.Context, in *pb.Block) (*pb.Empty, error) {
// 	return &pb.Empty{}, nil
// }
//
// func Listen() error {
// 	listen, err := net.Listen(network.TCP, config.Viper.Master.Endpoint)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}
// 	defer listen.Close()
//
// 	s := grpc.NewServer()
// 	defer s.GracefulStop()
//
// 	pb.RegisterMasterServer(s, &server{})
//
// 	if err := s.Serve(listen); err != nil {
// 		log.Error(err)
// 		return err
// 	}
//
// 	return nil
// }
