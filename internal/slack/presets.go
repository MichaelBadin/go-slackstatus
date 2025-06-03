package slack

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Preset struct {
	Text   string `json:"text" yaml:"text"`
	Emoji  string `json:"emoji" yaml:"emoji"`
	Expiry string `json:"expiry" yaml:"expiry"`
}

func LoadPresets(path string) (map[string]Preset, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	presets := make(map[string]Preset)
	switch {
	case strings.HasSuffix(path, ".json"):
		err = json.Unmarshal(data, &presets)
	case strings.HasSuffix(path, ".yaml"), strings.HasSuffix(path, ".yml"):
		err = yaml.Unmarshal(data, &presets)
	default:
		return nil, errors.New("unsupported file format (use .json or .yaml)")
	}

	return presets, err
}

func ParseDuration(dur string) int {
	if dur == "" || dur == "0s" {
		return 0
	}
	d, err := time.ParseDuration(dur)
	if err != nil {
		return 0
	}
	return int(time.Now().Add(d).Unix())
}

func ListPresets(path string) error {
	presets, err := LoadPresets(path)
	if err != nil {
		return err
	}

	names := make([]string, 0, len(presets))
	for name := range presets {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("Available Presets:")
	for _, name := range names {
		p := presets[name]
		fmt.Printf("  %-10s → %s %s (expires: %s)\n", name, p.Emoji, p.Text, p.Expiry)
	}
	return nil
}
