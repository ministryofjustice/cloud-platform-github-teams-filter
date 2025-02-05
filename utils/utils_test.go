package utils

import "testing"

func TestStripGithubPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "strip github prefix",
			input:    "github:team1",
			expected: "team1",
		},
		{
			name:     "strip github prefix",
			input:    "github:team2",
			expected: "team2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := StripGithubPrefix(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, actual)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "remove duplicates",
			input:    []string{"team1", "team1", "team2"},
			expected: []string{"team1", "team2"},
		},
		{
			name:     "remove duplicates",
			input:    []string{"team1", "team2", "team2", "team2", "team3"},
			expected: []string{"team1", "team2", "team3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := RemoveDuplicates(tt.input)
			if len(actual) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestFilterTeams(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		cluster  []string
		expected string
	}{
		{
			name:     "filter teams",
			input:    ":team1:team2:team3:",
			cluster:  []string{"team1", "team2"},
			expected: ":team1:team2:",
		},
		{
			name:     "filter teams",
			input:    ":team1:team2:team3:",
			cluster:  []string{"team1", "team2", "team3"},
			expected: ":team1:team2:team3:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FilterTeams(tt.input, tt.cluster)
			if actual != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, actual)
			}
		})
	}
}
