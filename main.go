package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	api "github.com/romanova09/bold-trunojoyo-scholarship-api/app"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/config"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/scheduler"
)

func main() {
	r := gin.Default()
	cfg := config.LoadDefault()
	scheduler := scheduler.NewJob()
	scheduler.StartCronJob(cfg)

	api := api.New()
	api.RegisterAPI(r)

	if err := r.Run(fmt.Sprintf("%s:%s", cfg.Server.BASEURL, cfg.Server.PORT)); err != nil {
		log.Fatalf("listen:%+s\n", err)
	}
	fmt.Printf("server is running at port: %v [env: %v, db: %v]", cfg.Server.PORT, cfg.Server.ENV, cfg.Database.Host)

	select {}
}
