package create_account

import (
	"fc-eda/internal/entity"
	"fc-eda/internal/gateway"
)

type CreateAccountInputDTO struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: a,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) error {
	account, err := uc.AccountGateway.FindByID(input.ID)
	if err != nil {
		accountNew := entity.NewAccount(input.ID, input.Balance)
		errSave := uc.AccountGateway.Save(accountNew)
		if errSave != nil {
			return err
		}
	} else {
		account.UpdateBalance(input.Balance)
		err = uc.AccountGateway.UpdateBalance(account)
		if err != nil {
			return err
		}
	}
	return nil
}
