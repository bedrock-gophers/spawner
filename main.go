package main

import (
	"github.com/bedrock-gophers/living/living"
	"github.com/bedrock-gophers/spawner/spawner"
	_ "github.com/bedrock-gophers/spawner/spawner"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sirupsen/logrus"
	"math/rand"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	conf, err := server.DefaultConfig().Config(log)
	if err != nil {
		log.Fatalln(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	srv.Listen()

	for srv.Accept(accept) {

	}
}

// Define a custom entity type for Enderman.
type entityTypeEnderman struct{}

// EncodeEntity ...
func (entityTypeEnderman) EncodeEntity() string {
	return "minecraft:enderman"
}

// BBox ...
func (entityTypeEnderman) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.3, 0, -0.3, 0.3, 2.9, 0.3)
}

func init() {
	spawner.RegisterEntityType(entityTypeEnderman{}, func(pos cube.Pos, w *world.World) world.Entity {
		enderman := living.NewLivingEntity(entityTypeEnderman{}, 40, 0.3, []item.Stack{item.NewStack(item.EnderPearl{}, rand.Intn(2)+1)}, &entity.MovementComputer{
			Gravity:           0.08,
			Drag:              0.02,
			DragBeforeGravity: true,
		}, pos.Vec3(), w)

		return enderman
	})
}

func accept(p *player.Player) {
	p.Inventory().AddItem(item.NewStack(spawner.SpawnEgg{Kind: entityTypeEnderman{}}, 1))
	p.Inventory().AddItem(item.NewStack(spawner.Spawner{}, 1))
	p.Inventory().AddItem(item.NewStack(item.Pickaxe{Tier: item.ToolTierGold}, 1))
	p.SetGameMode(world.GameModeSurvival)
}
