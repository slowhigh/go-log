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

func getNumDraws(year int) int {
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?year=%d", year)
	matches, err := getAllData[Match](url, 1)
	if err != nil {
		panic(err)
	}

	drawCount := 0
	for _, match := range *matches {
		if match.Team1Goals == match.Team2Goals {
			drawCount++
		}
	}

	return drawCount
}

func getAllData[T any](url string, curPageNum int) (*([]T), error) {
	resp, err := http.Get(url + fmt.Sprintf("&page=%d", curPageNum))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pageInfo := new(Pagination[[]T])
	err = json.Unmarshal(respBodyByte, pageInfo)
	if err != nil {
		return nil, err
	}

	if pageInfo.TotalPage == curPageNum {
		return &pageInfo.Data, nil
	}

	nextPageData, err := getAllData[T](url, curPageNum+1)
	if err != nil {
		return nil, err
	}

	pageInfo.Data = append(pageInfo.Data, (*nextPageData)...)

	return &pageInfo.Data, nil
}
