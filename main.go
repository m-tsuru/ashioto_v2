package main

import (
	"flag"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/m-tsuru/ashioto-api/lib/config"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the configuration file")
	// initDB := flag.Bool("init", false, "Initialize the database and create necessary tables")
	flag.Parse()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(cfg.Server.HostName + ":" + strconv.Itoa(cfg.Server.Port))

}
