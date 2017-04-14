package gochat

import (
	"strconv"
	"time"
	"net/http"
)

type Utils struct {

}

func (this Utils) getUnixMsTime() string {
	return strconv.FormatInt(time.Now().UnixNano() / 1000000, 10)
}

func (this Utils) getUnixTime() string {
	return strconv.FormatInt(time.Now().Unix() / 1000000, 10)
}

func (this Utils) cookies2Map(cookies []*http.Cookie) map[string]string {
	result := map[string]string{}
	for _, v := range cookies {
		result[v.Name] = v.Value
	}

	return result
}

func (this Utils) cookies2String(cookies []*http.Cookie) string{
	result := ""
	for _, v := range cookies {
		result += v.Name + "=" + v.Value + ";"
	}

	return result
}