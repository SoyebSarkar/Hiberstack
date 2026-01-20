package typesense

import (
	"fmt"
	"net/http"
)

func (c *Client) Delete(collection string) error {
	req, _ := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/collections/%s", c.BaseURL, collection),
		nil,
	)
	req.Header.Set("X-TYPESENSE-API-KEY", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("delete failed")
	}
	return nil
}
