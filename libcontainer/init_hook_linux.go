package libcontainer

import (
	"os"

	"github.com/opencontainers/runc/libcontainer/configs"
)

// InitHook, if set, runs in a container's init process once its root
// filesystem is final, and before the init drops capabilities and applies
// seccomp. It runs on the thread that later execs the container's
// process. keep are host files the init still opens through /proc after
// the hook (the exec fifo). Processes started in a running container do
// not run it. An error fails the container's start.
//
// It is for runtimes that embed libcontainer and confine the container's
// process in ways the OCI spec cannot express (Landlock, for one): set it
// before calling Init.
var InitHook func(c *configs.Config, keep []*os.File) error
