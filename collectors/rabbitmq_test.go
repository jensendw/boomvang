package collectors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type FakeRabbitMQ struct {
	Server string
	User   string
}

func (r *FakeRabbitMQ) GetQueueNames() ([]string, error) {
	return []string{"queueA", "queueB"}, nil

}

func TestGetQueueNames(t *testing.T) {
	f := &FakeRabbitMQ{Server: "server", User: "user"}
	fakeQueueNames, _ := f.GetQueueNames()
	assert.Equal(t, []string{"queueA", "queueB"}, fakeQueueNames, "fake queue names should be equal")
}

func TestUnmarshalQueueNames(t *testing.T) {
	multiQueueJSON := []byte(`[{"name":"QueueA"},{"name":"QueueB"}]`)
	multiQueueNames, _ := UnmarshalQueueNames(multiQueueJSON)
	assert.Equal(t, []string{"QueueA", "QueueB"}, multiQueueNames, "Should unmarshal more than one queue")

	singleQueueJSON := []byte(`[{"name":"QueueA"}]`)
	singleQueueNames, _ := UnmarshalQueueNames(singleQueueJSON)
	assert.Equal(t, []string{"QueueA"}, singleQueueNames, "Should unmarshal a single queue")

	badMetricJSON := []byte(`[{"name":"QueueA"},{"name":"QueueB"]`)
	_, err := UnmarshalQueueNames(badMetricJSON)
	assert.NotNil(t, err, "Should error if json cannot be parsed")
}
