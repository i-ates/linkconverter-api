// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import echo "github.com/labstack/echo/v4"
import mock "github.com/stretchr/testify/mock"

// DeepToUrlRouterInterface is an autogenerated mock type for the DeepToUrlRouterInterface type
type DeepToUrlRouterInterface struct {
	mock.Mock
}

// DeepToUrl provides a mock function with given fields: context
func (_m *DeepToUrlRouterInterface) DeepToUrl(context echo.Context) error {
	ret := _m.Called(context)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(context)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
