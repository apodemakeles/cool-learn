package main

import "github.com/sirupsen/logrus"


type service struct{
	logger logrus.FieldLogger
}

func (s *service) doSth(){
	s.logger.Infof("start to do..")
	// business code
}

func main() {
	logger := logrus.New()
	subLogger := logger.WithField("name", "member-A")
	svc := service{logger: subLogger}
	svc.doSth()
}
