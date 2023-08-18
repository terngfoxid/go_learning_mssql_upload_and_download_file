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

func UploadFile(req _Domain.RequestUploadFile) (res _Domain.Response, err error) {
	res.TransactionId = req.TransactionId
	res, err = _Repository.UploadFile(*&req)
	if err != nil {
		res.Message = err.Error()
		res.Code = "500"
		return res, err
	}
	res.Message = "Ok"
	res.Code = "200"
	return res, err
}
