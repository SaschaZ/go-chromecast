// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	application "github.com/vishen/go-chromecast/application"
	cast "github.com/vishen/go-chromecast/cast"

	mock "github.com/stretchr/testify/mock"

	net "net"
)

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// AddMessageFunc provides a mock function with given fields: f
func (_m *App) AddMessageFunc(f application.CastMessageFunc) {
	_m.Called(f)
}

// Close provides a mock function with given fields: stopMedia
func (_m *App) Close(stopMedia bool) error {
	ret := _m.Called(stopMedia)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(stopMedia)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Load provides a mock function with given fields: filenameOrUrl, contentType, transcode, detach, forceDetach
func (_m *App) Load(filenameOrUrl string, contentType string, transcode bool, detach bool, forceDetach bool) error {
	ret := _m.Called(filenameOrUrl, contentType, transcode, detach, forceDetach)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, bool) error); ok {
		r0 = rf(filenameOrUrl, contentType, transcode, detach, forceDetach)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadApp provides a mock function with given fields: appID, contentID
func (_m *App) LoadApp(appID string, contentID string) error {
	ret := _m.Called(appID, contentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(appID, contentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *App) Next() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Pause provides a mock function with given fields:
func (_m *App) Pause() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PlayableMediaType provides a mock function with given fields: filename
func (_m *App) PlayableMediaType(filename string) bool {
	ret := _m.Called(filename)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(filename)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PlayedItems provides a mock function with given fields:
func (_m *App) PlayedItems() map[string]application.PlayedItem {
	ret := _m.Called()

	var r0 map[string]application.PlayedItem
	if rf, ok := ret.Get(0).(func() map[string]application.PlayedItem); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]application.PlayedItem)
		}
	}

	return r0
}

// Previous provides a mock function with given fields:
func (_m *App) Previous() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueueLoad provides a mock function with given fields: filenames, contentType, transcode
func (_m *App) QueueLoad(filenames []string, contentType string, transcode bool) error {
	ret := _m.Called(filenames, contentType, transcode)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string, string, bool) error); ok {
		r0 = rf(filenames, contentType, transcode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Seek provides a mock function with given fields: value
func (_m *App) Seek(value int) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SeekFromStart provides a mock function with given fields: value
func (_m *App) SeekFromStart(value int) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SeekToTime provides a mock function with given fields: value
func (_m *App) SeekToTime(value float32) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(float32) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCacheDisabled provides a mock function with given fields: _a0
func (_m *App) SetCacheDisabled(_a0 bool) {
	_m.Called(_a0)
}

// SetConn provides a mock function with given fields: conn
func (_m *App) SetConn(conn cast.Conn) {
	_m.Called(conn)
}

// SetConnectionRetries provides a mock function with given fields: _a0
func (_m *App) SetConnectionRetries(_a0 int) {
	_m.Called(_a0)
}

// SetDebug provides a mock function with given fields: _a0
func (_m *App) SetDebug(_a0 bool) {
	_m.Called(_a0)
}

// SetIface provides a mock function with given fields: _a0
func (_m *App) SetIface(_a0 *net.Interface) {
	_m.Called(_a0)
}

// SetMuted provides a mock function with given fields: value
func (_m *App) SetMuted(value bool) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetServerPort provides a mock function with given fields: _a0
func (_m *App) SetServerPort(_a0 int) {
	_m.Called(_a0)
}

// SetVolume provides a mock function with given fields: value
func (_m *App) SetVolume(value float32) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(float32) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Skipad provides a mock function with given fields:
func (_m *App) Skipad() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Slideshow provides a mock function with given fields: filenames, duration, repeat
func (_m *App) Slideshow(filenames []string, duration int, repeat bool) error {
	ret := _m.Called(filenames, duration, repeat)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string, int, bool) error); ok {
		r0 = rf(filenames, duration, repeat)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: addr, port
func (_m *App) Start(addr string, port int) error {
	ret := _m.Called(addr, port)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(addr, port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status provides a mock function with given fields:
func (_m *App) Status() (*cast.Application, *cast.Media, *cast.Volume) {
	ret := _m.Called()

	var r0 *cast.Application
	if rf, ok := ret.Get(0).(func() *cast.Application); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cast.Application)
		}
	}

	var r1 *cast.Media
	if rf, ok := ret.Get(1).(func() *cast.Media); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cast.Media)
		}
	}

	var r2 *cast.Volume
	if rf, ok := ret.Get(2).(func() *cast.Volume); ok {
		r2 = rf()
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*cast.Volume)
		}
	}

	return r0, r1, r2
}

// Stop provides a mock function with given fields:
func (_m *App) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopMedia provides a mock function with given fields:
func (_m *App) StopMedia() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transcode provides a mock function with given fields: command, contentType
func (_m *App) Transcode(command string, contentType string) error {
	ret := _m.Called(command, contentType)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(command, contentType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unpause provides a mock function with given fields:
func (_m *App) Unpause() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields:
func (_m *App) Update() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}