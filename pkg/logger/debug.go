package logger

func (l logger) Debug(msg string, props ...ExtraProps) {
	l.log.Debug(msg, toZapFields(props...)...)
}
