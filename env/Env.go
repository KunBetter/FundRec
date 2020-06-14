package env

type Env int32

const (
	Test Env = 0
	Prod Env = 1
)

func GetCurEnv() Env {
	return Test
}
