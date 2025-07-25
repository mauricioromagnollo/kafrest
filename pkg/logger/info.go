package logger

func (l logger) Info(msg string, props ...ExtraProps) {
	l.log.Info(msg, toZapFields(props...)...)
}
