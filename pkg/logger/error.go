package logger

func (l logger) Error(msg string, props ...ExtraProps) {
	l.log.Error(msg, toZapFields(props...)...)
}
