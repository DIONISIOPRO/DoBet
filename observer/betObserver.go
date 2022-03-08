package observer

import (
	"gitthub.com/dionisiopro/dobet/models"
)

type BetConsumer struct {
	BetId string
}

type Consumers []BetConsumer

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
func (p *BetProvider) NotifyAll(Match_Result models.Match_Result, f func(string,models.Match_Result)) {
	for _, consumer := range p.Consumers {
	consumer.Update(Match_Result, f )
	}

}

func (c *BetConsumer) Update(Match_Result models.Match_Result,  f func(string,models.Match_Result)){
	f(c.BetId,Match_Result)
}