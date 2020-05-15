package server

import (
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"kawaii_search/files"
	"kawaii_search/searcher"
	"kawaii_search/storage"
	"net/http"
)

type Config struct {
	Address string
}

type Server struct {
	config        *Config
	echo          *echo.Echo
	hashesStorage *storage.Storage
	fileStorage   *files.FilesStorage
	searcher      *searcher.Searcher
}

func NewServer(config *Config, hashesStorage *storage.Storage, fileStorage *files.FilesStorage, searcher *searcher.Searcher) *Server {
	return &Server{
		config:        config,
		echo:          nil,
		hashesStorage: hashesStorage,
		fileStorage:   fileStorage,
		searcher:      searcher,
	}
}

func (s *Server) Start() {
	s.echo = echo.New()
	s.echo.HideBanner = true

	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Use(dependenciesMiddleware(s))

	s.echo.GET("/files/list", listFilesHandler)
	s.echo.POST("/search/file", searchByFile)
	s.echo.POST("/search/url", searchByUrl)

	s.echo.Logger.Fatal(s.echo.Start(s.config.Address))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func dependenciesMiddleware(s *Server) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("filesStorage", s.fileStorage)
			c.Set("hashesStorage", s.hashesStorage)
			c.Set("searcher", s.searcher)
			return next(c)
		}
	}
}

func getFileStorage(c echo.Context) (*files.FilesStorage, error) {
	m, ok := c.Get("filesStorage").(*files.FilesStorage)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Files Storage not found")
	}
	return m, nil
}

func getHashStorage(c echo.Context) (*storage.Storage, error) {
	m, ok := c.Get("hashesStorage").(*storage.Storage)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Hashes Storage not found")
	}
	return m, nil
}

func getSearcher(c echo.Context) (*searcher.Searcher, error) {
	m, ok := c.Get("searcher").(*searcher.Searcher)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Searcher not found")
	}
	return m, nil
}
