package logger

// Func represents a basic function that can be used for logging: it matching log.Printf and fmt.Printf.
type Func func(format string, args ...any)

// Prefix returns a logging Func that logs with a predefined prefix.
func Prefix(f Func, prefix string) Func {
	return func(format string, args ...any) { f(prefix+format, args...) }
}

// Discard does nothing.
func Discard(_ string, _ ...any) {}
