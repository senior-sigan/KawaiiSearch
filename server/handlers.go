package server

import (
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"kawaii_search/images"
	"net/http"
)

func listFilesHandler(c echo.Context) error {
	hs, err := getHashStorage(c)
	if err != nil {
		return err
	}

	files, err := hs.GetFiles()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, files)
}

func searchByFile(c echo.Context) error {
	searcher, err := getSearcher(c)
	if err != nil {
		return err
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	defer src.Close()
	imageRaw, err := ioutil.ReadAll(src)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	img, err := images.DecodeImage(imageRaw)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	similar, err := searcher.Find(img)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, similar)
}
func searchByUrl(c echo.Context) error {
	return fmt.Errorf("Not implemented")
}
