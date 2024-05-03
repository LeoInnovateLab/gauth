package facebook

type Scope struct {
	scope       string
	description string
	isDefault   bool
}

func (a *Scope) GetScope() string {
	return a.scope
}

func (a *Scope) IsDefault() bool {
	return a.isDefault
}

func (a *Scope) GetDefaultScopes() []string {
	scopes := []Scope{
		PublicProfile,
		Email,
		UserAgeRange,
		UserBirthday,
		UserFriends,
		UserGender,
		UserHometown,
		UserLikes,
		UserLink,
		UserLocation,
		UserPhotos,
		UserPosts,
		UserVideos,
		GroupsAccessMemberInfo,
		PublishToGroups,
	}

	var result []string
	for _, scope := range scopes {
		if scope.isDefault {
			result = append(result, scope.scope)
		}
	}

	return result
}

var (
	PublicProfile = Scope{
		scope:       "public_profile",
		description: "Provides access to the 'About Me' section of the profile",
		isDefault:   true,
	}
	Email = Scope{
		scope:       "email",
		description: "Provides access to the person's primary email address",
		isDefault:   false,
	}
	UserAgeRange = Scope{
		scope:       "user_age_range",
		description: "Provides access to the person's age range on Facebook",
		isDefault:   false,
	}
	UserBirthday = Scope{
		scope:       "user_birthday",
		description: "Provides access to the person's birthday",
		isDefault:   false,
	}
	UserFriends = Scope{
		scope:       "user_friends",
		description: "Provides access to the list of friends that also use your app",
		isDefault:   false,
	}
	UserGender = Scope{
		scope:       "user_gender",
		description: "Provides access to person's gender",
		isDefault:   false,
	}
	UserHometown = Scope{
		scope:       "user_hometown",
		description: "Provides access to the person's hometown",
		isDefault:   false,
	}
	UserLikes = Scope{
		scope:       "user_likes",
		description: "Provides access to the list of all of the pages the person has liked",
		isDefault:   false,
	}
	UserLink = Scope{
		scope:       "user_link",
		description: "Provides access to the person's personal website URL via the link field on the User object",
		isDefault:   false,
	}
	UserLocation = Scope{
		scope:       "user_location",
		description: "Provides access to the person's current city through the location field on the User object",
		isDefault:   false,
	}
	UserPhotos = Scope{
		scope:       "user_photos",
		description: "Provides access to the photos a person has uploaded or been tagged in",
		isDefault:   false,
	}
	UserPosts = Scope{
		scope:       "user_posts",
		description: "Provides access to the posts on a person's Timeline",
		isDefault:   false,
	}
	UserVideos = Scope{
		scope:       "user_videos",
		description: "Provides access to the videos a person has uploaded or been tagged in",
		isDefault:   false,
	}
	GroupsAccessMemberInfo = Scope{
		scope:       "groups_access_member_info",
		description: "Provides access to group member lists",
		isDefault:   false,
	}
	PublishToGroups = Scope{
		scope:       "publish_to_groups",
		description: "Enables your app to post content into a group on behalf of a user",
		isDefault:   false,
	}
)
