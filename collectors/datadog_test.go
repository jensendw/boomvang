package collectors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type FakeDataDogClient struct {
	URL    string
	APIKey string
	AppKey string
}

func (r *FakeDataDogClient) GetDDMetricNames() []string {
	return []string{"metricA", "metricB"}

}

func TestGetDDMetricNames(t *testing.T) {
	f := &FakeDataDogClient{URL: "server", APIKey: "apixxx", AppKey: "appxxx"}
	assert.Equal(t, []string{"metricA", "metricB"}, f.GetDDMetricNames(), "fake metric names should be equal")
}

func TestUnmarshalDatadogMetricNames(t *testing.T) {
	multiMetricJSON := []byte(`{"metrics":["aws.applicationelb.active_connection_count","aws.applicationelb.client_tlsnegotiation_error_count"],"from":"1512405473"}`)
	multiMetricNames, _ := UnmarshalDatadogMetricNames(multiMetricJSON)
	assert.Equal(t, []string{"aws.applicationelb.active_connection_count", "aws.applicationelb.client_tlsnegotiation_error_count"}, multiMetricNames, "Should unmarshal more than one metric name")

	singleMetricJSON := []byte(`{"metrics":["aws.applicationelb.active_connection_count"],"from":"1512405473"}`)
	singleMetricNames, _ := UnmarshalDatadogMetricNames(singleMetricJSON)
	assert.Equal(t, []string{"aws.applicationelb.active_connection_count"}, singleMetricNames, "Should unmarshal a single metric name")

	badMetricJSON := []byte(`{"metrics":["aws.applicationelb.active_connection_count"],"from:"1512405473"`)
	_, err := UnmarshalDatadogMetricNames(badMetricJSON)
	assert.NotNil(t, err, "Should error if json cannot be parsed")
}
