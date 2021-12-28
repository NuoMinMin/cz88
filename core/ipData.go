package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"cz88/config"

	"github.com/NuoMinMin/mahonia"
)

type IpInfoItem struct {
	IP      string `json:"ip"`
	Area    string `json:"area"`
	Isp     string `json:"isp"`
	IpStart int64  `json:"-"`
	IpEnd   int64  `json:"-"`
}

type IpData struct {
	List       []IpInfoItem
	ListLength int
	Cache      sync.Map
}

var ipData *IpData

// 加载 IP 地址库
func LoadIpData() *IpData {
	if ipData == nil {
		path := filepath.Dir(os.Args[0]) + "/" + config.GetInstance().CZip.FilePath
		fp, err := os.Open(path)

		if err != nil {
			panic(err)
		}

		defer fp.Close()
		decoder := mahonia.NewDecoder(config.GetInstance().CZip.Charset)
		if decoder == nil {
			panic("编码不存在")
		}

		ipData = &IpData{
			List: []IpInfoItem{},
		}

		reader := bufio.NewReader(fp)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				fmt.Println("IP 地址加载完成")
				break
			}
			if err != nil {
				panic(err)
			}

			lineSlice := strings.Fields(decoder.ConvertString(line))

			if len(lineSlice) < 3 || lineSlice[0] == "IP数据库共有数据" {
				continue
			}

			ipItem := IpInfoItem{
				Area:    lineSlice[2],
				Isp:     strings.Trim(fmt.Sprint(lineSlice[3:]), "[]"),
				IpStart: calcIP(lineSlice[0]),
				IpEnd:   calcIP(lineSlice[1]),
			}
			ipData.List = append(ipData.List, ipItem)
			ipData.ListLength++
		}

	}
	return ipData
}

// 匹配 IP 地址信息
func GetIpInfo(ip string) (res IpInfoItem) {
	ipData := LoadIpData()
	value, ok := ipData.Cache.Load(ip)
	if ok {
		return value.(IpInfoItem)
	}

	ipTarget := calcIP(ip)
	for _, item := range ipData.List {
		if ipTarget >= item.IpStart && ipTarget <= item.IpEnd {
			item.IP = ip
			res = item
			break
		}
	}

	ipData.Cache.Store(ip, res)
	return
}

func calcIP(ip string) int64 {
	x := strings.Split(ip, ".")
	if len(x) < 4 {
		return 0
	}
	b0, _ := strconv.ParseInt(x[0], 10, 0)
	b1, _ := strconv.ParseInt(x[1], 10, 0)
	b2, _ := strconv.ParseInt(x[2], 10, 0)
	b3, _ := strconv.ParseInt(x[3], 10, 0)
	return b0*16777216 + b1*65536 + b2*256 + b3*1
}

func CheckIP(ip string) bool {
	addr := strings.Trim(ip, " ")
	pattern := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(pattern, addr); match {
		return true
	}
	return false
}
