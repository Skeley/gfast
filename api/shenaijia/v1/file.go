/*
* @desc:文件上传
* @company:深爱家
* @Author:sk
* @Date:2025/03/11
 */

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadFileReq struct {
	g.Meta `path:"/upload" mime:"multipart/form-data" method:"post" tags:"文件管理" summary:"上传文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件`
}

type UploadFileRes struct {
	g.Meta `mime:"application/json"`
	URL    string `json:"url"`
}
