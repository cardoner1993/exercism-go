package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type TeamRecord struct {
	Name  string
	Stats *TeamStats
}

type TeamStats struct {
	MP int // Matches Played
	W  int // Wins
	D  int // Draws
	L  int // Losses
	P  int // Points
}

func Tally(reader io.Reader, writer io.Writer) error {
	// Create a string builder to accumulate results
	var result strings.Builder
	teams := make(map[string]*TeamStats)

	scanner := bufio.NewScanner(reader)

	// Read line by line
	for scanner.Scan() {
		line := scanner.Text() // This gives you each line as a string

		// Uncomment the next line for debugging purposes to see each line being processed
		// fmt.Printf("Processing line: '%s'\n", line)

		// Skip empty lines (including lines with only whitespace)
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Skip comment lines
		if strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ";")

		if len(parts) != 3 {
			return fmt.Errorf("invalid line format: expected 3 fields separated by ';', got %d in line: '%s'", len(parts), line)
		}

		team1 := parts[0]
		team2 := parts[1]
		outcome := parts[2]

		if teams[team1] == nil {
			teams[team1] = &TeamStats{}
		}

		if teams[team2] == nil {
			teams[team2] = &TeamStats{}
		}

		// Update matches played
		teams[team1].MP++
		teams[team2].MP++

		switch outcome {
		case "win":
			teams[team1].W++
			teams[team2].L++
			teams[team1].P += 3
		case "loss":
			teams[team1].L++
			teams[team2].W++
			teams[team2].P += 3
		case "draw":
			teams[team1].D++
			teams[team2].D++
			teams[team1].P++
			teams[team2].P++
		default:
			return fmt.Errorf("invalid outcome '%s': must be 'win', 'loss', or 'draw'", outcome)
		}
	}

	var teamRecords []TeamRecord
	for name, stats := range teams {
		teamRecords = append(teamRecords, TeamRecord{Name: name, Stats: stats})
	}

	sort.Slice(teamRecords, func(i, j int) bool {
		if teamRecords[i].Stats.P == teamRecords[j].Stats.P {
			return teamRecords[i].Name < teamRecords[j].Name
		}
		return teamRecords[i].Stats.P > teamRecords[j].Stats.P
	})

	if err := scanner.Err(); err != nil {
		return err
	}

	// Build output
	result.WriteString("Team                           | MP |  W |  D |  L |  P\n")

	for _, team := range teamRecords {
		result.WriteString(fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n",
			team.Name, team.Stats.MP, team.Stats.W, team.Stats.D, team.Stats.L, team.Stats.P))
	}

	// Write output.
	_, err := writer.Write([]byte(result.String()))
	return err
}
