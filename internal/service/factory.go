package service

type service struct {
}

func New() *service {
	//cc, err := grpc.Dial(config.GetInstance().MiddleServerRpc, grpc.WithInsecure())
	//if err != nil {
	//	panic(err)
	//}
	//rpcClient := middle.NewRpcClient(cc)
	//return &Service{
	//	middleDao: middle.NewMiddle(
	//		config.GetInstance().MiddleServerDomain,
	//		rpcClient,
	//	),
	//	rpcClient: rpcClient,
	//}
	return &service{}
}
