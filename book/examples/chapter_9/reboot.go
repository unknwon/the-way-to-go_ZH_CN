// reboot.go
// compile errors (Windows):
//undefined: syscall.SYS_REBOOT
// reboot.go:13: not enough arguments in call to syscall.Syscall
// Linux: compileert, uitvoeren met sudo ./6.out --> systeem herstart
package main

import (
	"syscall"
)

const LINUX_REBOOT_MAGIC1 uintptr = 0xfee1dead
const LINUX_REBOOT_MAGIC2 uintptr = 672274793
const LINUX_REBOOT_CMD_RESTART uintptr = 0x1234567

func main() {
	syscall.Syscall(syscall.SYS_REBOOT,
		LINUX_REBOOT_MAGIC1,
		LINUX_REBOOT_MAGIC2,
		LINUX_REBOOT_CMD_RESTART)
}
