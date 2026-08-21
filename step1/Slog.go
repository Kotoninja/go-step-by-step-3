package main

import (
	"log/slog"
)

func LogUserAction(logger *slog.Logger, user string, action string) {
	logger.Info("user action", "User", user, "Action", action)
}
