package fireparse

import (
	"fmt"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Mode struct {
	TeamSize     int           `yaml:"team_size"`
	DurationS    int           `yaml:"duration_s"`
	Position     []float64     `yaml:"position"`
	Levels       []string      `yaml:"levels"`
	Requirements []Requirement `yaml:"requirements"`
}

func (m *Mode) UnmarshalYAML(value *yaml.Node) error {
	if len(value.Content) < 4 {
		return fmt.Errorf("failed to parse mode: content len='%d'", len(value.Content))
	}
	// skipping first 2 nodes (as they are duel related ones)
	for idx := 2; idx < len(value.Content)-1; idx += 2 {
		str := value.Content[idx+1].Value
		tag := value.Content[idx+1].Tag
		content := value.Content[idx+1].Content

		switch value.Content[idx].Value {
		case "team_size":
			if tag == u32 {
				teamSize, err := strconv.Atoi(str)
				if err != nil {
					return err
				}
				m.TeamSize = teamSize
			}
		case "duration_s":
			if tag != u32 {
				continue
			}
			duration, err := strconv.Atoi(str)
			if err != nil {
				return err
			}
			m.DurationS = duration
		case "position":
			if tag != float3 {
				continue
			}
			if len(content) != 3 {
				return fmt.Errorf("should be array of 3 float points")
			}

			position := make([]float64, 0, 3)
			for _, pValue := range content {
				p, err := strconv.ParseFloat(pValue.Value, 64)
				if err != nil {
					return fmt.Errorf("failed to parse float value='%s'", pValue.Value)
				}
				position = append(position, p)
			}
			m.Position = position
		case "levels":
			for _, pValue := range content {
				m.Levels = append(m.Levels, pValue.Value)
			}
		case "requirements":
			for _, pValue := range content {
				var req Requirement
				if err := pValue.Decode(&req); err != nil {
					return err
				}
				m.Requirements = append(m.Requirements, req)
			}
		}
	}
	return nil
}
