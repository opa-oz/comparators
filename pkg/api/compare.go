package api

import (
	"comparators/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/opa-oz/hikaku"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// Resp If you put lower-case property, it won't be returned, because it's `private`
type Resp struct {
	Response bool    `json:"response"`
	Diff     float64 `json:"diff"`
	ID       int     `json:"id"`
}

// @BasePath /api

// ComparePreview godoc
// @Summary compare preview from S3 with preview from site
// @Schemes
// @Param type query string true "Type"
// @Param id query int true "ID"
// @Description Compare preview from S3 with preview from site
// @Tags api
// @Accept json
// @Produce json
// @Success 200 {object} api.Resp
// @Router /compare/preview [get]
func ComparePreview(c *gin.Context) {
	ctx := c.Request.Context()
	entityType := c.Query("type")
	entityID := c.Query("id")

	res := c.Writer

	if entityType == "" || !(entityType == "anime" || entityType == "manga") {
		http.Error(res, "Empty entityType provided", http.StatusBadRequest)
		return
	}
	entityIDInt, err := strconv.Atoi(entityID)
	if entityID == "" || err != nil {
		http.Error(res, "Empty entityID provided", http.StatusBadRequest)
		return
	}

	minioClient, err := utils.GetMinio(c)
	cfg, err := utils.GetConfig(c)

	prefix := "output/" + entityType + "/"

	if err != nil {
		http.Error(res, "Bad response value: "+err.Error(), http.StatusInternalServerError)
		return
	}

	dir, err := os.MkdirTemp("", entityType)
	if err != nil {
		http.Error(res, "Cannot create dir: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			http.Error(res, "Failed to clean /tmp: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}(dir)

	objectKey := fmt.Sprintf("%d.jpg", entityIDInt)
	path := filepath.Join(dir, objectKey)

	err = minioClient.FGetObject(ctx, cfg.BucketName, filepath.Join(prefix, objectKey), path, minio.GetObjectOptions{})
	if err != nil {
		http.Error(res, "Bad response value: "+err.Error(), http.StatusInternalServerError)
		return
	}

	secondPath := filepath.Join(dir, "compare-"+objectKey)
	err = utils.DownloadFile(secondPath, utils.GetUrl(entityType, entityIDInt, cfg.UrlPrefix))
	if err != nil {
		//http.Error(res, "Bad response value: "+err.Error(), http.StatusInternalServerError)
		response := Resp{Response: false, Diff: 1, ID: entityIDInt}

		res.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(res).Encode(response)
		if err != nil {
			http.Error(res, "Bad response value: "+err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	img1, err := utils.Open(path)
	if err != nil {
		http.Error(res, "Can't open S3 image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	img2, err := utils.Open(secondPath)
	if err != nil {
		http.Error(res, "Can't open CDN image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	compParams := hikaku.ComparisonParameters{BinsCount: 16, Threshold: 0.2}
	isImagesEqual, diff := hikaku.Compare(img1, img2, compParams)

	response := Resp{Response: isImagesEqual, Diff: diff, ID: entityIDInt}

	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		http.Error(res, "Bad response value: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
