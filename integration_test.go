package form3

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	testAccID string = uuid.New().String()
	testOrgID string = uuid.New().String()
)

func createTestAccount(t *testing.T) *AccountData {
	account := &AccountData{
		Type:           "accounts",
		ID:             testAccID,
		OrganisationID: testOrgID,
		Version:        0,
	}
	account.Attributes = &AccountAttributes{
		Country:      "FR",
		BaseCurrency: "EUR",
		BankID:       "AF14",
		BankIDCode:   "GBDSC",
	}
	return account
}

func TestCreateAccountService(t *testing.T) {
	service := LoadAccountsService(nil)
	account := createTestAccount(t)

	exptd := &AccountData{
		Type:           "accounts",
		ID:             testAccID,
		OrganisationID: testOrgID,
	}

	exptd.Attributes = &AccountAttributes{
		Country:      "FR",
		BaseCurrency: "EUR",
		BankID:       "AF14",
		BankIDCode:   "GBDSC",
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
	assert.Equal(t, exptd.ID, result.ID)
	assert.Equal(t, exptd.OrganisationID, result.OrganisationID)
	service.DeleteByID(ctx, account.ID, account.Version)
}

func TestGetAccountByID(t *testing.T) {
	service := LoadAccountsService(nil)
	account := createTestAccount(t)

	ctx := context.Background()
	result, resp, err := service.NewAccount(ctx, account)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, resp)

	resData, resResp, err := service.GetByID(ctx, account.ID)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resResp.StatusCode)
	assert.NotNil(t, resData)
	assert.NotNil(t, resResp)
	service.DeleteByID(ctx, account.ID, account.Version)
}

func TestDeleteByID(t *testing.T) {
	service := LoadAccountsService(nil)
	account := createTestAccount(t)

	ctx := context.Background()
	result, resp, err := service.NewAccount(ctx, account)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, resp)

	resResp, err := service.DeleteByID(ctx, account.ID, account.Version)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, resResp.StatusCode)
	assert.NotNil(t, resResp)

	_, resGetResp, err := service.GetByID(ctx, account.ID)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, resGetResp.StatusCode)
}
