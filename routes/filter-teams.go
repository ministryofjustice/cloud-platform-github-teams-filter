package routes

import (
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
	f "github.com/ministryofjustice/cloud-platform-github-teams-filter/pkg/filter_teams"
	"github.com/ministryofjustice/cloud-platform-github-teams-filter/utils"
	"k8s.io/client-go/kubernetes"
)

type FilterTeamsRequest struct {
	Teams string `json:"teams"`
}

type FilterTeamsResponse struct {
	FilteredTeams string `json:"filtered_teams"`
}

func InitFilterTeams(r *gin.Engine, clientset *kubernetes.Clientset) {
	r.POST("/filter-teams", authRequest(), func(c *gin.Context) {

		var req FilterTeamsRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		clusterTeams, err := f.GetClusterGithubTeams(clientset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		filteredTeams := utils.FilterTeams(req.Teams, clusterTeams)

		res := FilterTeamsResponse{
			FilteredTeams: filteredTeams,
		}

		c.JSON(http.StatusOK, res)

	})
}

func authRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		actualAPIKey := os.Getenv("API_KEY")
		APIKey := c.Request.Header.Get("X-API-Key")

		if APIKey != actualAPIKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	}
}
