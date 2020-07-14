package main

import (
	"drd-engineering/DRDWorkflow/routes"
)

func main() {
	// Start Server
	r := routes.GetInstance()
	r.Run(":8080")
}
