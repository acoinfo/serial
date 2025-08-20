package serial

import (
	"syscall"
)

func cfSetIspeed(termios *syscall.Termios, speed uint32) {
	// SylixOS has no Ispeed field in termios.
}

func cfSetOspeed(termios *syscall.Termios, speed uint32) {
	// SylixOS has no Ospeed field in termios.
}
