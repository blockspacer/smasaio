package firing

import (
	"github.com/20zinnm/smasaio/ecs"
	"github.com/20zinnm/smasaio/component"
)

type System struct {
	entities []entity
}

func (s *System) Update(dt float64) {
	for index, entity := range s.entities {

	}
}

func (s *System) Remove(id ecs.EntityID) {
	panic("implement me")
}

type entity struct {
	ecs.EntityID
	*component.Input
	*component.Cannon
}
