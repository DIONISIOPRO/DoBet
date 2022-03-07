package dto

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
	Odd   string `json:"odd"`
	Value string `json:"value"`
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
