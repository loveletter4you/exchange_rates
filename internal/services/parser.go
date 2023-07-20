package services

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/loveletter4you/exchange_rates/internal/model"
	"golang.org/x/text/encoding/charmap"
	"io"
)

func ParseDaily(data []byte) (*model.ValCurs, error) {
	var result *model.ValCurs
	reader := bytes.NewReader(data)
	d := xml.NewDecoder(reader)
	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}
	err := d.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
