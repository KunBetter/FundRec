package core

const (
	BaseFetchSpecTime = "0 0 8 * * ?"
)

type Fetch interface {
	Init()
	Process()
}
