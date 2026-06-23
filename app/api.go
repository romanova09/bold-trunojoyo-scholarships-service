package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/config"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/db"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/infra"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/src/donor"
)

type API struct {
	cfg *config.Config
	db  *pg.DB
}

func New() API {
	cfg := config.LoadDefault()
	db := db.NewGoPG(cfg)

	return API{
		cfg: cfg,
		db:  db,
	}
}
func (api API) RegisterAPI(r *gin.Engine) {
	var emailInfra infra.IEmail = infra.NewEmail(api.cfg)

	donorRepo := donor.NewRepository(api.db)
	registrationService := donor.NewService(api.cfg, donorRepo, emailInfra)
	donor.RegisterAPI(r, api.cfg, registrationService)
}
