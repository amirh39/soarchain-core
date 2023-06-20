package version

import (
	"fmt"
	"runtime"
)

// Build Info (set via linker flags)
var (
	AppVersion    = "v1.0.-2023"
	GitCommit     = "ge3089bea6"
	GitCommitDate = ""
)

func Version() string {
	return fmt.Sprintf(
		`Version: %s
Git Commit: %s
Git Commit Date: %s
Architecture: %s
Go Version: %s
Operating System: %s`,
		AppVersion,
		GitCommit,
		GitCommitDate,
		runtime.GOARCH,
		runtime.Version(),
		runtime.GOOS,
	)
}
