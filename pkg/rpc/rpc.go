package rpc

// func GetBrokerOpts(endpoint string) (*pb.BrokerOpts, error) {
// 	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}
// 	defer conn.Close()
//
// 	c := pb.NewTrackerClient(conn)
//
// 	r, err := c.GetBrokerOpts(context.Background(), &pb.Empty{})
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}
//
// 	return r, nil
// }

