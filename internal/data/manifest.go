package data

type RemoteStorage interface {
	Pull(remotePath string, localPath string) error
	Push(localPath string, remotePath string) error
}

type ScriptEngine interface {
	Run(scriptPath string, flags map[string]string) (map[string]interface{}, error)
}

type TemplateEngine interface {
	Run(filePath string, data map[string]interface{}) error
}

type LocalStorage interface {
	Put(key, value string) error
	Get(key string) (string, error)
}

type ConfigKeys string

const (
	GitUsername   ConfigKeys = "GIT_USERNAME"
	GitPassword   ConfigKeys = "GIT_PASSWORD"
	DefaultScript ConfigKeys = "DEFAULT_SCRIPT"
)
