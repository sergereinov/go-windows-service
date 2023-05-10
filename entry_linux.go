//go:build !windows

package service

import (
	"context"
	"os/signal"
	runtimeDebug "runtime/debug"
	"syscall"
)

/**************
 	Linux:
	For service's operations please use systemd service manager

	Install:
	sudo systemclt enable my-svc.service
	sudo systemctl start my-svc.service

	Uninstall:
	sudo systemctl stop my-svc.service
	sudo systemclt disable my-svc.service

--- /lib/systemd/system/my-svc.service sample file ---
[Unit]
Description=My service
ConditionPathExists=/home/user/my-svc/my-svc
After=network.target

[Service]
Type=simple
User=user
Group=user
LimitNOFILE=1024

Restart=always
RestartSec=5

WorkingDirectory=/home/user/my-svc
ExecStart=/home/user/my-svc/my-svc

[Install]
WantedBy=multi-user.target

*********************/

func (s Service) entryPoint(payload func(context.Context)) {
	defer func() {
		if x := recover(); x != nil {
			s.logger.Fatalf("panic: %v\n%v", x, string(runtimeDebug.Stack()))
		}
	}()

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	defer cancel()

	payload(ctx)
}
