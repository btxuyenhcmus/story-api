package controller

import (
	"net/http"

	"github.com/readtruyen/go-novelstory-api/domain"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

func (pc *ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	profile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (pc *ProfileController) Update(c *gin.Context) {
	var profile domain.UpdateProfile

	err := c.ShouldBind(&profile)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	err = pc.ProfileUsecase.UpdateProfile(c, userID, &profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Profile updated successfully",
	})
}
