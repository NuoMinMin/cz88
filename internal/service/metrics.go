package service

import (
	"runtime"

	"cz88/core"
)

type TMetricsNum struct {
	Data      int
	Cache     int
	Cpu       int
	Goroutine int
}

func GetMetricsNum() *TMetricsNum {
	ipData := core.LoadIpData()

	cacheNum := 0
	ipData.Cache.Range(func(_, _ interface{}) bool {
		cacheNum++
		return true
	})
	return &TMetricsNum{
		Data:      len(ipData.List),
		Cache:     cacheNum,
		Cpu:       runtime.NumCPU(),
		Goroutine: runtime.NumGoroutine(),
	}
}
