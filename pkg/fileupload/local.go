package fileupload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/lightsaid/hotel-bookings/pkg/random"
	"github.com/lightsaid/hotel-bookings/pkg/toolkit"
)

type LocalUploader struct {
	saveDir   string   // 保持文件地址
	allowExts []string // 允许上传类型
	maxMB     int      // 文件大小限制，单位：MB
}

var once sync.Once

var Local *LocalUploader

// NewLocalUploader 实现一个上传文件到本地 FileUploader 接口；
// saveDir 上传文件的目录；maxMB 文件最大限制, 默认是3MB, 超过30MB则设置30MB；allowEtxs 允许上传类型，如果没有则任意类型；
func NewLocalUploader(saveDir string, maxMB int, allowEtxs ...string) FileUploader {
	once.Do(func() {
		if !strings.HasSuffix(saveDir, "/") {
			saveDir = saveDir + "/"
		}

		if maxMB > 30 {
			maxMB = 30
		}

		if maxMB <= 0 {
			maxMB = 3
		}

		Local = &LocalUploader{
			saveDir:   saveDir,
			allowExts: allowEtxs,
			maxMB:     maxMB,
		}
	})

	return Local
}

// SaveFile 上传（保存）文件到本地
func (uploader *LocalUploader) SaveFile(fileHeader *multipart.FileHeader) (*UploadedFile, error) {
	var uploadedFile = new(UploadedFile)

	// 设置部分返回信息
	uploadedFile.FileSize = fileHeader.Size
	uploadedFile.OrgFileName = fileHeader.Filename

	maxBytes := int64(uploader.maxMB << 20)
	// 判断上传文件大小
	if fileHeader.Size > maxBytes {
		curMB := float64(fileHeader.Size) / 1024 / 1024
		return uploadedFile, fmt.Errorf("%w: 最大限制 %d MB, 当前 %f MB", ErrFileTooLarge, uploader.maxMB, curMB)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return uploadedFile, err
	}
	defer file.Close()

	allowed := false
	fileExt := filepath.Ext(fileHeader.Filename)
	if len(uploader.allowExts) > 0 {
		for _, ext := range uploader.allowExts {
			if strings.EqualFold(fileExt, ext) {
				allowed = true
			}
		}
	} else {
		allowed = true
	}

	if !allowed {
		return uploadedFile, fmt.Errorf("%w: %s", ErrrNotAllowExt, fileExt)
	}

	// 将名字哈希化
	newName, err := toolkit.TKs.MD5(fileHeader.Filename, []byte(random.RandomString(8))...)
	if err != nil {
		return uploadedFile, err
	}

	// 重新构造文件名
	filename := newName + random.RandomString(8) + fileExt

	uploadedFile.NewFileName = filename

	filename = uploader.saveDir + filename

	out, err := os.Create(filename)
	if err != nil {
		return uploadedFile, err
	}

	_, err = io.Copy(out, file)
	return uploadedFile, err
}
