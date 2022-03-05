package services

import "gitthub.com/dionisiopro/dobet/models"

type IBetProvider interface {
	AddConsumer(consumer IBetConsumer)
	DeleteConsumer(consumer IBetConsumer)
	NotifyAll(Match_Result models.Match_Result)
}

type IBetConsumer interface {
	Update(Match_Result models.Match_Result)
}

type BetConsumer struct {
	service BetService
	BetId string
}

type Consumers []BetConsumer

type BetProvider struct {
	Consumers map[string]IBetConsumer
	Match_id  string
}

func (p *BetProvider) AddConsumer(consumer BetConsumer) {
	p.Consumers[consumer.BetId] = consumer
}
	

func (p *BetProvider) DeleteConsumer(consumer BetConsumer) {
	delete(p.Consumers, consumer.BetId)

}
func (p *BetProvider) NotifyAll(Match_Result models.Match_Result) {
	for _, consumer := range p.Consumers {
	consumer.Update(Match_Result)
	}

}

func (c BetConsumer) Update(Match_Result models.Match_Result){
	c.service.ProcessBet(c.BetId,Match_Result)
}