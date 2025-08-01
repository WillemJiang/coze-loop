// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/coze-loop/backend/modules/data/infra/repo/dataset/mysql (interfaces: IItemDAO)
//
// Generated by this command:
//
//	mockgen -destination=mocks/item_dao.go -package=mocks . IItemDAO
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	db "github.com/coze-dev/coze-loop/backend/infra/db"
	entity "github.com/coze-dev/coze-loop/backend/modules/data/domain/dataset/entity"
	mysql "github.com/coze-dev/coze-loop/backend/modules/data/infra/repo/dataset/mysql"
	model "github.com/coze-dev/coze-loop/backend/modules/data/infra/repo/dataset/mysql/gorm_gen/model"
	pagination "github.com/coze-dev/coze-loop/backend/modules/data/pkg/pagination"
	gomock "go.uber.org/mock/gomock"
)

// MockIItemDAO is a mock of IItemDAO interface.
type MockIItemDAO struct {
	ctrl     *gomock.Controller
	recorder *MockIItemDAOMockRecorder
	isgomock struct{}
}

// MockIItemDAOMockRecorder is the mock recorder for MockIItemDAO.
type MockIItemDAOMockRecorder struct {
	mock *MockIItemDAO
}

// NewMockIItemDAO creates a new mock instance.
func NewMockIItemDAO(ctrl *gomock.Controller) *MockIItemDAO {
	mock := &MockIItemDAO{ctrl: ctrl}
	mock.recorder = &MockIItemDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIItemDAO) EXPECT() *MockIItemDAOMockRecorder {
	return m.recorder
}

// ArchiveItems mocks base method.
func (m *MockIItemDAO) ArchiveItems(ctx context.Context, spaceID, delVN int64, ids []int64, opt ...db.Option) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, spaceID, delVN, ids}
	for _, a := range opt {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ArchiveItems", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ArchiveItems indicates an expected call of ArchiveItems.
func (mr *MockIItemDAOMockRecorder) ArchiveItems(ctx, spaceID, delVN, ids any, opt ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, spaceID, delVN, ids}, opt...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArchiveItems", reflect.TypeOf((*MockIItemDAO)(nil).ArchiveItems), varargs...)
}

// ClearDataset mocks base method.
func (m *MockIItemDAO) ClearDataset(ctx context.Context, spaceID, datasetID, delVN int64, opt ...db.Option) ([]*entity.ItemIdentity, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, spaceID, datasetID, delVN}
	for _, a := range opt {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClearDataset", varargs...)
	ret0, _ := ret[0].([]*entity.ItemIdentity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClearDataset indicates an expected call of ClearDataset.
func (mr *MockIItemDAOMockRecorder) ClearDataset(ctx, spaceID, datasetID, delVN any, opt ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, spaceID, datasetID, delVN}, opt...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearDataset", reflect.TypeOf((*MockIItemDAO)(nil).ClearDataset), varargs...)
}

// CountItems mocks base method.
func (m *MockIItemDAO) CountItems(ctx context.Context, params *mysql.ListItemsParams, opt ...db.Option) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range opt {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountItems", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountItems indicates an expected call of CountItems.
func (mr *MockIItemDAOMockRecorder) CountItems(ctx, params any, opt ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, opt...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountItems", reflect.TypeOf((*MockIItemDAO)(nil).CountItems), varargs...)
}

// DeleteItems mocks base method.
func (m *MockIItemDAO) DeleteItems(ctx context.Context, spaceID int64, ids []int64, opt ...db.Option) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, spaceID, ids}
	for _, a := range opt {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteItems", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteItems indicates an expected call of DeleteItems.
func (mr *MockIItemDAOMockRecorder) DeleteItems(ctx, spaceID, ids any, opt ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, spaceID, ids}, opt...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItems", reflect.TypeOf((*MockIItemDAO)(nil).DeleteItems), varargs...)
}

// ListItems mocks base method.
func (m *MockIItemDAO) ListItems(ctx context.Context, params *mysql.ListItemsParams, opt ...db.Option) ([]*model.DatasetItem, *pagination.PageResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range opt {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListItems", varargs...)
	ret0, _ := ret[0].([]*model.DatasetItem)
	ret1, _ := ret[1].(*pagination.PageResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListItems indicates an expected call of ListItems.
func (mr *MockIItemDAOMockRecorder) ListItems(ctx, params any, opt ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, opt...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItems", reflect.TypeOf((*MockIItemDAO)(nil).ListItems), varargs...)
}

// MCreateItems mocks base method.
func (m *MockIItemDAO) MCreateItems(ctx context.Context, items []*model.DatasetItem, opt ...db.Option) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, items}
	for _, a := range opt {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MCreateItems", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MCreateItems indicates an expected call of MCreateItems.
func (mr *MockIItemDAOMockRecorder) MCreateItems(ctx, items any, opt ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, items}, opt...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MCreateItems", reflect.TypeOf((*MockIItemDAO)(nil).MCreateItems), varargs...)
}

// UpdateItem mocks base method.
func (m *MockIItemDAO) UpdateItem(ctx context.Context, item *model.DatasetItem, opt ...db.Option) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, item}
	for _, a := range opt {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateItem", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockIItemDAOMockRecorder) UpdateItem(ctx, item any, opt ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, item}, opt...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockIItemDAO)(nil).UpdateItem), varargs...)
}
