package form3

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createTestAccount(t *testing.T) *AccountData {
	account := &AccountData{
		ID:             uuid.New().String(),
		OrganisationID: uuid.New().String(),
		Attributes: &AccountAttributes{
			Country:      "FR",
			BaseCurrency: "EUR",
			BankID:       "789789789",
			BankIDCode:   "FR789",
			Name:         []string{"Jon Dee"},
		},
	}
	return account
}

func TestCreateAccountService(t *testing.T) {
	service := CreateAccountsService(nil)
	account := createTestAccount(t)

	exptd := &AccountData{
		ID:             account.ID,
		OrganisationID: account.OrganisationID,
		Attributes: &AccountAttributes{
			Country:      "FR",
			BaseCurrency: "EUR",
			BankID:       "789789789",
			BankIDCode:   "FR789",
			Name:         []string{"Jon Dee"},
		},
	}

	ctx := context.Background()
	result, resp, err := service.NewAccount(ctx, account)
	assert.Nil(t, err)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, exptd, result)
}
