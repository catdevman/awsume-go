package hooks

type PreCollectProfilesHook interface {
	PreCollectProfiles()
}

type CollectProfilesHook interface {
	CollectProfiles()
}

type PostCollectProfilesHook interface {
	PostCollectProfiles()
}
