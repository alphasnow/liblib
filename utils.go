package liblib

import (
	"crypto/rand"
	"github.com/alphasnow/liblib/consts"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	result := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, charsetLength)
		result[i] = charset[index.Int64()]
	}
	return string(result)
}

func getEndpointByTemplate(template string) string {
	switch template {
	case consts.FlagshipF1TextToImage:
		return consts.Text2Img
	case consts.FlagshipF1ImageToImage:
		return consts.Img2Img
	case consts.Classic15AndXLTextToImage:
		return consts.Text2Img
	case consts.Classic15AndXLImageToImage:
		return consts.Img2Img
	case consts.ControlnetLocalRedraw:
		return consts.Text2Img
	case consts.ImageToImageLocalRedraw:
		return consts.Img2Img
	case consts.InstantIDFaceSwap:
		return consts.Img2Img
	case consts.SimpleFlagshipF1TextToImage:
		return consts.Text2ImgUltra
	case consts.SimpleFlagshipF1ImageToImage:
		return consts.Img2ImgUltra
	case consts.SimpleClassicXLTextToImage:
		return consts.Text2ImgClassic
	case consts.SimpleClassicXLImageToImage:
		return consts.Img2ImgClassic
	default:
		return ""
	}
}
