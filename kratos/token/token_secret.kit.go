package tokenutil

// GenSecret ...
func GenSecret(prefix, secret string) []byte {
	return []byte(prefix + secret)
}

// GenSecretForDefault ...
func GenSecretForDefault(secret string) []byte {
	return GenSecret(KeyPrefixDefault, secret)
}

// GenSecretForApp ...
func GenSecretForApp(secret string) []byte {
	return GenSecret(KeyPrefixApp, secret)
}

// GenSecretForService ...
func GenSecretForService(secret string) []byte {
	return GenSecret(KeyPrefixService, secret)
}

// GenSecretForAdmin ...
func GenSecretForAdmin(secret string) []byte {
	return GenSecret(KeyPrefixAdmin, secret)
}

// GenSecretForApi ...
func GenSecretForApi(secret string) []byte {
	return GenSecret(KeyPrefixApi, secret)
}

// GenSecretForWeb ...
func GenSecretForWeb(secret string) []byte {
	return GenSecret(KeyPrefixWeb, secret)
}
