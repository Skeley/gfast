// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
)

type (
	IFile interface {
		Upload(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error)
	}
)

var (
	localFile IFile
)

func File() IFile {
	if localFile == nil {
		panic("implement not found for interface IFile, forgot register?")
	}
	return localFile
}

func RegisterFile(i IFile) {
	localFile = i
}
