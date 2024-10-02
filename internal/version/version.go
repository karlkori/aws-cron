package version

import (
	"fmt"
	"runtime"
)

var (
	// These variables are replaced by ldflags at build time
	gitVersion = "v0.0.0-main"
	gitCommit  = ""
	buildDate  = "1970-01-01T00:00:00Z"
)

type VersionInfo struct {
	GitVersion string `json:"gitVersion" yaml:"gitVersion"`
	GitCommit  string `json:"gitCommit" yaml:"gitCommit"`
	BuildDate  string `json:"buildDate" yaml:"buildDate"`
	GoVersion  string `json:"goVersion" yaml:"goVersion"`
	Compiler   string `json:"compiler" yaml:"compiler"`
	Platform   string `json:"platform" yaml:"platform"`
}

func Get() *VersionInfo {
	// These variables typically come from -ldflags settings and in
	// their absence fallback to the constants above
	return &VersionInfo{
		GitVersion: gitVersion,
		GitCommit:  gitCommit,
		BuildDate:  buildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
