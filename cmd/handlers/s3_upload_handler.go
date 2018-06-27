package handlers

import (
	"fmt"
	"github.com/ducnt114/testprj/drivers/mongo"
	"github.com/ducnt114/testprj/entity"
	"github.com/ducnt114/testprj/utils"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// S3UploadHandler --
type S3UploadHandler struct {
	MongoConn *mongo.MongoConnection
}

type s3UploadResponse struct {
	Success  bool   `json:"success"`
	ImageURL string `json:"image_url"`
}

func (h *S3UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("Error: ", err)
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	// copy to local storage
	storageTmpDir := viper.GetString("storage.temp_dir")
	storageFileName := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
	storageFilePath := fmt.Sprintf("%s/%s", storageTmpDir, storageFileName)
	f, err := os.OpenFile(storageFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("Error: ", err)
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()
	io.Copy(f, file)

	// then upload to s3
	s3Region := viper.GetString("s3.region")
	s3ID := viper.GetString("s3.secret_id")
	s3Secret := viper.GetString("s3.secret_key")
	s3Token := viper.GetString("s3.secret_token")

	s3Client, err := utils.NewS3Client(s3Region, s3ID, s3Secret, s3Token)
	if err != nil {
		log.Println("Error when init aws s3 client, detail: ", err)
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	fileURL, err := s3Client.UploadFile(storageFilePath, viper.GetString("s3.bucket"))
	if err != nil {
		log.Println("Error when upload file to aws s3, detail: ", err)
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// store metadata to mongodb
	collection := h.MongoConn.Session.DB(viper.GetString("mongo.db")).C(viper.GetString("mongo.collection"))
	err = collection.Insert(&entity.FileUploadMetaData{
		FileName:  storageFileName,
		FileSize:  handler.Size,
		FileURL:   fileURL,
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		log.Println("Error when save metadata to mongo, detail: ", err)
	}

	// remove local file in temporary dir, ignore even if error occur
	os.Remove(storageFilePath)

	// return success
	utils.ResponseJSON(w, &s3UploadResponse{Success: true, ImageURL: fileURL})
}
