package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/STAMBOULI-ABDELKARIM/car_repair_shop/util"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func createRandomCustomer(t *testing.T) Customer {

	arg := CreateCustomerParams{
		FullName:    util.RandomName(),
		PhoneNumber: util.RandomPhone(),
	}

	customer, err := testQueries.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)

	require.Equal(t, arg.FullName, customer.FullName)
	require.Equal(t, arg.PhoneNumber, customer.PhoneNumber)

	require.NotZero(t, customer.ID)
	require.NotZero(t, customer.CreatedAt)

	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestGetCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	customer2, err := testQueries.GetCustomer(context.Background(), customer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, customer1)

	require.Equal(t, customer1.ID, customer2.ID)
	require.Equal(t, customer1.FullName, customer2.FullName)
	require.Equal(t, customer1.PhoneNumber, customer2.PhoneNumber)
	require.WithinDuration(t, customer1.CreatedAt, customer2.CreatedAt, time.Second)
}

func TestUpdateCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)

	arg := UpdateCustomerParams{
		ID:          customer1.ID,
		FullName:    util.RandomName(),
		PhoneNumber: util.RandomPhone(),
	}

	customer2, err := testQueries.UpdateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.ID, customer2.ID)
	require.Equal(t, arg.FullName, customer2.FullName)
	require.Equal(t, arg.PhoneNumber, customer2.PhoneNumber)
	require.WithinDuration(t, customer1.CreatedAt, customer2.CreatedAt, time.Second)
}

func TestDeleteCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	err := testQueries.DeleteCustomer(context.Background(), customer1.ID)
	require.NoError(t, err)

	customer2, err := testQueries.GetCustomer(context.Background(), customer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, customer2)
}

func TestListCustomers(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomCustomer(t)
	}

	arg := ListCustomersParams{
		Limit:  5,
		Offset: 0,
	}

	customers, err := testQueries.ListCustomers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customers)

	for _, customer := range customers {
		require.NotEmpty(t, customer)
	}
}
