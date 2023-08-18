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

func UploadFile(req _Domain.RequestUploadFile) (res _Domain.Response, err error) {
	tsql := "insert_upload_record"

	dbRes := new(_Domain.ResponseInsertLog)

	result, err := Config.DB.Raw(tsql, req.DocTypeId, req.GroupId, req.Filename).Rows()
	if err != nil {
		res.Message = "Server can't insert record to DB"
		res.Code = "500"
		return res, err
	}

	Config.DB.ScanRows(result, &dbRes.Id)

	tsql = "insert_log"
	_, err = Config.DB.Raw(tsql, req.TransactionId, "upload", dbRes.Id).Rows()
	if err != nil {
		res.Message = "Server can't insert upload log to DB"
		res.Code = "500"
		return res, nil
	}

	return res, err
}
