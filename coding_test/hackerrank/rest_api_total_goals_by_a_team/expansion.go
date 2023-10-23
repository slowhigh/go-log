package rest_api_total_goals_by_a_team

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
)

type Matches struct {
	Competition string `json:"competition"`
	Year        int32  `json:"year"`
	Round       string `json:"round"`
	Team1       string `json:"team1"`
	Team2       string `json:"team2"`
	Team1Goals  string `json:"team1goals"`
	Team2Goals  string `json:"team2goals"`
}

type FootballMatch struct {
	Page       int64     `json:"page"`
	PerPage    int64     `json:"per_page"`
	Total      int64     `json:"total"`
	TotalPages int32     `json:"total_pages"`
	DataArr    []Matches `json:"data"`
}

type Team struct {
	Name       string
	TotalGoals int32
	Records    []Record
}

type Record struct {
	Year  int32
	Goals int32
}

func GetAllTeamTotalGoals() (*[]Team, error) {
	teams := make([]Team, 0)
	
	matches, err := getAllFootballMatches()
	if err != nil {
		return nil, err
	}

	teamMap := getTeamMap(matches)

	for name, recordMap := range teamMap {
		team := Team{Name: name, Records: make([]Record, 0)}
		totalGoals := int32(0)

		for year, goals := range recordMap {
			team.Records = append(team.Records, Record{Year: year, Goals: goals})
			totalGoals += goals
		}

		team.TotalGoals = totalGoals
		teams = append(teams, team)
	}

	sort.Slice(teams, func(i, j int) bool { return teams[i].TotalGoals > teams[j].TotalGoals })

	return &teams, nil
}

func getAllFootballMatches() ([]Matches, error) {
	matches := make([]Matches, 0)

	getFootballMatches := func(page int32) (*FootballMatch, error) {
		resp, err := http.Get(fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?page=%d", page))
		if err != nil {
			return nil, err
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		footballMatch := new(FootballMatch)
		err = json.Unmarshal(body, &footballMatch)
		if err != nil {
			return nil, err
		}

		return footballMatch, nil
	}

	firstPn, err := getFootballMatches(1)
	if err != nil {
		return nil, err
	}

	matches = append(matches, firstPn.DataArr...)
	totalPages := firstPn.TotalPages
	goroutineCount := int((totalPages - 1) / 2)
	done := make(chan []Matches)

	for i := int32(2); i <= totalPages; i += 2 {
		go func(page int32) {
			matches := make([]Matches, 0)

			pn, err := getFootballMatches(page)
			if err != nil {
				done <- nil
			}
			if pn.DataArr != nil {
				matches = append(matches, pn.DataArr...)
			}

			pn, err = getFootballMatches(page + 1)
			if err != nil {
				done <- nil
			}
			if pn.DataArr != nil {
				matches = append(matches, pn.DataArr...)
			}

			done <- matches
		}(i)
	}

	for i := 0; i < goroutineCount; i++ {
		matches = append(matches, <-done...)
	}

	return matches, nil
}

// return [name][year]goals
func getTeamMap(matches []Matches) map[string]map[int32]int32 {
	teamMap := make(map[string]map[int32]int32)

	stackGoals := func(name string, year, goals int32) {
		if _, ok := teamMap[name]; !ok {
			teamMap[name] = make(map[int32]int32)
		}
		if _, ok := teamMap[name][year]; !ok {
			teamMap[name][year] = goals
		} else {
			teamMap[name][year] += goals
		}
	}

	for _, match := range matches {
		team1Goals, err := strconv.Atoi(match.Team1Goals)
		if err == nil {
			stackGoals(match.Team1, match.Year, int32(team1Goals))
		}

		team2Goals, err := strconv.Atoi(match.Team2Goals)
		if err == nil {
			stackGoals(match.Team2, match.Year, int32(team2Goals))
		}
	}

	return teamMap
}
