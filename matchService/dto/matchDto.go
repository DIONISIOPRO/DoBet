package dto

type MatchDto struct {
	Errors []interface{} `json:"errors"`
	Get    string        `json:"get"`
	Paging struct {
		Current int64 `json:"current"`
		Total   int64 `json:"total"`
	} `json:"paging"`
	Parameters struct {
		Live string `json:"live"`
	} `json:"parameters"`
	Response []struct {
		Fixture Fixture `json:"fixture"`
		Goals   Goals   `json:"goals"`
		League  League  `json:"league"`
		Score   Score   `json:"score"`
		Teams   Teams   `json:"teams"`
	} `json:"response"`
	Results int64 `json:"results"`
}

type Fixture struct {
	Date      string      `json:"date"`
	ID        int64       `json:"id"`
	Periods   Periods     `json:"periods"`
	Referee   interface{} `json:"referee"`
	Status    Status      `json:"status"`
	Timestamp int64       `json:"timestamp"`
	Timezone  string      `json:"timezone"`
	Venue     Venue       `json:"venue"`
}

type Periods struct {
	First  int64       `json:"first"`
	Second interface{} `json:"second"`
}

type Status struct {
	Elapsed int64  `json:"elapsed"`
	Long    string `json:"long"`
	Short   string `json:"short"`
}

type Venue struct {
	City string `json:"city"`
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type League struct {
	Country string `json:"country"`
	Flag    string `json:"flag"`
	ID      int64  `json:"id"`
	Logo    string `json:"logo"`
	Name    string `json:"name"`
	Round   string `json:"round"`
	Season  int64  `json:"season"`
}

type Goals struct {
	Away int64 `json:"away"`
	Home int64 `json:"home"`
}

type Score struct {
	Extratime struct {
		Away interface{} `json:"away"`
		Home interface{} `json:"home"`
	} `json:"extratime"`
	Fulltime struct {
		Away interface{} `json:"away"`
		Home interface{} `json:"home"`
	} `json:"fulltime"`
	Halftime struct {
		Away int64 `json:"away"`
		Home int64 `json:"home"`
	} `json:"halftime"`
	Penalty struct {
		Away interface{} `json:"away"`
		Home interface{} `json:"home"`
	} `json:"penalty"`
}

type Teams struct {
	Away struct {
		ID     int64  `json:"id"`
		Logo   string `json:"logo"`
		Name   string `json:"name"`
		Winner bool   `json:"winner"`
	} `json:"away"`
	Home struct {
		ID     int64  `json:"id"`
		Logo   string `json:"logo"`
		Name   string `json:"name"`
		Winner bool   `json:"winner"`
	} `json:"home"`
}
