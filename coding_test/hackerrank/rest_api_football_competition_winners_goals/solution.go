package rest_api_football_competition_winners_goals

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"strings"
)

const baseURL = "https://jsonmock.hackerrank.com/api"

type Pagination[T any] struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_pages"`
	PageData  []T `json:"data"`
}

type Competition struct {
	Name     string `json:"name"`
	Country  string `json:"country"`
	Year     int32  `json:"year"`
	Winner   string `json:"winner"`
	RunnerUp string `json:"runnerup"`
}

type Match struct {
	Team1      string `json:"team1"`
	Team2      string `json:"team2"`
	Team1Goals string `json:"team1goals"`
	Team2Goals string `json:"team2goals"`
}

type Response[T any] struct {
	Body *T
	Err  error
}

func getWinnerTotalGoals(league string, year int) int {
	winnerTotalGoals := 0
	league = strings.ReplaceAll(league, " ", "%20")

	ch := make(chan Response[Pagination[Competition]], 1)
	go getData[Pagination[Competition]](getFootballCompetitionsURL(league, int32(year)), ch)

	res := <-ch
	if res.Err != nil {
		panic(res.Err)
	}

	if len(res.Body.PageData) == 0 {
		return 0
	}

	winner := res.Body.PageData[0].Winner

	matches1 := make([]Match, 0)
	matches2 := make([]Match, 0)

	ch1 := make(chan Response[Pagination[Match]], 1)
	ch2 := make(chan Response[Pagination[Match]], 1)
	go getData[Pagination[Match]](getFootballMatchesURL(league, int32(year), 1, winner, 1), ch1)
	go getData[Pagination[Match]](getFootballMatchesURL(league, int32(year), 2, winner, 1), ch2)

	res1 := <-ch1
	if res1.Err != nil {
		panic(res1.Err)
	}
	res2 := <-ch2
	if res2.Err != nil {
		panic(res2.Err)
	}

	matches1 = append(matches1, res1.Body.PageData...)
	matches2 = append(matches2, res2.Body.PageData...)

	totalPage1 := int32(res1.Body.TotalPage)
	ch1 = make(chan Response[Pagination[Match]], totalPage1-1)
	for i := int32(2); i <= totalPage1; i++ {
		go getData[Pagination[Match]](getFootballMatchesURL(league, int32(year), 1, winner, i), ch1)
	}

	totalPage2 := int32(res2.Body.TotalPage)
	ch2 = make(chan Response[Pagination[Match]], totalPage2-1)
	for i := int32(2); i <= totalPage2; i++ {
		go getData[Pagination[Match]](getFootballMatchesURL(league, int32(year), 2, winner, i), ch2)
	}

	for i := int32(2); i <= totalPage1; i++ {
		matches1 = append(matches1, (<-ch1).Body.PageData...)
	}

	for i := int32(2); i <= totalPage2; i++ {
		matches2 = append(matches2, (<-ch2).Body.PageData...)
	}

	for _, match := range matches1 {
		goals, err := strconv.Atoi(match.Team1Goals)
		if err != nil {
			panic(err)
		}

		winnerTotalGoals += goals
	}

	for _, match := range matches2 {
		goals, err := strconv.Atoi(match.Team2Goals)
		if err != nil {
			panic(err)
		}

		winnerTotalGoals += goals
	}

	return winnerTotalGoals
}

func getFootballCompetitionsURL(name string, year int32) string {
	return baseURL + fmt.Sprintf("/football_competitions?name=%s&year=%d", name, year)
}

func getFootballMatchesURL(competition string, year int32, teamNum int32, teamName string, page int32) string {
	return baseURL + fmt.Sprintf("/football_matches?competition=%s&year=%d&team%d=%s&page=%d", competition, year, teamNum, teamName, page)
}

func getData[T any](url string, ch chan (Response[T])) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- Response[T]{Body: nil, Err: err}
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- Response[T]{Body: nil, Err: err}
		return
	}

	var body *T
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		ch <- Response[T]{Body: nil, Err: err}
		return
	}

	ch <- Response[T]{Body: body, Err: nil}
}
