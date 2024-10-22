package usecases

import (
	"github.com/aSel1x/Gin_Template/adapters"
)

type Usecases struct {
	*UserUsecase
}

func NewUsecases(adapters adapters.Adapters, appSecret string) *Usecases {
	uc := NewUserUsecase(adapters, appSecret)
	return &Usecases{
		UserUsecase: uc,
	}
}
