package typesense

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Export(collection string) (io.ReadCloser, error) {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/collections/%s/documents/export", c.BaseURL, collection),
		nil,
	)
	req.Header.Set("X-TYPESENSE-API-KEY", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		resp.Body.Close()
		return nil, fmt.Errorf("export failed")
	}

	return resp.Body, nil
}
