package logger

func (l logger) Warn(msg string, props ...ExtraProps) {
	l.log.Warn(msg, toZapFields(props...)...)
}
