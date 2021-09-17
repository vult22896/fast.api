package services

import "fast.bibabo.vn/database"

var caching = database.GetInstanceRedis().Caching()
var db = database.GetInstanceMysql().Connect()
