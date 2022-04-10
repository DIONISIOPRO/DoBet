package event

import (
	"github/namuethopro/dobet-user/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventPublisher(t *testing.T) {
	t.Run("emptly queue", func(t *testing.T) {
		publisher := NewRabbitMQEventPublisher(nil)
		err := publisher.Publish("", domain.AddMoneyEvent{Amount: 10, UserId: "id"})
		assert.NotNil(t, err)
		assert.Equal(t, "invalid parameters", err.Error())
	})
	
	t.Run("event nil", func(t *testing.T) {
		publisher := NewRabbitMQEventPublisher(nil)
		err := publisher.Publish("some", nil)
		assert.NotNil(t, err)
		assert.Equal(t, "invalid parameters", err.Error())
	})
}
