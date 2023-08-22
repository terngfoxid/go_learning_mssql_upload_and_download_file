package repository

import (
	"go-mssql-upload/Config"
	_Domain "go-mssql-upload/domain"
)

func GetUploadLists() (res _Domain.ResponseGetUploadLists, err error) {
	tsql := "SELECT * FROM upload"
	//uploadList := new([]_Domain.upload)
	tp, err := Config.DB.Raw(tsql).Rows()
	if err != nil {
		res.Message = "[Internal Server Error] Server can't get data from DB"
		res.Code = "500"
		return res, err
	}
	Config.DB.ScanRows(tp, &res.UploadLists)
	if len(res.UploadLists) > 1 {
		res.UploadLists = res.UploadLists[1:]
	}
	return res, err
}

func UploadFile(req _Domain.RequestUploadFile) (fileId int, err error) {
	tsql := "insert_upload_record"

	result, err := Config.DB.Raw(tsql, req.DocTypeId, req.GroupId, req.Filename).Rows()
	if err != nil {
		return fileId, err
	}

	Config.DB.ScanRows(result, &fileId)

	tsql = "insert_log"
	_, err = Config.DB.Raw(tsql, req.TransactionId, "upload", fileId).Rows()
	if err != nil {
		return fileId, nil
	}

	return fileId, err
}
