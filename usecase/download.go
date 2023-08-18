package usecase

import (
	_Domain "go-mssql-upload/domain"
	_Repository "go-mssql-upload/repository"
)

func DownloadFileLog(req _Domain.RequestDownloadFile) (res _Domain.Response, err error) {
	res, err = _Repository.DownloadFileLog(req)
	if err != nil {
		res.Message = err.Error()
		res.Code = "500"
		return res, err
	}
	res.Message = "Ok"
	res.Code = "200"
	return res, err
}

func SearchById(FileId string) (res _Domain.Upload, err error) {
	res, err = _Repository.SearchById(FileId)
	return res, err
}

func DownloadFileById(transId string, FileId string) (res _Domain.Response, err error) {
	res, err = _Repository.DownloadFileLogById(transId, FileId)
	return res, err
}
