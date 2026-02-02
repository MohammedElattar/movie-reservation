package logger

type LogField struct {
	Key   string `json:"key"`
	Value any `json:"value"`
}

func String(key, value string) LogField {
	return LogField{Key: key, Value: value}
}

func Int(key string, value int) LogField {
	return LogField{Key: key, Value: value}
}

func Error(err error) LogField {
	if err == nil {
		return LogField{
			Key:   "error",
			Value: nil,
		}
	}

	return LogField{
		Key: "error",
		Value: err.Error(),
	}
}
