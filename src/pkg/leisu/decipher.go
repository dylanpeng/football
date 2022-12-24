package leisu

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"football/common"
	"io"
	"net/url"
	"strconv"
	"strings"
)

func Rot(cipher string, i int32) (result string) {
	t := Roott(cipher, i)
	result = Pushmsg(t)
	return
}

func Roott(data string, i int32) (result string) {
	for _, item := range data {
		o := item
		if item >= 65 && item <= 90 {
			o = (item-65-1*i+26)%26 + 65
		}
		if item >= 97 && item <= 122 {
			o = (item-97-1*i+26)%26 + 97
		}

		result += string(o)
	}

	return
}

func Pushmsg(data string) (result string) {
	deByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		common.Logger.Infof("decode failed. err: %s", err)
		return
	}

	zlibByte := DoZlibUnCompress(deByte)

	a := base64.StdEncoding.EncodeToString(zlibByte)

	b, err := base64.StdEncoding.DecodeString(a)
	if err != nil {
		common.Logger.Infof("decode failed. err: %s", err)
		return
	}

	c := string(b)
	d := strings.Replace(c, "%u", "\\u", -1)
	e, err := url.QueryUnescape(d)
	if err != nil {
		common.Logger.Infof("QueryUnescape failed. err: %s", err)
		return
	}

	f, err := zhToUnicode([]byte(e))
	if err != nil {
		common.Logger.Infof("zhToUnicode failed. err: %s", err)
		return
	}

	result = string(f)

	return
}

func zhToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

// 进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}
