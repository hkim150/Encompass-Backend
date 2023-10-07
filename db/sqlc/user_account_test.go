package db

import (
	"context"
	"database/sql"
	"encompass/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUserAccount(t *testing.T) (UserAccount, func(t *testing.T)) {
	arg := CreateUserAccountParams{
		Username: util.RandomUsername(),
		Email:    util.RandomEmail(),
	}

	user_account, err := testQueries.CreateUserAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user_account)

	require.Equal(t, user_account.Username, arg.Username)
	require.Equal(t, user_account.Email, arg.Email)

	require.NotZero(t, user_account.UserAccountID)
	require.NotZero(t, user_account.CreateTime)

	return user_account, func(t *testing.T) {
		err = testQueries.DeleteUserAccount(context.Background(), user_account.UserAccountID)
		require.NoError(t, err)

		user_account, err = testQueries.GetUserAccount(context.Background(), user_account.UserAccountID)
		require.Error(t, err)
		require.EqualError(t, err, sql.ErrNoRows.Error())
		require.Empty(t, user_account)
	}
}

func TestGetUserAccount(t *testing.T) {
	user_account, deleteRandomUserAccount := createRandomUserAccount(t)
	defer deleteRandomUserAccount(t)

	get_account, err := testQueries.GetUserAccount(context.Background(), user_account.UserAccountID)
	require.NoError(t, err)
	require.NotEmpty(t, get_account)

	require.EqualValues(t, user_account, get_account)
}

func TestListUserAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		_, deleteRandomUserAccount := createRandomUserAccount(t)
		defer deleteRandomUserAccount(t)
	}

	arg := ListUserAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListUserAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
}

func TestUpdateUserAccount(t *testing.T) {
	user_account, deleteRandomUserAccount := createRandomUserAccount(t)
	defer deleteRandomUserAccount(t)

	new_username := util.RandomUsername()
	new_email := util.RandomEmail()
	arg := UpdateUserAccountParams{
		UserAccountID: user_account.UserAccountID,
		Username:      new_username,
		Email:         new_email,
	}

	updated_account, err := testQueries.UpdateUserAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updated_account)

	require.Equal(t, user_account.UserAccountID, updated_account.UserAccountID)
	require.Equal(t, user_account.CreateTime, updated_account.CreateTime)

	require.Equal(t, new_username, updated_account.Username)
	require.Equal(t, new_email, updated_account.Email)
}
