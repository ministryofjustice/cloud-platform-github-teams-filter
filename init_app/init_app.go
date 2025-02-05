package init_app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ministryofjustice/cloud-platform-github-teams-filter/routes"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func InitGin(ginMode string) *gin.Engine {
	gin.SetMode(ginMode)

	r := gin.New()

	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error building in-cluster config: %s", err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %s", err.Error())
	}

	routes.InitLogger(r)

	routes.InitRouter(r, clientset)

	return r
}
