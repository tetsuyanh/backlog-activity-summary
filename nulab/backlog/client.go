package backlog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tetsuyanh/backlog-activity-summary/nulab/backlog/model"
)

const (
	TIMEOUT_MSEC = 3000
)

var (
	team  string
	token string
)

type (
	// Client represents client.
	Client struct {
		team  string
		token string
	}
)

// Setup is setup
func Setup(teamIn, tokenIn string) {
	team = teamIn
	token = tokenIn
}

// NewClient returns Client.
func NewClient() *Client {
	return &Client{
		team:  team,
		token: token,
	}
}

func (c *Client) Myself() (*model.User, error) {
	url := c.buildURL("users/myself", nil)
	u := &model.User{}
	return u, c.call(http.MethodGet, url, nil, u)
}

func (c *Client) RecentlyActivities(userID int, params *Params) ([]*model.Activity, error) {
	url := c.buildURL(fmt.Sprintf("users/%d/activities", userID), params)
	a := []*model.Activity{}
	return a, c.call(http.MethodGet, url, nil, &a)
}

func (c *Client) Stars(userID int, params *Params) ([]*model.Star, error) {
	url := c.buildURL(fmt.Sprintf("users/%d/stars", userID), params)
	s := []*model.Star{}
	return s, c.call(http.MethodGet, url, nil, &s)
}

func (c *Client) StarsCount(userID int) (*model.Count, error) {
	url := c.buildURL(fmt.Sprintf("users/%d/stars/count", userID), nil)
	cnt := &model.Count{}
	return cnt, c.call(http.MethodGet, url, nil, cnt)
}

// call is to request by net/http.
func (c *Client) call(method, url string, data []byte, result interface{}) error {
	req, errReq := http.NewRequest(method, url, bytes.NewReader(data))
	if errReq != nil {
		return errReq
	}

	ctx, cancel := context.WithTimeout(req.Context(), TIMEOUT_MSEC*time.Millisecond)
	defer cancel()
	req = req.WithContext(ctx)

	resp, errDo := http.DefaultClient.Do(req)
	if errDo != nil {
		return errDo
	}
	defer resp.Body.Close()

	if errParse := c.parseError(resp); errParse != nil {
		return errParse
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

func (c *Client) buildURL(path string, params *Params) string {
	vals := url.Values{}
	vals.Set("apiKey", c.token)
	if params != nil {
		for k, v := range params.Map() {
			vals.Set(k, v)
		}
	}
	u := url.URL{
		Scheme:   "https",
		Host:     fmt.Sprintf("%s.backlog.jp", c.team),
		Path:     fmt.Sprintf("api/v2/%s", path),
		RawQuery: vals.Encode(),
	}
	return u.String()
}

func (c *Client) parseError(resp *http.Response) error {
	if resp.StatusCode < 400 {
		return nil
	}

	if resp.StatusCode < 500 {
		es := &model.Errors{}
		errDecode := json.NewDecoder(resp.Body).Decode(es)
		if errDecode != nil && errDecode.Error() != "EOF" {
			return errDecode
		}
		return es
	}

	return fmt.Errorf("server response %d", resp.StatusCode)
}
