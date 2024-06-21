package application_test

import (
	"github.com/codeedu/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "teste"
	product.Status = application.DISABLED
	product.Price = 100

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()
	require.Equal(t, "The price must be greater than 0", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "teste"
	product.Status = application.DISABLED
	product.Price = 100

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0

	err = product.Disable()
	require.Nil(t, err)
}

func TestIsValid_Product(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "teste"
	product.Status = application.DISABLED
	product.Price = 100

	isValid, err := product.IsValid()
	require.Nil(t, err)
	require.True(t, isValid)
}
