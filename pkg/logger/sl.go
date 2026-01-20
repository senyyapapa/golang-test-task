package sl

import (
	"log/slog"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func Msg(msg string) slog.Attr {
	return slog.Attr{
		Key:   "message",
		Value: slog.StringValue(msg),
	}
}
