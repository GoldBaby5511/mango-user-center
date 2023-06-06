package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	// Demo: pkg.Http.PostJson("http://xxx.com/yy", map[string]string{"a": "1"}).Do()
	Http myHttp
	// 默认超时时间
	httpDefaultTimeout = time.Second * 5
)

// func init() {
// 	go func() {
// 		time.Sleep(time.Second)
// 		Http.Get("http://127.0.0.1:19956/ping", nil).Do()
// 	}()
// }

type myHttp struct{}

func (myHttp) Get(url string, params map[string]string) *httpBuilder {
	var r httpBuilder
	r.Method = http.MethodGet
	if len(params) > 0 {
		r.Url = url + "?" + string(r.BuildQuery(params))
	} else {
		r.Url = url
	}
	return &r
}

func (myHttp) PostJson(url string, bodyData interface{}) *httpBuilder {
	var r httpBuilder
	r.Method = http.MethodPost
	r.Url = url
	r.Header = map[string]string{"Content-Type": "application/json"}
	b, _ := json.Marshal(bodyData)
	r.Body = bytes.NewReader(b)
	r.RequestBodyJson = b
	return &r
}

func (myHttp) PostForm(url string, bodyData map[string]string) *httpBuilder {
	var r httpBuilder
	r.Method = http.MethodPost
	r.Url = url
	r.Header = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	bodyDataFormat := r.BuildQuery(bodyData)
	r.Body = bytes.NewReader(bodyDataFormat)
	r.RequestBodyString = string(bodyDataFormat)
	return &r
}

type httpBuilder struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Timeout time.Duration     `json:"-"`
	Header  map[string]string `json:"header,omitempty"`
	Body    io.Reader         `json:"-"`
	Start   time.Time         `json:"-"`
	httpLog
}

type httpLog struct {
	Cost              int64           `json:"cost"`
	RequestBodyString string          `json:"req_string,omitempty"`
	RequestBodyJson   json.RawMessage `json:"req_json,omitempty"`
	ResponseStatus    int             `json:"res_status"`
	ResponseBody      json.RawMessage `json:"res_body"`
}

func (r *httpBuilder) BuildQuery(params map[string]string) []byte {
	if len(params) == 0 {
		return nil
	}

	var buffer bytes.Buffer
	var i int
	for k := range params {
		buffer.WriteString(k)
		buffer.WriteByte('=')
		buffer.WriteString(params[k])
		if i != len(params)-1 {
			buffer.WriteByte('&')
		}
		i++
	}

	return buffer.Bytes()
}

func (r *httpBuilder) SetTimeout(timeout time.Duration) *httpBuilder {
	r.Timeout = timeout
	return r
}

func (r *httpBuilder) AddHeader(key, val string) *httpBuilder {
	if len(r.Header) == 0 {
		r.Header = make(map[string]string)
	}
	r.Header[key] = val
	return r
}

func (r *httpBuilder) Do() ([]byte, error) {
	var (
		req    *http.Request
		client *http.Client
		resp   *http.Response
		err    error
	)
	defer r.writeLog(&err)

	r.Start = time.Now()
	req, err = http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		return nil, err
	}

	for k, v := range r.Header {
		req.Header.Set(k, v)
	}
	if r.Timeout == 0 {
		r.Timeout = httpDefaultTimeout
	}

	// 发送请求
	client = &http.Client{Timeout: r.Timeout}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r.ResponseStatus = resp.StatusCode
	r.ResponseBody, err = ioutil.ReadAll(resp.Body)

	return r.ResponseBody, err
}

func (r *httpBuilder) DoTo(to interface{}) error {
	b, err := r.Do()
	if err != nil {
		return err
	}
	if to != nil {
		return json.Unmarshal(b, to)
	}
	return nil
}

func (r *httpBuilder) writeLog(err *error) {
	r.Cost = time.Since(r.Start).Milliseconds()
	info, _ := json.Marshal(&r.httpLog)
	l := logrus.WithField("HTTPLOG", json.RawMessage(info))
	if *err != nil {
		l.Error(*err)
	} else {
		l.Info()
	}
}
