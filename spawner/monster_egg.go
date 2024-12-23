package spawner

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type SpawnEgg struct {
	// Kind is the Kind that will be spawned when the egg is used.
	Kind world.EntityType
}

func (m SpawnEgg) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) bool {
	_, ok := tx.Block(pos).(Spawner)
	if ok {
		return false
	}
	ent := newEntities[m.Kind.EncodeEntity()](pos.Add(cube.Pos{0, 1, 0}), tx)
	ctx.SubtractFromCount(1)
	tx.AddEntity(ent)

	return true
}

func (m SpawnEgg) EncodeItem() (name string, meta int16) {
	return m.Kind.EncodeEntity() + "_spawn_egg", 0
}
