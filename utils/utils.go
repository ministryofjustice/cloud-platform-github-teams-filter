package utils

import (
	"fmt"
	"strings"
)

// Strip github prefix
func StripGithubPrefix(name string) string {
	return strings.Replace(name, "github:", "", -1)
}

// Deduplicate teams
func RemoveDuplicates(allTeams []string) []string {
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

// Filter out teams that are not present in the cluster teams slice
func FilterTeams(input string, cluster []string) string {

	// Convert input teams string to slice
	inputSlice := strings.Split(input, ":")
	fmt.Println(inputSlice)

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

	// Convert filtered slice back to string
	return ":" + strings.Join(result, ":") + ":"
}
