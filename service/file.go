package service

import (
	"mime/multipart"
	"order/db"
	"order/model"
)

func UploadFile(file *multipart.FileHeader, path string, status string) {
	parseBool := false
	if status == "on" {
		parseBool = true
	}
	if parseBool == true {
		db.DB.Model(model.File{}).Where("status = 1").Update("status", 0)
	}
	db.DB.Model(model.File{}).Create(&model.File{
		Name:   file.Filename,
		File:   path,
		Status: parseBool,
	})
}

func File() (files []model.File) {
	db.DB.Model(model.File{}).Find(&files)
	return files
}

func GetFile() (file model.File) {
	db.DB.Model(model.File{}).Where("status = 1").First(&file)
	return file
}
