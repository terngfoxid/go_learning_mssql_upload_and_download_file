package usecase

import (
	_Domain "go-mssql-upload/domain"
	_Repository "go-mssql-upload/repository"
)

func GetUploadLists(req _Domain.RequestGetUploadLists) (res _Domain.ResponseGetUploadLists, err error) {
	res.TransactionId = req.TransactionId
	resData, err := _Repository.GetUploadLists()
	if err != nil {
		res.Message = resData.Message
		res.Code = resData.Code
	} else {
		res.Message = "[Success] Get Upload List Success"
		res.Code = "200"
		res.UploadLists = resData.UploadLists
	}
	return res, err
}

func UploadFile(req _Domain.RequestUploadFile) (fileId int, err error) {

	fileId, err = _Repository.UploadFile(req)
	if err != nil {
		return fileId, err
	}
	return fileId, err
}
