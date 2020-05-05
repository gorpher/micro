package main

import (
	eureka "github.com/gorpher/go-plugins/registry/eurekav1/v2"
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/micro/v2/api"
	"github.com/micro/micro/v2/plugin"
)

func init() {
	//
	api.Register(plugin.NewPlugin(
		plugin.WithFlag(&cli.StringFlag{
			Name:    "registry_auth_username",
			Usage:   "eureka registry basicAuth username",
			EnvVars: []string{"MICRO_REGISTRY_AUTH_USERNAME"},
		}, &cli.StringFlag{
			Name:    "registry_auth_password",
			Usage:   "eureka registry basicAuth  password",
			EnvVars: []string{"MICRO_REGISTRY_AUTH_PASSWORD"},
		}),
		plugin.WithInit(func(ctx *cli.Context) error {
			username := ctx.String("registry_auth_username")
			password := ctx.String("registry_auth_password")
			regName := ctx.String("registry")
			if regName != "" && regName == "eureka" && username != "" && password != "" {
				log.Debugf("eureka registry：BasicAuth info（username:%s,password:%s）", username, password)
				registry.DefaultRegistry.Init(eureka.BasicAuth(username, password))
			}
			return nil
		}),
	))
}
