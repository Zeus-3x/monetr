package platypus

import (
	"context"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jarcoal/httpmock"
	"github.com/monetr/monetr/server/config"
	"github.com/monetr/monetr/server/internal/fixtures"
	"github.com/monetr/monetr/server/internal/mock_plaid"
	"github.com/monetr/monetr/server/internal/testutils"
	"github.com/monetr/monetr/server/models"
	"github.com/monetr/monetr/server/repository"
	"github.com/monetr/monetr/server/secrets"
	"github.com/plaid/plaid-go/v30/plaid"
	"github.com/stretchr/testify/assert"
)

func TestPlaidClient_GetAccount(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		clock := clock.New()

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		log := testutils.GetLog(t)
		kms := secrets.NewPlaintextKMS()
		db := testutils.GetPgDatabase(t)
		user, _ := fixtures.GivenIHaveABasicAccount(t, clock)
		plaidLink := fixtures.GivenIHaveAPlaidLink(t, clock, user)

		secret, err := repository.NewSecretsRepository(
			log,
			clock,
			db,
			kms,
			plaidLink.AccountId,
		).Read(context.Background(), plaidLink.PlaidLink.SecretId)
		assert.NoError(t, err, "must be able to read the secret")

		account := mock_plaid.BankAccountFixture(t)

		mock_plaid.MockGetAccounts(t, []plaid.AccountBase{
			account,
		})

		client := NewPlaid(log, clock, secrets.NewPlaintextKMS(), nil, config.Plaid{
			ClientID:     gofakeit.UUID(),
			ClientSecret: gofakeit.UUID(),
			Environment:  plaid.Sandbox,
		})

		link := &models.Link{
			LinkId:    "link_foo",
			AccountId: user.AccountId,
		}

		platypus, err := client.NewClient(
			context.Background(),
			link,
			secret.Value,
			gofakeit.UUID(),
		)
		assert.NoError(t, err, "should create platypus")
		assert.NotNil(t, platypus, "should not be nil")

		accounts, err := platypus.GetAccounts(context.Background(), account.GetAccountId())
		assert.NoError(t, err, "should not return an error retrieving accounts")
		assert.NotEmpty(t, accounts, "should return some accounts")
	})
}

func TestPlaidClient_GetAllTransactions(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		clock := clock.New()

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		log := testutils.GetLog(t)
		kms := secrets.NewPlaintextKMS()
		db := testutils.GetPgDatabase(t)
		user, _ := fixtures.GivenIHaveABasicAccount(t, clock)
		plaidLink := fixtures.GivenIHaveAPlaidLink(t, clock, user)

		secret, err := repository.NewSecretsRepository(
			log,
			clock,
			db,
			kms,
			plaidLink.AccountId,
		).Read(context.Background(), plaidLink.PlaidLink.SecretId)
		assert.NoError(t, err, "must be able to read the secret")

		account := mock_plaid.BankAccountFixture(t)

		end := time.Now()
		start := end.Add(-7 * 24 * time.Hour)
		mock_plaid.MockGetRandomTransactions(t, start, end, 5000, []string{
			account.GetAccountId(),
		})

		platypus := NewPlaid(log, clock, kms, db, config.Plaid{
			ClientID:     gofakeit.UUID(),
			ClientSecret: gofakeit.UUID(),
			Environment:  plaid.Sandbox,
		})

		link := &models.Link{
			LinkId:    "link_foo",
			AccountId: user.AccountId,
		}

		client, err := platypus.NewClient(
			context.Background(),
			link,
			secret.Value,
			gofakeit.UUID(),
		)
		assert.NoError(t, err, "should create platypus")
		assert.NotNil(t, client, "should not be nil")

		transactions, err := client.GetAllTransactions(
			context.Background(),
			start,
			end,
			[]string{
				account.GetAccountId(),
			},
		)
		assert.NoError(t, err, "should not return an error")
		assert.NotEmpty(t, transactions, "should return a few transactions")
		assert.Equal(t, map[string]int{
			"POST https://sandbox.plaid.com/transactions/get": 11,
		}, httpmock.GetCallCountInfo(), "API calls should match")
	})
}

func TestPlaidClient_UpdateItem(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		clock := clock.New()

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		log := testutils.GetLog(t)
		kms := secrets.NewPlaintextKMS()
		db := testutils.GetPgDatabase(t)
		user, _ := fixtures.GivenIHaveABasicAccount(t, clock)
		plaidLink := fixtures.GivenIHaveAPlaidLink(t, clock, user)

		secret, err := repository.NewSecretsRepository(
			log,
			clock,
			db,
			kms,
			plaidLink.AccountId,
		).Read(context.Background(), plaidLink.PlaidLink.SecretId)
		assert.NoError(t, err, "must be able to read the secret")

		mock_plaid.MockCreateLinkToken(t)

		platypus := NewPlaid(log, clock, kms, db, config.Plaid{
			ClientID:     gofakeit.UUID(),
			ClientSecret: gofakeit.UUID(),
			Environment:  plaid.Sandbox,
			OAuthDomain:  "localhost",
		})

		link := &models.Link{
			LinkId:    "link_foo",
			AccountId: user.AccountId,
		}

		client, err := platypus.NewClient(
			context.Background(),
			link,
			secret.Value,
			gofakeit.UUID(),
		)
		assert.NoError(t, err, "should create client")
		assert.NotNil(t, client, "should not be nil")

		linkToken, err := client.UpdateItem(context.Background(), false)
		assert.NoError(t, err, "should not return an error creating an update link token")
		assert.NotEmpty(t, linkToken.Token(), "must not be empty")
	})
}
