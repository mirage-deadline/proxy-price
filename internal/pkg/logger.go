package pkg

import (
	"log/slog"
	"os"
	"time"
)

func customHandler() *slog.JSONHandler {
	return slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "time" {
				a.Value = slog.StringValue(time.Now().UTC().String())
			}
			return a
		},
	})
}

func NewLogger() *slog.Logger {
	return slog.New(customHandler())
}
