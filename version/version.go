package version

import (
	"fmt"
	"runtime"
)

// Build Info
var (
	AppVersion = "0.1.0"
	AppName    = "soarchain"
)

func Version() string {
	return fmt.Sprintf(
		`Version: %s
AppName: %s
Architecture: %s
Go Version: %s
Operating System: %s`,
		AppVersion,
		AppName,
		runtime.GOARCH,
		runtime.Version(),
		runtime.GOOS,
	)
}
