// Code generated by mockery v1.0.0. DO NOT EDIT.

package report

import (
	all "github.com/goharbor/harbor/src/pkg/scan/all"
	mock "github.com/stretchr/testify/mock"

	scan "github.com/goharbor/harbor/src/pkg/scan/dao/scan"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Create provides a mock function with given fields: r
func (_m *Manager) Create(r *scan.Report) (string, error) {
	ret := _m.Called(r)

	var r0 string
	if rf, ok := ret.Get(0).(func(*scan.Report) string); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan.Report) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByDigests provides a mock function with given fields: digests
func (_m *Manager) DeleteByDigests(digests ...string) error {
	_va := make([]interface{}, len(digests))
	for _i := range digests {
		_va[_i] = digests[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(digests...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: uuid
func (_m *Manager) Get(uuid string) (*scan.Report, error) {
	ret := _m.Called(uuid)

	var r0 *scan.Report
	if rf, ok := ret.Get(0).(func(string) *scan.Report); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan.Report)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBy provides a mock function with given fields: digest, registrationUUID, mimeTypes
func (_m *Manager) GetBy(digest string, registrationUUID string, mimeTypes []string) ([]*scan.Report, error) {
	ret := _m.Called(digest, registrationUUID, mimeTypes)

	var r0 []*scan.Report
	if rf, ok := ret.Get(0).(func(string, string, []string) []*scan.Report); ok {
		r0 = rf(digest, registrationUUID, mimeTypes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*scan.Report)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, []string) error); ok {
		r1 = rf(digest, registrationUUID, mimeTypes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStats provides a mock function with given fields: requester
func (_m *Manager) GetStats(requester string) (*all.Stats, error) {
	ret := _m.Called(requester)

	var r0 *all.Stats
	if rf, ok := ret.Get(0).(func(string) *all.Stats); ok {
		r0 = rf(requester)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*all.Stats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(requester)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReportData provides a mock function with given fields: uuid, _a1, rev
func (_m *Manager) UpdateReportData(uuid string, _a1 string, rev int64) error {
	ret := _m.Called(uuid, _a1, rev)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(uuid, _a1, rev)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateScanJobID provides a mock function with given fields: trackID, jobID
func (_m *Manager) UpdateScanJobID(trackID string, jobID string) error {
	ret := _m.Called(trackID, jobID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(trackID, jobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatus provides a mock function with given fields: trackID, status, rev
func (_m *Manager) UpdateStatus(trackID string, status string, rev int64) error {
	ret := _m.Called(trackID, status, rev)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(trackID, status, rev)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}