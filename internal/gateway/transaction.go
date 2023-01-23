package gateway

import "github.com.br/thiago-batista-da-silva-oliveira/fc-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
