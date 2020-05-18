package docker

import (
	"encoding/json"
	"net/http"
)

// ContainerChanges returns changes in the filesystem of the given container.
//
// See https://goo.gl/15KKzh for more details.
func (c *Client) ContainerChanges(id string) ([]Change, error) {
	path := "/api/endpoints/1/docker/containers/" + id + "/changes"
	resp, err := c.do(http.MethodGet, path, doOptions{})
	if err != nil {
		if e, ok := err.(*Error); ok && e.Status == http.StatusNotFound {
			return nil, &NoSuchContainer{ID: id}
		}
		return nil, err
	}
	defer resp.Body.Close()
	var changes []Change
	if err := json.NewDecoder(resp.Body).Decode(&changes); err != nil {
		return nil, err
	}
	return changes, nil
}
