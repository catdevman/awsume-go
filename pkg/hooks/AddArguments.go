package hooks

type PreAddArgumentsHook interface {
	PreAddArguments()
}

type AddArgumentsHook interface {
	AddArguments()
}

type PostAddArgumentsHook interface {
	PostAddArguments()
}
