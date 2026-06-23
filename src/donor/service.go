package donor

import (
	"context"
	"fmt"

	"github.com/romanova09/bold-trunojoyo-scholarship-api/config"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/helpers"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/infra"
)

type Service struct {
	cfg   *config.Config
	repo  IRepository
	email infra.IEmail
}

func NewService(cfg *config.Config, repo *Repository, email infra.IEmail) *Service {
	return &Service{
		cfg:   cfg,
		repo:  repo,
		email: email,
	}
}

func (s *Service) WelcomeEmail(ctx context.Context, req RequestWelcomeEmail) (BaseResponse, error) {
	templatePath := "assets/template/email/donatur_registration.html"
	for _, to := range req.To {
		templateData := struct {
			DONOR_NAME     string
			BANK_NAME      string
			ACCOUNT_NAME   string
			ACCOUNT_NUMBER string
		}{
			DONOR_NAME:     to.Nama,
			BANK_NAME:      s.cfg.Bank.BankName,
			ACCOUNT_NAME:   s.cfg.Bank.AccountName,
			ACCOUNT_NUMBER: s.cfg.Bank.AccountNumber,
		}
		parseBody, err := helpers.ParseHTMLTemplate(templatePath, templateData)
		if err != nil {
			return BaseResponse{}, err
		}
		arrEmail := []string{to.Email}
		err = s.email.SendEmail(arrEmail, EMAIL_SUBJECT, string(parseBody))
		if err != nil {
			return BaseResponse{}, err
		}
	}
	return BaseResponse{Ok: true}, nil
}

func (s *Service) ReportDonation(ctx context.Context, req RequestDonationReport) (BaseResponse, error) {
	templatePath := "assets/template/email/donation_report.html"
	templateData := struct {
		TOTAL_DONATION string
		STUDENT_NAME   string
		PRODI          string
		ANGKATAN       string
		TOTAL_UKT      string
		LINK           string
	}{
		TOTAL_DONATION: req.TotalDonasi,
		STUDENT_NAME:   req.Penerima,
		PRODI:          req.Prodi,
		ANGKATAN:       req.Angkatan,
		TOTAL_UKT:      req.TotalUKT,
		LINK:           s.cfg.Bank.DonationLink,
	}
	parseBody, err := helpers.ParseHTMLTemplate(templatePath, templateData)
	if err != nil {
		return BaseResponse{}, err
	}
	donors, err := s.repo.GetAllDonor(ctx)
	if err != nil {
		fmt.Println("[Error] get donor: ", err)
	}
	for _, donor := range donors {
		arrEmail := []string{donor.Email}
		err = s.email.SendEmail(arrEmail, DONATION_REPORT, string(parseBody))
		if err != nil {
			return BaseResponse{}, err
		}
	}
	return BaseResponse{Ok: true}, nil
}

func (s *Service) ApplicantAnnouncement(ctx context.Context, req RequestApplicantAnnouncement) (BaseResponse, error) {
	for _, to := range req.To {
		templatePath := "assets/template/email/announcement_failed_applicant.html"
		if to.Success {
			templatePath = "assets/template/email/announcement_success_applicant.html"
		}
		templateData := struct {
			NAME string
		}{
			NAME: to.Nama,
		}
		parseBody, err := helpers.ParseHTMLTemplate(templatePath, templateData)
		if err != nil {
			return BaseResponse{}, err
		}
		arrEmail := []string{to.Email}
		err = s.email.SendEmail(arrEmail, EMAIL_SUBJECT_APPLICANT, string(parseBody))
		if err != nil {
			return BaseResponse{}, err
		}
	}
	return BaseResponse{Ok: true}, nil
}
