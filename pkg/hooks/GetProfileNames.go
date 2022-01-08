package hooks

type PreGetProfileNamesHook interface {
	PreGetProfileNames()
}

type GetProfileNamesHooks interface {
	GetProfileNames()
}

type PostGetProfileNamesHook interface {
	PostGetProfileNames()
}
