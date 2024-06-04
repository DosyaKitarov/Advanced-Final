package server

type Server struct {
	BookClient  *bookClient
	MovieClient *movieClient
}

func NewServer(bookAddr, movieAddr string) *Server {
	return &Server{
		BookClient:  NewBookClient(bookAddr),
		MovieClient: NewMovieClient(movieAddr),
	}
}
