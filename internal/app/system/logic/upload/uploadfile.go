package upload

import "github.com/gogf/gf/v2/net/ghttp"

func Upload(r *ghttp.Request) {
	files := r.GetUploadFiles("upload-file")
	names, err := files.Save("/tmp/")
	if err != nil {
		r.Response.WriteExit(err)
	}
	r.Response.WriteExit("upload successfully: ", names)
}
