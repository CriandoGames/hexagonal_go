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
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{
		ProductRepository: persistence,
	}
	// Act

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
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	// Act

	service := application.ProductService{ProductRepository: persistence}

	result, err := service.Create("Product 1", 10.0)
	// 	Assert
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestPruductService_EnableAndDisable(t *testing.T) {

	// Arrange

	ctrl := gomock.NewController(t)
	//defer prolonga a vida do ctrl até o final do teste
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)

	product.EXPECT().Enable().Return(nil).AnyTimes()
	product.EXPECT().Disable().Return(nil).AnyTimes()

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	// Act

	service := application.ProductService{ProductRepository: persistence}

	result, err := service.Enable(product)
	// 	Assert
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	// 	Assert
	require.Nil(t, err)
	require.Equal(t, product, result)

}
