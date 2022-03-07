package dto


type BookMakerDto struct {
	Errors []interface{} `json:"errors"`
	Get    string        `json:"get"`
	Paging struct {
		Current int64 `json:"current"`
		Total   int64 `json:"total"`
	} `json:"paging"`
	Parameters []interface{} `json:"parameters"`
	Response   []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"response"`
	Results int64 `json:"results"`
}
