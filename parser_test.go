package fireparse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		Name     string
		Err      string
		Data     string
		Expected *Config
	}{
		{
			Name: "success",
			Data: `
header:
  type: !<Enum> GameModes GameLevels
content:
  duel: !<Mode>
  team_size: !<U32> 1
  duration_s: !<U32> 60
  position: !<Float3> [1.0, 2.0, 3.0]
  requirements:
    - Player: !<Requirement>
      min_battles: !<U32> 4
    - Squad: !<Requirement> 
      min_battles: !<U32> 4
  levels:
    - levels/level_1 
    - levels/level_2 
    - levels/level_3
`,
			Expected: &Config{
				Header: Header{
					Type: map[EModes]struct{}{GameModes: {}, GameLevels: {}},
				},
				Content: Content{
					Duel: Mode{
						TeamSize:  1,
						DurationS: 60,
						Position:  []float64{1.0, 2.0, 3.0},
						Levels:    []string{"levels/level_1", "levels/level_2", "levels/level_3"},
						Requirements: []Requirement{
							{
								Name:       "Player",
								MinBattles: 4,
							},
							{
								Name:       "Squad",
								MinBattles: 4,
							},
						},
					},
				},
			},
		},
		{
			Name: "invalid enum",
			Err:  "invalid enum type",
			Data: `
header:
  type: !<Enum> GameNodes
content:
  duel: !<Mode>
  team_size: !<U32> 1
  duration_s: !<U32> 60
  position: !<Float3> [1.0, 2.0, 3.0]
  requirements:
    - Player: !<Requirement>
      min_battles: !<U32> 4
    - Squad: !<Requirement> 
      min_battles: !<U32> 4
  levels:
    - levels/level_1 
    - levels/level_2 
    - levels/level_3
`,
		},
		{
			Name: "too much points",
			Err:  "should be array of 3 float points",
			Data: `
header:
  type: !<Enum> GameModes
content:
  duel: !<Mode>
  team_size: !<U32> 1
  duration_s: !<U32> 60
  position: !<Float3> [1.0, 2.0, 3.0, 1.0]
  requirements:
    - Player: !<Requirement>
      min_battles: !<U32> 4
    - Squad: !<Requirement> 
      min_battles: !<U32> 4
  levels:
    - levels/level_1 
    - levels/level_2 
    - levels/level_3
`,
		},
		{
			Name: "lacks requirement",
			Err:  "should be !<Requirement> tag",
			Data: `
header:
  type: !<Enum> GameModes
content:
  duel: !<Mode>
  team_size: !<U32> 1
  duration_s: !<U32> 60
  position: !<Float3> [1.0, 2.0, 3.0]
  requirements:
    - Player:
      min_battles: !<U32> 4
  levels:
    - levels/level_1 
    - levels/level_2 
    - levels/level_3
`,
		},
	}

	for _, tt := range tests {
		cfg, err := Parse(tt.Data)
		if tt.Err == "" {
			require.NoError(t, err)
			require.Equal(t, tt.Expected, cfg)
		} else {
			require.Error(t, err)
			require.ErrorContains(t, err, tt.Err)
		}
	}
}
