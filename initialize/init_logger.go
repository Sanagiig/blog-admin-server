package initialize

import (
	"github.com/sirupsen/logrus"
	"go-blog/global"
	"go-blog/global/settings"
	"io"
	"os"
	"path"
	"strings"
)

func InitLogger() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	logDirStr := settings.LogDir
	if strings.HasPrefix(settings.LogDir, ".") {
		logDirStr = path.Join(pwd, logDirStr)
	}

	file, err := os.OpenFile(logDirStr, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
		return
	}

	multiWrite := io.MultiWriter(os.Stderr, file)

	global.Log = logrus.New()
	global.Log.SetOutput(multiWrite)
	global.Log.SetLevel(logrus.WarnLevel)
}
