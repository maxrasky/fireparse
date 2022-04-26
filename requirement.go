package fireparse

import (
	"fmt"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Requirement struct {
	Name       string
	MinBattles uint32 `yaml:"min_battles"`
}

func (r *Requirement) UnmarshalYAML(value *yaml.Node) error {
	if len(value.Content) < 2 {
		return fmt.Errorf("failed to parse requirements: content len='%d'", len(value.Content))
	}

	tag := value.Content[1].Tag
	if tag != requirement {
		return fmt.Errorf("%s should be !<Requirement> tag", tag)
	}
	r.Name = value.Content[0].Value

	for idx := 2; idx < len(value.Content)-1; idx += 2 {
		sBattles := value.Content[idx+1].Value
		tag = value.Content[idx+1].Tag

		switch value.Content[idx].Value {
		case "min_battles":
			if tag == u32 {
				minBattles, err := strconv.Atoi(sBattles)
				if err != nil {
					return fmt.Errorf("failed to convert to int value='%s'", sBattles)
				}
				r.MinBattles = uint32(minBattles)
			}
		}
	}
	return nil
}
