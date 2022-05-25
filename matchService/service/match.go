package service
 import(
	 "github.com/chromedp/chromedp"
 )
type MatchRepository interface {
}
type MatchService struct {
	repository MatchRepository
}

func NewMatchService(repository MatchRepository) *MatchService {
	return &MatchService{
		repository: repository,
	}
}


func (s MatchService) createMatch(){
	
}
