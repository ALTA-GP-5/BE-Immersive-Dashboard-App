// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mentorentity "immersive/domains/mentor/entities"

	mock "github.com/stretchr/testify/mock"
)

// MentorRepoMock is an autogenerated mock type for the IMentorRepo type
type MentorRepoMock struct {
	mock.Mock
}

// Delete provides a mock function with given fields: mentor
func (_m *MentorRepoMock) Delete(mentor mentorentity.MentorEntity) error {
	ret := _m.Called(mentor)

	var r0 error
	if rf, ok := ret.Get(0).(func(mentorentity.MentorEntity) error); ok {
		r0 = rf(mentor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: mentor
func (_m *MentorRepoMock) GetAll(mentor mentorentity.MentorEntity) ([]mentorentity.MentorEntity, error) {
	ret := _m.Called(mentor)

	var r0 []mentorentity.MentorEntity
	if rf, ok := ret.Get(0).(func(mentorentity.MentorEntity) []mentorentity.MentorEntity); ok {
		r0 = rf(mentor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentorentity.MentorEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mentorentity.MentorEntity) error); ok {
		r1 = rf(mentor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: mentor
func (_m *MentorRepoMock) GetById(mentor mentorentity.MentorEntity) (mentorentity.MentorEntity, error) {
	ret := _m.Called(mentor)

	var r0 mentorentity.MentorEntity
	if rf, ok := ret.Get(0).(func(mentorentity.MentorEntity) mentorentity.MentorEntity); ok {
		r0 = rf(mentor)
	} else {
		r0 = ret.Get(0).(mentorentity.MentorEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mentorentity.MentorEntity) error); ok {
		r1 = rf(mentor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: mentor
func (_m *MentorRepoMock) Insert(mentor mentorentity.MentorEntity) error {
	ret := _m.Called(mentor)

	var r0 error
	if rf, ok := ret.Get(0).(func(mentorentity.MentorEntity) error); ok {
		r0 = rf(mentor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: mentor
func (_m *MentorRepoMock) Update(mentor mentorentity.MentorEntity) error {
	ret := _m.Called(mentor)

	var r0 error
	if rf, ok := ret.Get(0).(func(mentorentity.MentorEntity) error); ok {
		r0 = rf(mentor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMentorRepoMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewMentorRepoMock creates a new instance of MentorRepoMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMentorRepoMock(t mockConstructorTestingTNewMentorRepoMock) *MentorRepoMock {
	mock := &MentorRepoMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
