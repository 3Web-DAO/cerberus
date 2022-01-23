package models

import "github.com/go-xorm/xorm"
import _"github.com/go-sql-driver/mysql"

var Engine *xorm.Engine

func Init(dataSourceName string) {
    var err error
    Engine, err = xorm.NewEngine("mysql", dataSourceName)
    if err != nil {
        panic(err)
    }
}