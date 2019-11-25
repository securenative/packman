package business

import "github.com/securenative/packman/internal/data"

type configService struct {
	localStorage data.LocalStorage
}

func NewConfigService(localStorage data.LocalStorage) *configService {
	return &configService{localStorage: localStorage}
}

func (this *configService) SetAuth(username string, password string) error {
	if uErr := this.localStorage.Put(string(data.GitUsername), username); uErr != nil {
		return uErr
	}

	if pErr := this.localStorage.Put(string(data.GitPassword), password); pErr != nil {
		return pErr
	}

	return nil
}

func (this *configService) SetDefaultEngine(command string) error {
	if err := this.localStorage.Put(string(data.DefaultScript), command); err != nil {
		return err
	}
	return nil
}
