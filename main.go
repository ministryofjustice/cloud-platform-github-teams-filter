package main

import (
	"log"

	"github.com/ministryofjustice/cloud-platform-github-teams-filter/init_app"
)

func main() {
	ginMode := init_app.InitEnvVars()

	r := init_app.InitGin(ginMode)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

}
