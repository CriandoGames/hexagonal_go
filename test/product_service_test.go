package application_test

import (
	"github/criandogames/go-hexagonal/application"
	mock_application "github/criandogames/go-hexagonal/application/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestPruductService_Get(t *testing.T) {

	// Arrange

	ctrl := gomock.NewController(t)
	//defer prolonga a vida do ctrl até o final do teste
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	// Act
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{
		ProductRepository: persistence,
	}

	result, err := service.Get("123")
	// 	Assert
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestPruductService_Save(t *testing.T) {

	// Arrange

	ctrl := gomock.NewController(t)
	//defer prolonga a vida do ctrl até o final do teste
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	// Act
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{ProductRepository: persistence}

	result, err := service.Create("Product 1", 10.0)
	// 	Assert
	require.Nil(t, err)
	require.Equal(t, product, result)

}
