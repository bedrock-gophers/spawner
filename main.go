package main

import (
	"github.com/bedrock-gophers/living/living"
	"github.com/bedrock-gophers/spawner/spawner"
	_ "github.com/bedrock-gophers/spawner/spawner"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"log"
	"log/slog"
)

func main() {
	chat.Global.Subscribe(chat.StdoutSubscriber{})

	conf, err := server.DefaultConfig().Config(slog.Default())
	if err != nil {
		log.Fatalln(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	srv.Listen()

	for p := range srv.Accept() {
		p.Inventory().AddItem(item.NewStack(spawner.SpawnEgg{Kind: entityTypeEnderman{}}, 1))
		p.Inventory().AddItem(item.NewStack(spawner.Spawner{}, 1))
		p.Inventory().AddItem(item.NewStack(item.Pickaxe{Tier: item.ToolTierGold}, 1))
		p.SetGameMode(world.GameModeSurvival)
	}
}

// Define a custom entity type for Enderman.
type entityTypeEnderman struct {
	living.NopLivingType
}

// EncodeEntity ...
func (entityTypeEnderman) EncodeEntity() string {
	return "minecraft:enderman"
}

// BBox ...
func (entityTypeEnderman) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.3, 0, -0.3, 0.3, 2.9, 0.3)
}

func init() {
	spawner.RegisterEntityType(entityTypeEnderman{}, func(pos cube.Pos, tx *world.Tx) *world.EntityHandle {
		opts := world.EntitySpawnOpts{
			Position: pos.Vec3(),
		}

		conf := living.Config{
			EntityType: entityTypeEnderman{},
			Drops: []living.Drop{
				living.NewDrop(item.EnderPearl{}, 0, 2),
			},
		}
		return opts.New(conf.EntityType, conf)
	})
}
