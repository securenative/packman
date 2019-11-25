package business

type TemplatingService interface {
	Pack(remoteUrl string, packagePath string) error
	Unpack(remoteUtl string, packagePath string, flags map[string]string) error
	Render(templatePath string, packagePath string, flags map[string]string) error
}

type ConfigService interface {
	SetAuth(username string, password string) error
	SetDefaultEngine(command string) error
}
