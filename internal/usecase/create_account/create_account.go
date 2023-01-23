package createaccount

import (
	"github.com.br/thiago-batista-da-silva-oliveira/fc-ms-wallet/internal/entity"
	"github.com.br/thiago-batista-da-silva-oliveira/fc-ms-wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGatewat  gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: a,
		ClientGatewat:  c,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := uc.ClientGatewat.Get(input.ClientID)
	if err != nil {
		return nil, err
	}
	account := entity.NewAccount(client)
	accountErr := uc.AccountGateway.Save(account)
	if accountErr != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
