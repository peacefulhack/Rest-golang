package mock

import (
	"github.com/stretchr/testify/mock"
	"time"
)

type RedisMock struct {
	mock.Mock
}

func (m *RedisMock) Set(key string, val interface{}, duration time.Duration) error {
	call := m.Called(key, val, duration)
	if call.Error(0) != nil {
		return call.Error(0)
	}
	return nil
}

func (m *RedisMock) Get(key string) (string, bool, error) {
	call := m.Called(key)
	return call.String(0), call.Bool(1), call.Error(2)
}

func (m *RedisMock) Del(key string) error {
	call := m.Called(key)
	if call.Error(0) != nil {
		return call.Error(0)
	}
	return nil
}
