package spawner

import (
	"github.com/bedrock-gophers/living/living"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type handler struct {
	living.NopHandler
	e *living.Living

	stack int

	s *Spawner
}

func (h *handler) HandleHurt(ctx *event.Context, damage float64, src world.DamageSource) {
	w := h.e.World()
	pos := h.e.Position()

	if h.e.Health()-damage <= 0 && h.s.stacked {
		if h.stack <= 1 {
			return
		}

		ctx.Cancel()
		h.e.Heal(h.e.MaxHealth(), effect.InstantHealingSource{})
		h.stack--
		h.e.SetNameTag(text.Colourf("<yellow>%dx</yellow>", h.stack))
		h.e.TriggerLastAttack()

		for _, drop := range h.e.Drops() {
			w.AddEntity(entity.NewItem(drop, pos))
		}

		if s, ok := src.(entity.AttackDamageSource); ok {
			h.e.KnockBack(s.Attacker.Position(), 0.4, 0.4)

			for _, v := range w.Viewers(pos) {
				v.ViewEntityAction(h.e, entity.HurtAction{})
			}
		}
	}
}
