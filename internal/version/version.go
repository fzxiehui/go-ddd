package version

var (
	// 通过 -ldflags 注入
	Version   = "dev"
	GitHash   = "none"
	BuildTime = "unknown"
)

type Info struct {
	Version   string `json:"version"`
	GitHash   string `json:"git_hash"`
	BuildTime string `json:"build_time"`
}

func Get() Info {
	return Info{
		Version:   Version,
		GitHash:   GitHash,
		BuildTime: BuildTime,
	}
}
