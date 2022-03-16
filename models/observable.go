package models


type BetConsumer struct {
	BetId string
}

type BetProvider struct {
	Consumers map[string]BetConsumer
	Match_id  string
}

func (p *BetProvider) AddConsumer(consumer BetConsumer) {
	p.Consumers[consumer.BetId] = consumer
}
	

func (p *BetProvider) DeleteConsumer(consumer BetConsumer) {
	delete(p.Consumers, consumer.BetId)
}
func (p *BetProvider) NotifyAll(Match_Result Match_Result, f func(string, Match_Result) (error)) {
	for _, consumer := range p.Consumers {
	consumer.Update(Match_Result, f )
	}

}

func (c *BetConsumer) Update(Match_Result Match_Result,  f func(string, Match_Result) error){
	 f(c.BetId,Match_Result)
}