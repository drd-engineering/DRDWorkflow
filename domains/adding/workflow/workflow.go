package workflow

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SaveWorkflow(c *gin.Context) {
	fmt.Print("Save Workflow")
	c.JSON(200, "")
}
