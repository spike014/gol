# gol

[![Go Report Card](https://goreportcard.com/badge/github.com/yaoming00/gol)](https://goreportcard.com/report/github.com/yaoming00/gol)
![CI Actions](https://github.com/spike014/gol/actions/workflows/main.yml/badge.svg??branch=master)
![Test Actions](https://github.com/spike014/gol/actions/workflows/test.yml/badge.svg??branch=master)
[![codecov](https://codecov.io/gh/spike014/gol/branch/master/graph/badge.svg?token=62JSDGNHB5)](https://codecov.io/gh/spike014/gol)

gol is log, log for gin and gorm.

- Rotating logs by date

---

go 1.16+

## Installation
```bash
go get github.com/yaoming00/gol 
```
## Usage
### define logger
```golang
var (
	ServerLogger io.Writer
	SQLLogger    io.Writer
	ErrorLogger  io.Writer
)
```

### log for gin
```golang
    runMode := "debug"
	gin.SetMode(runMode)
	gin.DisableConsoleColor()

	gin.DefaultWriter = gol.SetLogger(runMode, ServerLogger)
	gin.DefaultErrorWriter = gol.SetLogger(runMode, ServerLogger)
```

#### log error by log.Println()
```golang
// main
	log.SetOutput(
		gol.SetLogger(
			runMode,
			ServerLogger,
		),
	)
```

### log for gorm
```golang
func init() {
	SQLLogger = gol.NewLogger(${path + filename})
	getDB()
}

func getDB() {
	runMode := "debug"
	sqlLogWriter := gol.SetLogger(runMode, SQLLogger)

	newLogger := logger.New(
		log.New(sqlLogWriter, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.LogLevel(${logLevel}),
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	var err error
	db, err = gorm.Open(mysql.Open(config.DBLink), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Println("Get SQL connection ERROR:", err)
	}
}
```

## Thanks

- [file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs)
