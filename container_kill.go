package docker

import (
	"context"
	"net/http"
)

// KillContainerOptions represents the set of options that can be used in a
// call to KillContainer.
//
// See https://goo.gl/JnTxXZ for more details.
type KillContainerOptions struct {
	// The ID of the container.
	ID string `qs:"-"`

	// The signal to send to the container. When omitted, Docker server
	// will assume SIGKILL.
	Signal  Signal
	Context context.Context
}

// KillContainer sends a signal to a container, returning an error in case of
// failure.
//
// See https://goo.gl/JnTxXZ for more details.
func (c *Client) KillContainer(opts KillContainerOptions) error {
	path := "/api/endpoints/1/docker/containers/" + opts.ID + "/kill" + "?" + queryString(opts)
	resp, err := c.do(http.MethodPost, path, doOptions{context: opts.Context})
	if err != nil {
		e, ok := err.(*Error)
		if !ok {
			return err
		}
		switch e.Status {
		case http.StatusNotFound:
			return &NoSuchContainer{ID: opts.ID}
		case http.StatusConflict:
			return &ContainerNotRunning{ID: opts.ID}
		default:
			return err
		}
	}
	resp.Body.Close()
	return nil
}
