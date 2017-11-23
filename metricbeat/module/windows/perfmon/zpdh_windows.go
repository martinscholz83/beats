// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package perfmon

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modpdh = windows.NewLazySystemDLL("pdh.dll")

	procPdhOpenQueryW                   = modpdh.NewProc("PdhOpenQueryW")
	procPdhAddEnglishCounterW           = modpdh.NewProc("PdhAddEnglishCounterW")
	procPdhCollectQueryData             = modpdh.NewProc("PdhCollectQueryData")
	procPdhGetFormattedCounterValue     = modpdh.NewProc("PdhGetFormattedCounterValue")
	procPdhGetFormattedCounterArrayW    = modpdh.NewProc("PdhGetFormattedCounterArrayW")
	procPdhGetRawCounterValue           = modpdh.NewProc("PdhGetRawCounterValue")
	procPdhGetRawCounterArrayW          = modpdh.NewProc("PdhGetRawCounterArrayW")
	procPdhCalculateCounterFromRawValue = modpdh.NewProc("PdhCalculateCounterFromRawValue")
	procPdhFormatFromRawValue           = modpdh.NewProc("PdhFormatFromRawValue")
	procPdhCloseQuery                   = modpdh.NewProc("PdhCloseQuery")
)

func _PdhOpenQuery(dataSource *uint16, userData uintptr, query *PdhQueryHandle) (errcode error) {
	r0, _, _ := syscall.Syscall(procPdhOpenQueryW.Addr(), 3, uintptr(unsafe.Pointer(dataSource)), uintptr(userData), uintptr(unsafe.Pointer(query)))
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhAddCounter(query PdhQueryHandle, counterPath string, userData uintptr, counter *PdhCounterHandle) (errcode error) {
	var _p0 *uint16
	_p0, errcode = syscall.UTF16PtrFromString(counterPath)
	if errcode != nil {
		return
	}
	return __PdhAddCounter(query, _p0, userData, counter)
}

func __PdhAddCounter(query PdhQueryHandle, counterPath *uint16, userData uintptr, counter *PdhCounterHandle) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhAddEnglishCounterW.Addr(), 4, uintptr(query), uintptr(unsafe.Pointer(counterPath)), uintptr(userData), uintptr(unsafe.Pointer(counter)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhCollectQueryData(query PdhQueryHandle) (errcode error) {
	r0, _, _ := syscall.Syscall(procPdhCollectQueryData.Addr(), 1, uintptr(query), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhGetFormattedCounterValue(counter PdhCounterHandle, format PdhCounterFormat, counterType *uint32, value *PdhCounterValue) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhGetFormattedCounterValue.Addr(), 4, uintptr(counter), uintptr(format), uintptr(unsafe.Pointer(counterType)), uintptr(unsafe.Pointer(value)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhGetFormattedCounterArray(counter PdhCounterHandle, format PdhCounterFormat, bufferSize *uint32, bufferCount *uint32, itemBuffer *byte) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhGetFormattedCounterArrayW.Addr(), 5, uintptr(counter), uintptr(format), uintptr(unsafe.Pointer(bufferSize)), uintptr(unsafe.Pointer(bufferCount)), uintptr(unsafe.Pointer(itemBuffer)), 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhGetRawCounterValue(counter PdhCounterHandle, counterType *uint32, value *PdhRawCounter) (errcode error) {
	r0, _, _ := syscall.Syscall(procPdhGetRawCounterValue.Addr(), 3, uintptr(counter), uintptr(unsafe.Pointer(counterType)), uintptr(unsafe.Pointer(value)))
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhGetRawCounterArray(counter PdhCounterHandle, bufferSize *uint32, bufferCount *uint32, itemBuffer *byte) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhGetRawCounterArrayW.Addr(), 4, uintptr(counter), uintptr(unsafe.Pointer(bufferSize)), uintptr(unsafe.Pointer(bufferCount)), uintptr(unsafe.Pointer(itemBuffer)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhCalculateCounterFromRawValue(counter PdhCounterHandle, format PdhCounterFormat, rawValue1 *PdhRawCounter, rawValue2 *PdhRawCounter, value *PdhCounterValue) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhCalculateCounterFromRawValue.Addr(), 5, uintptr(counter), uintptr(format), uintptr(unsafe.Pointer(rawValue1)), uintptr(unsafe.Pointer(rawValue2)), uintptr(unsafe.Pointer(value)), 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhFormatFromRawValue(counterType uint32, format PdhCounterFormat, timeBase *uint64, rawValue1 *PdhRawCounter, rawValue2 *PdhRawCounter, value *PdhCounterValue) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhFormatFromRawValue.Addr(), 6, uintptr(counterType), uintptr(format), uintptr(unsafe.Pointer(timeBase)), uintptr(unsafe.Pointer(rawValue1)), uintptr(unsafe.Pointer(rawValue2)), uintptr(unsafe.Pointer(value)))
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhCloseQuery(query PdhQueryHandle) (errcode error) {
	r0, _, _ := syscall.Syscall(procPdhCloseQuery.Addr(), 1, uintptr(query), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}
