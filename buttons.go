package vrcinput

import (
	"errors"
	"time"

	"github.com/hypebeast/go-osc/osc"
)

type MoveDirection string

const (
	MoveForward  MoveDirection = "MoveForward"
	MoveBackward MoveDirection = "MoveBackward"
	MoveLeft     MoveDirection = "MoveLeft"
	MoveRight    MoveDirection = "MoveRight"
)

func (c *localOscClient) Move(direction MoveDirection, b bool) error {
	message := osc.NewMessage("/input/"+string(direction), b)
	err := c.Send(message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	return nil
}

type LookDirection string

const (
	LookLeft  LookDirection = "LookLeft"
	LookRight LookDirection = "LookRight"
)

func (c *localOscClient) Look(direction LookDirection, b bool) error {
	message := osc.NewMessage("/input/"+string(direction), b)
	err := c.Send(message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	return nil
}

func (c *localOscClient) Jump() error {
	var jumping bool
	for i := 0; i < 2; i++ {
		jumping = !jumping
		message := osc.NewMessage("/input/Jump", jumping)
		err := c.Send(message)
		if err != nil {
			return errors.New("failed to send message: " + err.Error())
		}

		// Sometimes it doesn't work if we don't sleep for a bit
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

// TODO: Items, PanicButton, QuickMenu, and Voice
