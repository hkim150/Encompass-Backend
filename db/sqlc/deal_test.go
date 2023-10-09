package db

import (
	"context"
	"database/sql"
	"encompass/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomDeal(t *testing.T) (Deal, func(t *testing.T)) {
	account, deleteRandomUserAccount := createRandomUserAccount(t)
	arg := CreateDealParams{
		Author:       account.UserAccountID,
		StoreName:    util.RandomString(util.RandomNumber(5, 10)),
		Description:  util.RandomString(util.RandomNumber(10, 200)),
		RegularPrice: sql.NullString{String: util.RandomPrice(50, 100), Valid: true},
		SalePrice:    sql.NullString{String: util.RandomPrice(1, 50), Valid: true},
	}

	deal, err := testQueries.CreateDeal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, deal)

	require.Equal(t, arg.Author, deal.Author)
	require.Equal(t, arg.StoreName, deal.StoreName)
	require.Equal(t, arg.Description, deal.Description)
	require.Equal(t, arg.RegularPrice, deal.RegularPrice)
	require.Equal(t, arg.SalePrice, deal.SalePrice)

	require.NotZero(t, deal.DealID)
	require.NotZero(t, deal.CreateTime)

	return deal, func(t *testing.T) {
		defer deleteRandomUserAccount(t)
		err = testQueries.DeleteDeal(context.Background(), deal.DealID)
		require.NoError(t, err)

		deal, err = testQueries.GetDeal(context.Background(), deal.DealID)
		require.Error(t, err)
		require.EqualError(t, err, sql.ErrNoRows.Error())
		require.Empty(t, deal)
	}
}

func TestGetDeal(t *testing.T) {
	deal, deleteRandomDeal := createRandomDeal(t)
	defer deleteRandomDeal(t)

	get_deal, err := testQueries.GetDeal(context.Background(), deal.DealID)
	require.NoError(t, err)
	require.NotEmpty(t, get_deal)

	require.EqualValues(t, deal, get_deal)
}

func TestListDeal(t *testing.T) {
	for i := 0; i < 10; i++ {
		_, deleteRandomDeal := createRandomDeal(t)
		defer deleteRandomDeal(t)
	}

	arg := ListDealsParams{
		Limit:  5,
		Offset: 5,
	}

	deals, err := testQueries.ListDeals(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, deals, 5)
}

func TestUpdateDeal(t *testing.T) {
	deal, deleteRandomDeal := createRandomDeal(t)
	defer deleteRandomDeal(t)

	arg := UpdateDealParams{
		DealID:       deal.DealID,
		StoreName:    util.RandomString(util.RandomNumber(5, 10)),
		Description:  util.RandomString(util.RandomNumber(10, 200)),
		RegularPrice: sql.NullString{String: util.RandomPrice(50, 100), Valid: true},
		SalePrice:    sql.NullString{String: util.RandomPrice(1, 50), Valid: true},
	}

	updated_deal, err := testQueries.UpdateDeal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updated_deal)

	require.Equal(t, deal.DealID, updated_deal.DealID)
	require.Equal(t, deal.Upvote, updated_deal.Upvote)
	require.Equal(t, deal.CreateTime, updated_deal.CreateTime)

	require.Equal(t, arg.StoreName, updated_deal.StoreName)
	require.Equal(t, arg.Description, updated_deal.Description)
	require.Equal(t, arg.RegularPrice, updated_deal.RegularPrice)
	require.Equal(t, arg.SalePrice, updated_deal.SalePrice)
}
