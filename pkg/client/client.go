package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
)

type slackThatClient struct {
	httpClient *http.Client
	baseURL *url.URL
}

func (client *slackThatClient) getURL(endpoint string) string {
	urlCopy := &url.URL{}
	*urlCopy = *client.baseURL

	urlCopy.Path = path.Join(urlCopy.Path, endpoint)
	return urlCopy.String()
}

func (client *slackThatClient) request(method string, endpoint string, reader io.Reader, v interface{}) error {
	request, err := http.NewRequest(method, client.getURL(endpoint), reader)
	if err != nil {
		return err
	}
	resp, err := client.httpClient.Do(request)
	if err != nil {
		return err
	}

	if err := validateResponse(resp); err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(v)
}


// PostMessage makes a POST request to the slack_that API to send a message to a channel or multiple users in
// a workspace.
func (client *slackThatClient) PostMessage(options ...PostMessageOptions) (map[string]interface{}, error) {
	params := defaultPostMessageParameters()
	for _, opt := range options {
		opt(params)
	}

	data := make(map[string]interface{})
	if err := client.request(http.MethodPost, "/", params, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// GetHealth makes a GET request to the slack_that API's health endpoint.
func (client *slackThatClient) GetHealth() (map[string]interface{}, error) {
	data := make(map[string]interface{})
	if err := client.request(http.MethodGet, "/health", nil, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// NewClient returns a SlackThat client ready to make requests to the base_url passed.
func NewClient(baseURL string, baseClient *http.Client) (SlackThatClient, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &slackThatClient{
		httpClient: baseClient,
		baseURL: parsedURL,
	}, nil
}