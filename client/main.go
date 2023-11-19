package main

import (
	"log"
	"time"
)

func main() {
	var args arguments
	args.Read()
	if args.Version {
		printVersion()
		return
	}

	var conf config
	err := conf.Read(args.Path)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		ipList, err := getNetworkAddrs(conf.IPPrefix)
		if err != nil || len(ipList) == 0 {
			log.Println("网络位址取得失败")
		} else {
			address := ipList[0]
			log.Println("开始上报网络位址", address)
			err := reportNetworkAddrs(conf.Timeout, conf.DeviceTag, conf.ReportURI, address)
			if err != nil {
				log.Println("网络位址上报失败")
			} else {
				log.Println("网络位址上报成功", conf.DeviceTag)
			}
		}

		time.Sleep(time.Second * time.Duration(conf.Interval))
	}
}
