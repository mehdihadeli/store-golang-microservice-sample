// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	context "context"

	orders_service "github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/internal/orders/contracts/proto/service_clients"
	mock "github.com/stretchr/testify/mock"
)

// OrdersServiceServer is an autogenerated mock type for the OrdersServiceServer type
type OrdersServiceServer struct {
	mock.Mock
}

type OrdersServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *OrdersServiceServer) EXPECT() *OrdersServiceServer_Expecter {
	return &OrdersServiceServer_Expecter{mock: &_m.Mock}
}

// CreateOrder provides a mock function with given fields: _a0, _a1
func (_m *OrdersServiceServer) CreateOrder(_a0 context.Context, _a1 *orders_service.CreateOrderReq) (*orders_service.CreateOrderRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *orders_service.CreateOrderRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.CreateOrderReq) (*orders_service.CreateOrderRes, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.CreateOrderReq) *orders_service.CreateOrderRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orders_service.CreateOrderRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *orders_service.CreateOrderReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrdersServiceServer_CreateOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrder'
type OrdersServiceServer_CreateOrder_Call struct {
	*mock.Call
}

// CreateOrder is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *orders_service.CreateOrderReq
func (_e *OrdersServiceServer_Expecter) CreateOrder(_a0 interface{}, _a1 interface{}) *OrdersServiceServer_CreateOrder_Call {
	return &OrdersServiceServer_CreateOrder_Call{Call: _e.mock.On("CreateOrder", _a0, _a1)}
}

func (_c *OrdersServiceServer_CreateOrder_Call) Run(run func(_a0 context.Context, _a1 *orders_service.CreateOrderReq)) *OrdersServiceServer_CreateOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*orders_service.CreateOrderReq))
	})
	return _c
}

func (_c *OrdersServiceServer_CreateOrder_Call) Return(_a0 *orders_service.CreateOrderRes, _a1 error) *OrdersServiceServer_CreateOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrdersServiceServer_CreateOrder_Call) RunAndReturn(run func(context.Context, *orders_service.CreateOrderReq) (*orders_service.CreateOrderRes, error)) *OrdersServiceServer_CreateOrder_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrderByID provides a mock function with given fields: _a0, _a1
func (_m *OrdersServiceServer) GetOrderByID(_a0 context.Context, _a1 *orders_service.GetOrderByIDReq) (*orders_service.GetOrderByIDRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *orders_service.GetOrderByIDRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.GetOrderByIDReq) (*orders_service.GetOrderByIDRes, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.GetOrderByIDReq) *orders_service.GetOrderByIDRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orders_service.GetOrderByIDRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *orders_service.GetOrderByIDReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrdersServiceServer_GetOrderByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrderByID'
type OrdersServiceServer_GetOrderByID_Call struct {
	*mock.Call
}

// GetOrderByID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *orders_service.GetOrderByIDReq
func (_e *OrdersServiceServer_Expecter) GetOrderByID(_a0 interface{}, _a1 interface{}) *OrdersServiceServer_GetOrderByID_Call {
	return &OrdersServiceServer_GetOrderByID_Call{Call: _e.mock.On("GetOrderByID", _a0, _a1)}
}

func (_c *OrdersServiceServer_GetOrderByID_Call) Run(run func(_a0 context.Context, _a1 *orders_service.GetOrderByIDReq)) *OrdersServiceServer_GetOrderByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*orders_service.GetOrderByIDReq))
	})
	return _c
}

func (_c *OrdersServiceServer_GetOrderByID_Call) Return(_a0 *orders_service.GetOrderByIDRes, _a1 error) *OrdersServiceServer_GetOrderByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrdersServiceServer_GetOrderByID_Call) RunAndReturn(run func(context.Context, *orders_service.GetOrderByIDReq) (*orders_service.GetOrderByIDRes, error)) *OrdersServiceServer_GetOrderByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrders provides a mock function with given fields: _a0, _a1
func (_m *OrdersServiceServer) GetOrders(_a0 context.Context, _a1 *orders_service.GetOrdersReq) (*orders_service.GetOrdersRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *orders_service.GetOrdersRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.GetOrdersReq) (*orders_service.GetOrdersRes, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.GetOrdersReq) *orders_service.GetOrdersRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orders_service.GetOrdersRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *orders_service.GetOrdersReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrdersServiceServer_GetOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrders'
type OrdersServiceServer_GetOrders_Call struct {
	*mock.Call
}

// GetOrders is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *orders_service.GetOrdersReq
func (_e *OrdersServiceServer_Expecter) GetOrders(_a0 interface{}, _a1 interface{}) *OrdersServiceServer_GetOrders_Call {
	return &OrdersServiceServer_GetOrders_Call{Call: _e.mock.On("GetOrders", _a0, _a1)}
}

func (_c *OrdersServiceServer_GetOrders_Call) Run(run func(_a0 context.Context, _a1 *orders_service.GetOrdersReq)) *OrdersServiceServer_GetOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*orders_service.GetOrdersReq))
	})
	return _c
}

func (_c *OrdersServiceServer_GetOrders_Call) Return(_a0 *orders_service.GetOrdersRes, _a1 error) *OrdersServiceServer_GetOrders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrdersServiceServer_GetOrders_Call) RunAndReturn(run func(context.Context, *orders_service.GetOrdersReq) (*orders_service.GetOrdersRes, error)) *OrdersServiceServer_GetOrders_Call {
	_c.Call.Return(run)
	return _c
}

// SubmitOrder provides a mock function with given fields: _a0, _a1
func (_m *OrdersServiceServer) SubmitOrder(_a0 context.Context, _a1 *orders_service.SubmitOrderReq) (*orders_service.SubmitOrderRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *orders_service.SubmitOrderRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.SubmitOrderReq) (*orders_service.SubmitOrderRes, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.SubmitOrderReq) *orders_service.SubmitOrderRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orders_service.SubmitOrderRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *orders_service.SubmitOrderReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrdersServiceServer_SubmitOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubmitOrder'
type OrdersServiceServer_SubmitOrder_Call struct {
	*mock.Call
}

// SubmitOrder is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *orders_service.SubmitOrderReq
func (_e *OrdersServiceServer_Expecter) SubmitOrder(_a0 interface{}, _a1 interface{}) *OrdersServiceServer_SubmitOrder_Call {
	return &OrdersServiceServer_SubmitOrder_Call{Call: _e.mock.On("SubmitOrder", _a0, _a1)}
}

func (_c *OrdersServiceServer_SubmitOrder_Call) Run(run func(_a0 context.Context, _a1 *orders_service.SubmitOrderReq)) *OrdersServiceServer_SubmitOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*orders_service.SubmitOrderReq))
	})
	return _c
}

func (_c *OrdersServiceServer_SubmitOrder_Call) Return(_a0 *orders_service.SubmitOrderRes, _a1 error) *OrdersServiceServer_SubmitOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrdersServiceServer_SubmitOrder_Call) RunAndReturn(run func(context.Context, *orders_service.SubmitOrderReq) (*orders_service.SubmitOrderRes, error)) *OrdersServiceServer_SubmitOrder_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateShoppingCart provides a mock function with given fields: _a0, _a1
func (_m *OrdersServiceServer) UpdateShoppingCart(_a0 context.Context, _a1 *orders_service.UpdateShoppingCartReq) (*orders_service.UpdateShoppingCartRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *orders_service.UpdateShoppingCartRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.UpdateShoppingCartReq) (*orders_service.UpdateShoppingCartRes, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *orders_service.UpdateShoppingCartReq) *orders_service.UpdateShoppingCartRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orders_service.UpdateShoppingCartRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *orders_service.UpdateShoppingCartReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrdersServiceServer_UpdateShoppingCart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateShoppingCart'
type OrdersServiceServer_UpdateShoppingCart_Call struct {
	*mock.Call
}

// UpdateShoppingCart is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *orders_service.UpdateShoppingCartReq
func (_e *OrdersServiceServer_Expecter) UpdateShoppingCart(_a0 interface{}, _a1 interface{}) *OrdersServiceServer_UpdateShoppingCart_Call {
	return &OrdersServiceServer_UpdateShoppingCart_Call{Call: _e.mock.On("UpdateShoppingCart", _a0, _a1)}
}

func (_c *OrdersServiceServer_UpdateShoppingCart_Call) Run(run func(_a0 context.Context, _a1 *orders_service.UpdateShoppingCartReq)) *OrdersServiceServer_UpdateShoppingCart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*orders_service.UpdateShoppingCartReq))
	})
	return _c
}

func (_c *OrdersServiceServer_UpdateShoppingCart_Call) Return(_a0 *orders_service.UpdateShoppingCartRes, _a1 error) *OrdersServiceServer_UpdateShoppingCart_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrdersServiceServer_UpdateShoppingCart_Call) RunAndReturn(run func(context.Context, *orders_service.UpdateShoppingCartReq) (*orders_service.UpdateShoppingCartRes, error)) *OrdersServiceServer_UpdateShoppingCart_Call {
	_c.Call.Return(run)
	return _c
}

// NewOrdersServiceServer creates a new instance of OrdersServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrdersServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrdersServiceServer {
	mock := &OrdersServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
