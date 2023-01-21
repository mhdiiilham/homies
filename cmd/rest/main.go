package main

import (
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mhdiiilham/homies/common"
	"github.com/mhdiiilham/homies/delivery/rest"
	"github.com/sirupsen/logrus"
)

func main() {
	cfgFile := flag.String("cfg", "config.yaml", "set configuration file")
	flag.Parse()

	logrus.Infof("server running with configuration file: %s", *cfgFile)
	cfg, err := common.ReadConfig(*cfgFile)
	if err != nil {
		logrus.Fatalf("failed to read configuration file: %v", err)
	}

	app := fiber.New()
	app.Use(rest.PreRequest())
	app.Use(logger.New(logger.Config{
		Format:     "time=${time} request-id=${requestid} pid=${pid} status=${status} method=${method} path=${path}\n",
		TimeFormat: time.RFC822Z,
		TimeZone:   "Asia/Jakarta",
	}))
	app.Use(recover.New())

	apiV1 := app.Group("api/v1")

	// register handlers
	rest.RootHandler(apiV1)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		logrus.Info("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(cfg.GetPort()); err != nil {
		logrus.Fatalf("failed to listen on port %d; err=%v", cfg.Port, err)
	}

	logrus.Info("running cleanup task...")
}
