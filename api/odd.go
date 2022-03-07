package api


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
	Response []struct {
		Bookmakers []struct {
			Bets []struct {
				ID     int64  `json:"id"`
				Name   string `json:"name"`
				Values []struct {
					Odd   string `json:"odd"`
					Value string `json:"value"`
				} `json:"values"`
			} `json:"bets"`
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"bookmakers"`
		Fixture struct {
			Date      string `json:"date"`
			ID        int64  `json:"id"`
			Timestamp int64  `json:"timestamp"`
			Timezone  string `json:"timezone"`
		} `json:"fixture"`
		League struct {
			Country string `json:"country"`
			Flag    string `json:"flag"`
			ID      int64  `json:"id"`
			Logo    string `json:"logo"`
			Name    string `json:"name"`
			Season  int64  `json:"season"`
		} `json:"league"`
		Update string `json:"update"`
	} `json:"response"`
	Results int64 `json:"results"`
}
