package filter_teams

import (
	"context"
	"strings"

	"github.com/ministryofjustice/cloud-platform-github-teams-filter/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetClusterGithubTeams(clientset *kubernetes.Clientset) ([]string, error) {
	var allTeams []string
	roleBindings, err := clientset.RbacV1().RoleBindings("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, rb := range roleBindings.Items {
		for _, subject := range rb.Subjects {
			if strings.Contains(subject.Name, "github:") {
				allTeams = append(allTeams, utils.StripGithubPrefix(subject.Name))
			}
		}
	}

	dedupTeams := utils.RemoveDuplicates(allTeams)
	return dedupTeams, nil
}
