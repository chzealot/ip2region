package ip2region

import (
	"fmt"
)

type Location struct {
	Region   string
	Province string
	City     string
	Area     string
	Name     string
}

func NewLocation() *Location {
	return &Location{}
}

func NewLocationFromCnFull(name string) *Location {
	return &Location{Name: name}
}

func (loc *Location) String() string {
	return fmt.Sprintf("Location(Region=%s, Province=%s, City=%s, Area=%s, Name=%s)",
		loc.Region, loc.Province, loc.City, loc.Area, loc.Name)
}
