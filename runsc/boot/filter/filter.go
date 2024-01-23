// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package filter defines all syscalls the sandbox is allowed to make
// to the host, and installs seccomp filters to prevent prohibited
// syscalls in case it's compromised.
package filter

import (
	"github.com/ttpreport/gvisor-ligolo/pkg/log"
	"github.com/ttpreport/gvisor-ligolo/pkg/seccomp"
	"github.com/ttpreport/gvisor-ligolo/pkg/sentry/devices/accel"
	"github.com/ttpreport/gvisor-ligolo/pkg/sentry/devices/nvproxy"
	"github.com/ttpreport/gvisor-ligolo/pkg/sentry/platform"
)

// Options are seccomp filter related options.
type Options struct {
	Platform              platform.Platform
	HostNetwork           bool
	HostNetworkRawSockets bool
	HostFilesystem        bool
	ProfileEnable         bool
	NVProxy               bool
	TPUProxy              bool
	ControllerFD          int
}

// Install seccomp filters based on the given platform.
func Install(opt Options) error {
	s := allowedSyscalls
	s.Merge(controlServerFilters(opt.ControllerFD))

	// Set of additional filters used by -race and -msan. Returns empty
	// when not enabled.
	s.Merge(instrumentationFilters())

	if opt.HostNetwork {
		if opt.HostNetworkRawSockets {
			Report("host networking (with raw sockets) enabled: syscall filters less restrictive!")
		} else {
			Report("host networking enabled: syscall filters less restrictive!")
		}
		s.Merge(hostInetFilters(opt.HostNetworkRawSockets))
	}
	if opt.ProfileEnable {
		Report("profile enabled: syscall filters less restrictive!")
		s.Merge(profileFilters())
	}
	if opt.HostFilesystem {
		Report("host filesystem enabled: syscall filters less restrictive!")
		s.Merge(hostFilesystemFilters())
	}
	if opt.NVProxy {
		Report("Nvidia GPU driver proxy enabled: syscall filters less restrictive!")
		s.Merge(nvproxy.Filters())
	}
	if opt.TPUProxy {
		Report("TPU device proxy enabled: syscall filters less restrictive!")
		s.Merge(accel.Filters())
	}

	s.Merge(opt.Platform.SyscallFilters())

	return seccomp.Install(s, seccomp.DenyNewExecMappings)
}

// Report writes a warning message to the log.
func Report(msg string) {
	log.Warningf("*** SECCOMP WARNING: %s", msg)
}
