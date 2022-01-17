package form3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL   = "http://localhost:8080"
	defaultUserAgent = "form3"
	mediaTypeJson    = "application/json"
)

//Underline type for services
type service struct {
	client *Client
}

type Response struct {
	*http.Response
}

// API requests can throw errors. Those are wrapped in ErrorResponse
type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"error_message"`
}

func (r *ErrorResponse) Error() string {

	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

type Links struct {
	Self string `json:"self"`
}

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	userAgent  string
	common     service
	Account    *AccountService
}

func NewClient() *Client {
	client := http.DefaultClient
	url, _ := url.Parse(defaultBaseURL)
	c := &Client{httpClient: client, baseURL: url, userAgent: defaultUserAgent}

	c.common.client = c
	c.Account = (*AccountService)(&c.common)

	return c
}

// Set the base URL of the account service. If not the defalut is used
func (c *Client) WithBaseURL(bu string) *Client {
	u, err := url.Parse(bu)
	if err == nil {
		c.baseURL = u
	}

	return c
}

// Set the user agent. If not the defalut is used
func (c *Client) WithUserAgent(ua string) *Client {
	c.userAgent = ua
	return c
}

// Set the http client. If not the defalut is used
func (c *Client) WithHttpClient(client *http.Client) *Client {
	c.httpClient = client
	return c
}

func (c *Client) GET(ctx context.Context, url string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	var req *http.Request

	req, err = http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", mediaTypeJson)
	req.Header.Set("User-Agent", c.userAgent)

	return req, nil
}

func (c *Client) POST(url string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	var req *http.Request

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", mediaTypeJson)

	return req, nil
}

func (c *Client) DELETE(url string) (*http.Request, error) {
	u, err := c.baseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	var req *http.Request

	req, err = http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", mediaTypeJson)

	return req, nil
}

// Call the API and returns the response. JSON decoded response is stored in the value pointed by v.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	response := Response{Response: resp}

	respErr := CheckResponse(resp)

	if respErr != nil {
		return &response, respErr
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}

	return &response, err
}

// Check the API response whether it is success of fail. Status code between 200 and 299,
// including both values are considerd as successful response.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			errorResponse.Message = string(data)
		}
	}

	return errorResponse
}
