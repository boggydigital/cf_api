package cf_trace

import (
	"github.com/boggydigital/cf_api/urls"
	"net/url"
)

func TraceUrl() *url.URL {
	return &url.URL{
		Scheme: urls.HTTPS,
		Host:   urls.WwwCloudflareHost,
		Path:   urls.TracePath,
	}
}
