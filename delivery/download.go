package delivery

import (
	"fmt"
	_Domain "go-mssql-upload/domain"
	_Usecase "go-mssql-upload/usecase"
	"net/http"
	"os"
	"strconv"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DownloadFile(c *gin.Context) {
	transId := uuid.New().String()

	var req = new(_Domain.RequestDownloadFile)
	req.TransactionId = transId
	req.Filename = c.Param("fileName")
	req.DocTypeId, _ = strconv.Atoi(c.Param("docTypeId"))
	req.GroupId, _ = strconv.Atoi(c.Param("groupId"))

	prefixPath := fmt.Sprintf("./FileSave/%s/%s", c.Param("docTypeId"), c.Param("groupId"))
	filepath := fmt.Sprintf("%s/%s", prefixPath, req.Filename)
	if _, err := os.Stat(prefixPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, _Domain.Response{
			TransactionId: transId,
			Message:       "[Not Found Error] Dir not Found",
			Code:          "404",
		})
		return
	}

	entries, err := os.ReadDir(prefixPath)
	if err != nil || len(entries) == 0 {
		c.JSON(http.StatusBadRequest, _Domain.Response{
			TransactionId: transId,
			Message:       "[Internal Server Error] Server can't read any files in Dir",
			Code:          "500",
		})
		return
	}

	for _, entry := range entries {
		if entry.Name() == req.Filename {

			_, err := _Usecase.DownloadFileLog(*req)
			if err != nil {
				//log insert error
				err = nil
			}

			mtype, err := mimetype.DetectFile(filepath)

			if err != nil {
				c.JSON(http.StatusInternalServerError, _Domain.Response{
					TransactionId: transId,
					Message:       "[Internal Server Error]: Server can't read File",
					Code:          "500",
				})
				return
			}

			c.Header("Content-type", mtype.String()+"; charset=utf-8")
			c.Header("Content-Disposition", "attachment; filename="+entry.Name())

			c.File(prefixPath + "/" + entry.Name())
			return
		}
	}

	c.JSON(http.StatusNotFound, _Domain.Response{
		TransactionId: transId,
		Message:       "[Not Found Error] File not Found",
		Code:          "404",
	})
}

func DownloadFileById(c *gin.Context) {
	transId := uuid.New().String()

	var req = new(_Domain.RequestDownloadFile)
	req.TransactionId = transId

	FileId := c.Param("fileId")

	res, err := _Usecase.SearchById(FileId)

	if err != nil {
		c.JSON(http.StatusNotFound, _Domain.Response{
			TransactionId: transId,
			Message:       "[Not Found Error] Search File by Id not Found",
			Code:          "404",
		})
		return
	}

	if res.DocTypeId == -1 {
		c.JSON(http.StatusNotFound, _Domain.Response{
			TransactionId: transId,
			Message:       "[Not Found Error] Search File by Id not Found",
			Code:          "404",
		})
		return
	}

	prefixPath := fmt.Sprintf("./FileSave/%s/%s", strconv.Itoa(res.DocTypeId), strconv.Itoa(res.GroupId))
	filepath := fmt.Sprintf("%s/%s", prefixPath, res.FileName)
	if _, err := os.Stat(prefixPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, _Domain.Response{
			TransactionId: transId,
			Message:       "[Not Found Error] Dir not Found",
			Code:          "404",
		})
		return
	}

	entries, err := os.ReadDir(prefixPath)
	if err != nil || len(entries) == 0 {
		c.JSON(http.StatusBadRequest, _Domain.Response{
			TransactionId: transId,
			Message:       "[Internal Server Error] Server can't read any files in Dir",
			Code:          "500",
		})
		return
	}

	for _, entry := range entries {
		if entry.Name() == res.FileName {

			_, err := _Usecase.DownloadFileById(transId, FileId)
			if err != nil {
				//log insert error
				err = nil
			}

			mtype, err := mimetype.DetectFile(filepath)

			if err != nil {
				c.JSON(http.StatusInternalServerError, _Domain.Response{
					TransactionId: transId,
					Message:       "[Internal Server Error]: Server can't read File",
					Code:          "500",
				})
				return
			}

			c.Header("Content-type", mtype.String()+"; charset=utf-8")
			c.Header("Content-Disposition", "attachment; filename="+entry.Name())

			c.File(prefixPath + "/" + entry.Name())
			return
		}
	}

	c.JSON(http.StatusNotFound, _Domain.Response{
		TransactionId: transId,
		Message:       "[Not Found Error] File not Found",
		Code:          "404",
	})
}
