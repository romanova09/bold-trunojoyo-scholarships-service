package donor

type RequestWelcomeEmail struct {
	To []struct {
		Email string `json:"email" required:"true"`
		Nama  string `json:"nama" required:"true"`
	} `json:"to" required:"true"`
}
type RequestApplicantAnnouncement struct {
	To []struct {
		Email   string `json:"email" required:"true"`
		Nama    string `json:"nama" required:"true"`
		Success bool   `json:"success" required:"true"`
	} `json:"to" required:"true"`
}
type BaseResponse struct {
	Ok   bool        `json:"success"`
	Data interface{} `json:"data"`
}

type RequestDonationReport struct {
	TotalDonasi string `json:"total_donasi" required:"true"`
	Penerima    string `json:"penerima" required:"true"`
	Prodi       string `json:"prodi" required:"true"`
	Angkatan    string `json:"angkatan" required:"true"`
	TotalUKT    string `json:"total_ukt" required:"true"`
}
