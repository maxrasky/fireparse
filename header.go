package fireparse

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Header struct {
	Type map[EModes]struct{} `yaml:"type"`
}

func (h *Header) UnmarshalYAML(value *yaml.Node) error {
	if h.Type == nil {
		h.Type = make(map[EModes]struct{})
	}

	if len(value.Content) >= 2 {
		tag := value.Content[1].ShortTag()
		if tag == enum {
			tagValue := value.Content[1].Value
			arr := strings.Split(tagValue, " ")
			for _, v := range arr {
				emode := EModes(v)
				if err := emode.Validate(); err != nil {
					return fmt.Errorf("invalid enum type: %w", err)
				}
				h.Type[emode] = struct{}{}
			}
		}
		return nil
	}

	var name map[string]interface{}
	if err := value.Decode(&name); err != nil {
		return err
	}
	for _, v := range name {
		if sv, ok := v.(string); ok {
			emode := EModes(sv)
			if err := emode.Validate(); err != nil {
				return fmt.Errorf("invalid enum type: %w", err)
			}
			h.Type[emode] = struct{}{}
		}
	}
	return nil
}
