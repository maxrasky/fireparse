package fireparse

import "fmt"

type EModes string

const (
	u32         = "U32"
	float3      = "Float3"
	enum        = "Enum"
	requirement = "Requirement"

	GameModes      EModes = "GameModes"
	GameLevels     EModes = "GameLevels"
	GameCharacters EModes = "GameCharacters"
)

func (e EModes) Validate() error {
	switch e {
	case GameModes, GameLevels, GameCharacters:
		return nil
	default:
		return fmt.Errorf("can be one of: [%s,%s,%s]", GameModes, GameLevels, GameCharacters)
	}
}
