package donor

import (
	"github.com/gin-gonic/gin"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/config"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/helpers"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/helpers/response"
)

func RegisterAPI(r *gin.Engine, cfg *config.Config, service IService) {
	registrationGroup := r.Group("/registrations")
	handler := handler{cfg: cfg, service: service}
	{
		registrationGroup.POST("/welcome-email", handler.welcomeEmail)
		registrationGroup.POST("/donation-report", handler.donationReport)
		registrationGroup.POST("/applicant-failed", handler.failedApplicant)
	}

}

type handler struct {
	cfg     *config.Config
	service IService
}

func (h handler) welcomeEmail(c *gin.Context) {
	ctx := c.Request.Context()
	req := RequestWelcomeEmail{}
	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := helpers.Validate(req); err != nil {
		response.BadRequest(c, err)
		return
	}

	resp, err := h.service.WelcomeEmail(ctx, req)
	if err != nil {
		response.ErrInternalServerError(c, err)
		return
	}

	response.Success(c, resp)
}

func (h handler) failedApplicant(c *gin.Context) {
	ctx := c.Request.Context()
	req := RequestApplicantAnnouncement{}
	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := helpers.Validate(req); err != nil {
		response.BadRequest(c, err)
		return
	}

	resp, err := h.service.ApplicantAnnouncement(ctx, req)
	if err != nil {
		response.ErrInternalServerError(c, err)
		return
	}

	response.Success(c, resp)
}

func (h handler) donationReport(c *gin.Context) {
	ctx := c.Request.Context()
	req := RequestDonationReport{}
	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := helpers.Validate(req); err != nil {
		response.BadRequest(c, err)
		return
	}

	resp, err := h.service.ReportDonation(ctx, req)
	if err != nil {
		response.ErrInternalServerError(c, err)
		return
	}

	response.Success(c, resp)
}
