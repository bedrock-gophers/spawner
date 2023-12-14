package spawner

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/model"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"math/rand"
	"time"
)

// Spawner represents a spawner block that spawns entities at a certain rate.
type Spawner struct {
	e   func(mgl64.Vec3, *world.World) world.Entity
	pos mgl64.Vec3
	w   *world.World

	maxEntities int

	rate time.Duration
	c    chan struct{}
}

// New returns a new spawner for the entity passed.
func New(e func(mgl64.Vec3, *world.World) world.Entity, pos mgl64.Vec3, w *world.World, rate time.Duration, maxEntities int) *Spawner {
	s := &Spawner{
		e:   e,
		pos: pos,
		w:   w,

		maxEntities: maxEntities,
		rate:        rate,
	}
	go s.tick()
	return s
}

// EncodeBlock ...
func (s *Spawner) EncodeBlock() (string, map[string]any) {
	return "minecraft:mob_spawner", nil
}

var h = rand.Uint64()

// Hash ...
func (*Spawner) Hash() uint64 {
	return h
}

// Model ...
func (*Spawner) Model() world.BlockModel {
	return model.Solid{}
}

func (s *Spawner) tick() {
	t := time.NewTicker(s.rate)
	for {
		select {
		case <-t.C:
			s.spawn()
		case <-s.c:
			t.Stop()
			return
		}
	}
}

func (s *Spawner) spawn() {
	p1, p2 := s.pos.Add(mgl64.Vec3{-16, -16, -16}), s.pos.Add(mgl64.Vec3{16, 16, 16})
	x0, y0, z0, x1, y1, z1 := p1.X(), p1.Y(), p1.Z(), p2.X(), p2.Y(), p2.Z()

	if len(s.w.EntitiesWithin(cube.Box(x0, y0, z0, x1, y1, z1), func(entity world.Entity) bool {
		_, ok := entity.(*player.Player)
		return !ok
	})) <= 0 {
		return
	}

	b := s.w.Block(cube.PosFromVec3(s.pos))
	if b != s {
		s.c <- struct{}{}
		return
	}

	var pos mgl64.Vec3

	if rand.Float64() > 0.5 {
		pos = s.pos.Add(mgl64.Vec3{rand.Float64() * 1.5, 1, rand.Float64() * 1.5})
	} else {
		pos = s.pos.Sub(mgl64.Vec3{-rand.Float64() * -1.5, -1, -rand.Float64() * -1.5})
	}

	e := s.e(pos, s.w)

	if len(s.w.EntitiesWithin(cube.Box(x0, y0, z0, x1, y1, z1), func(entity world.Entity) bool {
		return entity.Type() != e.Type()
	})) >= s.maxEntities {
		return
	}

	s.w.AddEntity(e)
}
