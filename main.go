package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"

	"go.uber.org/zap"

	"workflow/initialize"
)

func main() {

	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "端口号")
	flag.Parse()
	app := fiber.New()

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	initialize.InitRoute(app)

	zap.S().Info("ip:%s", *IP)
	zap.S().Info("port:%s", *Port)

	zap.S().Fatal(app.Listen(":8090"))
}
