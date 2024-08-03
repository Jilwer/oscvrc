package vrcosc

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
}

type OutputConfig struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}

func ReadAvatarParamConfig(avatarId, userId string) (AvatarParamConfig, error) {

	var path string

	user, err := user.Current()
	if err != nil {
		return AvatarParamConfig{}, fmt.Errorf("failed to get current user: %w", err)
	}

	switch runtime.GOOS {
	case "windows":
		path = fmt.Sprintf(`%s\AppData\LocalLow\VRChat\VRChat\OSC\%s\Avatars\%s.json`, user.HomeDir, userId, avatarId)
	case "linux":
		path = fmt.Sprintf(`%s/.local/share/Steam/steamapps/compatdata/438100/pfx/drive_c/users/steamuser/AppData/LocalLow/VRChat/VRChat/OSC/%s/Avatars/%s.json`, user.HomeDir, userId, avatarId)
	default:
		return AvatarParamConfig{}, errors.New("unsupported operating system")
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

	return avatarParamConfig, nil
}
