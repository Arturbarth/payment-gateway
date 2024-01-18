package process_transaction

import (
	"github.com/Arturbarth/payment-gateway/domain/entity"
	"github.com/Arturbarth/payment-gateway/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDTOInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	_, invalidCC := entity.NewCreditCard(input.CreditCardNumber,
		input.CreditCardName,
		input.CreditCardExpirationMonth,
		input.CreditCardExpirationYear,
		input.CreditCardCVV)
	if invalidCC != nil {
		err := p.Repository.Insert(transaction.ID,
			transaction.AccountID,
			transaction.Amount,
			entity.REJECTED,
			invalidCC.Error())
		if err != nil {
			return TransactionDtoOutput{}, err
		}
		output := TransactionDtoOutput{
			ID:           transaction.ID,
			Status:       entity.REJECTED,
			ErrorMessage: invalidCC.Error(),
		}
		return output, nil
	}
	return TransactionDtoOutput{}, nil
}
