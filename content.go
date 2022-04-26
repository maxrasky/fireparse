package fireparse

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Content struct {
	Duel Mode `yaml:"duel"`
}

func (c *Content) UnmarshalYAML(value *yaml.Node) error {
	if len(value.Content) >= 2 && value.Content[1].Tag == "Mode" {
		var m Mode
		if err := value.Decode(&m); err != nil {
			return err
		}
		c.Duel = m
		return nil
	}

	return fmt.Errorf("failed to parse duel: should have !<Mode> tag")
}
