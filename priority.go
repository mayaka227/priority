package priority

import (
	"golang.org/x/sys/windows"
)

const (
	ABOVE    uint32 = 0x00008000
	BELOW    uint32 = 0x00004000
	HIGH     uint32 = 0x00000080
	IDLE     uint32 = 0x00000040
	NORMAL   uint32 = 0x00000020
	REALTIME uint32 = 0x00000100
)

var (
	kernel32                            *windows.LazyDLL
	getCurrentProcess, setPriorityClass *windows.LazyProc
)

func init() {
	kernel32 = windows.NewLazySystemDLL("Kernel32.dll")
	getCurrentProcess = kernel32.NewProc("GetCurrentProcess")
	setPriorityClass = kernel32.NewProc("SetPriorityClass")
}

func _getCurrentProcess() (hProcess uintptr) {
	hProcess, _, _ = getCurrentProcess.Call()
	return
}

func _setPriorityClass(hProcess uintptr, dwPriorityClass uint32) (err error) {
	r1, _, err := setPriorityClass.Call(hProcess, uintptr(dwPriorityClass))
	if int32(r1) == 0 {
		return
	}
	return nil
}

func Set(dwPriorityClass uint32) (err error) {
	return _setPriorityClass(_getCurrentProcess(), dwPriorityClass)
}
