package models

type Page struct{
	Current int `json:"current""`  //当前页
	Size int `json:"size"`//每页条数
	Total int `json:"total"` //总条数
	Pages int `json:"pages"`//总页数
	Records interface{} `json:"records"`//结果集
}


