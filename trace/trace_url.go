package trace

import (
	"github.com/boggydigital/cf_api/urls"
	"net/url"
)

func Url() *url.URL {
	return &url.URL{
		Scheme: urls.HTTPS,
		Host:   urls.WwwCloudflareHost,
		Path:   urls.TracePath,
	}
}
