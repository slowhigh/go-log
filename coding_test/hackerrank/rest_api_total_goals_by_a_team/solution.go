package restApiTotalGoalsByATeam

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const baseURL = "https://jsonmock.hackerrank.com/api/football_matches"

type Pagination[T any] struct {
	Page       int32 `json:"page"`
	PerPage    int32 `json:"per_page"`
	Total      int64 `json:"total"`
	TotalPages int32 `json:"total_pages"`
	PageData   []T   `json:"data"`
}

type Match struct {
	Team1      string `json:"team1"`
	Team2      string `json:"team2"`
	Team1Goals string `json:"team1goals"`
	Team2Goals string `json:"team2goals"`
}

type Response[T any] struct {
	Data *T
	Err  error
}

func GetTotalGoals(team string, year int32) int32 {
	totalGoals := int32(0)

	ch1 := make(chan (Response[Pagination[Match]]), 1)
	ch2 := make(chan (Response[Pagination[Match]]), 1)

	go getData[Pagination[Match]](getFootballMatchURL(year, 1, team, 1), ch1)
	go getData[Pagination[Match]](getFootballMatchURL(year, 2, team, 1), ch2)

	result1 := <-ch1
	if result1.Err != nil {
		panic(result1.Err)
	}

	result2 := <-ch2
	if result2.Err != nil {
		panic(result2.Err)
	}

	matchesArr1 := make([]Match, 0)
	matchesArr2 := make([]Match, 0)

	matchesArr1 = append(matchesArr1, result1.Data.PageData...)
	matchesArr2 = append(matchesArr2, result2.Data.PageData...)

	ch1 = make(chan (Response[Pagination[Match]]), result1.Data.TotalPages-1)
	ch2 = make(chan (Response[Pagination[Match]]), result2.Data.TotalPages-1)

	for i := int32(2); i <= result1.Data.TotalPages; i++ {
		go getData[Pagination[Match]](getFootballMatchURL(year, 1, team, i), ch1)
	}

	for i := int32(2); i <= result2.Data.TotalPages; i++ {
		go getData[Pagination[Match]](getFootballMatchURL(year, 2, team, i), ch2)
	}

	for i := int32(2); i <= result1.Data.TotalPages; i++ {
		result := <-ch1
		if result.Err != nil {
			panic(result.Err)
		}

		matchesArr1 = append(matchesArr1, result.Data.PageData...)
	}

	for i := int32(2); i <= result2.Data.TotalPages; i++ {
		result := <-ch2
		if result.Err != nil {
			panic(result.Err)
		}

		matchesArr2 = append(matchesArr2, result.Data.PageData...)
	}

	for _, match := range matchesArr1 {
		goals, err := strconv.Atoi(match.Team1Goals)
		if err != nil {
			panic(err)
		}

		totalGoals += int32(goals)
	}

	for _, match := range matchesArr2 {
		goals, err := strconv.Atoi(match.Team2Goals)
		if err != nil {
			panic(err)
		}

		totalGoals += int32(goals)
	}

	return totalGoals
}

func getFootballMatchURL(year int32, teamNum int32, teamName string, page int32) string {
	return baseURL + fmt.Sprintf("?year=%d&team%d=%s&page=%d", year, teamNum, teamName, page)
}

func getData[T any](url string, ch chan (Response[T])) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- Response[T]{Data: nil, Err: err}
		return
	}

	defer resp.Body.Close()

	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- Response[T]{Data: nil, Err: err}
		return
	}

	var data *T
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		ch <- Response[T]{Data: nil, Err: err}
		return
	}

	ch <- Response[T]{Data: data, Err: nil}
}

func GetTotalGoals2(team string, year int32) int32 {
	totalGoals := int32(0)
	matchesArr1 := make([]Match, 0)
	matchesArr2 := make([]Match, 0)

	res1, err := getData2[Pagination[Match]](getFootballMatchURL2(year, 1, team, 1))
	if err != nil {
		panic(err)
	}
	matchesArr1 = append(matchesArr1, res1.PageData...)

	res2, err := getData2[Pagination[Match]](getFootballMatchURL2(year, 2, team, 1))
	if err != nil {
		panic(err)
	}
	matchesArr2 = append(matchesArr2, res2.PageData...)

	for i := int32(2); i <= res1.TotalPages; i++ {
		res, err := getData2[Pagination[Match]](getFootballMatchURL2(year, 1, team, i))
		if err != nil {
			panic(err)
		}
		matchesArr1 = append(matchesArr1, res.PageData...)
	}

	for i := int32(2); i <= res2.TotalPages; i++ {
		res, err := getData2[Pagination[Match]](getFootballMatchURL2(year, 2, team, i))
		if err != nil {
			panic(err)
		}
		matchesArr2 = append(matchesArr1, res.PageData...)
	}

	for _, match := range matchesArr1 {
		goals, err := strconv.Atoi(match.Team1Goals)
		if err != nil {
			panic(err)
		}

		totalGoals += int32(goals)
	}

	for _, match := range matchesArr2 {
		goals, err := strconv.Atoi(match.Team2Goals)
		if err != nil {
			panic(err)
		}

		totalGoals += int32(goals)
	}

	return totalGoals
}

func getFootballMatchURL2(year int32, teamNum int32, teamName string, page int32) string {
	return baseURL + fmt.Sprintf("?year=%d&team%d=%s&page=%d", year, teamNum, teamName, page)
}

func getData2[T any](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data *T
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
