package rest_api_total_goals_by_a_team

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type Pagination struct {
	Page       float64   `json:"page"`
	PerPage    float64   `json:"per_page"`
	Total      float64   `json:"total"`
	TotalPages float64   `json:"total_pages"`
	DataArr    []Matches `json:"data"`
}

func GetTotalGoals(team string, year int32) int32 {
	totalGoals := int32(0)

	getGoals := func(teamNumber int) int32 {
		page := 1
		totalPages := 1
		goals := int32(0)
		for ; page <= totalPages; page++ {
			resp, err := http.Get(fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?year=%d&team%d=%s&page=%d", year, teamNumber, team, page))
			if err != nil {
				return 0
			}

			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return 0
			}

			pn := new(Pagination)
			err = json.Unmarshal(body, &pn)
			if err != nil {
				return 0
			}

			totalPages = int(pn.TotalPages)

			for _, data := range pn.DataArr {
				if teamNumber == 1 {
					if goal, err := strconv.Atoi(data.Team1Goals); err == nil {
						goals += int32(goal)
					}
				} else {
					if goal, err := strconv.Atoi(data.Team2Goals); err == nil {
						goals += int32(goal)
					}
				}
			}
		}

		return goals
	}

	totalGoals += getGoals(1)
	totalGoals += getGoals(2)

	return totalGoals
}
