package shared

import (
	"net/http"
	urlpkg "net/url"
)

func Url(r *http.Request, path string, params map[string]string) string {
	endpoint := r.URL.Path

	if params == nil {
		params = map[string]string{}
	}

	params["controller"] = path

	url := endpoint + query(params)

	return url
}

func query(queryData map[string]string) string {
	queryString := ""

	if len(queryData) > 0 {
		v := urlpkg.Values{}
		for key, value := range queryData {
			v.Set(key, value)
		}
		queryString += "?" + httpBuildQuery(v)
	}

	return queryString
}

func httpBuildQuery(queryData urlpkg.Values) string {
	return queryData.Encode()
}
