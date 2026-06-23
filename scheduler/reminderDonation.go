package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/romanova09/bold-trunojoyo-scholarship-api/helpers"
)

func (s Job) ReminderDonation() {
	ctx := context.Background()
	_, _, date := time.Now().Date()
	donors, err := s.donorRepo.GetDonorReminder(ctx, date)
	if err != nil {
		fmt.Println("[Error] get donor: ", err)
	}
	for _, donor := range donors {
		err = s.SendReminderEmail(ctx, donor.Email, donor.NickName)
		fmt.Println("Error SendReminderEmail: ", err)
	}
}

func (s *Job) SendReminderEmail(ctx context.Context, email string, name string) error {
	templatePath := "assets/template/email/donatur_reminder.html"
	templateData := struct {
		DONOR_NAME     string
		BANK_NAME      string
		ACCOUNT_NAME   string
		ACCOUNT_NUMBER string
	}{
		DONOR_NAME:     name,
		BANK_NAME:      s.cfg.Bank.BankName,
		ACCOUNT_NAME:   s.cfg.Bank.AccountName,
		ACCOUNT_NUMBER: s.cfg.Bank.AccountNumber,
	}
	parseBody, err := helpers.ParseHTMLTemplate(templatePath, templateData)
	if err != nil {
		return err
	}
	arrEmail := []string{email}
	err = s.email.SendEmail(arrEmail, fmt.Sprintf("Pengingat: Waktu Donasi Telah Tiba, %v!", name), string(parseBody))
	if err != nil {
		fmt.Printf("Error send email [email: %v ,err: %v", email, err)
	}
	return err
}
