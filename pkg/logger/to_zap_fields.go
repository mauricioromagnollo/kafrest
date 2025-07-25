package logger

import "go.uber.org/zap"

func toZapFields(props ...ExtraProps) []zap.Field {
	fields := make([]zap.Field, 0, len(props))
	for _, p := range props {
		for k, v := range p {
			fields = append(fields, zap.Any(k, v))
		}
	}
	return fields
}
