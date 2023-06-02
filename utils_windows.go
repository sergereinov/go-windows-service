//go:build windows

package service

import (
	"os"

	"golang.org/x/sys/windows/svc"
)

func IsDebugMode() bool {
	isInService, err := svc.IsWindowsService()
	if isInService || err != nil {
		return false
	}

	if len(os.Args) < 2 {
		return false
	}

	if os.Args[1] != "/d" {
		return false
	}

	return true
}

func AdditionalArgs() []string {
	if IsDebugMode() {
		return os.Args[2:] //skip ["<exec>", "/d"]
	}

	if len(os.Args) > 0 {
		return os.Args[1:] //skip ["<exec>"]
	}

	return []string{} //no args
}
