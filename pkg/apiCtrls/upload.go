package apiCtrls

import (
	"github.com/hassannmoussaa/pill.go/uploader"
	"github.com/valyala/fasthttp"
)

type UploadCtrl struct {
	*APICtrl
}

func (this *UploadCtrl) Upload(requestCtx *fasthttp.RequestCtx) {
	if multipartForm, err := requestCtx.MultipartForm(); err == nil {
		multipleUpload := uploader.MultipleUpload{WithCrop: true, FormData: multipartForm, FilesInputName: "files", FileType: "image|document|pdf", ImageCategory: "file", ImageSizes: []string{"small", "medium"}}
		multipleUpload.SetUploadDir("files")
		var files []string
		err, files = multipleUpload.Upload()
		if err == nil {
			if files != nil && len(files) > 0 {
				for i, _ := range files {
					files[i] = multipleUpload.UrlOfFile(files[i])
				}
				this.Success(requestCtx, files, "files_was_uploaded_successfully")
			} else {
				this.ValidationError(requestCtx, nil, "please_select_a_file")
			}
		} else {
			this.ValidationError(requestCtx, nil, err.Error())
		}
	} else {
		this.ServerError(requestCtx, "server_error")
	}
}
