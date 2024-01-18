package process_transaction

import (
	"github.com/Arturbarth/payment-gateway/domain/entity"
	mock_repository "github.com/Arturbarth/payment-gateway/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	rejected = "rejected"
	aproved  = "aproved"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionDTOInput{
		ID:                        "123",
		AccountID:                 "123",
		CreditCardNumber:          "40000000000000000",
		CreditCardName:            "Jose da Silva",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    100,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "123",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID,
			input.Amount, rejected,
			expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
