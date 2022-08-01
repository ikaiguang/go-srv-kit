package tokenutil

import "strconv"

// AssembleCacheID ...
func AssembleCacheID(prefix string, id uint64) string {
	return DefaultCachePrefix + prefix + strconv.FormatUint(id, 10)
}

// AssembleCacheIDForDefault ...
func AssembleCacheIDForDefault(id uint64) string {
	return AssembleCacheID(KeyPrefixDefault, id)
}

// AssembleCacheIDForApp ...
func AssembleCacheIDForApp(id uint64) string {
	return AssembleCacheID(KeyPrefixApp, id)
}

// AssembleCacheIDForService ...
func AssembleCacheIDForService(id uint64) string {
	return AssembleCacheID(KeyPrefixService, id)
}

// AssembleCacheIDForAdmin ...
func AssembleCacheIDForAdmin(id uint64) string {
	return AssembleCacheID(KeyPrefixAdmin, id)
}

// AssembleCacheIDForApi ...
func AssembleCacheIDForApi(id uint64) string {
	return AssembleCacheID(KeyPrefixApi, id)
}

// AssembleCacheIDForWeb ...
func AssembleCacheIDForWeb(id uint64) string {
	return AssembleCacheID(KeyPrefixWeb, id)
}
