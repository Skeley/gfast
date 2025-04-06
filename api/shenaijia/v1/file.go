/*
* @desc:文件上传
* @company:深爱家
* @Author:sk
* @Date:2025/03/11
 */

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UploadFileReq struct {
	g.Meta   `path:"/upload" tags:"文件管理" method:"post" summary:"上传文件"`
	FileName string `json:"fileName"`
	Content  []byte `json:"content"`
}

type UploadFileRes struct {
	g.Meta `mime:"application/json"`
	URL    string `json:"url"`
}
