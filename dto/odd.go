package dto

import "time"

type OddsDto1 struct {
	Get        string `json:"get"`
	Parameters struct {
		Page   string `json:"page"`
		League string `json:"league"`
		Season string `json:"season"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		League struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Country string `json:"country"`
			Logo    string `json:"logo"`
			Flag    string `json:"flag"`
			Season  int    `json:"season"`
		} `json:"league"`
		Fixture struct {
			ID        int       `json:"id"`
			Timezone  string    `json:"timezone"`
			Date      time.Time `json:"date"`
			Timestamp int       `json:"timestamp"`
		} `json:"fixture"`
		Update     time.Time `json:"update"`
		Bookmakers []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Bets []struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Values []struct {
					Value interface{} `json:"value"`
					Odd   string `json:"odd"`
				} `json:"values"`
			} `json:"bets"`
		} `json:"bookmakers"`
	} `json:"response"`
}

type OddsDto struct {
	Errors []interface{} `json:"errors"`
	Get    string        `json:"get"`
	Paging struct {
		Current int64 `json:"current"`
		Total   int64 `json:"total"`
	} `json:"paging"`
	Parameters struct {
		Bookmaker string `json:"bookmaker"`
		Fixture   string `json:"fixture"`
	} `json:"parameters"`
	Response []OddResponse `json:"response"`
	Results  int64         `json:"results"`
}

type OddResponse struct {
	Bookmakers []BookMaker `json:"bookmakers"`
	Fixture    OddFixture  `json:"fixture"`
	League     OddLeague   `json:"league"`
	Update     string      `json:"update"`
}

type BookMaker struct {
	Bets []OddBet `json:"bets"`
	ID   int64    `json:"id"`
	Name string   `json:"name"`
}

type OddBet struct {
	ID     int64      `json:"id"`
	Name   string     `json:"name"`
	Values []OddValue `json:"values"`
}

type OddValue struct {
	Value interface{} `json:"value"`
	Odd   string `json:"odd"`
}

type OddFixture struct {
	Date      string `json:"date"`
	ID        int64  `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Timezone  string `json:"timezone"`
}

type OddLeague struct {
	Country string `json:"country"`
	Flag    string `json:"flag"`
	ID      int64  `json:"id"`
	Logo    string `json:"logo"`
	Name    string `json:"name"`
	Season  int64  `json:"season"`
}
