package http

import (
	"cz88/pb"
	"fmt"
	"time"

	"cz88/config"
	"cz88/internal/service"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	xtime "github.com/go-kratos/kratos/pkg/time"
)

func New() {
	r := bm.DefaultServer(&bm.ServerConfig{
		Timeout: xtime.Duration(time.Second * 10),
	})

	pb.RegisterAppBMServer(r, service.New())
	fmt.Println("http server addr: ", config.GetInstance().Http)
	if err := r.Run(config.GetInstance().Http); err != nil {
		panic(err)
	}
}
