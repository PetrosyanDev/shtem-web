// Erik Petrosyan ©
package main

import (
	"context"
	"log"
	"net/http"
	embd "shtem-web"
	"shtem-web/sources/internal/adapters/web"
	"shtem-web/sources/internal/adapters/web/handlers"
	"shtem-web/sources/internal/core/services"
	"shtem-web/sources/internal/repositories"
	"shtem-web/sources/internal/system"
	"sync"
)

func main() {
	appCtx, cancelAppCtx := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	go system.HandleSysCalls(cancelAppCtx)

	log.Println("loading configs")
	// TODO
	// cfg, err := configs.NewConfigs(os.Args)
	// if err != nil {
	// 	log.Fatalf("unable to load configs (%v)", err)
	// }

	embeds := embd.NewEMBD()
	opts := []web.WEBServerOpt{}

	log.Println("init repositories")
	templatesRepo, err := repositories.NewHTMLTemplates(embeds.Templates)
	if err != nil {
		log.Fatalf("failed to load templates (%v)", err)
	}

	log.Println("init services")
	webService, err := services.NewWEBService(embeds.Assets, templatesRepo)
	if err != nil {
		log.Fatalf("failed to create WEB service (%v)", err)
	}

	log.Println("init handlers")
	webHandler := handlers.NewWEBHandler(webService)

	webRouter := web.NewWEBRouter(webHandler)
	webApp, err := web.NewWEBServer(webRouter, opts...)
	if err != nil {
		log.Fatalf("failed to create WEB server (%v)", err)
	}

	toStop := []system.Service{webApp}
	wg.Add(len(toStop))
	go system.HandleGracefullExit(appCtx, wg, toStop...)

	if err := webApp.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to run WEB server (%v)", err)
	}
	wg.Wait()
}
