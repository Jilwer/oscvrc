package oscvrc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/user"
	"runtime"
)

type AvatarParamConfig struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Parameters []ParameterConfig `json:"parameters"`
}

type ParameterConfig struct {
	Name   string       `json:"name"`
	Input  InputConfig  `json:"input,omitempty"`
	Output OutputConfig `json:"output"`
}

type InputConfig struct {
	Address string `json:"address"`
	Type    string `json:"type"`
	client  *Client
}

type OutputConfig struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}

func (ic *InputConfig) setInputClient(c *Client) {
	ic.client = c
}

// ReadAvatarParamConfig reads the avatar parameter configuration from the specified file.
// The returned struct should be used in conjunction with the client.SendMessage function
// to send messages to the VRChat avatar parameters. 
//The `path` can be specified to read from a custom location. If `path` is empty, it will default to the standard VRChat
// path based on the operating system. It should point to ~/AppData/LocalLow/VRChat/VRChat/OSC/ 
func (c *Client) ReadAvatarParamConfig(avatarId, userId string, path string) (AvatarParamConfig, error) {

	user, err := user.Current()
	if err != nil {
		return AvatarParamConfig{}, fmt.Errorf("failed to get current user: %w", err)
	}

	if path == "" {
		switch runtime.GOOS {
		case "windows":
			path = fmt.Sprintf(`%s\AppData\LocalLow\VRChat\VRChat\OSC\%s\Avatars\%s.json`, user.HomeDir, userId, avatarId)
		case "linux":
			path = fmt.Sprintf(`%s/.local/share/Steam/steamapps/compatdata/438100/pfx/drive_c/users/steamuser/AppData/LocalLow/VRChat/VRChat/OSC/%s/Avatars/%s.json`, user.HomeDir, userId, avatarId)
		default:
			return AvatarParamConfig{}, errors.New("unsupported operating system or unknown path")
		}
	} else {
		switch runtime.GOOS {
		case "windows":
			path = fmt.Sprintf(`%s\%s\Avatars\%s.json`, path, userId, avatarId)
		case "linux":
			path = fmt.Sprintf(`%s/%s/Avatars/%s.json`, path, userId, avatarId)
		default:
			return AvatarParamConfig{}, errors.New("unsupported operating system or unknown path")
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return AvatarParamConfig{}, fmt.Errorf("failed to read file: %w", err)
	}

	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf")) // remove BOM

	var avatarParamConfig AvatarParamConfig
	err = json.Unmarshal(data, &avatarParamConfig)
	if err != nil {
		return AvatarParamConfig{}, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	for i := range avatarParamConfig.Parameters {
		avatarParamConfig.Parameters[i].Input.setInputClient(c)
	}

	return avatarParamConfig, nil
}

func (i *InputConfig) Send(value ...interface{}) error {
	if i.client == nil {
		return errors.New("client not set")
	}

	err := i.client.SendMessage(i.Address, value...)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
