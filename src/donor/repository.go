package donor

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/model"
)

type Repository struct {
	db *pg.DB
}

// NewRepository creates a new user repository
func NewRepository(db *pg.DB) *Repository {
	return &Repository{db}
}

func (r Repository) GetDonorReminder(ctx context.Context, dateReminder int) ([]model.Donor, error) {
	donors := []model.Donor{}
	err := r.db.Model(&donors).
		Where("date_reminder = ?", dateReminder).
		Select()
	if err != nil {
		return []model.Donor{}, err
	}
	return donors, err
}

func (r Repository) GetAllDonor(ctx context.Context) ([]model.Donor, error) {
	donors := []model.Donor{}
	err := r.db.Model(&donors).
		Select()
	if err != nil {
		return []model.Donor{}, err
	}
	return donors, err
}
