package api


type TimeDto struct {
	Errors []interface{} `json:"errors"`
	Get    string        `json:"get"`
	Paging struct {
		Current int64 `json:"current"`
		Total   int64 `json:"total"`
	} `json:"paging"`
	Parameters struct {
		ID string `json:"id"`
	} `json:"parameters"`
	Response []struct {
		Team struct {
			Code     string `json:"code"`
			Country  string `json:"country"`
			Founded  int64  `json:"founded"`
			ID       int64  `json:"id"`
			Logo     string `json:"logo"`
			Name     string `json:"name"`
			National bool   `json:"national"`
		} `json:"team"`
		Venue struct {
			Address  string `json:"address"`
			Capacity int64  `json:"capacity"`
			City     string `json:"city"`
			ID       int64  `json:"id"`
			Image    string `json:"image"`
			Name     string `json:"name"`
			Surface  string `json:"surface"`
		} `json:"venue"`
	} `json:"response"`
	Results int64 `json:"results"`
}
