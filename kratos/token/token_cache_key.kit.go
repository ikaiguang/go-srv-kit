package tokenutil

// AssembleCacheKey ...
func AssembleCacheKey(prefix, identifier string) string {
	return DefaultCachePrefix + prefix + identifier
}

// AssembleCacheKeyForDefault ...
func AssembleCacheKeyForDefault(identifier string) string {
	return AssembleCacheKey(KeyPrefixDefault, identifier)
}

// AssembleCacheKeyForApp ...
func AssembleCacheKeyForApp(identifier string) string {
	return AssembleCacheKey(KeyPrefixApp, identifier)
}

// AssembleCacheKeyForService ...
func AssembleCacheKeyForService(identifier string) string {
	return AssembleCacheKey(KeyPrefixService, identifier)
}

// AssembleCacheKeyForAdmin ...
func AssembleCacheKeyForAdmin(identifier string) string {
	return AssembleCacheKey(KeyPrefixAdmin, identifier)
}

// AssembleCacheKeyForApi ...
func AssembleCacheKeyForApi(identifier string) string {
	return AssembleCacheKey(KeyPrefixApi, identifier)
}

// AssembleCacheKeyForWeb ...
func AssembleCacheKeyForWeb(identifier string) string {
	return AssembleCacheKey(KeyPrefixWeb, identifier)
}
