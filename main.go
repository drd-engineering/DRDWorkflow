package main

import (
	"github.com/drd-engineering/DRDWorkflow/db"
	"github.com/drd-engineering/DRDWorkflow/domains/adding"
	"github.com/drd-engineering/DRDWorkflow/routes"
)

func main() {
	db.GetDb()
	// Add Specific router group to main router
	adding.InitiateRoutes()

	// Start Server
	r := routes.GetInstance()
	r.Run(":8080")
}
