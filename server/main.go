package main

import (
	"log"
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

	log.Printf("服务器启动在 %s\n", conf.Listen)
	startServer(conf.Listen, &serverOptions{
		Gzip:       GZIP_LEVEL,
		CORS:       conf.CORS,
		WebPrefix:  WEB_PREFIX,
		DeviceLife: conf.Life,
		DeviceList: []map[string]any{},
	})
}
