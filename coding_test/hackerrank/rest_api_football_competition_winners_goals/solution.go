package rest_api_football_competition_winners_goals

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Pagination[T any] struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_pages"`
	Data      []T `json:"data"`
}

type Competition struct {
	Name     string  `json:"name"`
	Country  string  `json:"country"`
	Year     float64 `json:"year"`
	Winner   string  `json:"winner"`
	Runnerup string  `json:"runnerup"`
}

type Matches struct {
	Competition string `json:"competition"`
	Year        int32  `json:"year"`
	Round       string `json:"round"`
	Team1       string `json:"team1"`
	Team2       string `json:"team2"`
	Team1Goals  string `json:"team1goals"`
	Team2Goals  string `json:"team2goals"`
}

func getWinnerTotalGoals(league string, year int) int {
	league = strings.ReplaceAll(league, " ", "%20")
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_competitions?name=%s&year=%d", league, year)
	competition, err := getOneData[Competition](url)
	if err != nil {
		panic(err)
	}

	winnerTotalGoals := 0

	url = fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?competition=%s&year=%d&team%d=%s", league, year, 1, competition.Winner)
	matches1, err := getAllData[Matches](url, 1)
	if err != nil {
		panic(err)
	}

	for _, match := range *matches1 {
		if goals, err := strconv.Atoi(match.Team1Goals); err == nil {
			winnerTotalGoals += goals
		}
	}

	url = fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?competition=%s&year=%d&team%d=%s", league, year, 2, competition.Winner)
	matches2, err := getAllData[Matches](url, 1)
	if err != nil {
		panic(err)
	}

	for _, match := range *matches2 {
		if goals, err := strconv.Atoi(match.Team2Goals); err == nil {
			winnerTotalGoals += goals
		}
	}

	return winnerTotalGoals
}

func getOneData[T any](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pn := new(Pagination[T])
	err = json.Unmarshal(body, &pn)
	if err != nil {
		return nil, err
	}

	if len(pn.Data) < 1 {
		return nil, nil
	}

	return &pn.Data[0], nil
}

func getAllData[T any](url string, curPage int) (*([]T), error) {
	resp, err := http.Get(url + fmt.Sprintf("&page=%d", curPage))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pn := new(Pagination[T])
	err = json.Unmarshal(body, &pn)
	if err != nil {
		return nil, err
	}

	if pn.TotalPage == curPage {
		return &pn.Data, nil
	}

	nextData, err := getAllData[T](url, curPage+1)
	if err != nil {
		return nil, err
	}

	pn.Data = append(pn.Data, (*nextData)...)

	return &pn.Data, nil
}
