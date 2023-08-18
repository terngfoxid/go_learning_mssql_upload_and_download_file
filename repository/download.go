package repository

import (
	"go-mssql-upload/Config"
	_Domain "go-mssql-upload/domain"
	"strconv"
)

func DownloadFileLog(req _Domain.RequestDownloadFile) (res _Domain.Response, err error) {
	dbRes := new(_Domain.ResponseInsertLog)
	tsql := "select ID from upload where (DOC_TYPE_ID = " + strconv.Itoa(req.DocTypeId) + ") and (GROUP_ID = " + strconv.Itoa(req.GroupId) + ") and (FILE_NAME ='" + req.Filename + "')"
	result, err := Config.DB.Raw(tsql).Rows()
	if err != nil {
		res.Message = "Server can't get File Id"
		res.Code = "500"
		err = nil
	}
	Config.DB.ScanRows(result, &dbRes.Id)

	tsql = "insert_log"
	_, err = Config.DB.Raw(tsql, req.TransactionId, "download", dbRes.Id).Rows()
	if err != nil {
		res.Message = "Server can't insert download log to DB"
		res.Code = "500"
		return res, nil
	}

	return res, err
}

func SearchById(FileId string) (res _Domain.Upload, err error) {
	dbRes := new(_Domain.ResponseGetUploadLists)
	tsql := "select * from upload where Id =" + FileId
	result, err := Config.DB.Raw(tsql).Rows()
	if err != nil {
		return res, err
	}
	Config.DB.ScanRows(result, &dbRes.UploadLists)
	if len(dbRes.UploadLists) > 1 {
		res = dbRes.UploadLists[1]
	} else {
		res.DocTypeId = -1
	}

	return res, err
}

func DownloadFileLogById(transId string, FileId string) (res _Domain.Response, err error) {

	tsql := "insert_log"
	_, err = Config.DB.Raw(tsql, transId, "download", FileId).Rows()
	if err != nil {
		res.Message = "Server can't insert download log to DB"
		res.Code = "500"
		return res, nil
	}

	return res, err
}
