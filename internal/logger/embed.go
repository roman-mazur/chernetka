package logger

type LogEmbed struct {
	LogDebug bool

	logf Func
}

func Embed(logf Func) LogEmbed {
	return LogEmbed{logf: logf}
}

func (le *LogEmbed) Logf(format string, args ...any) {
	if le.logf == nil {
		return
	}
	le.logf(format, args...)
}

func (le *LogEmbed) Debugf(fmt string, args ...any) {
	if !le.LogDebug {
		return
	}
	le.Logf(fmt, args...)
}

func (le *LogEmbed) EmbeddedLogger() *LogEmbed { return le }
