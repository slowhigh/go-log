package rest_api_number_of_drawn_matches

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pagination[T any] struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_pages"`
	Data      T   `json:"data"`
}

type Match struct {
	Competition string `json:"Competition"`
	Year        int    `json:"year"`
	Round       string `json:"round"`
	Team1       string `json:"team1"`
	Team2       string `json:"team2"`
	Team1Goals  string `json:"team1goals"`
	Team2Goals  string `json:"team2goals"`
}

type Result[T any] struct {
	Data  T
	Error error
}

func getNumDraws(year int) int {
	matches := make([]Match, 0)
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?year=%d", year)

	resp, err := http.Get(url + "&page=1")
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var pn Pagination[[]Match]
	if err := json.Unmarshal(body, &pn); err != nil {
		panic(err)
	}

	totalPage := pn.TotalPage
	matches = append(matches, pn.Data...)
	chArr := make(chan Result[[]Match], totalPage-1)

	for i := 2; i <= totalPage; i++ {
		go getData[Match](url, i, chArr)
	}

	for i := 2; i <= totalPage; i++ {
		res := <-chArr
		if res.Error != nil {
			panic(res.Error)
		}
		
		matches = append(matches, res.Data...)
	}

	drawCount := 0
	for _, match := range matches {
		if match.Team1Goals == match.Team2Goals {
			drawCount++
		}
	}

	return drawCount
}

func getData[T any](url string, curPageNum int, resChan chan (Result[[]T])) {
	resp, err := http.Get(url + fmt.Sprintf("&page=%d", curPageNum))
	if err != nil {
		resChan <- Result[[]T]{Data: []T{}, Error: err}
	}

	defer resp.Body.Close()

	respBodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		resChan <- Result[[]T]{Data: []T{}, Error: err}
	}

	pageInfo := new(Pagination[[]T])
	err = json.Unmarshal(respBodyByte, pageInfo)
	if err != nil {
		resChan <- Result[[]T]{Data: []T{}, Error: err}
	}

	resChan <- Result[[]T]{Data: pageInfo.Data, Error: nil}
}
