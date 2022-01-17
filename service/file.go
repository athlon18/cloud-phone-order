package service

import (
	"io/ioutil"
	"mime/multipart"
	"order/db"
	"order/model"
	"os"
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
	if err := db.DB.Model(model.File{}).Find(&files).Error; err != nil {
		return
	}
	for index, item := range files {
		if _, err := os.Stat(item.File); err == nil {
			files[index].IsExist = true
		}
	}
	return files
}

func GetFile() (file model.File) {
	if err := db.DB.Model(model.File{}).Where("status = 1").First(&file).Error; err != nil {
		return
	}
	return file
}

func FileList() (list, all []interface{}) {
	files, _ := ioutil.ReadDir("./upload")
	for _, f := range files {
		all = append(all, f.Name())
		if f.IsDir() {
			list = append(list, f.Name())
		}
	}
	return
}

func DeleteFile(id string) error {
	return db.DB.Where("id =?", id).Delete(new(model.File)).Error
}
