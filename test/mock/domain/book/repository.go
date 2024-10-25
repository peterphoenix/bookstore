// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package book_mock is a generated GoMock package.
package book_mock

import (
	pagination "bookstore/common/pagination"
	model "bookstore/domain/book/model"
	model0 "bookstore/domain/common/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetBookDetails mocks base method.
func (m *MockRepository) GetBookDetails(ctx context.Context, id []string) ([]model.BookDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookDetails", ctx, id)
	ret0, _ := ret[0].([]model.BookDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookDetails indicates an expected call of GetBookDetails.
func (mr *MockRepositoryMockRecorder) GetBookDetails(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookDetails", reflect.TypeOf((*MockRepository)(nil).GetBookDetails), ctx, id)
}

// GetBooksList mocks base method.
func (m *MockRepository) GetBooksList(ctx context.Context, pagination pagination.Pagination) ([]model0.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooksList", ctx, pagination)
	ret0, _ := ret[0].([]model0.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooksList indicates an expected call of GetBooksList.
func (mr *MockRepositoryMockRecorder) GetBooksList(ctx, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooksList", reflect.TypeOf((*MockRepository)(nil).GetBooksList), ctx, pagination)
}
