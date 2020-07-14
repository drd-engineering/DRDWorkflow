package main

import (
	"github.com/drd-engineering/DRDWorkflow/db"
	"github.com/drd-engineering/DRDWorkflow/routes"
)

func main() {
	db.GetDb()
	// Start Server
	r := routes.GetInstance()
	r.Run(":8080")
}
