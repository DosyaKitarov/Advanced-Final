package movie

type Movie struct {
	ID       uint32  "json:\"_id\" bson:\"_id\""
	Title    string  "json:\"title\" bson:\"title\""
	Director string  "json:\"director\" bson:\"director\""
	Rating   float64 "json:\"rating\" bson:\"rating\""
}
