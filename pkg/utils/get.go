package utils

import (
	"comparators/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func GetMinio(c *gin.Context) (*minio.Client, error) {
	r := c.Value("Minio")

	if r == nil {
		err := fmt.Errorf("could not retrieve Minio")
		return nil, err
	}

	rdb, ok := r.(*minio.Client)
	if !ok {
		err := fmt.Errorf("variable Minio has wrong type")
		return nil, err
	}

	return rdb, nil
}

func GetConfig(c *gin.Context) (*config.Environment, error) {
	r := c.Value("Config")

	if r == nil {
		err := fmt.Errorf("could not retrieve Config")
		return nil, err
	}

	rdb, ok := r.(*config.Environment)
	if !ok {
		err := fmt.Errorf("variable Config has wrong type")
		return nil, err
	}

	return rdb, nil
}

func GetUrl(entityType string, id int, urlPrefix string) string {
	if entityType == "manga" {
		return getMangaUrl(id, urlPrefix)
	}

	if entityType == "anime" {
		return getAnimeUrl(id, urlPrefix)
	}

	panic("OMG SOMETHING BAD")
}

func getMangaUrl(id int, urlPrefix string) string {
	return fmt.Sprintf("%s/manga/%d.jpg", urlPrefix, id)
}

func getAnimeUrl(id int, urlPrefix string) string {
	return fmt.Sprintf("%s/%d.jpg", urlPrefix, id)
}
