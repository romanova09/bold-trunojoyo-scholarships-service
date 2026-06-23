package infra

type IEmail interface {
	SendEmail(to []string, subject, body string) error
}
