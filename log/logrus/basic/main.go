package main

import "github.com/sirupsen/logrus"

func main() {
	// 全局单例
	logrus.Info("this is a message #1")
	logrus.Warningf("this is a %s #2", "message")
	logrus.Errorf("this is a message %s", "#3")

	// 新建Logger实例
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	logger.Info("this is a message #4")
	logger.SetLevel(logrus.PanicLevel)
	logger.Info("this is a message #5") //不会输出
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.InfoLevel)
	logger.Info("this is a message #6") //输出调用者信息

	logger.WithField("name", "cool-learn"). // 结构化日志
		WithField("author", "apodemakeles").
		Info("this is a message #7")
}
