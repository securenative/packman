package internal

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

// Takes a template and a data structure
// will expand the template based on the provided data
type TemplateEngine interface {
	Render(templateText string, data interface{}) (string, error)
}

// Will run a script file with the provided arguments
type ScriptEngine interface {
	Run(scriptFile string, args []string) error
}
