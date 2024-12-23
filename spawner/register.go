package spawner

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/creative"
	"github.com/df-mc/dragonfly/server/world"
)

var (
	newEntities = map[string]func(cube.Pos, *world.Tx) *world.EntityHandle{}
	entities    = map[string]world.EntityType{}
)

func RegisterEntityType(kind world.EntityType, newEnt func(cube.Pos, *world.Tx) *world.EntityHandle) {
	newEntities[kind.EncodeEntity()] = newEnt
	entities[kind.EncodeEntity()] = kind

	world.RegisterItem(SpawnEgg{Kind: kind})
	creative.RegisterItem(item.NewStack(SpawnEgg{Kind: kind}, 1))
}

func init() {
	world.RegisterBlock(Spawner{})
	world.RegisterItem(Spawner{})
	creative.RegisterItem(item.NewStack(Spawner{
		MaxNearbyEntities:   6,
		MaxSpawnDelay:       800,
		MinSpawnDelay:       200,
		RequiredPlayerRange: 16,
		Movable:             true,
		Delay:               20,
		SpawnCount:          4,
		SpawnRange:          4,
	}, 1))
}
