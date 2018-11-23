package store

// IStore is the interface implemented by the specific variables stores
type IStore interface {
	Set(environment, variable, value string) error
	Get(environment, variable string) (string, error)
}
