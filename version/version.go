package version

var (
	version = ""
	commit  = ""
)

func Version() string {
	return version
}

func Commit() string {
	return commit
}
