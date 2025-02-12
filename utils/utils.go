package utils

import (
	"strings"
)

func StripGithubPrefix(name string) string {
	return strings.Replace(name, "github:", "", -1)
}

func DeduplicateTeams(allTeams []string) []string {
	seen := map[string]bool{}
	dedupTeams := []string{}

	for _, team := range allTeams {
		if !seen[team] {
			seen[team] = true
			dedupTeams = append(dedupTeams, team)
		}
	}

	return dedupTeams
}

func FilterTeams(input string, cluster []string) string {

	inputSlice := strings.Split(input, ":")
	clusterMap := map[string]bool{}

	for _, team := range cluster {
		clusterMap[team] = true
	}

	var result []string
	for _, team := range inputSlice {
		if clusterMap[team] {
			result = append(result, team)
		}
	}

	return ":" + strings.Join(result, ":") + ":"
}
