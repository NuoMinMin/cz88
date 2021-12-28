package service

import (
	"context"
	"cz88/pb"
	"net/http"
	"strconv"

	"cz88/core"

	"github.com/go-kratos/kratos/pkg/ecode"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
)

func (p *service) MyAddress(ctx context.Context, _ *google_protobuf1.Empty) (resp *pb.AddressResp, err error) {
	c, ok := ctx.(*bm.Context)
	if !ok {
		return nil, ecode.Error(ecode.RequestErr, "IP地址参数错误")
	}

	ip := c.RemoteIP()
	if len(ip) == 0 {
		return nil, ecode.Error(ecode.RequestErr, "IP地址参数错误")
	}

	data := core.GetIpInfo(ip)
	return &pb.AddressResp{
		Ip:   data.IP,
		Area: data.Area,
		Isp:  data.Isp,
	}, nil
}

func (p *service) Address(ctx context.Context, req *pb.AddressReq) (resp *pb.AddressResp, err error) {
	if req.Ip == "" {
		return p.MyAddress(ctx, nil)
	}

	if !core.CheckIP(req.Ip) {
		return nil, ecode.Error(ecode.RequestErr, "IP地址参数错误")
	}

	data := core.GetIpInfo(req.Ip)

	return &pb.AddressResp{
		Ip:   data.IP,
		Area: data.Area,
		Isp:  data.Isp,
	}, nil
}

func Metrics(ctx context.Context) {
	metricsNum := GetMetricsNum()
	// IP 库总条目数指标
	str := "ip_service_data_num{service=\"cz88\"} " + strconv.Itoa(metricsNum.Data)
	// 当前进程内缓存 IP 数指标
	str += "\nip_service_cache_num{service=\"cz88\"} " + strconv.Itoa(metricsNum.Cache)
	// 当前程序的协程数
	str += "\nip_service_cpu_num{service=\"cz88\"} " + strconv.Itoa(metricsNum.Cpu)
	str += "\nip_service_goroutine_num{service=\"cz88\"} " + strconv.Itoa(metricsNum.Goroutine)

	ctx.(*bm.Context).String(http.StatusOK, str)
}
