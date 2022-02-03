package hooks

import "github.com/catdevman/awsume-go/shared"

type PreCollectProfilesHook interface {
	PreCollectProfiles()
}

type CollectProfilesHook interface {
	CollectProfiles() shared.Profiles
}

type PostCollectProfilesHook interface {
	PostCollectProfiles()
}
