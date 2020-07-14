package adding

import (
	"fmt"

	"github.com/drd-engineering/DRDWorkflow/db"
	"github.com/drd-engineering/DRDWorkflow/domains/adding/workflow"
	"github.com/drd-engineering/DRDWorkflow/routes"

	"github.com/gin-gonic/gin"
)

func InitiateRoutes() {
	r := routes.GetInstance()
	apiRoutes := r.Group("/api/v1")
	{
		routeForAdding := apiRoutes.Group("/adding")

		routeForAdding.Use(ApplicationAuthorizationRequired())
		{
			routeForAdding.POST("/save-workflow", workflow.SaveWorkflow)
		}
	}
}

func ApplicationAuthorizationRequired(auths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header["Drd-Api-Key"]

		if len(apiKey) < 1 {
			fmt.Println(apiKey)
			c.AbortWithStatus(401)
			return
		}

		if apiKey[0] == "" {
			fmt.Println(apiKey[0])
			c.AbortWithStatus(401)
			return
		}

		dbInstance := db.GetDb()
		var token db.AppToken
		if err := dbInstance.Where("token = ? AND is_active", apiKey[0]).First(&token).Error; err != nil {
			c.AbortWithStatus(401)
			fmt.Println(err)
			return
		}
		fmt.Println(token)
		c.Next()
	}
}
