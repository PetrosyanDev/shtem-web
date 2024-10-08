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
	storageclient "shtem-web/sources/internal/clients/storage"
	telegramclient "shtem-web/sources/internal/clients/telegram"
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
	shtemsDB := postgresrepository.NewShtemsDB(appCtx, postgresDB)
	categoriesDB := postgresrepository.NewCategoriesDB(appCtx, postgresDB)
	emailsDB := postgresrepository.NewEmailsDB(appCtx, postgresDB)

	log.Println("init repositories")

	templatesRepo, err := repositories.NewHTMLTemplates(embeds.Templates)
	if err != nil {
		log.Fatalf("failed to load templates (%v)", err)
	}

	questionsRepository := repositories.NewQuestionsRepository(questionsDB)
	shtemsRepository := repositories.NewShtemsRepository(shtemsDB)
	categoriesRepository := repositories.NewCategoriesRepository(categoriesDB)
	emailsRepository := repositories.NewEmailsRepository(emailsDB)

	log.Println("init clients")
	storageClient, err := storageclient.NewStorageClient(appCtx, cfg)
	if err != nil {
		log.Fatalf("failed to connect with Storage (%v)", err)
	}
	telegramClient, err := telegramclient.NewTelegamClient(cfg)
	if err != nil {
		log.Fatalf("failed to create Telegram client (%v)", err)
	}

	log.Println("init services")
	filesService := services.NewFilesService(storageClient)
	webService, err := services.NewWEBService(embeds.Assets, embeds.Uploads, templatesRepo)
	if err != nil {
		log.Fatalf("failed to create WEB service (%v)", err)
	}

	questionsService := services.NewQuestionsService(questionsRepository)
	shtemsService := services.NewShtemsService(shtemsRepository)
	categoriesService := services.NewCategoriesService(categoriesRepository, shtemsRepository)
	emailsService := services.NewEmailsService(emailsRepository)

	log.Println("init handlers")
	webHandler := handlers.NewWEBHandler(webService, telegramClient, questionsService, shtemsService, categoriesService, emailsService, filesService)

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
