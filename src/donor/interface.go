package donor

import (
	"context"

	"github.com/romanova09/bold-trunojoyo-scholarship-api/model"
)

type IService interface {
	WelcomeEmail(ctx context.Context, req RequestWelcomeEmail) (BaseResponse, error)
	ApplicantAnnouncement(ctx context.Context, req RequestApplicantAnnouncement) (BaseResponse, error)
	ReportDonation(ctx context.Context, req RequestDonationReport) (BaseResponse, error)
}

type IRepository interface {
	GetDonorReminder(ctx context.Context, dateReminder int) ([]model.Donor, error)
	GetAllDonor(ctx context.Context) ([]model.Donor, error)
}
