package file

import (
	"context"
	"fmt"
	"os"
	"testing"

	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
)

func Test_sFile_Upload(t *testing.T) {
	b, _ := os.ReadFile("/Users/skyler/Documents/LFS.pdf")

	req := &v1.UploadFileReq{}
	req.Content = b

	ctx := context.Background()
	s := New()
	res, err := s.Upload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.URL)
}
