package domain

import "mime/multipart"

//db struct
type Upload struct {
	Id         uint   `gorm:"column:ID"`
	DocTypeId  int    `gorm:"column:DOC_TYPE_ID"`
	GroupId    int    `gorm:"column:GROUP_ID"`
	FileName   string `gorm:"column:FILE_NAME"`
	UploadTime string `gorm:"column:UPLOAD_TIME"`
}

type actionLog struct {
	Id            uint   `gorm:"column:ID"`
	TransactionId string `gorm:"column:TRANS_ID"`
	ActionType    string `gorm:"column:ACTION_TYPE"`
	FileId        int    `gorm:"column:FILE_ID"`
	TimeStamp     string `gorm:"column:TIMESTAMP"`
}

// req
type RequestGetUploadLists struct {
	TransactionId string
}

type RequestUploadFile struct {
	TransactionId string
	File          *multipart.FileHeader
	Filename      string
	DocTypeId     int
	GroupId       int
	FinalPath     string
}

type RequestDownloadFile struct {
	TransactionId string
	Filename      string
	DocTypeId     int
	GroupId       int
}

type RequestInsertLog struct {
	ActionLog actionLog
}

// res
type Response struct {
	TransactionId string `json:"transactionId"`
	Message       string `json:"msg"`
	Code          string `json:"code"`
}

type ResponseGetUploadLists struct {
	TransactionId string   `json:"transactionId"`
	Message       string   `json:"msg"`
	Code          string   `json:"code"`
	UploadLists   []Upload `json:"uploadLists"`
}

type ResponseUploadFile struct {
	TransactionId string   `json:"transactionId"`
	Message       []string `json:"msg"`
	Code          []string `json:"code"`
	Filepath      []string `json:"filepath"`
}

type ResponseInsertLog struct {
	Id uint `gorm:"column:ID"`
}
