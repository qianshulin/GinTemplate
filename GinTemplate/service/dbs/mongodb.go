package dbs

import (
	"github.com/qiniu/qmgo/options"
	"github.com/yqboy/qmg"
)

var Mongo *qmg.Model

func InitMongo(url, database string, opts ...options.ClientOptions) {
	Mongo = qmg.MustNewModel(url, database, opts...)
}
