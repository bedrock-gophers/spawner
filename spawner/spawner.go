package spawner

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/model"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"math/rand"
)

type Spawner struct {
	EntityType          world.EntityType
	Delay               int
	Movable             bool
	RequiredPlayerRange int
	MaxNearbyEntities   int
	MaxSpawnDelay       int
	MinSpawnDelay       int
	SpawnCount          int
	SpawnRange          int

	pos cube.Pos
}

func (s Spawner) DecodeNBT(data map[string]any) any {
	s.Delay = int(data["Delay"].(int16))
	s.Movable = data["isMovable"].(byte) == 1
	s.RequiredPlayerRange = int(data["RequiredPlayerRange"].(int16))
	s.MaxNearbyEntities = int(data["MaxNearbyEntities"].(int16))
	s.MaxSpawnDelay = int(data["MaxSpawnDelay"].(int16))
	s.MinSpawnDelay = int(data["MinSpawnDelay"].(int16))
	s.SpawnCount = int(data["SpawnCount"].(int16))
	s.SpawnRange = int(data["SpawnRange"].(int16))

	if id := data["EntityIdentifier"].(string); id != "" {
		s.EntityType = entities[id]
	}
	return s
}

func (s Spawner) EncodeNBT() map[string]any {
	var entityID string
	if s.EntityType != nil {
		entityID = s.EntityType.EncodeEntity()
	}
	return map[string]any{
		"Delay":               int16(s.Delay),
		"DisplayEntityHeight": float32(1),
		"DisplayEntityWidth":  float32(1),
		"EntityIdentifier":    entityID,
		"MaxNearbyEntities":   int16(s.MaxNearbyEntities),
		"MaxSpawnDelay":       int16(s.MaxSpawnDelay),
		"MinSpawnDelay":       int16(s.MinSpawnDelay),
		"RequiredPlayerRange": int16(s.RequiredPlayerRange),
		"SpawnCount":          int16(s.SpawnCount),
		"SpawnRange":          int16(s.SpawnRange),
		"id":                  "MobSpawner",
		"isMovable":           boolToByte(s.Movable),
		"x":                   int32(s.pos.X()),
		"y":                   int32(s.pos.Y()),
		"z":                   int32(s.pos.Z()),
	}
}

func boolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func (s Spawner) Tick(_ int64, pos cube.Pos, w *world.World) {
	s.pos = pos
	s.Delay--

	if s.Delay > 0 {
		w.SetBlock(pos, s, nil)
		return
	}

	minRange := s.RequiredPlayerRange
	p1, p2 := s.pos.Add(cube.Pos{-minRange, -minRange, -minRange}), s.pos.Add(cube.Pos{minRange, minRange, minRange})
	x0, y0, z0, x1, y1, z1 := float64(p1.X()), float64(p1.Y()), float64(p1.Z()), float64(p2.X()), float64(p2.Y()), float64(p2.Z())

	if len(w.EntitiesWithin(cube.Box(x0, y0, z0, x1, y1, z1), func(entity world.Entity) bool {
		_, ok := entity.(*player.Player)
		return !ok
	})) <= 0 {
		return
	}

	if len(w.EntitiesWithin(cube.Box(x0, y0, z0, x1, y1, z1), func(entity world.Entity) bool {
		return entity.Type() != s.EntityType
	})) >= s.MaxNearbyEntities {
		return
	}
	s.SpawnCount = rand.Intn(4)
	blockPos := pos.Vec3()

	for i := 0; i < s.SpawnCount; i++ {
		var spawnPos mgl64.Vec3

		if rand.Float64() > 0.5 {
			spawnPos = blockPos.Add(mgl64.Vec3{rand.Float64() * 1.5, 1, rand.Float64() * 1.5})
		} else {
			spawnPos = blockPos.Sub(mgl64.Vec3{-rand.Float64() * -1.5, -1, -rand.Float64() * -1.5})
		}
		w.AddEntity(newEntities[s.EntityType.EncodeEntity()](cube.PosFromVec3(spawnPos)))
	}

	s.Delay = rand.Intn(s.MaxSpawnDelay-s.MinSpawnDelay) + s.MinSpawnDelay
	w.SetBlock(pos, s, nil)
}

func (s Spawner) EncodeItem() (name string, meta int16) {
	return "minecraft:mob_spawner", 0
}

func (s Spawner) EncodeBlock() (string, map[string]any) {
	return "minecraft:mob_spawner", nil
}

func (s Spawner) Model() world.BlockModel {
	return model.Solid{}
}
