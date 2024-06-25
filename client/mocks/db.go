// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/sourcenetwork/defradb/client"

	datastore "github.com/sourcenetwork/defradb/datastore"

	event "github.com/sourcenetwork/defradb/event"

	go_datastore "github.com/ipfs/go-datastore"

	immutable "github.com/sourcenetwork/immutable"

	mock "github.com/stretchr/testify/mock"

	model "github.com/lens-vm/lens/host-go/config/model"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

type DB_Expecter struct {
	mock *mock.Mock
}

func (_m *DB) EXPECT() *DB_Expecter {
	return &DB_Expecter{mock: &_m.Mock}
}

// AddPolicy provides a mock function with given fields: ctx, policy
func (_m *DB) AddPolicy(ctx context.Context, policy string) (client.AddPolicyResult, error) {
	ret := _m.Called(ctx, policy)

	if len(ret) == 0 {
		panic("no return value specified for AddPolicy")
	}

	var r0 client.AddPolicyResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (client.AddPolicyResult, error)); ok {
		return rf(ctx, policy)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) client.AddPolicyResult); ok {
		r0 = rf(ctx, policy)
	} else {
		r0 = ret.Get(0).(client.AddPolicyResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, policy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_AddPolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddPolicy'
type DB_AddPolicy_Call struct {
	*mock.Call
}

// AddPolicy is a helper method to define mock.On call
//   - ctx context.Context
//   - policy string
func (_e *DB_Expecter) AddPolicy(ctx interface{}, policy interface{}) *DB_AddPolicy_Call {
	return &DB_AddPolicy_Call{Call: _e.mock.On("AddPolicy", ctx, policy)}
}

func (_c *DB_AddPolicy_Call) Run(run func(ctx context.Context, policy string)) *DB_AddPolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_AddPolicy_Call) Return(_a0 client.AddPolicyResult, _a1 error) *DB_AddPolicy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_AddPolicy_Call) RunAndReturn(run func(context.Context, string) (client.AddPolicyResult, error)) *DB_AddPolicy_Call {
	_c.Call.Return(run)
	return _c
}

// AddSchema provides a mock function with given fields: _a0, _a1
func (_m *DB) AddSchema(_a0 context.Context, _a1 string) ([]client.CollectionDescription, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddSchema")
	}

	var r0 []client.CollectionDescription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]client.CollectionDescription, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []client.CollectionDescription); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.CollectionDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_AddSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSchema'
type DB_AddSchema_Call struct {
	*mock.Call
}

// AddSchema is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *DB_Expecter) AddSchema(_a0 interface{}, _a1 interface{}) *DB_AddSchema_Call {
	return &DB_AddSchema_Call{Call: _e.mock.On("AddSchema", _a0, _a1)}
}

func (_c *DB_AddSchema_Call) Run(run func(_a0 context.Context, _a1 string)) *DB_AddSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_AddSchema_Call) Return(_a0 []client.CollectionDescription, _a1 error) *DB_AddSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_AddSchema_Call) RunAndReturn(run func(context.Context, string) ([]client.CollectionDescription, error)) *DB_AddSchema_Call {
	_c.Call.Return(run)
	return _c
}

// AddView provides a mock function with given fields: ctx, gqlQuery, sdl, transform
func (_m *DB) AddView(ctx context.Context, gqlQuery string, sdl string, transform immutable.Option[model.Lens]) ([]client.CollectionDefinition, error) {
	ret := _m.Called(ctx, gqlQuery, sdl, transform)

	if len(ret) == 0 {
		panic("no return value specified for AddView")
	}

	var r0 []client.CollectionDefinition
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, immutable.Option[model.Lens]) ([]client.CollectionDefinition, error)); ok {
		return rf(ctx, gqlQuery, sdl, transform)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, immutable.Option[model.Lens]) []client.CollectionDefinition); ok {
		r0 = rf(ctx, gqlQuery, sdl, transform)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.CollectionDefinition)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, immutable.Option[model.Lens]) error); ok {
		r1 = rf(ctx, gqlQuery, sdl, transform)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_AddView_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddView'
type DB_AddView_Call struct {
	*mock.Call
}

// AddView is a helper method to define mock.On call
//   - ctx context.Context
//   - gqlQuery string
//   - sdl string
//   - transform immutable.Option[model.Lens]
func (_e *DB_Expecter) AddView(ctx interface{}, gqlQuery interface{}, sdl interface{}, transform interface{}) *DB_AddView_Call {
	return &DB_AddView_Call{Call: _e.mock.On("AddView", ctx, gqlQuery, sdl, transform)}
}

func (_c *DB_AddView_Call) Run(run func(ctx context.Context, gqlQuery string, sdl string, transform immutable.Option[model.Lens])) *DB_AddView_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(immutable.Option[model.Lens]))
	})
	return _c
}

func (_c *DB_AddView_Call) Return(_a0 []client.CollectionDefinition, _a1 error) *DB_AddView_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_AddView_Call) RunAndReturn(run func(context.Context, string, string, immutable.Option[model.Lens]) ([]client.CollectionDefinition, error)) *DB_AddView_Call {
	_c.Call.Return(run)
	return _c
}

// BasicExport provides a mock function with given fields: ctx, config
func (_m *DB) BasicExport(ctx context.Context, config *client.BackupConfig) error {
	ret := _m.Called(ctx, config)

	if len(ret) == 0 {
		panic("no return value specified for BasicExport")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *client.BackupConfig) error); ok {
		r0 = rf(ctx, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_BasicExport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BasicExport'
type DB_BasicExport_Call struct {
	*mock.Call
}

// BasicExport is a helper method to define mock.On call
//   - ctx context.Context
//   - config *client.BackupConfig
func (_e *DB_Expecter) BasicExport(ctx interface{}, config interface{}) *DB_BasicExport_Call {
	return &DB_BasicExport_Call{Call: _e.mock.On("BasicExport", ctx, config)}
}

func (_c *DB_BasicExport_Call) Run(run func(ctx context.Context, config *client.BackupConfig)) *DB_BasicExport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*client.BackupConfig))
	})
	return _c
}

func (_c *DB_BasicExport_Call) Return(_a0 error) *DB_BasicExport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_BasicExport_Call) RunAndReturn(run func(context.Context, *client.BackupConfig) error) *DB_BasicExport_Call {
	_c.Call.Return(run)
	return _c
}

// BasicImport provides a mock function with given fields: ctx, filepath
func (_m *DB) BasicImport(ctx context.Context, filepath string) error {
	ret := _m.Called(ctx, filepath)

	if len(ret) == 0 {
		panic("no return value specified for BasicImport")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, filepath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_BasicImport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BasicImport'
type DB_BasicImport_Call struct {
	*mock.Call
}

// BasicImport is a helper method to define mock.On call
//   - ctx context.Context
//   - filepath string
func (_e *DB_Expecter) BasicImport(ctx interface{}, filepath interface{}) *DB_BasicImport_Call {
	return &DB_BasicImport_Call{Call: _e.mock.On("BasicImport", ctx, filepath)}
}

func (_c *DB_BasicImport_Call) Run(run func(ctx context.Context, filepath string)) *DB_BasicImport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_BasicImport_Call) Return(_a0 error) *DB_BasicImport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_BasicImport_Call) RunAndReturn(run func(context.Context, string) error) *DB_BasicImport_Call {
	_c.Call.Return(run)
	return _c
}

// Blockstore provides a mock function with given fields:
func (_m *DB) Blockstore() datastore.DAGStore {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Blockstore")
	}

	var r0 datastore.DAGStore
	if rf, ok := ret.Get(0).(func() datastore.DAGStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.DAGStore)
		}
	}

	return r0
}

// DB_Blockstore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Blockstore'
type DB_Blockstore_Call struct {
	*mock.Call
}

// Blockstore is a helper method to define mock.On call
func (_e *DB_Expecter) Blockstore() *DB_Blockstore_Call {
	return &DB_Blockstore_Call{Call: _e.mock.On("Blockstore")}
}

func (_c *DB_Blockstore_Call) Run(run func()) *DB_Blockstore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_Blockstore_Call) Return(_a0 datastore.DAGStore) *DB_Blockstore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Blockstore_Call) RunAndReturn(run func() datastore.DAGStore) *DB_Blockstore_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *DB) Close() {
	_m.Called()
}

// DB_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type DB_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *DB_Expecter) Close() *DB_Close_Call {
	return &DB_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *DB_Close_Call) Run(run func()) *DB_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_Close_Call) Return() *DB_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *DB_Close_Call) RunAndReturn(run func()) *DB_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Events provides a mock function with given fields:
func (_m *DB) Events() *event.Bus {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Events")
	}

	var r0 *event.Bus
	if rf, ok := ret.Get(0).(func() *event.Bus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*event.Bus)
		}
	}

	return r0
}

// DB_Events_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Events'
type DB_Events_Call struct {
	*mock.Call
}

// Events is a helper method to define mock.On call
func (_e *DB_Expecter) Events() *DB_Events_Call {
	return &DB_Events_Call{Call: _e.mock.On("Events")}
}

func (_c *DB_Events_Call) Run(run func()) *DB_Events_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_Events_Call) Return(_a0 *event.Bus) *DB_Events_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Events_Call) RunAndReturn(run func() *event.Bus) *DB_Events_Call {
	_c.Call.Return(run)
	return _c
}

// ExecRequest provides a mock function with given fields: ctx, request
func (_m *DB) ExecRequest(ctx context.Context, request string) *client.RequestResult {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ExecRequest")
	}

	var r0 *client.RequestResult
	if rf, ok := ret.Get(0).(func(context.Context, string) *client.RequestResult); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.RequestResult)
		}
	}

	return r0
}

// DB_ExecRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecRequest'
type DB_ExecRequest_Call struct {
	*mock.Call
}

// ExecRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - request string
func (_e *DB_Expecter) ExecRequest(ctx interface{}, request interface{}) *DB_ExecRequest_Call {
	return &DB_ExecRequest_Call{Call: _e.mock.On("ExecRequest", ctx, request)}
}

func (_c *DB_ExecRequest_Call) Run(run func(ctx context.Context, request string)) *DB_ExecRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_ExecRequest_Call) Return(_a0 *client.RequestResult) *DB_ExecRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_ExecRequest_Call) RunAndReturn(run func(context.Context, string) *client.RequestResult) *DB_ExecRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllIndexes provides a mock function with given fields: _a0
func (_m *DB) GetAllIndexes(_a0 context.Context) (map[string][]client.IndexDescription, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetAllIndexes")
	}

	var r0 map[string][]client.IndexDescription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[string][]client.IndexDescription, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[string][]client.IndexDescription); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]client.IndexDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetAllIndexes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllIndexes'
type DB_GetAllIndexes_Call struct {
	*mock.Call
}

// GetAllIndexes is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *DB_Expecter) GetAllIndexes(_a0 interface{}) *DB_GetAllIndexes_Call {
	return &DB_GetAllIndexes_Call{Call: _e.mock.On("GetAllIndexes", _a0)}
}

func (_c *DB_GetAllIndexes_Call) Run(run func(_a0 context.Context)) *DB_GetAllIndexes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DB_GetAllIndexes_Call) Return(_a0 map[string][]client.IndexDescription, _a1 error) *DB_GetAllIndexes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetAllIndexes_Call) RunAndReturn(run func(context.Context) (map[string][]client.IndexDescription, error)) *DB_GetAllIndexes_Call {
	_c.Call.Return(run)
	return _c
}

// GetCollectionByName provides a mock function with given fields: _a0, _a1
func (_m *DB) GetCollectionByName(_a0 context.Context, _a1 string) (client.Collection, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetCollectionByName")
	}

	var r0 client.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (client.Collection, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) client.Collection); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetCollectionByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionByName'
type DB_GetCollectionByName_Call struct {
	*mock.Call
}

// GetCollectionByName is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *DB_Expecter) GetCollectionByName(_a0 interface{}, _a1 interface{}) *DB_GetCollectionByName_Call {
	return &DB_GetCollectionByName_Call{Call: _e.mock.On("GetCollectionByName", _a0, _a1)}
}

func (_c *DB_GetCollectionByName_Call) Run(run func(_a0 context.Context, _a1 string)) *DB_GetCollectionByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_GetCollectionByName_Call) Return(_a0 client.Collection, _a1 error) *DB_GetCollectionByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetCollectionByName_Call) RunAndReturn(run func(context.Context, string) (client.Collection, error)) *DB_GetCollectionByName_Call {
	_c.Call.Return(run)
	return _c
}

// GetCollections provides a mock function with given fields: _a0, _a1
func (_m *DB) GetCollections(_a0 context.Context, _a1 client.CollectionFetchOptions) ([]client.Collection, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetCollections")
	}

	var r0 []client.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, client.CollectionFetchOptions) ([]client.Collection, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, client.CollectionFetchOptions) []client.Collection); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, client.CollectionFetchOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetCollections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollections'
type DB_GetCollections_Call struct {
	*mock.Call
}

// GetCollections is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 client.CollectionFetchOptions
func (_e *DB_Expecter) GetCollections(_a0 interface{}, _a1 interface{}) *DB_GetCollections_Call {
	return &DB_GetCollections_Call{Call: _e.mock.On("GetCollections", _a0, _a1)}
}

func (_c *DB_GetCollections_Call) Run(run func(_a0 context.Context, _a1 client.CollectionFetchOptions)) *DB_GetCollections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.CollectionFetchOptions))
	})
	return _c
}

func (_c *DB_GetCollections_Call) Return(_a0 []client.Collection, _a1 error) *DB_GetCollections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetCollections_Call) RunAndReturn(run func(context.Context, client.CollectionFetchOptions) ([]client.Collection, error)) *DB_GetCollections_Call {
	_c.Call.Return(run)
	return _c
}

// GetSchemaByVersionID provides a mock function with given fields: _a0, _a1
func (_m *DB) GetSchemaByVersionID(_a0 context.Context, _a1 string) (client.SchemaDescription, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetSchemaByVersionID")
	}

	var r0 client.SchemaDescription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (client.SchemaDescription, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) client.SchemaDescription); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(client.SchemaDescription)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetSchemaByVersionID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSchemaByVersionID'
type DB_GetSchemaByVersionID_Call struct {
	*mock.Call
}

// GetSchemaByVersionID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *DB_Expecter) GetSchemaByVersionID(_a0 interface{}, _a1 interface{}) *DB_GetSchemaByVersionID_Call {
	return &DB_GetSchemaByVersionID_Call{Call: _e.mock.On("GetSchemaByVersionID", _a0, _a1)}
}

func (_c *DB_GetSchemaByVersionID_Call) Run(run func(_a0 context.Context, _a1 string)) *DB_GetSchemaByVersionID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_GetSchemaByVersionID_Call) Return(_a0 client.SchemaDescription, _a1 error) *DB_GetSchemaByVersionID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetSchemaByVersionID_Call) RunAndReturn(run func(context.Context, string) (client.SchemaDescription, error)) *DB_GetSchemaByVersionID_Call {
	_c.Call.Return(run)
	return _c
}

// GetSchemas provides a mock function with given fields: _a0, _a1
func (_m *DB) GetSchemas(_a0 context.Context, _a1 client.SchemaFetchOptions) ([]client.SchemaDescription, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetSchemas")
	}

	var r0 []client.SchemaDescription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, client.SchemaFetchOptions) ([]client.SchemaDescription, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, client.SchemaFetchOptions) []client.SchemaDescription); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.SchemaDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, client.SchemaFetchOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_GetSchemas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSchemas'
type DB_GetSchemas_Call struct {
	*mock.Call
}

// GetSchemas is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 client.SchemaFetchOptions
func (_e *DB_Expecter) GetSchemas(_a0 interface{}, _a1 interface{}) *DB_GetSchemas_Call {
	return &DB_GetSchemas_Call{Call: _e.mock.On("GetSchemas", _a0, _a1)}
}

func (_c *DB_GetSchemas_Call) Run(run func(_a0 context.Context, _a1 client.SchemaFetchOptions)) *DB_GetSchemas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.SchemaFetchOptions))
	})
	return _c
}

func (_c *DB_GetSchemas_Call) Return(_a0 []client.SchemaDescription, _a1 error) *DB_GetSchemas_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_GetSchemas_Call) RunAndReturn(run func(context.Context, client.SchemaFetchOptions) ([]client.SchemaDescription, error)) *DB_GetSchemas_Call {
	_c.Call.Return(run)
	return _c
}

// Headstore provides a mock function with given fields:
func (_m *DB) Headstore() go_datastore.Read {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Headstore")
	}

	var r0 go_datastore.Read
	if rf, ok := ret.Get(0).(func() go_datastore.Read); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(go_datastore.Read)
		}
	}

	return r0
}

// DB_Headstore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Headstore'
type DB_Headstore_Call struct {
	*mock.Call
}

// Headstore is a helper method to define mock.On call
func (_e *DB_Expecter) Headstore() *DB_Headstore_Call {
	return &DB_Headstore_Call{Call: _e.mock.On("Headstore")}
}

func (_c *DB_Headstore_Call) Run(run func()) *DB_Headstore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_Headstore_Call) Return(_a0 go_datastore.Read) *DB_Headstore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Headstore_Call) RunAndReturn(run func() go_datastore.Read) *DB_Headstore_Call {
	_c.Call.Return(run)
	return _c
}

// LensRegistry provides a mock function with given fields:
func (_m *DB) LensRegistry() client.LensRegistry {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LensRegistry")
	}

	var r0 client.LensRegistry
	if rf, ok := ret.Get(0).(func() client.LensRegistry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.LensRegistry)
		}
	}

	return r0
}

// DB_LensRegistry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LensRegistry'
type DB_LensRegistry_Call struct {
	*mock.Call
}

// LensRegistry is a helper method to define mock.On call
func (_e *DB_Expecter) LensRegistry() *DB_LensRegistry_Call {
	return &DB_LensRegistry_Call{Call: _e.mock.On("LensRegistry")}
}

func (_c *DB_LensRegistry_Call) Run(run func()) *DB_LensRegistry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_LensRegistry_Call) Return(_a0 client.LensRegistry) *DB_LensRegistry_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_LensRegistry_Call) RunAndReturn(run func() client.LensRegistry) *DB_LensRegistry_Call {
	_c.Call.Return(run)
	return _c
}

// MaxTxnRetries provides a mock function with given fields:
func (_m *DB) MaxTxnRetries() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MaxTxnRetries")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// DB_MaxTxnRetries_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MaxTxnRetries'
type DB_MaxTxnRetries_Call struct {
	*mock.Call
}

// MaxTxnRetries is a helper method to define mock.On call
func (_e *DB_Expecter) MaxTxnRetries() *DB_MaxTxnRetries_Call {
	return &DB_MaxTxnRetries_Call{Call: _e.mock.On("MaxTxnRetries")}
}

func (_c *DB_MaxTxnRetries_Call) Run(run func()) *DB_MaxTxnRetries_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_MaxTxnRetries_Call) Return(_a0 int) *DB_MaxTxnRetries_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_MaxTxnRetries_Call) RunAndReturn(run func() int) *DB_MaxTxnRetries_Call {
	_c.Call.Return(run)
	return _c
}

// NewConcurrentTxn provides a mock function with given fields: _a0, _a1
func (_m *DB) NewConcurrentTxn(_a0 context.Context, _a1 bool) (datastore.Txn, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for NewConcurrentTxn")
	}

	var r0 datastore.Txn
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool) (datastore.Txn, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool) datastore.Txn); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.Txn)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_NewConcurrentTxn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewConcurrentTxn'
type DB_NewConcurrentTxn_Call struct {
	*mock.Call
}

// NewConcurrentTxn is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 bool
func (_e *DB_Expecter) NewConcurrentTxn(_a0 interface{}, _a1 interface{}) *DB_NewConcurrentTxn_Call {
	return &DB_NewConcurrentTxn_Call{Call: _e.mock.On("NewConcurrentTxn", _a0, _a1)}
}

func (_c *DB_NewConcurrentTxn_Call) Run(run func(_a0 context.Context, _a1 bool)) *DB_NewConcurrentTxn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool))
	})
	return _c
}

func (_c *DB_NewConcurrentTxn_Call) Return(_a0 datastore.Txn, _a1 error) *DB_NewConcurrentTxn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_NewConcurrentTxn_Call) RunAndReturn(run func(context.Context, bool) (datastore.Txn, error)) *DB_NewConcurrentTxn_Call {
	_c.Call.Return(run)
	return _c
}

// NewTxn provides a mock function with given fields: _a0, _a1
func (_m *DB) NewTxn(_a0 context.Context, _a1 bool) (datastore.Txn, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for NewTxn")
	}

	var r0 datastore.Txn
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool) (datastore.Txn, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool) datastore.Txn); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.Txn)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_NewTxn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewTxn'
type DB_NewTxn_Call struct {
	*mock.Call
}

// NewTxn is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 bool
func (_e *DB_Expecter) NewTxn(_a0 interface{}, _a1 interface{}) *DB_NewTxn_Call {
	return &DB_NewTxn_Call{Call: _e.mock.On("NewTxn", _a0, _a1)}
}

func (_c *DB_NewTxn_Call) Run(run func(_a0 context.Context, _a1 bool)) *DB_NewTxn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool))
	})
	return _c
}

func (_c *DB_NewTxn_Call) Return(_a0 datastore.Txn, _a1 error) *DB_NewTxn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_NewTxn_Call) RunAndReturn(run func(context.Context, bool) (datastore.Txn, error)) *DB_NewTxn_Call {
	_c.Call.Return(run)
	return _c
}

// PatchCollection provides a mock function with given fields: _a0, _a1
func (_m *DB) PatchCollection(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for PatchCollection")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_PatchCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchCollection'
type DB_PatchCollection_Call struct {
	*mock.Call
}

// PatchCollection is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *DB_Expecter) PatchCollection(_a0 interface{}, _a1 interface{}) *DB_PatchCollection_Call {
	return &DB_PatchCollection_Call{Call: _e.mock.On("PatchCollection", _a0, _a1)}
}

func (_c *DB_PatchCollection_Call) Run(run func(_a0 context.Context, _a1 string)) *DB_PatchCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_PatchCollection_Call) Return(_a0 error) *DB_PatchCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_PatchCollection_Call) RunAndReturn(run func(context.Context, string) error) *DB_PatchCollection_Call {
	_c.Call.Return(run)
	return _c
}

// PatchSchema provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *DB) PatchSchema(_a0 context.Context, _a1 string, _a2 immutable.Option[model.Lens], _a3 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for PatchSchema")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, immutable.Option[model.Lens], bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_PatchSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchSchema'
type DB_PatchSchema_Call struct {
	*mock.Call
}

// PatchSchema is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 immutable.Option[model.Lens]
//   - _a3 bool
func (_e *DB_Expecter) PatchSchema(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *DB_PatchSchema_Call {
	return &DB_PatchSchema_Call{Call: _e.mock.On("PatchSchema", _a0, _a1, _a2, _a3)}
}

func (_c *DB_PatchSchema_Call) Run(run func(_a0 context.Context, _a1 string, _a2 immutable.Option[model.Lens], _a3 bool)) *DB_PatchSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(immutable.Option[model.Lens]), args[3].(bool))
	})
	return _c
}

func (_c *DB_PatchSchema_Call) Return(_a0 error) *DB_PatchSchema_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_PatchSchema_Call) RunAndReturn(run func(context.Context, string, immutable.Option[model.Lens], bool) error) *DB_PatchSchema_Call {
	_c.Call.Return(run)
	return _c
}

// Peerstore provides a mock function with given fields:
func (_m *DB) Peerstore() datastore.DSBatching {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Peerstore")
	}

	var r0 datastore.DSBatching
	if rf, ok := ret.Get(0).(func() datastore.DSBatching); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.DSBatching)
		}
	}

	return r0
}

// DB_Peerstore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Peerstore'
type DB_Peerstore_Call struct {
	*mock.Call
}

// Peerstore is a helper method to define mock.On call
func (_e *DB_Expecter) Peerstore() *DB_Peerstore_Call {
	return &DB_Peerstore_Call{Call: _e.mock.On("Peerstore")}
}

func (_c *DB_Peerstore_Call) Run(run func()) *DB_Peerstore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_Peerstore_Call) Return(_a0 datastore.DSBatching) *DB_Peerstore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Peerstore_Call) RunAndReturn(run func() datastore.DSBatching) *DB_Peerstore_Call {
	_c.Call.Return(run)
	return _c
}

// PrintDump provides a mock function with given fields: ctx
func (_m *DB) PrintDump(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for PrintDump")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_PrintDump_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PrintDump'
type DB_PrintDump_Call struct {
	*mock.Call
}

// PrintDump is a helper method to define mock.On call
//   - ctx context.Context
func (_e *DB_Expecter) PrintDump(ctx interface{}) *DB_PrintDump_Call {
	return &DB_PrintDump_Call{Call: _e.mock.On("PrintDump", ctx)}
}

func (_c *DB_PrintDump_Call) Run(run func(ctx context.Context)) *DB_PrintDump_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DB_PrintDump_Call) Return(_a0 error) *DB_PrintDump_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_PrintDump_Call) RunAndReturn(run func(context.Context) error) *DB_PrintDump_Call {
	_c.Call.Return(run)
	return _c
}

// Root provides a mock function with given fields:
func (_m *DB) Root() datastore.RootStore {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Root")
	}

	var r0 datastore.RootStore
	if rf, ok := ret.Get(0).(func() datastore.RootStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.RootStore)
		}
	}

	return r0
}

// DB_Root_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Root'
type DB_Root_Call struct {
	*mock.Call
}

// Root is a helper method to define mock.On call
func (_e *DB_Expecter) Root() *DB_Root_Call {
	return &DB_Root_Call{Call: _e.mock.On("Root")}
}

func (_c *DB_Root_Call) Run(run func()) *DB_Root_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DB_Root_Call) Return(_a0 datastore.RootStore) *DB_Root_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Root_Call) RunAndReturn(run func() datastore.RootStore) *DB_Root_Call {
	_c.Call.Return(run)
	return _c
}

// SetActiveSchemaVersion provides a mock function with given fields: _a0, _a1
func (_m *DB) SetActiveSchemaVersion(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SetActiveSchemaVersion")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_SetActiveSchemaVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetActiveSchemaVersion'
type DB_SetActiveSchemaVersion_Call struct {
	*mock.Call
}

// SetActiveSchemaVersion is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *DB_Expecter) SetActiveSchemaVersion(_a0 interface{}, _a1 interface{}) *DB_SetActiveSchemaVersion_Call {
	return &DB_SetActiveSchemaVersion_Call{Call: _e.mock.On("SetActiveSchemaVersion", _a0, _a1)}
}

func (_c *DB_SetActiveSchemaVersion_Call) Run(run func(_a0 context.Context, _a1 string)) *DB_SetActiveSchemaVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DB_SetActiveSchemaVersion_Call) Return(_a0 error) *DB_SetActiveSchemaVersion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_SetActiveSchemaVersion_Call) RunAndReturn(run func(context.Context, string) error) *DB_SetActiveSchemaVersion_Call {
	_c.Call.Return(run)
	return _c
}

// SetMigration provides a mock function with given fields: _a0, _a1
func (_m *DB) SetMigration(_a0 context.Context, _a1 client.LensConfig) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SetMigration")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.LensConfig) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_SetMigration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetMigration'
type DB_SetMigration_Call struct {
	*mock.Call
}

// SetMigration is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 client.LensConfig
func (_e *DB_Expecter) SetMigration(_a0 interface{}, _a1 interface{}) *DB_SetMigration_Call {
	return &DB_SetMigration_Call{Call: _e.mock.On("SetMigration", _a0, _a1)}
}

func (_c *DB_SetMigration_Call) Run(run func(_a0 context.Context, _a1 client.LensConfig)) *DB_SetMigration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.LensConfig))
	})
	return _c
}

func (_c *DB_SetMigration_Call) Return(_a0 error) *DB_SetMigration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_SetMigration_Call) RunAndReturn(run func(context.Context, client.LensConfig) error) *DB_SetMigration_Call {
	_c.Call.Return(run)
	return _c
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
