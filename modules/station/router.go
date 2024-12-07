package station

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izsal/mrt-schedules/common/response"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/station")
	station.GET("", func(c *gin.Context) {
		GetAllStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStation()
	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			},
		)
		return

	}

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Successfully get all station",
		Data:    datas,
	})
}