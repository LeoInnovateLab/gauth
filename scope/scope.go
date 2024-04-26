package scope

type AuthScope interface {
	// GetScope Retrieve scopes corresponding to the actual code scopes used by various platforms
	GetScope() string
	// IsDefault Determine whether the current scope is enabled by default on various platforms.
	IsDefault() bool
	GetDefaultScopes() []string
}
