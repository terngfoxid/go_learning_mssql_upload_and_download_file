package delivery

import (
	"errors"
	"fmt"
	_Domain "go-mssql-upload/domain"
	_Usecase "go-mssql-upload/usecase"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUploadLists(c *gin.Context) {
	transId := uuid.New().String()
	req := new(_Domain.RequestGetUploadLists)
	req.TransactionId = transId

	res, err := _Usecase.GetUploadLists(*req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, _Domain.Response{
			TransactionId: res.TransactionId,
			Message:       res.Message,
			Code:          res.Code,
		})
		return
	}

	c.JSON(http.StatusOK, _Domain.ResponseGetUploadLists{
		TransactionId: res.TransactionId,
		Message:       res.Message,
		Code:          res.Code,
		UploadLists:   res.UploadLists,
	})
}

func UploadFile(c *gin.Context) {
	var messageRes []string
	var filepathRes []string
	var codeRes []string
	var fileIdRes []int

	transId := uuid.New().String()
	docTypeId, _ := strconv.Atoi(c.Param("docTypeId"))
	groupId, _ := strconv.Atoi(c.Param("groupId"))

	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, _Domain.Response{
			TransactionId: transId,
			Message:       "[Bad Request] Server can't receive any files",
			Code:          "400",
		})
		return
	}

	files := form.File["file"]

	for _, file := range files {
		var req = new(_Domain.RequestUploadFile)

		req.TransactionId = transId
		req.DocTypeId = docTypeId
		req.GroupId = groupId

		if file != nil && file.Size > 0 {
			req.File = file
			Filename := time.Now().Format("2006-01-02T15-04-05") + "-" + file.Filename
			req.Filename = Filename
			req.FinalPath, req.FileId, err = SaveFile(*req)

			if err != nil {
				messageRes = append(messageRes, "Upload File ["+file.Filename+"]: [Internal Server Error] "+err.Error())
				filepathRes = append(filepathRes, "Error")
				fileIdRes = append(fileIdRes, -1)
				codeRes = append(codeRes, "500")
				continue
			}

			messageRes = append(messageRes, "Upload File ["+file.Filename+"]: [Success]")
			filepathRes = append(filepathRes, req.FinalPath)
			fileIdRes = append(fileIdRes, req.FileId)
			codeRes = append(codeRes, "200")

		} else {
			messageRes = append(messageRes, "Upload File ["+file.Filename+"]: [Bad Request Error] File is null or zero")
			filepathRes = append(filepathRes, "Error")
			fileIdRes = append(fileIdRes, -1)
			codeRes = append(codeRes, "400")
		}
	}
	if messageRes == nil || codeRes == nil || filepathRes == nil {
		c.JSON(http.StatusBadRequest, _Domain.Response{
			TransactionId: transId,
			Message:       "[Bad Request] Server can't receive any files",
			Code:          "400",
		})
		return
	}

	c.JSON(http.StatusMultiStatus, _Domain.ResponseUploadFile{
		TransactionId: transId,
		Message:       messageRes,
		Code:          codeRes,
		Filepath:      filepathRes,
		FileId:        fileIdRes,
	})
}

// func for save file
func SaveFile(req _Domain.RequestUploadFile) (filePath string, fileId int, err error) {
	Path := strconv.Itoa(req.DocTypeId) + "/" + strconv.Itoa(req.GroupId)
	prefixPath := fmt.Sprintf("./FileSave/%s", Path)

	if _, err := os.Stat(prefixPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(prefixPath, os.ModePerm)
		if err != nil {
			return filePath, fileId, errors.New("Error Process Create File Path, " + err.Error())
		}
		err = nil
	}

	filePath = fmt.Sprintf("%s/%s", prefixPath, req.Filename)

	// Source File
	src, err := req.File.Open()
	if src != nil {
		defer func(src multipart.File) {
			err := src.Close()
			if err != nil {
				fmt.Println("Src Close Error.")
			}
		}(src)
	}

	if err != nil {
		return filePath, fileId, errors.New("Error Process Open File, " + err.Error())
	}

	// Destination File
	dst, err := os.Create(filePath)
	if dst != nil {
		defer func(dst *os.File) {
			err := dst.Close()
			if err != nil {
				fmt.Println("Dst Close Error.")
			}
		}(dst)
	}

	if err != nil {
		return filePath, fileId, errors.New("Error Process Create File, " + err.Error())
	}

	// Copy source file to destination file
	_, err = io.Copy(dst, src)
	if err != nil {
		return filePath, fileId, errors.New("Error Process Copy File, " + err.Error())
	}

	// Insert new record

	fileId, err = _Usecase.UploadFile(req)

	if err != nil {
		//remove that file record error
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			_ = os.Remove(filePath)
		}

		return filePath, fileId, errors.New("Error Process Insert Upload File Record, " + err.Error())
	}

	return filePath, fileId, nil
}
