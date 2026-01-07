package engine

import (
	"os/exec"
)

// CollectEventLogs stub: on Windows we can export system/application logs via wevtutil (requires privileges)
func CollectEventLogs() error {
	// example: wevtutil epl System system.evtx
	cmd := exec.Command("wevtutil", "epl", "System", "system.evtx")
	return cmd.Run()
}