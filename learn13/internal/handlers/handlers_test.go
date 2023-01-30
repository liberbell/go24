package handlers

type postData struct {
	key   string
	value string
}

var theTests []struct {
	name   string
	url    string
	method string
	params []postData
}
