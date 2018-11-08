package hello

import (
	"github.com/go-modules-by-example-forks/incomplete"
	"github.com/kr/pretty"
	"github.com/sirupsen/logrus"
)

const (
	IncompleteFieldKeyMsg = incomplete.FieldKeyMsg

	FieldKeyMsg = logrus.FieldKeyMsg
)

var (
	IncompleteLogger = incomplete.Logger
	Logger           pretty.Logfer
)
