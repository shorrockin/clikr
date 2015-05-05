package main

import "fmt"

type SceneObject struct {
	ObjectID    string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

var SceneObjects = createSampleObjects()

func createSampleObject(mp map[string]*SceneObject, name string) {
	id := fmt.Sprint(NextId())

	mp[id] = &SceneObject{
		ObjectID:    id,
		Name:        "name",
		Image:       fmt.Sprintf("http://placehold.it/140&text=%v", name),
		Description: "Lorem ipsum dolor sit amet, placerat laoreet commodo nec id. Ut ultrices accumsan odio magna risus, duis dolorem turpis, faucibus convallis fringilla orci sit erat, lacus vel venenatis soluta duis convallis eleifend, nullam repellendus natoque phasellus interdum. Dui vel dolor vitae nunc posuere mauris, ut lacus dolor, malesuada orci bibendum incididunt mi, aenean cras id, arcu facilisis tellus quis facilisis. Vel aliquet adipiscing, a massa suscipit aliquam id felis gravida.",
	}
}

func createSampleObjects() map[string]*SceneObject {
	out := make(map[string]*SceneObject)
	createSampleObject(out, "Glasses")
	createSampleObject(out, "Shirt")
	createSampleObject(out, "Picture")
	createSampleObject(out, "Beer")
	createSampleObject(out, "Dress")
	createSampleObject(out, "Dog")
	createSampleObject(out, "Pants")
	createSampleObject(out, "Slacks")
	createSampleObject(out, "Tights")
	createSampleObject(out, "Earing")
	createSampleObject(out, "Necklace")
	createSampleObject(out, "Watch")
	createSampleObject(out, "Top")
	createSampleObject(out, "Shoes")
	createSampleObject(out, "Tie")
	return out
}
