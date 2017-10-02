// Created by cgo -godefs - DO NOT EDIT
// cgo.exe -godefs defs_services_windows.go

package services

type ServiceErrno uintptr

const (
	SERVICE_ERROR_ACCESS_DENIED           ServiceErrno = 0x5
	SERVICE_ERROR_MORE_DATA               ServiceErrno = 0xea
	SERVICE_ERROR_INVALID_PARAMETER       ServiceErrno = 0x57
	SERVICE_ERROR_INVALID_HANDLE          ServiceErrno = 0x6
	SERVICE_ERROR_INVALID_LEVEL           ServiceErrno = 0x7c
	SERVICE_ERROR_SHUTDOWN_IN_PROGRESS    ServiceErrno = 0x45b
	SERVICE_ERROR_DATABASE_DOES_NOT_EXIST ServiceErrno = 0x429
)

type ServiceErrorControl uint32

const (
	SERVICE_ERROR_CRITICAL ServiceErrno = 0x3
	SERVICE_ERROR_IGNORE   ServiceErrno = 0x0
	SERVICE_ERROR_NORMAL   ServiceErrno = 0x1
	SERVICE_ERROR_SEVERE   ServiceErrno = 0x2
)

var serviceErrors = map[ServiceErrno]struct{}{
	SERVICE_ERROR_ACCESS_DENIED:           struct{}{},
	SERVICE_ERROR_MORE_DATA:               struct{}{},
	SERVICE_ERROR_INVALID_PARAMETER:       struct{}{},
	SERVICE_ERROR_INVALID_HANDLE:          struct{}{},
	SERVICE_ERROR_INVALID_LEVEL:           struct{}{},
	SERVICE_ERROR_SHUTDOWN_IN_PROGRESS:    struct{}{},
	SERVICE_ERROR_DATABASE_DOES_NOT_EXIST: struct{}{},
	SERVICE_ERROR_CRITICAL:                struct{}{},
	SERVICE_ERROR_IGNORE:                  struct{}{},
	SERVICE_ERROR_NORMAL:                  struct{}{},
	SERVICE_ERROR_SEVERE:                  struct{}{},
}

type ServiceType uint32

const (
	ServiceDriver ServiceType = 0xb

	ServiceFileSystemDriver ServiceType = 0x2

	ServiceKernelDriver ServiceType = 0x1

	ServiceWin32 ServiceType = 0x30

	ServiceWin32OwnProcess ServiceType = 0x10

	ServiceWin32Shareprocess  ServiceType = 0x20
	ServiceInteractiveProcess ServiceType = 0x100
)

type ServiceState uint32

const (
	ServiceContinuePending ServiceState = 0x5
	ServicePausePending    ServiceState = 0x6
	ServicePaused          ServiceState = 0x7
	ServiceRuning          ServiceState = 0x4
	ServiceStartPending    ServiceState = 0x2
	ServiceStopPending     ServiceState = 0x3
	ServiceStopped         ServiceState = 0x1
)

type ServiceEnumState uint32

const (
	ServiceActive ServiceEnumState = 0x1

	ServiceInActive ServiceEnumState = 0x2

	ServiceStateAll ServiceEnumState = 0x3
)

type ServiceSCMAccessRight uint32

const (
	ScManagerAllAccess ServiceSCMAccessRight = 0xf003f

	ScManagerConnect ServiceSCMAccessRight = 0x1

	ScManagerEnumerateService ServiceSCMAccessRight = 0x4

	ScManagerQueryLockStatus ServiceSCMAccessRight = 0x10
)

type ServiceAccessRight uint32

const (
	ServiceAllAccess ServiceAccessRight = 0xf01ff

	ServcieChangeConfig ServiceAccessRight = 0x2

	ServiceEnumerateDependents ServiceAccessRight = 0x8

	ServiceInterrogate ServiceAccessRight = 0x80

	ServicePauseContinue ServiceAccessRight = 0x40

	ServiceQueryConfig ServiceAccessRight = 0x1

	ServiceQueryStatus ServiceAccessRight = 0x4

	ServiceStart ServiceAccessRight = 0x10

	ServiceStop ServiceAccessRight = 0x20

	ServiceUserDefinedControl ServiceAccessRight = 0x100
)

type ServiceInfoLevel uint32

const (
	ScEnumProcessInfo ServiceInfoLevel = 0x0
)

type ServiceStartType uint32

const (
	ServiceAuotStart ServiceStartType = 0x2

	ServiceBootStart ServiceStartType = 0x0

	ServiceDemandStart ServiceStartType = 0x3

	ServiceDisabled ServiceStartType = 0x4

	ServcieSystemStart ServiceStartType = 0x1
)

type ServiceStatusProcess struct {
	DwServiceType             uint32
	DwCurrentState            uint32
	DwControlsAccepted        uint32
	DwWin32ExitCode           uint32
	DwServiceSpecificExitCode uint32
	DwCheckPoint              uint32
	DwWaitHint                uint32
	DwProcessId               uint32
	DwServiceFlags            uint32
}

type EnumServiceStatusProcess struct {
	LpServiceName        *int8
	LpDisplayName        *int8
	ServiceStatusProcess ServiceStatusProcess
}
