package api

import (
	"fmt"
	"net"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/pairecservice"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")

	defaultHttpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			MaxConnsPerHost:       100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
)

// APIClient manages communication with the Pairec Experiment Restful Api API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	*pairecservice.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	region string

	// API Services
	ExperimentApi *ExperimentApiService

	ExperimentGroupApi *ExperimentGroupApiService

	ExperimentRoomApi *ExperimentRoomApiService

	LayerApi *LayerApiService

	SceneApi *SceneApiService

	ParamApi *ParamApiService

	CrowdApi *CrowdApiService

	FlowCtrlApi *FlowCtrlApiService
}

type service struct {
	client     *APIClient
	instanceId string
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(instanceId, region, accessId, accessKey string) (*APIClient, error) {
	client, err := pairecservice.NewClientWithAccessKey(region, accessId, accessKey)
	if err != nil {
		return nil, err
	}
	c := &APIClient{
		Client: client,
		region: region,
	}
	c.common.client = c
	c.common.instanceId = instanceId

	// API Services
	c.ExperimentApi = (*ExperimentApiService)(&c.common)
	c.ExperimentGroupApi = (*ExperimentGroupApiService)(&c.common)
	c.ExperimentRoomApi = (*ExperimentRoomApiService)(&c.common)
	c.LayerApi = (*LayerApiService)(&c.common)
	c.SceneApi = (*SceneApiService)(&c.common)
	c.ParamApi = (*ParamApiService)(&c.common)
	c.CrowdApi = (*CrowdApiService)(&c.common)
	c.FlowCtrlApi = (*FlowCtrlApiService)(&c.common)

	return c, nil
}

func (c APIClient) GetDomain() string {
	return fmt.Sprintf("pairecservice.%s.aliyuncs.com", c.region)
}
func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}
