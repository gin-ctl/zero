package bootstrap

import (
	"zero/package/get"
	"zero/package/logger"
)

func SetupLogger() {
	logger.InitLogger(
		get.String("log.filename"),
		get.Int("log.max_size"),
		get.Int("log.max_backup"),
		get.Int("log.max_age"),
		get.Bool("log.compress"),
		get.String("log.type"),
		get.String("log.level"),
	)
}
