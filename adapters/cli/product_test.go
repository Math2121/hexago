package cli_test

import (
	"fmt"
	"testing"

	"github.com/Math2121/hexago/adapters/cli"
	mock_application "github.com/Math2121/hexago/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T){
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	productName := "Product Teste"
	productPrice := 25
	productStatus := "enabled"
	productId := "abs"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(float64(productPrice)).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Create(productName, float64(productPrice)).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID #%s with the name #%s", productId, productName)

	result, err := cli.Run(service, "create", "", productName, float64(productPrice))
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID #%s with the name #%s and status %s has been enabled", productId, productName, productStatus)

	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID #%s with the name #%s and status %s has been disabled", productId, productName, productStatus)

	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID #%s with the name #%s and status %s", productId, productName, productStatus)
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}