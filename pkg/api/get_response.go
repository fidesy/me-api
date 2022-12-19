package api

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"net/http"
)

func (c *Client) GetResponse(ctx context.Context, url string, dataStruct interface{}) error {
	req, _ := http.NewRequest("GET", url, nil)
	req = req.WithContext(ctx)
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyData, &dataStruct)
	if err != nil {
		// Unmarshal error
		// Return response error message from the ME API
		return errors.New(string(bodyData))
	}

	return nil
}
