package spawner

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/item"
)

// newBreakInfo creates a BreakInfo struct with the properties passed. The XPDrops field is 0 by default. The blast
// resistance is set to the block's hardness*5 by default.
func newBreakInfo(hardness float64, harvestable func(item.Tool) bool, effective func(item.Tool) bool, drops func(item.Tool, []item.Enchantment) []item.Stack) block.BreakInfo {
	return block.BreakInfo{
		Hardness:        hardness,
		BlastResistance: hardness * 5,
		Harvestable:     harvestable,
		Effective:       effective,
		Drops:           drops,
	}
}
