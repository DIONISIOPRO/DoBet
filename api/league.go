package api

type LeagueDto struct {
	MyError []interface{} `json:"errors"`
	Get    string        `json:"get"`
	Paging struct {
		Current int64 `json:"current"`
		Total   int64 `json:"total"`
	} `json:"paging"`
	Parameters struct {
		ID string `json:"id"`
	} `json:"parameters"`
	Response []struct {
		Country struct {
			Code string `json:"code"`
			Flag string `json:"flag"`
			Name string `json:"name"`
		} `json:"country"`
		League struct {
			ID   int64  `json:"id"`
			Logo string `json:"logo"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"league"`
		Seasons []struct {
			Coverage struct {
				Fixtures struct {
					Events             bool `json:"events"`
					Lineups            bool `json:"lineups"`
					StatisticsFixtures bool `json:"statistics_fixtures"`
					StatisticsPlayers  bool `json:"statistics_players"`
				} `json:"fixtures"`
				Injuries    bool `json:"injuries"`
				Odds        bool `json:"odds"`
				Players     bool `json:"players"`
				Predictions bool `json:"predictions"`
				Standings   bool `json:"standings"`
				TopAssists  bool `json:"top_assists"`
				TopCards    bool `json:"top_cards"`
				TopScorers  bool `json:"top_scorers"`
			} `json:"coverage"`
			Current bool   `json:"current"`
			End     string `json:"end"`
			Start   string `json:"start"`
			Year    int64  `json:"year"`
		} `json:"seasons"`
	} `json:"response"`
	Results int64 `json:"results"`
}