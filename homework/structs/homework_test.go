package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		copy(person.name[:], name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.manaHealth[0] = byte(mana & 0b1111_1111)
		person.manaHealth[1] |= byte(mana >> 8)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.manaHealth[2] = byte(health & 0b1111_1111)
		person.manaHealth[1] |= byte(health>>8) << 4
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.strengthExperienceRespectLevel |= uint16(respect) << 12
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.strengthExperienceRespectLevel |= uint16(strength) << 8
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.strengthExperienceRespectLevel |= uint16(experience) << 4
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.strengthExperienceRespectLevel |= uint16(level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.paramsPersonType |= PersonHasHouse << 4
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.paramsPersonType |= PersonHasGun << 4
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.paramsPersonType |= PersonHasFamilty << 4
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.paramsPersonType += uint8(personType)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

const (
	PersonHasHouse   = 1
	PersonHasGun     = 2
	PersonHasFamilty = 4
)

type GamePerson struct {
	name                           [42]byte
	strengthExperienceRespectLevel uint16 // на каждое значение по 4 бита
	gold                           uint32
	x                              int32
	y                              int32
	z                              int32
	manaHealth                     [3]byte // первые 12 бит для маны, 4 бита 2-го байта для двух старших битов здоровья
	paramsPersonType               uint8   // первые 4 бита для типа, вторые 4 бита для параметров
}

func NewGamePerson(options ...Option) GamePerson {
	gamePerson := GamePerson{}

	for _, option := range options {
		option(&gamePerson)
	}

	return gamePerson
}

func (p *GamePerson) Name() string {
	return string(p.name[:])
}

func (p *GamePerson) X() int {
	return int(p.x)
}

func (p *GamePerson) Y() int {
	return int(p.y)
}

func (p *GamePerson) Z() int {
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int(p.manaHealth[0]) | int(p.manaHealth[1]&0b1111)<<8
}

func (p *GamePerson) Health() int {
	return int(p.manaHealth[2]) | int(p.manaHealth[1]>>4)<<8
}

func (p *GamePerson) Respect() int {
	return int(p.strengthExperienceRespectLevel >> 4 & 0b1111)
}

func (p *GamePerson) Strength() int {
	return int(p.strengthExperienceRespectLevel >> 12 & 0b1111)
}

func (p *GamePerson) Experience() int {
	return int(p.strengthExperienceRespectLevel >> 8 & 0b1111)
}

func (p *GamePerson) Level() int {
	return int(p.strengthExperienceRespectLevel & 0b1111)
}

func (p *GamePerson) HasHouse() bool {
	if p.paramsPersonType>>4&PersonHasHouse != 0 {
		return true
	}
	return false
}

func (p *GamePerson) HasGun() bool {
	if p.paramsPersonType>>4&PersonHasGun != 0 {
		return true
	}
	return false
}

func (p *GamePerson) HasFamilty() bool {
	if (p.paramsPersonType>>4)&PersonHasFamilty != 0 {
		return true
	}
	return false
}

func (p *GamePerson) Type() int {
	return int(p.paramsPersonType & 0b111)
}

func TestGamePerson(t *testing.T) {
	assert.LessOrEqual(t, unsafe.Sizeof(GamePerson{}), uintptr(64))

	const x, y, z = math.MinInt32, math.MaxInt32, 0
	const name = "aaaaaaaaaaaaa_bbbbbbbbbbbbb_cccccccccccccc"
	const personType = BuilderGamePersonType
	const gold = math.MaxInt32
	const mana = 1000
	const health = 1000
	const respect = 10
	const strength = 10
	const experience = 10
	const level = 10

	options := []Option{
		WithName(name),
		WithCoordinates(x, y, z),
		WithGold(gold),
		WithMana(mana),
		WithHealth(health),
		WithRespect(respect),
		WithStrength(strength),
		WithExperience(experience),
		WithLevel(level),
		WithHouse(),
		WithFamily(),
		WithType(personType),
	}

	person := NewGamePerson(options...)
	assert.Equal(t, name, person.Name())
	assert.Equal(t, x, person.X())
	assert.Equal(t, y, person.Y())
	assert.Equal(t, z, person.Z())
	assert.Equal(t, gold, person.Gold())
	assert.Equal(t, mana, person.Mana())
	assert.Equal(t, health, person.Health())
	assert.Equal(t, respect, person.Respect())
	assert.Equal(t, strength, person.Strength())
	assert.Equal(t, experience, person.Experience())
	assert.Equal(t, level, person.Level())
	assert.True(t, person.HasHouse())
	assert.True(t, person.HasFamilty())
	assert.False(t, person.HasGun())
	assert.Equal(t, personType, person.Type())
}
