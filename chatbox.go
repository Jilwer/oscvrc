package oscvrc
import (
	"errors"
)

// Chat is a function that sends a message into the VRChat text chat.
// If b is True, send the text in s immediately, bypassing the keyboard.
// If b is False, open the keyboard and populate it with the provided text.
// n is an additional bool parameter that when set to False will not trigger the notification SFX (defaults to True if not specified).
func (c *Client) Chat(s string, b, n bool) error {
	const address = "/chatbox/input"
	if len(s) > 144 {
		return errors.New("message is too long")
	}

	err := c.SendMessage(address, s, b, n)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}

	return nil
}
