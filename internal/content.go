package internal

import (
	"encoding/json"
	"errors"
	"io"
)

type content struct {
	Meta Meta            `json:"meta"`
	Data json.RawMessage `json:"data"`
}

func Parse(body io.Reader, v any) error {
	buf, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	var c content
	if err := json.Unmarshal(buf, &c); err != nil {
		return err
	}

	if c.Meta.Code >= 300 {
		return errors.New(c.Meta.Message)
	}

	if len(c.Data) == 0 {
		return nil
	}

	return json.Unmarshal(c.Data, &v)
}
