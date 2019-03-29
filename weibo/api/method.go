package api

import (
	"WeiPro/weibo/app"
	"WeiPro/weibo/code"
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

//get 方法
func GetMethond(args map[string]string, urls string) (res interface{}, err error) {
	URL, err := url.Parse(urls)
	if err != nil {
		return
	}
	query := URL.Query()
	for key, val := range args {
		query.Add(key, val)
	}
	URL.RawQuery = query.Encode()
	req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		return
	}
	resp, err := Do(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if resp.StatusCode != 200 {
		app.Logger.Warn().Str("response", string(b)).Msg("received error response.")
		var result code.APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//post方法
func PostMethod(request interface{}, URL string) (err error) {
	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	app.Logger.Debug().Str("url", URL).Msg("post data.")
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewReader(b))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	resp, err := Do(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if resp.StatusCode != 200 {
		app.Logger.Warn().Str("response", string(b)).Msg("received error response.")
		var result code.APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}

	return
}
