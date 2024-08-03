package oscvrc

import (
	"errors"

	"github.com/hypebeast/go-osc/osc"
)

// SendMessage is a helper function that sends an OSC message to VRChat.
func (c *Client) SendMessage(address string, value ...interface{}) error {

	// validate the value types
	for _, v := range value {
		switch v.(type) {
		case int32:
		case float32:
		case bool:
		case string:
		default:
			return errors.New("unsupported value type")
		}
	}

	message := osc.NewMessage(address, value...)
	err := c.Send(message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	return nil
}
