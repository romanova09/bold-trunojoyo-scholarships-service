package scheduler

import (
	"fmt"

	cron "github.com/robfig/cron/v3"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/config"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/db"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/infra"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/src/donor"
)

type Job struct {
	cfg       *config.Config
	cron      *cron.Cron
	donorRepo *donor.Repository
	email     infra.IEmail
}

func NewJob() Job {
	cfg := config.LoadDefault()
	db := db.NewGoPG(cfg)
	donorRepo := donor.NewRepository(db)
	var emailInfra infra.IEmail = infra.NewEmail(cfg)
	c := cron.New()
	return Job{
		cfg:       cfg,
		cron:      c,
		donorRepo: donorRepo,
		email:     emailInfra,
	}
}

func (job Job) StartCronJob(cfg *config.Config) {
	_, err := job.cron.AddFunc("0 17 3 * *", job.ReminderDonation)
	if err != nil {
		fmt.Printf("Gagal menambahkan cron [job: %v, error: %v]", "Reminder Donation", err)
	}

	job.cron.Start()
}
