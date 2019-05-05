package data

// A place we can store packages
type Backend interface {
	Push(name string, source string) error
	Pull(name string, destination string) error
	ConfigKey() string
}

// A simple kv store to load configuration
type ConfigStore interface {
	Put(key string, value interface{}) error
	Get(key string, valueOut interface{}) bool
}
