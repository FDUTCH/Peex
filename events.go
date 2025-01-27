// Package peex
// This file was generated using the event generator. Do not edit.
package peex

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"net"
	"time"
)

type eventId uint

const (
	eventMove eventId = iota
	eventJump
	eventTeleport
	eventChangeWorld
	eventToggleSprint
	eventToggleSneak
	eventChat
	eventFoodLoss
	eventHeal
	eventHurt
	eventDeath
	eventRespawn
	eventSkinChange
	eventFireExtinguish
	eventStartBreak
	eventBlockBreak
	eventBlockPlace
	eventBlockPick
	eventItemUse
	eventItemUseOnBlock
	eventItemUseOnEntity
	eventItemRelease
	eventItemConsume
	eventAttackEntity
	eventExperienceGain
	eventPunchAir
	eventSignEdit
	eventLecternPageTurn
	eventItemDamage
	eventItemPickup
	eventHeldSlotChange
	eventItemDrop
	eventTransfer
	eventCommandExecution
	eventQuit
	eventDiagnostics
)

// getHandlerEvents returns which events a handler implements. Since it is impossible to distinguish actually imlemented
// methods from ones embedded using player.NopHandler, it is recommended to not embed it at all. Most Peex handlers
// won't implement player.Handler!
func getHandlerEvents(h Handler) map[eventId]struct{} {
	m := make(map[eventId]struct{})

	if _, ok := h.(eventMoveHandler); ok {
		m[eventMove] = struct{}{}
	}
	if _, ok := h.(eventJumpHandler); ok {
		m[eventJump] = struct{}{}
	}
	if _, ok := h.(eventTeleportHandler); ok {
		m[eventTeleport] = struct{}{}
	}
	if _, ok := h.(eventChangeWorldHandler); ok {
		m[eventChangeWorld] = struct{}{}
	}
	if _, ok := h.(eventToggleSprintHandler); ok {
		m[eventToggleSprint] = struct{}{}
	}
	if _, ok := h.(eventToggleSneakHandler); ok {
		m[eventToggleSneak] = struct{}{}
	}
	if _, ok := h.(eventChatHandler); ok {
		m[eventChat] = struct{}{}
	}
	if _, ok := h.(eventFoodLossHandler); ok {
		m[eventFoodLoss] = struct{}{}
	}
	if _, ok := h.(eventHealHandler); ok {
		m[eventHeal] = struct{}{}
	}
	if _, ok := h.(eventHurtHandler); ok {
		m[eventHurt] = struct{}{}
	}
	if _, ok := h.(eventDeathHandler); ok {
		m[eventDeath] = struct{}{}
	}
	if _, ok := h.(eventRespawnHandler); ok {
		m[eventRespawn] = struct{}{}
	}
	if _, ok := h.(eventSkinChangeHandler); ok {
		m[eventSkinChange] = struct{}{}
	}
	if _, ok := h.(eventFireExtinguishHandler); ok {
		m[eventFireExtinguish] = struct{}{}
	}
	if _, ok := h.(eventStartBreakHandler); ok {
		m[eventStartBreak] = struct{}{}
	}
	if _, ok := h.(eventBlockBreakHandler); ok {
		m[eventBlockBreak] = struct{}{}
	}
	if _, ok := h.(eventBlockPlaceHandler); ok {
		m[eventBlockPlace] = struct{}{}
	}
	if _, ok := h.(eventBlockPickHandler); ok {
		m[eventBlockPick] = struct{}{}
	}
	if _, ok := h.(eventItemUseHandler); ok {
		m[eventItemUse] = struct{}{}
	}
	if _, ok := h.(eventItemUseOnBlockHandler); ok {
		m[eventItemUseOnBlock] = struct{}{}
	}
	if _, ok := h.(eventItemUseOnEntityHandler); ok {
		m[eventItemUseOnEntity] = struct{}{}
	}
	if _, ok := h.(eventItemReleaseHandler); ok {
		m[eventItemRelease] = struct{}{}
	}
	if _, ok := h.(eventItemConsumeHandler); ok {
		m[eventItemConsume] = struct{}{}
	}
	if _, ok := h.(eventAttackEntityHandler); ok {
		m[eventAttackEntity] = struct{}{}
	}
	if _, ok := h.(eventExperienceGainHandler); ok {
		m[eventExperienceGain] = struct{}{}
	}
	if _, ok := h.(eventPunchAirHandler); ok {
		m[eventPunchAir] = struct{}{}
	}
	if _, ok := h.(eventSignEditHandler); ok {
		m[eventSignEdit] = struct{}{}
	}
	if _, ok := h.(eventLecternPageTurnHandler); ok {
		m[eventLecternPageTurn] = struct{}{}
	}
	if _, ok := h.(eventItemDamageHandler); ok {
		m[eventItemDamage] = struct{}{}
	}
	if _, ok := h.(eventItemPickupHandler); ok {
		m[eventItemPickup] = struct{}{}
	}
	if _, ok := h.(eventHeldSlotChangeHandler); ok {
		m[eventHeldSlotChange] = struct{}{}
	}
	if _, ok := h.(eventItemDropHandler); ok {
		m[eventItemDrop] = struct{}{}
	}
	if _, ok := h.(eventTransferHandler); ok {
		m[eventTransfer] = struct{}{}
	}
	if _, ok := h.(eventCommandExecutionHandler); ok {
		m[eventCommandExecution] = struct{}{}
	}
	if _, ok := h.(eventQuitHandler); ok {
		m[eventQuit] = struct{}{}
	}
	if _, ok := h.(eventDiagnosticsHandler); ok {
		m[eventDiagnostics] = struct{}{}
	}
	return m
}

var allEvents = map[string]eventId{
	"eventMove":             eventMove,
	"eventJump":             eventJump,
	"eventTeleport":         eventTeleport,
	"eventChangeWorld":      eventChangeWorld,
	"eventToggleSprint":     eventToggleSprint,
	"eventToggleSneak":      eventToggleSneak,
	"eventChat":             eventChat,
	"eventFoodLoss":         eventFoodLoss,
	"eventHeal":             eventHeal,
	"eventHurt":             eventHurt,
	"eventDeath":            eventDeath,
	"eventRespawn":          eventRespawn,
	"eventSkinChange":       eventSkinChange,
	"eventFireExtinguish":   eventFireExtinguish,
	"eventStartBreak":       eventStartBreak,
	"eventBlockBreak":       eventBlockBreak,
	"eventBlockPlace":       eventBlockPlace,
	"eventBlockPick":        eventBlockPick,
	"eventItemUse":          eventItemUse,
	"eventItemUseOnBlock":   eventItemUseOnBlock,
	"eventItemUseOnEntity":  eventItemUseOnEntity,
	"eventItemRelease":      eventItemRelease,
	"eventItemConsume":      eventItemConsume,
	"eventAttackEntity":     eventAttackEntity,
	"eventExperienceGain":   eventExperienceGain,
	"eventPunchAir":         eventPunchAir,
	"eventSignEdit":         eventSignEdit,
	"eventLecternPageTurn":  eventLecternPageTurn,
	"eventItemDamage":       eventItemDamage,
	"eventItemPickup":       eventItemPickup,
	"eventHeldSlotChange":   eventHeldSlotChange,
	"eventItemDrop":         eventItemDrop,
	"eventTransfer":         eventTransfer,
	"eventCommandExecution": eventCommandExecution,
	"eventQuit":             eventQuit,
	"eventDiagnostics":      eventDiagnostics,
}

type eventMoveHandler interface {
	HandleMove(ctx *player.Context, newPos mgl64.Vec3, newRot cube.Rotation)
}

type eventJumpHandler interface {
	HandleJump(p *player.Player)
}

type eventTeleportHandler interface {
	HandleTeleport(ctx *player.Context, pos mgl64.Vec3)
}

type eventChangeWorldHandler interface {
	HandleChangeWorld(p *player.Player, before, after *world.World)
}

type eventToggleSprintHandler interface {
	HandleToggleSprint(ctx *player.Context, after bool)
}

type eventToggleSneakHandler interface {
	HandleToggleSneak(ctx *player.Context, after bool)
}

type eventChatHandler interface {
	HandleChat(ctx *player.Context, message *string)
}

type eventFoodLossHandler interface {
	HandleFoodLoss(ctx *player.Context, from int, to *int)
}

type eventHealHandler interface {
	HandleHeal(ctx *player.Context, health *float64, src world.HealingSource)
}

type eventHurtHandler interface {
	HandleHurt(ctx *player.Context, damage *float64, immune bool, attackImmunity *time.Duration, src world.DamageSource)
}

type eventDeathHandler interface {
	HandleDeath(p *player.Player, src world.DamageSource, keepInv *bool)
}

type eventRespawnHandler interface {
	HandleRespawn(p *player.Player, pos *mgl64.Vec3, w **world.World)
}

type eventSkinChangeHandler interface {
	HandleSkinChange(ctx *player.Context, skin *skin.Skin)
}

type eventFireExtinguishHandler interface {
	HandleFireExtinguish(ctx *player.Context, pos cube.Pos)
}

type eventStartBreakHandler interface {
	HandleStartBreak(ctx *player.Context, pos cube.Pos)
}

type eventBlockBreakHandler interface {
	HandleBlockBreak(ctx *player.Context, pos cube.Pos, drops *[]item.Stack, xp *int)
}

type eventBlockPlaceHandler interface {
	HandleBlockPlace(ctx *player.Context, pos cube.Pos, b world.Block)
}

type eventBlockPickHandler interface {
	HandleBlockPick(ctx *player.Context, pos cube.Pos, b world.Block)
}

type eventItemUseHandler interface {
	HandleItemUse(ctx *player.Context)
}

type eventItemUseOnBlockHandler interface {
	HandleItemUseOnBlock(ctx *player.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3)
}

type eventItemUseOnEntityHandler interface {
	HandleItemUseOnEntity(ctx *player.Context, e world.Entity)
}

type eventItemReleaseHandler interface {
	HandleItemRelease(ctx *player.Context, item item.Stack, dur time.Duration)
}

type eventItemConsumeHandler interface {
	HandleItemConsume(ctx *player.Context, item item.Stack)
}

type eventAttackEntityHandler interface {
	HandleAttackEntity(ctx *player.Context, e world.Entity, force, height *float64, critical *bool)
}

type eventExperienceGainHandler interface {
	HandleExperienceGain(ctx *player.Context, amount *int)
}

type eventPunchAirHandler interface {
	HandlePunchAir(ctx *player.Context)
}

type eventSignEditHandler interface {
	HandleSignEdit(ctx *player.Context, pos cube.Pos, frontSide bool, oldText, newText string)
}

type eventLecternPageTurnHandler interface {
	HandleLecternPageTurn(ctx *player.Context, pos cube.Pos, oldPage int, newPage *int)
}

type eventItemDamageHandler interface {
	HandleItemDamage(ctx *player.Context, i item.Stack, damage int)
}

type eventItemPickupHandler interface {
	HandleItemPickup(ctx *player.Context, i *item.Stack)
}

type eventHeldSlotChangeHandler interface {
	HandleHeldSlotChange(ctx *player.Context, from, to int)
}

type eventItemDropHandler interface {
	HandleItemDrop(ctx *player.Context, s item.Stack)
}

type eventTransferHandler interface {
	HandleTransfer(ctx *player.Context, addr *net.UDPAddr)
}

type eventCommandExecutionHandler interface {
	HandleCommandExecution(ctx *player.Context, command cmd.Command, args []string)
}

type eventQuitHandler interface {
	HandleQuit(p *player.Player)
}

type eventDiagnosticsHandler interface {
	HandleDiagnostics(p *player.Player, d session.Diagnostics)
}

func (h *Session) HandleMove(ctx *player.Context, newPos mgl64.Vec3, newRot cube.Rotation) {
	h.handleEvent(eventMove, func(h Handler) {
		h.(eventMoveHandler).HandleMove(ctx, newPos, newRot)
	})
}

func (h *Session) HandleJump(p *player.Player) {
	h.handleEvent(eventJump, func(h Handler) {
		h.(eventJumpHandler).HandleJump(p)
	})
}

func (h *Session) HandleTeleport(ctx *player.Context, pos mgl64.Vec3) {
	h.handleEvent(eventTeleport, func(h Handler) {
		h.(eventTeleportHandler).HandleTeleport(ctx, pos)
	})
}

func (h *Session) HandleChangeWorld(p *player.Player, before, after *world.World) {
	h.handleEvent(eventChangeWorld, func(h Handler) {
		h.(eventChangeWorldHandler).HandleChangeWorld(p, before, after)
	})
}

func (h *Session) HandleToggleSprint(ctx *player.Context, after bool) {
	h.handleEvent(eventToggleSprint, func(h Handler) {
		h.(eventToggleSprintHandler).HandleToggleSprint(ctx, after)
	})
}

func (h *Session) HandleToggleSneak(ctx *player.Context, after bool) {
	h.handleEvent(eventToggleSneak, func(h Handler) {
		h.(eventToggleSneakHandler).HandleToggleSneak(ctx, after)
	})
}

func (h *Session) HandleChat(ctx *player.Context, message *string) {
	h.handleEvent(eventChat, func(h Handler) {
		h.(eventChatHandler).HandleChat(ctx, message)
	})
}

func (h *Session) HandleFoodLoss(ctx *player.Context, from int, to *int) {
	h.handleEvent(eventFoodLoss, func(h Handler) {
		h.(eventFoodLossHandler).HandleFoodLoss(ctx, from, to)
	})
}

func (h *Session) HandleHeal(ctx *player.Context, health *float64, src world.HealingSource) {
	h.handleEvent(eventHeal, func(h Handler) {
		h.(eventHealHandler).HandleHeal(ctx, health, src)
	})
}

func (h *Session) HandleHurt(ctx *player.Context, damage *float64, immune bool, attackImmunity *time.Duration, src world.DamageSource) {
	h.handleEvent(eventHurt, func(h Handler) {
		h.(eventHurtHandler).HandleHurt(ctx, damage, immune, attackImmunity, src)
	})
}

func (h *Session) HandleDeath(p *player.Player, src world.DamageSource, keepInv *bool) {
	h.handleEvent(eventDeath, func(h Handler) {
		h.(eventDeathHandler).HandleDeath(p, src, keepInv)
	})
}

func (h *Session) HandleRespawn(p *player.Player, pos *mgl64.Vec3, w **world.World) {
	h.handleEvent(eventRespawn, func(h Handler) {
		h.(eventRespawnHandler).HandleRespawn(p, pos, w)
	})
}

func (h *Session) HandleSkinChange(ctx *player.Context, skin *skin.Skin) {
	h.handleEvent(eventSkinChange, func(h Handler) {
		h.(eventSkinChangeHandler).HandleSkinChange(ctx, skin)
	})
}

func (h *Session) HandleFireExtinguish(ctx *player.Context, pos cube.Pos) {
	h.handleEvent(eventFireExtinguish, func(h Handler) {
		h.(eventFireExtinguishHandler).HandleFireExtinguish(ctx, pos)
	})
}

func (h *Session) HandleStartBreak(ctx *player.Context, pos cube.Pos) {
	h.handleEvent(eventStartBreak, func(h Handler) {
		h.(eventStartBreakHandler).HandleStartBreak(ctx, pos)
	})
}

func (h *Session) HandleBlockBreak(ctx *player.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	h.handleEvent(eventBlockBreak, func(h Handler) {
		h.(eventBlockBreakHandler).HandleBlockBreak(ctx, pos, drops, xp)
	})
}

func (h *Session) HandleBlockPlace(ctx *player.Context, pos cube.Pos, b world.Block) {
	h.handleEvent(eventBlockPlace, func(h Handler) {
		h.(eventBlockPlaceHandler).HandleBlockPlace(ctx, pos, b)
	})
}

func (h *Session) HandleBlockPick(ctx *player.Context, pos cube.Pos, b world.Block) {
	h.handleEvent(eventBlockPick, func(h Handler) {
		h.(eventBlockPickHandler).HandleBlockPick(ctx, pos, b)
	})
}

func (h *Session) HandleItemUse(ctx *player.Context) {
	h.handleEvent(eventItemUse, func(h Handler) {
		h.(eventItemUseHandler).HandleItemUse(ctx)
	})
}

func (h *Session) HandleItemUseOnBlock(ctx *player.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	h.handleEvent(eventItemUseOnBlock, func(h Handler) {
		h.(eventItemUseOnBlockHandler).HandleItemUseOnBlock(ctx, pos, face, clickPos)
	})
}

func (h *Session) HandleItemUseOnEntity(ctx *player.Context, e world.Entity) {
	h.handleEvent(eventItemUseOnEntity, func(h Handler) {
		h.(eventItemUseOnEntityHandler).HandleItemUseOnEntity(ctx, e)
	})
}

func (h *Session) HandleItemRelease(ctx *player.Context, item item.Stack, dur time.Duration) {
	h.handleEvent(eventItemRelease, func(h Handler) {
		h.(eventItemReleaseHandler).HandleItemRelease(ctx, item, dur)
	})
}

func (h *Session) HandleItemConsume(ctx *player.Context, item item.Stack) {
	h.handleEvent(eventItemConsume, func(h Handler) {
		h.(eventItemConsumeHandler).HandleItemConsume(ctx, item)
	})
}

func (h *Session) HandleAttackEntity(ctx *player.Context, e world.Entity, force, height *float64, critical *bool) {
	h.handleEvent(eventAttackEntity, func(h Handler) {
		h.(eventAttackEntityHandler).HandleAttackEntity(ctx, e, force, height, critical)
	})
}

func (h *Session) HandleExperienceGain(ctx *player.Context, amount *int) {
	h.handleEvent(eventExperienceGain, func(h Handler) {
		h.(eventExperienceGainHandler).HandleExperienceGain(ctx, amount)
	})
}

func (h *Session) HandlePunchAir(ctx *player.Context) {
	h.handleEvent(eventPunchAir, func(h Handler) {
		h.(eventPunchAirHandler).HandlePunchAir(ctx)
	})
}

func (h *Session) HandleSignEdit(ctx *player.Context, pos cube.Pos, frontSide bool, oldText, newText string) {
	h.handleEvent(eventSignEdit, func(h Handler) {
		h.(eventSignEditHandler).HandleSignEdit(ctx, pos, frontSide, oldText, newText)
	})
}

func (h *Session) HandleLecternPageTurn(ctx *player.Context, pos cube.Pos, oldPage int, newPage *int) {
	h.handleEvent(eventLecternPageTurn, func(h Handler) {
		h.(eventLecternPageTurnHandler).HandleLecternPageTurn(ctx, pos, oldPage, newPage)
	})
}

func (h *Session) HandleItemDamage(ctx *player.Context, i item.Stack, damage int) {
	h.handleEvent(eventItemDamage, func(h Handler) {
		h.(eventItemDamageHandler).HandleItemDamage(ctx, i, damage)
	})
}

func (h *Session) HandleItemPickup(ctx *player.Context, i *item.Stack) {
	h.handleEvent(eventItemPickup, func(h Handler) {
		h.(eventItemPickupHandler).HandleItemPickup(ctx, i)
	})
}

func (h *Session) HandleHeldSlotChange(ctx *player.Context, from, to int) {
	h.handleEvent(eventHeldSlotChange, func(h Handler) {
		h.(eventHeldSlotChangeHandler).HandleHeldSlotChange(ctx, from, to)
	})
}

func (h *Session) HandleItemDrop(ctx *player.Context, s item.Stack) {
	h.handleEvent(eventItemDrop, func(h Handler) {
		h.(eventItemDropHandler).HandleItemDrop(ctx, s)
	})
}

func (h *Session) HandleTransfer(ctx *player.Context, addr *net.UDPAddr) {
	h.handleEvent(eventTransfer, func(h Handler) {
		h.(eventTransferHandler).HandleTransfer(ctx, addr)
	})
}

func (h *Session) HandleCommandExecution(ctx *player.Context, command cmd.Command, args []string) {
	h.handleEvent(eventCommandExecution, func(h Handler) {
		h.(eventCommandExecutionHandler).HandleCommandExecution(ctx, command, args)
	})
}

func (h *Session) HandleQuit(p *player.Player) {
	h.handleEvent(eventQuit, func(h Handler) {
		h.(eventQuitHandler).HandleQuit(p)
	})
	h.doQuit()
}

func (h *Session) HandleDiagnostics(p *player.Player, d session.Diagnostics) {
	h.handleEvent(eventDiagnostics, func(h Handler) {
		h.(eventDiagnosticsHandler).HandleDiagnostics(p, d)
	})
}
