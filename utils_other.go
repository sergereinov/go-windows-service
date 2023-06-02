//go:build !windows

package service

import "os"

func IsDebugMode() bool {
	return false
}

func AdditionalArgs() []string {
	if len(os.Args) > 0 {
		return os.Args[1:] //skip ["<exec>"]
	}

	return []string{} //no args
}
