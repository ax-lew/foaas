package foaas

import (
	"encoding/json"
	"fmt"
	"github.com/ax-lew/foaas/domain/model"
	"github.com/ax-lew/foaas/errors"
	"github.com/ax-lew/foaas/logger"
	"io/ioutil"
	"net/http"
)

const messageFrom = "Axel"

type Client struct {
	host   string
	client *http.Client
}

func NewClient(host string, client *http.Client) *Client {
	return &Client{host: host, client: client}
}

func (c *Client) GetFuckOff(userID string) (*model.Response, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.host, "bday", userID, messageFrom)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewInternalError(fmt.Sprintf("unexpected status code from foaas: %d", resp.StatusCode))
	}
	httpBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		logger.Logger.Errorf("error while closing request body %s", err.Error())
	}

	response := &model.Response{}
	err = json.Unmarshal(httpBody, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
