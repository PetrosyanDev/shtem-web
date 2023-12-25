// Erik Petrosyan ©
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	embd "shtem-web"
	"shtem-web/sources/internal/adapters/web"
	"shtem-web/sources/internal/adapters/web/handlers"
	postgresclient "shtem-web/sources/internal/clients/postgres"
	"shtem-web/sources/internal/configs"
	"shtem-web/sources/internal/core/services"
	"shtem-web/sources/internal/repositories"
	postgresrepository "shtem-web/sources/internal/repositories/postgres"
	"shtem-web/sources/internal/system"
	"sync"
)

func main() {
	appCtx, cancelAppCtx := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	go system.HandleSysCalls(cancelAppCtx)

	log.Println("loading configs")
	cfg, err := configs.NewConfigs(os.Args)
	if err != nil {
		log.Fatalf("unable to load configs (%v)", err)
	}

	embeds := embd.NewEMBD()
	opts := []web.WEBServerOpt{}

	log.Println("init databases")
	postgresDB, err := postgresclient.NewPostgresDBConn(appCtx, cfg)
	if err != nil {
		log.Fatalf("failed to connect with PostgresDB (%v)", err)
	}
	questionsDB := postgresrepository.NewQuestionsDB(appCtx, postgresDB)

	log.Println("init repositories")

	templatesRepo, err := repositories.NewHTMLTemplates(embeds.Templates)
	if err != nil {
		log.Fatalf("failed to load templates (%v)", err)
	}

	questionsRepository := repositories.NewQuestionsRepository(questionsDB)

	log.Println("init services")
	webService, err := services.NewWEBService(embeds.Assets, templatesRepo)
	if err != nil {
		log.Fatalf("failed to create WEB service (%v)", err)
	}

	questionsService := services.NewQuestionsService(questionsRepository)

	log.Println("init handlers")
	webHandler := handlers.NewWEBHandler(webService, questionsService)

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
