package main

import "fmt"

type SceneObject struct {
	ObjectID    uint64 `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

var SceneObjects = createSampleObjects()

func createSampleObject(ary []SceneObject, index int, name string) {
	ary[index] = SceneObject{
		ObjectID:    NextId(),
		Name:        name,
		Image:       fmt.Sprintf("http://placehold.it/140&text=%v", name),
		Description: "Lorem ipsum dolor sit amet, placerat laoreet commodo nec id. Ut ultrices accumsan odio magna risus, duis dolorem turpis, faucibus convallis fringilla orci sit erat, lacus vel venenatis soluta duis convallis eleifend, nullam repellendus natoque phasellus interdum. Dui vel dolor vitae nunc posuere mauris, ut lacus dolor, malesuada orci bibendum incididunt mi, aenean cras id, arcu facilisis tellus quis facilisis. Vel aliquet adipiscing, a massa suscipit aliquam id felis gravida. Pretium montes rutrum sociis magna ante, magna praesent integer morbi sem diam rhoncus, egestas inceptos ornare sapien libero fusce dui. Tristique qui et, non pede hendrerit massa, sed magna justo turpis quis nunc sit, erat vestibulum etiam sed aliquet.",
	}
}

func createSampleObjects() []SceneObject {
	out := make([]SceneObject, 15)
	createSampleObject(out, 0, "Glasses")
	createSampleObject(out, 1, "Shirt")
	createSampleObject(out, 2, "Picture")
	createSampleObject(out, 3, "Beer")
	createSampleObject(out, 4, "Dress")
	createSampleObject(out, 5, "Dog")
	createSampleObject(out, 6, "Pants")
	createSampleObject(out, 7, "Slacks")
	createSampleObject(out, 8, "Tights")
	createSampleObject(out, 9, "Earing")
	createSampleObject(out, 10, "Necklace")
	createSampleObject(out, 11, "Watch")
	createSampleObject(out, 12, "Top")
	createSampleObject(out, 13, "Shoes")
	createSampleObject(out, 14, "Tie")
	return out
}
