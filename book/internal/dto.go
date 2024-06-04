package bookinfo

type Book struct {
	ID     uint32 "json:\"_id\" bson:\"_id\""
	Title  string "json:\"title\" bson:\"title\""
	Author string "json:\"author\" bson:\"author\""
}
