package collectors

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RabbitMQClient struct {
	URL      string
	Username string
	Password string
}

type QueuesOutput []struct {
	Messages               int    `json:"messages"`
	MessagesReady          int    `json:"messages_ready"`
	MessagesUnacknowledged int    `json:"messages_unacknowledged"`
	Policy                 string `json:"policy"`
	ExclusiveConsumerTag   string `json:"exclusive_consumer_tag"`
	Consumers              int    `json:"consumers"`
	Memory                 int    `json:"memory"`
	State                  string `json:"state"`
	Name                   string `json:"name"`
	Vhost                  string `json:"vhost"`
	Durable                bool   `json:"durable"`
	AutoDelete             bool   `json:"auto_delete"`
	Node                   string `json:"node"`
}

type RabbitMQCollector interface {
	GetQueueNames() ([]string, error)
}

func (r RabbitMQClient) GetQueueNames() ([]string, error) {
	manager := r.URL
	client := &http.Client{}
	req, _ := http.NewRequest("GET", manager+"/api/queues", nil)
	req.SetBasicAuth(string(r.Username), r.Password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	output, _ := ioutil.ReadAll(resp.Body)
	queueNames, err := UnmarshalQueueNames(output)
	if err != nil {
		return nil, err
	}
	return queueNames, nil
}

func UnmarshalQueueNames(queueJSON []byte) ([]string, error) {
	queueNames := []string{}
	rabbitstuff := QueuesOutput{}

	err := json.Unmarshal(queueJSON, &rabbitstuff)
	if err != nil {
		return nil, err
	}

	for _, queue := range rabbitstuff {
		queueNames = append(queueNames, queue.Name)
	}
	Logger.Infof("RabbitMQ collector retrieved %v queues", len(queueNames))
	return queueNames, nil
}

func RunRabbitMQCollector(r RabbitMQCollector) {
	metricNames, err := r.GetQueueNames()
	if err != nil {
		Logger.Errorf("Unable to ger RabbitMQ queue names: %s", err)
	}
	Logger.Infof("%v", metricNames)

}
