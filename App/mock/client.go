package mock

import (
	"github.com/stretchr/testify/mock"
	"restArchitecture/mikail/App/models"
)

type ClientMock struct {
	mock.Mock
}

func (m *ClientMock) InsertProductList(request *models.ProductList) error {
	call := m.Called(request)
	if call.Error(0) != nil {
		return call.Error(0)
	}
	return nil
}
func (m *ClientMock) SortDateProductList(order string) ([]*models.SortProductResponse, error) {
	call := m.Called(order)
	if call.Error(1) != nil {
		return nil, call.Error(1)
	}
	return call.Get(0).([]*models.SortProductResponse), nil
}
func (m *ClientMock) SortPriceProductList(order string) ([]*models.SortProductResponse, error) {
	call := m.Called(order)
	if call.Error(1) != nil {
		return nil, call.Error(1)
	}
	return call.Get(0).([]*models.SortProductResponse), nil
}
func (m *ClientMock) SortNameProductList(order string) ([]*models.SortProductResponse, error) {
	call := m.Called(order)
	if call.Error(1) != nil {
		return nil, call.Error(1)
	}
	return call.Get(0).([]*models.SortProductResponse), nil
}
