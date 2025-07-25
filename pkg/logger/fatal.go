package logger

func (l logger) Fatal(msg string, props ...ExtraProps) {
	l.log.Fatal(msg, toZapFields(props...)...)
}
