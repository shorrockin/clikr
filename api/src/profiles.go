package main

type UserProfile struct {
	ID    string
	Likes []*SceneObject
}

func (u UserProfile) AlreadyLikes(o *SceneObject) bool {
	for _, l := range u.Likes {
		if l == o {
			return true
		}
	}
	return false
}

var UserProfiles = make(map[string]*UserProfile)

func GetUserProfile(user string) *UserProfile {
	profile, ok := UserProfiles[user]

	if !ok {
		profile = &UserProfile{user, make([]*SceneObject, 0)}
		UserProfiles[user] = profile
	}

	return profile
}

func LikeObject(user string, object *SceneObject) {
	profile := GetUserProfile(user)

	if !profile.AlreadyLikes(object) {
		profile.Likes = append(profile.Likes, object)
	}

	// perhaps we should restrict them from strengthening a bond more than
	// once? helps with testing though.
	for _, liked := range profile.Likes {
		if liked != object {
			StrengthenBond(liked, object)
		}
	}

}
