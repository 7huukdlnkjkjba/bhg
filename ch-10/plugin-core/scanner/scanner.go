package scanner

// Scanner定义一个所有检查都遵循的接口
type Checker interface {
	Check(host string, port uint64) *Result
}

// Result定义检查的结果
type Result struct {
	Vulnerable bool
	Details    string
}
