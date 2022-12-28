package utilities

import (
	"image"
	_ "image/jpeg"
	"os"
	"strings"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func DecodeQR(path string) ([]string, error) {
	// open and decode image file
	file, _ := os.Open(path)
	img, _, _ := image.Decode(file)

	// prepare BinaryBitmap
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return nil, err
	}

	results := strings.Split(strings.ReplaceAll(result.String(), " ", ""), ">")
	//surname, firstName, dateOfBirth, pollingStationCode, dateOfRegistration, idNumber := results[0], results[1], results[2], results[3], results[4], results[5]

	return results, nil
}
