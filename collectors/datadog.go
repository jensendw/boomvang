package collectors

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type DatadogClient struct {
	APIKey string
	AppKey string
	URL    string
}

type DatadogCollector interface {
	GetDDMetricNames() []string
	pushDDMetricNames() error
}

type DatadogMetricNames struct {
	Metrics []string `json:"metrics"`
	From    string   `json:"from"`
}

func (d DatadogClient) GetDDMetricNames() ([]string, error) {
	req, err := http.NewRequest("GET", d.URL+"/v1/metrics", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("api_key", d.APIKey)
	q.Add("application_key", d.AppKey)
	q.Add("from", strconv.FormatInt(time.Now().Add(-24*time.Hour).Unix(), 10))
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	output, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	metricNames, err := UnmarshalDatadogMetricNames(output)
	if err != nil {
		return nil, err
	}

	return metricNames, nil
}

func UnmarshalDatadogMetricNames(metricsJSON []byte) ([]string, error) {
	metricNames := DatadogMetricNames{}

	err := json.Unmarshal(metricsJSON, &metricNames)
	if err != nil {
		return nil, err
	}

	Logger.Infof("Datadog collector retrieved %v metric names", len(metricNames.Metrics))
	return metricNames.Metrics, nil
}

func RunDDCollector(d DatadogClient) {
	metricNames, err := d.GetDDMetricNames()
	if err != nil {
		Logger.Errorf("Unable to get Datadog metric names: %s", err)
	}
	Logger.Infof("%v", metricNames)
	//d.pushDDMetricNames(AllMetricNames{DataDog: metricNames})
}
