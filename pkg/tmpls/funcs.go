package tmpls

import "net/url"

func GetStaticFileURLPath(path string, opts ...string) string {
	Url, err := url.Parse(staticFilesURLPath)
	if err == nil {
		Url.Path += path
		parameters := url.Values{}
		if opts != nil {
			if len(opts) > 0 {
				parameters.Add("app_version", opts[0])
			}
		}
		Url.RawQuery = parameters.Encode()
		return Url.String()
	}
	return ""
}

func GetUploadURLPath(path string, opts ...string) string {
	Url, err := url.Parse(uploadsURLPath)
	if err == nil {
		Url.Path += path
		parameters := url.Values{}
		if opts != nil {
			if len(opts) > 0 {
				parameters.Add("app_version", opts[0])
			}
		}
		Url.RawQuery = parameters.Encode()
		return Url.String()
	}
	return ""
}
