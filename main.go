package main

import (
	"flag"
	"go.uber.org/zap"
	"workflow/initialize"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "端口号")

	flag.Parse()

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	zap.S().Info("ip:%s", *IP)
	zap.S().Info("port:%s", *Port)
}
