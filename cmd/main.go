package main

import (
	"github.com/carrot-systems/cs-core/src/adapters/rest"
	"github.com/carrot-systems/cs-core/src/config"
	"github.com/carrot-systems/cs-core/src/core/usecases"
	configurationClient "github.com/carrot-systems/csl-configuration-client"
	discoveryClient "github.com/carrot-systems/csl-discovery-client"
	env "github.com/carrot-systems/csl-env"
	"log"
)

func main() {
	env.LoadEnv()

	discovery := discoveryClient.NewClient()
	err := discovery.Register("core")
	if err != nil {
		log.Println(err.Error())
		return
	}
	configuration := configurationClient.NewClient(discovery)
	err = configuration.LoadConfiguration()
	if err != nil {
		log.Println(err.Error())
		return
	}

	ginConfiguration := config.LoadGinConfiguration()

	usecasesHandler := usecases.NewInteractor()

	restServer := rest.NewServer(ginConfiguration)
	routesHandler := rest.NewRouter(usecasesHandler)

	rest.SetRoutes(restServer.Router, routesHandler)
	restServer.Start()
}
