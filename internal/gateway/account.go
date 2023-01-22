package gateway

import "github.com.br/thiago-batista-da-silva-oliveira/fc-ms-wallet/internal/entity"

type AccountGatewat interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
