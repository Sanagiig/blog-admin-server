package service

import (
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/cdn"
	"github.com/qiniu/go-sdk/v7/storage"
	"go-blog/global/settings"
	"go-blog/model/request"
	"go-blog/model/response"
	"io"
	"mime/multipart"
	url2 "net/url"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

var (
	PubTypes = []string{"profile"}
	PriTypes = []string{"article"}
)

var (
	pubKVDomain = settings.QiNiuPubKVDomain
	priKVDomain = settings.QiNiuPriKVDomain
	pubBucket   = settings.QiNiuPubBucket
	priBucket   = settings.QiNiuPriBucket
	ak          = settings.QiNiuAccessKey
	sk          = settings.QiNiuSecretKey
	mac         = qbox.NewMac(ak, sk)
)

func genScope(upType string, username string, isPub bool) string {
	sb := strings.Builder{}
	if isPub {
		sb.WriteString(pubBucket)
	} else {
		sb.WriteString(priBucket)
	}
	switch upType {
	case "profile":
		sb.WriteString(":profile/")
		sb.WriteString(username + "_")
		sb.WriteString(time.Now().Format("2006-01-02 03:04:05"))
	}

	return sb.String()
}

func genUrl(domain string, key string) string {
	sb := strings.Builder{}
	sb.WriteString(domain)
	if !strings.HasSuffix(domain, "/") && !strings.HasPrefix(key, "/") {
		sb.WriteString("/")
	}
	sb.WriteString(key)
	return sb.String()
}

func DisposeResCallbackConfig(params *request.UploadAuthBody, putPolicy *storage.PutPolicy) {
	switch params.UploadType {
	case "profile":
		key := strings.Split(putPolicy.Scope, ":")[1]
		putPolicy.CallbackURL = genUrl(settings.Domain, "/api/upload/refreshCDN")
		putPolicy.CallbackBody = fmt.Sprintf(`{"urls":["%s"]}`, genUrl(pubKVDomain, key))
	}
}

func GetUploadAuth(params *request.UploadAuthBody) (*response.UploadAuthData, error) {
	var scope string
	username := params.Username
	upType := params.UploadType
	pubPos := sort.SearchStrings(PubTypes, upType)
	priPos := sort.SearchStrings(PriTypes, upType)
	putPolicy := storage.PutPolicy{}
	preReturnBody := `"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)"`
	key := ""

	switch {
	case pubPos != len(PubTypes):
		scope = genScope(upType, username, true)
		key = strings.SplitN(scope, ":", 2)[1]
		url := genUrl(pubKVDomain, key)
		putPolicy.ReturnBody = fmt.Sprintf(`{%s,"url":"%s"}`, preReturnBody, url)
	case priPos != len(PriTypes):
		scope = genScope(upType, username, false)
		putPolicy.ReturnBody = fmt.Sprintf(`{%s,"url":"%s"}`, preReturnBody, priKVDomain)
	default:
		return nil, errors.New("生成上传验证信息失败")
	}

	putPolicy.Scope = scope
	//DisposeResCallbackConfig(params, &putPolicy)
	return &response.UploadAuthData{
		UpToken: putPolicy.UploadToken(mac),
		Key:     key,
	}, nil
}

func RefreshCdn(urls []string) (cdn.RefreshResp, error) {
	cdnManager := cdn.NewCdnManager(mac)
	return cdnManager.RefreshUrls(urls)
}

func UploadFile(typ string, username string, file *multipart.FileHeader) (url string, err error) {
	var dir = path.Join(settings.StaticDir, typ)
	var dst string

	switch typ {
	case "profile":
		dst = path.Join(dir, username)
	default:
		dst = path.Join(dir, username)
	}

	if _, err = os.Stat(dir); err != nil {
		err = os.MkdirAll(dir, 0666)
		if err != nil {
			return
		}
	}

	if _, err = os.Stat(dst); err == nil {
		err = os.Remove(dst)
		if err != nil {
			return
		}
	}

	f, err := os.Create(dst)
	defer f.Close()
	if err != nil {
		return
	}

	srcFile, err := file.Open()
	defer srcFile.Close()
	if err != nil {
		return
	}

	_, err = io.Copy(f, srcFile)
	if err != nil {
		return
	}

	url, err = url2.JoinPath(settings.Domain, "static", typ, username)
	return
}
