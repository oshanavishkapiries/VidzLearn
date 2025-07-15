package user

import (
	"github.com/Cenzios/pf-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	// Simulated data
	profile := map[string]string{"name": "Ruvinda"}
	response.Success(c, profile, "User profile fetched")
}
