package main

import (
	"github.com/bedrock-gophers/spawner/spawner"
	_ "github.com/bedrock-gophers/spawner/spawner"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sirupsen/logrus"
)

func init() {
	spawner.RegisterEntityType(player.Type{}, func(pos cube.Pos) world.Entity {
		return player.New("test", skin.Skin{}, pos.Vec3())
	})
}

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

func accept(p *player.Player) {
}
