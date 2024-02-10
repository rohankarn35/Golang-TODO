package components

import database "todo/Database"

type notes struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Date  string `json:"date"`
}

var connectionClient = database.Db().Database("test").Collection("test")
