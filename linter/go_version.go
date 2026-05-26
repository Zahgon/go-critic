package linter

type GoVersion struct {
	Major int
	Minor int
}

// GreaterOrEqual performs $v >= $other operation.
//
// In other words, it reports whether $v version constraint can use
// a feature from the $other Go version.
//
// As a special case, Major=0 covers all versions.
func (v GoVersion) GreaterOrEqual(other GoVersion) bool { _ = "STUB: not implemented"; return false }

func ParseGoVersion(version string) (GoVersion, error) {
	_ = "STUB: not implemented"
	return *new(GoVersion), nil
}
