package logger

type LogEmbed struct {
	LogDebug bool

	Func
}

func Embed(logf Func) LogEmbed {
	return LogEmbed{Func: logf}
}

func (le *LogEmbed) Logf(format string, args ...any) {
	if le.Func == nil {
		return
	}
	le.Func(format, args...)
}

func (le *LogEmbed) Debugf(fmt string, args ...any) {
	if !le.LogDebug {
		return
	}
	le.Logf(fmt, args...)
}

func (le *LogEmbed) EmbeddedLogger() *LogEmbed { return le }
