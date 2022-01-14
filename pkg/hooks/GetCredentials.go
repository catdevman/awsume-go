package hooks

type PreGetCredentialsHook interface {
	PreGetCredentials()
}

type GetCredentialsHook interface {
	GetCredentials()
}

type PostGetCredentialsHook interface {
	PostGetCredentials()
}

type PreGetCredentialsWithSAMLHook interface {
	PreGetCredentialsWithSAML()
}

type GetCredentialsWithSAMLHook interface {
	GetCredentialsWithSAMLHook()
}

type PostGetCredentialsWithSAMLHook interface {
	PostGetCredentialsWithSAML()
}

type PreGetCredentialsWithWebIdentityHook interface {
	PreGetCredentialsWithWebIdentity()
}

type GetCredentialsWithWebIdentityHook interface {
	GetCredentialsWithWebIdentity()
}

type PostGetCredentialsWithWebIdentityHook interface {
	PostGetCredentialsWithWebIdentity()
}
