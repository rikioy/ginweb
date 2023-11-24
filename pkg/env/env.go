package env

import "os"

const (
	ENV_KEY    = "ginweb_env"
	ENV_DEV    = "dev"
	ENV_TEST   = "test"
	ENV_GRAY   = "gray"
	ENV_ONLINE = ""
)

var (
	env *Env
)

type Env struct {
	env string
}

func init() {
	env = &Env{env: os.Getenv(ENV_KEY)}
}

func GetEnvString() string {
	return env.env
}

func SetEnvKey(key string) {
	env = &Env{env: os.Getenv(key)}
}

func IsDev() bool {
	return env.env == ENV_DEV
}

func IsTest() bool {
	return env.env == ENV_TEST
}

func IsGray() bool {
	return env.env == ENV_GRAY
}

func IsOnline() bool {
	return env.env == ENV_ONLINE
}

func IsDevTest() bool {
	if env.env == ENV_DEV || env.env == ENV_TEST {
		return true
	}
	return false
}
