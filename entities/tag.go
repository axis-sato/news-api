package entities

import "fmt"

type Tag struct {
	ID   uint
	Name string
}

func (t Tag) String() string {
	return fmt.Sprintf("{id: %v, name: %v}", t.ID, t.Name)
}

func NewTag(id uint, name string) *Tag {
	return &Tag{
		ID:   id,
		Name: name,
	}
}
