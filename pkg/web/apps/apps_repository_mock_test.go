// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
	"sync"
)

var (
	lockRepositoryMockDisableAllAppVersionsByAppID sync.RWMutex
	lockRepositoryMockGetAppByID                   sync.RWMutex
	lockRepositoryMockGetAppVersionsByAppID        sync.RWMutex
	lockRepositoryMockGetApps                      sync.RWMutex
	lockRepositoryMockUpdateAppVersions            sync.RWMutex
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//     func TestSomethingThatUsesRepository(t *testing.T) {
//
//         // make and configure a mocked Repository
//         mockedRepository := &RepositoryMock{
//             DisableAllAppVersionsByAppIDFunc: func(appID string, message string) error {
// 	               panic("mock out the DisableAllAppVersionsByAppID method")
//             },
//             GetAppByIDFunc: func(ID string) (*models.App, error) {
// 	               panic("mock out the GetAppByID method")
//             },
//             GetAppVersionsByAppIDFunc: func(ID string) (*[]models.Version, error) {
// 	               panic("mock out the GetAppVersionsByAppID method")
//             },
//             GetAppsFunc: func() (*[]models.App, error) {
// 	               panic("mock out the GetApps method")
//             },
//             UpdateAppVersionsFunc: func(versions []models.Version) error {
// 	               panic("mock out the UpdateAppVersions method")
//             },
//         }
//
//         // use mockedRepository in code that requires Repository
//         // and then make assertions.
//
//     }
type RepositoryMock struct {
	// DisableAllAppVersionsByAppIDFunc mocks the DisableAllAppVersionsByAppID method.
	DisableAllAppVersionsByAppIDFunc func(appID string, message string) error

	// GetAppByIDFunc mocks the GetAppByID method.
	GetAppByIDFunc func(ID string) (*models.App, error)

	// GetAppVersionsByAppIDFunc mocks the GetAppVersionsByAppID method.
	GetAppVersionsByAppIDFunc func(ID string) (*[]models.Version, error)

	// GetAppsFunc mocks the GetApps method.
	GetAppsFunc func() (*[]models.App, error)

	// UpdateAppVersionsFunc mocks the UpdateAppVersions method.
	UpdateAppVersionsFunc func(versions []models.Version) error

	// calls tracks calls to the methods.
	calls struct {
		// DisableAllAppVersionsByAppID holds details about calls to the DisableAllAppVersionsByAppID method.
		DisableAllAppVersionsByAppID []struct {
			// AppID is the appID argument value.
			AppID string
			// Message is the message argument value.
			Message string
		}
		// GetAppByID holds details about calls to the GetAppByID method.
		GetAppByID []struct {
			// ID is the ID argument value.
			ID string
		}
		// GetAppVersionsByAppID holds details about calls to the GetAppVersionsByAppID method.
		GetAppVersionsByAppID []struct {
			// ID is the ID argument value.
			ID string
		}
		// GetApps holds details about calls to the GetApps method.
		GetApps []struct {
		}
		// UpdateAppVersions holds details about calls to the UpdateAppVersions method.
		UpdateAppVersions []struct {
			// Versions is the versions argument value.
			Versions []models.Version
		}
	}
}

// DisableAllAppVersionsByAppID calls DisableAllAppVersionsByAppIDFunc.
func (mock *RepositoryMock) DisableAllAppVersionsByAppID(appID string, message string) error {
	if mock.DisableAllAppVersionsByAppIDFunc == nil {
		panic("RepositoryMock.DisableAllAppVersionsByAppIDFunc: method is nil but Repository.DisableAllAppVersionsByAppID was just called")
	}
	callInfo := struct {
		AppID   string
		Message string
	}{
		AppID:   appID,
		Message: message,
	}
	lockRepositoryMockDisableAllAppVersionsByAppID.Lock()
	mock.calls.DisableAllAppVersionsByAppID = append(mock.calls.DisableAllAppVersionsByAppID, callInfo)
	lockRepositoryMockDisableAllAppVersionsByAppID.Unlock()
	return mock.DisableAllAppVersionsByAppIDFunc(appID, message)
}

// DisableAllAppVersionsByAppIDCalls gets all the calls that were made to DisableAllAppVersionsByAppID.
// Check the length with:
//     len(mockedRepository.DisableAllAppVersionsByAppIDCalls())
func (mock *RepositoryMock) DisableAllAppVersionsByAppIDCalls() []struct {
	AppID   string
	Message string
} {
	var calls []struct {
		AppID   string
		Message string
	}
	lockRepositoryMockDisableAllAppVersionsByAppID.RLock()
	calls = mock.calls.DisableAllAppVersionsByAppID
	lockRepositoryMockDisableAllAppVersionsByAppID.RUnlock()
	return calls
}

// GetAppByID calls GetAppByIDFunc.
func (mock *RepositoryMock) GetAppByID(ID string) (*models.App, error) {
	if mock.GetAppByIDFunc == nil {
		panic("RepositoryMock.GetAppByIDFunc: method is nil but Repository.GetAppByID was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: ID,
	}
	lockRepositoryMockGetAppByID.Lock()
	mock.calls.GetAppByID = append(mock.calls.GetAppByID, callInfo)
	lockRepositoryMockGetAppByID.Unlock()
	return mock.GetAppByIDFunc(ID)
}

// GetAppByIDCalls gets all the calls that were made to GetAppByID.
// Check the length with:
//     len(mockedRepository.GetAppByIDCalls())
func (mock *RepositoryMock) GetAppByIDCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	lockRepositoryMockGetAppByID.RLock()
	calls = mock.calls.GetAppByID
	lockRepositoryMockGetAppByID.RUnlock()
	return calls
}

// GetAppVersionsByAppID calls GetAppVersionsByAppIDFunc.
func (mock *RepositoryMock) GetAppVersionsByAppID(ID string) (*[]models.Version, error) {
	if mock.GetAppVersionsByAppIDFunc == nil {
		panic("RepositoryMock.GetAppVersionsByAppIDFunc: method is nil but Repository.GetAppVersionsByAppID was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: ID,
	}
	lockRepositoryMockGetAppVersionsByAppID.Lock()
	mock.calls.GetAppVersionsByAppID = append(mock.calls.GetAppVersionsByAppID, callInfo)
	lockRepositoryMockGetAppVersionsByAppID.Unlock()
	return mock.GetAppVersionsByAppIDFunc(ID)
}

// GetAppVersionsByAppIDCalls gets all the calls that were made to GetAppVersionsByAppID.
// Check the length with:
//     len(mockedRepository.GetAppVersionsByAppIDCalls())
func (mock *RepositoryMock) GetAppVersionsByAppIDCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	lockRepositoryMockGetAppVersionsByAppID.RLock()
	calls = mock.calls.GetAppVersionsByAppID
	lockRepositoryMockGetAppVersionsByAppID.RUnlock()
	return calls
}

// GetApps calls GetAppsFunc.
func (mock *RepositoryMock) GetApps() (*[]models.App, error) {
	if mock.GetAppsFunc == nil {
		panic("RepositoryMock.GetAppsFunc: method is nil but Repository.GetApps was just called")
	}
	callInfo := struct {
	}{}
	lockRepositoryMockGetApps.Lock()
	mock.calls.GetApps = append(mock.calls.GetApps, callInfo)
	lockRepositoryMockGetApps.Unlock()
	return mock.GetAppsFunc()
}

// GetAppsCalls gets all the calls that were made to GetApps.
// Check the length with:
//     len(mockedRepository.GetAppsCalls())
func (mock *RepositoryMock) GetAppsCalls() []struct {
} {
	var calls []struct {
	}
	lockRepositoryMockGetApps.RLock()
	calls = mock.calls.GetApps
	lockRepositoryMockGetApps.RUnlock()
	return calls
}

// UpdateAppVersions calls UpdateAppVersionsFunc.
func (mock *RepositoryMock) UpdateAppVersions(versions []models.Version) error {
	if mock.UpdateAppVersionsFunc == nil {
		panic("RepositoryMock.UpdateAppVersionsFunc: method is nil but Repository.UpdateAppVersions was just called")
	}
	callInfo := struct {
		Versions []models.Version
	}{
		Versions: versions,
	}
	lockRepositoryMockUpdateAppVersions.Lock()
	mock.calls.UpdateAppVersions = append(mock.calls.UpdateAppVersions, callInfo)
	lockRepositoryMockUpdateAppVersions.Unlock()
	return mock.UpdateAppVersionsFunc(versions)
}

// UpdateAppVersionsCalls gets all the calls that were made to UpdateAppVersions.
// Check the length with:
//     len(mockedRepository.UpdateAppVersionsCalls())
func (mock *RepositoryMock) UpdateAppVersionsCalls() []struct {
	Versions []models.Version
} {
	var calls []struct {
		Versions []models.Version
	}
	lockRepositoryMockUpdateAppVersions.RLock()
	calls = mock.calls.UpdateAppVersions
	lockRepositoryMockUpdateAppVersions.RUnlock()
	return calls
}