package main

import (
	"math"
	"sort"

	"github.com/shorrockin/clikr/log"
)

type Recommendations struct {
	Recommendations []SceneObject `json:"recommendation"`
}

type Bond struct {
	With     *SceneObject
	Strength uint
}

// creates a type that we can use to sort an array of bonds
type ByStrength []*Bond

func (s ByStrength) Len() int {
	return len(s)
}
func (s ByStrength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByStrength) Less(i, j int) bool {
	return s[i].Strength > s[j].Strength
}

var Bonds map[*SceneObject][]*Bond = make(map[*SceneObject][]*Bond)

// retrieves a bond between two objects, creating it if it
// doesn't exist.
func GetBond(from *SceneObject, to *SceneObject) *Bond {
	bonds, ok := Bonds[from]

	// if no bonds are created yet then initialize that array
	if !ok {
		bonds = make([]*Bond, 0)
		Bonds[from] = bonds
	}

	// locate the bond within the bond array and strength in
	for _, bond := range bonds {
		if bond.With == to {
			return bond
		}
	}

	// bond not in array, create it and return it
	newBond := &Bond{With: to, Strength: 0}
	bonds = append(bonds, newBond)
	Bonds[from] = bonds

	return newBond
}

// strengthens the bond between two objects and returns the value
// of that strength.
func StrengthenBond(from *SceneObject, to *SceneObject) uint {
	bond1 := GetBond(from, to)
	bond2 := GetBond(to, from)

	bond1.Strength = bond1.Strength + 1
	bond2.Strength = bond2.Strength + 1

	if bond1.Strength != bond2.Strength {
		panic("bonds became out of sync, strengths expected to be equal but were not")
	}

	// keep our bond arrays in nice sorted order
	sort.Sort(ByStrength(Bonds[from]))
	sort.Sort(ByStrength(Bonds[to]))

	return bond1.Strength
}

func RegisterObjectClick(variables map[string]string) interface{} {
	objectId := variables["object_id"]
	userId := variables["X-Clikr-User"]

	sceneObject, ok := SceneObjects[objectId]
	if !ok {
		log.Warn("attempt to register object click with invalid object id: '%v'", objectId)
		return []*SceneObject{}
	}

	log.Debug("user '%v' register click on object '%v:%v'", userId, sceneObject.Name, objectId)

	// mark the object as liked for this user, will strength bond between other
	// liked objects.
	LikeObject(userId, sceneObject)

	bonds, ok := Bonds[sceneObject]
	if !ok {
		log.Debug("no bonds available for object: '%v'", objectId)
		return []*SceneObject{}
	}

	maxSize := 3 // maximum number of recommendations to return
	size := int(math.Min(float64(len(bonds)), float64(maxSize)))
	out := make([]*SceneObject, size, size)

	for i, b := range bonds {
		log.Debug("recommending '%v' with bond strength '%v'", b.With.Name, b.Strength)
		out[i] = b.With
		if i == (maxSize - 1) {
			break
		}
	}

	return out
}
