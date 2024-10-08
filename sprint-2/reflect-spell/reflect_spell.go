package reflect_spell

import (
	"fmt"
	"reflect"
)

type Spell interface {
	// название заклинания
	Name() string
	// характеристика, на которую воздействует
	Char() string
	// количественное значение
	Value() int
}

// CastReceiver — если объект удовлетворяет этом интерфейсу, то заклинание применяется через него
type CastReceiver interface {
	ReceiveSpell(s Spell)
}

func CastToAll(spell Spell, objects []interface{}) {
	for _, obj := range objects {
		CastTo(spell, obj)
	}
}

func CastTo(spell Spell, object interface{}) {

	if receiver, ok := object.(CastReceiver); ok {
		receiver.ReceiveSpell(spell)
		fmt.Println("object receiver")
		return
	}

	val := reflect.ValueOf(object)

	if val.Kind() != reflect.Ptr {
		fmt.Println("object not ptr")
	}

	val = val.Elem()

	if val.Kind() != reflect.Struct {
		fmt.Println("object not struct")
		return
	}

	field := val.FieldByName(spell.Char())
	if !field.IsValid() {
		fmt.Println("not valid")
		return
	}

	if !field.CanSet() {
		fmt.Println("cant set field")
		return
	}
	switch field.Kind() {

	case reflect.Int:
		field.SetInt(field.Int() + int64(spell.Value()))
	default:
		fmt.Println("unknown field kind")
		return
	}

}

type spell struct {
	name string
	char string
	val  int
}

func newSpell(name string, char string, val int) Spell {
	return &spell{name: name, char: char, val: val}
}

func (s spell) Name() string {
	return s.name
}

func (s spell) Char() string {
	return s.char
}

func (s spell) Value() int {
	return s.val
}

type Player struct {
	// nolint: unused
	name   string
	health int
}

func (p *Player) ReceiveSpell(s Spell) {
	if s.Char() == "Health" {
		p.health += s.Value()
	}
}

type Zombie struct {
	Health int
}

type Daemon struct {
	Health int
}

type Orc struct {
	Health int
}

type Wall struct {
	Durability int
}
