package application_test

import (
	"github/criandogames/go-hexagonal/application"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestApplicationProduct_Enabled(t *testing.T) {
	// Arrange
	product := application.Product{
		Name:  "Product 1",
		Price: 10,
	}
	// Act
	err := product.Enable()
	// Assert
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())

}

func TestApplicationProduct_Disabled(t *testing.T) {
	// Arrange
	product := application.Product{
		Name:  "Product 1",
		Price: 0,
	}
	// Act
	err := product.Disable()
	// Assert
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())

}

func TestApplicationProduct_IsValid(t *testing.T) {
	// Arrange
	product := application.Product{
		Name:  "Product 1",
		Price: 10,
	}

	// Act
	_, err := product.IsValid()

	// Assert
	require.Nil(t, err)

	// Act
	product.Status = "invalida"
	_, err = product.IsValid()

	// Assert
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	// Act
	product.Status = application.ENABLED
	_, err = product.IsValid()

	// Assert

	require.Nil(t, err)

	// Act
	product.Price = -10
	_, err = product.IsValid()

	// Assert
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestApplicationProduct_Gets(t *testing.T) {
	// Arrange
	product := application.Product{
		Id:     uuid.NewV4().String(),
		Name:   "Product 1",
		Price:  10,
		Status: application.DISABLED,
	}

	// Act
	id := product.GetId()
	name := product.GetName()
	price := product.GetPrice()
	status := product.GetStatus()

	// Assert
	require.Equal(t, product.Id, id)
	require.Equal(t, product.Name, name)
	require.Equal(t, product.Price, price)
	require.Equal(t, product.Status, status)

}
