package fileupload

import (
	"errors"
	"mime/multipart"
)

var (
	ErrrNotAllowExt = errors.New("不支持文件类型")
	ErrFileTooLarge = errors.New("文件太大")
)

// FileUploader 保存文件接口
type FileUploader interface {
	SaveFile(fileHeader *multipart.FileHeader) (*UploadedFile, error)
}

// UploadedFile 已经上传文件返回结果
type UploadedFile struct {
	OrgFileName string `json:"org_filename"` // 原文件名
	NewFileName string `json:"new_filename"` // 新文件名
	FileSize    int64  `json:"file_size"`    // 文件大小
}

func IsUploaderError(err error) bool {
	if errors.Is(err, ErrrNotAllowExt) || errors.Is(err, ErrFileTooLarge) {
		return true
	}
	return false
}
