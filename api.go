package liblib

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/alphasnow/liblib/consts"
	"github.com/alphasnow/liblib/schema"
	"io"
	"net/http"
	"time"
)

type API struct {
	secretKey string
	accessKey string
}

func NewAPI(ak, sk string) *API {
	return &API{secretKey: sk, accessKey: ak}
}

func (a *API) Text2Img(req *schema.Text2ImgReq) (result *schema.GenerateResp, err error) {
	var reqJSON []byte
	if reqJSON, err = json.Marshal(req); err != nil {
		return
	}
	return a.Generate(consts.Text2Img, reqJSON)
}
func (a *API) Img2Img(req *schema.Img2ImgReq) (result *schema.GenerateResp, err error) {
	var reqJSON []byte
	if reqJSON, err = json.Marshal(req); err != nil {
		return
	}
	return a.Generate(consts.Img2Img, reqJSON)
}
func (a *API) Text2ImgUltra(req *schema.Text2ImgUltraReq) (result *schema.GenerateResp, err error) {
	var reqJSON []byte
	if reqJSON, err = json.Marshal(req); err != nil {
		return
	}
	return a.Generate(consts.Text2ImgUltra, reqJSON)
}
func (a *API) Img2ImgUltra(req *schema.Img2ImgUltraReq) (result *schema.GenerateResp, err error) {
	var reqJSON []byte
	if reqJSON, err = json.Marshal(req); err != nil {
		return
	}
	return a.Generate(consts.Img2ImgUltra, reqJSON)
}
func (a *API) Text2ImgClassic(req *schema.Text2ImgClassicReq) (result *schema.GenerateResp, err error) {
	var reqJSON []byte
	if reqJSON, err = json.Marshal(req); err != nil {
		return
	}
	return a.Generate(consts.Text2ImgClassic, reqJSON)
}
func (a *API) Img2ImgClassic(req *schema.Img2ImgClassicReq) (result *schema.GenerateResp, err error) {
	var reqJSON []byte
	if reqJSON, err = json.Marshal(req); err != nil {
		return
	}
	return a.Generate(consts.Img2ImgClassic, reqJSON)
}

func (a *API) Generate(endpoint string, req []byte) (resp *schema.GenerateResp, err error) {
	respJSON, err := a.request(endpoint, req)
	if err != nil {
		return nil, err
	}
	resp = new(schema.GenerateResp)
	if err = json.Unmarshal(respJSON, resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code=%d,msg=%s", resp.Code, resp.Msg)
	}
	return resp, nil
}

func (a *API) Status(req *schema.StatusReq) (result *schema.StatusResp, err error) {
	var dataJSON []byte
	if dataJSON, err = json.Marshal(req); err != nil {
		return
	}
	respJSON, err := a.request(consts.Status, dataJSON)
	if err != nil {
		return nil, err
	}
	resp := new(schema.StatusResp)
	if err = json.Unmarshal(respJSON, resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code=%d,msg=%s", resp.Code, resp.Msg)
	}
	return resp, nil
}

func (a *API) request(endpoint string, dataJSON []byte) (result []byte, err error) {
	url := a.getURL(endpoint)
	var req *http.Request
	var resp *http.Response
	if req, err = http.NewRequest("POST", url, bytes.NewBuffer(dataJSON)); err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	result, err = io.ReadAll(resp.Body)
	return result, err
}

func (a *API) getURL(endpoint string) string {
	timestamp := time.Now().UnixNano() / 1_000_000
	nonce := randomString(16)
	signature := a.hashSk(endpoint, timestamp, nonce)
	return fmt.Sprintf("%s%s?AccessKey=%s&Signature=%s&Timestamp=%d&SignatureNonce=%s",
		consts.OpenApi, endpoint, a.accessKey, signature, timestamp, nonce)
}

func (a *API) hashSk(endpoint string, timestamp int64, nonce string) string {
	data := fmt.Sprintf("%s&%d&%s", endpoint, timestamp, nonce)
	s := base64.RawURLEncoding.EncodeToString(a.hmacSha1(data))
	return s
}
func (a *API) hmacSha1(data string) []byte {
	h := hmac.New(sha1.New, []byte(a.secretKey))
	h.Write([]byte(data))
	return h.Sum(nil)
}
