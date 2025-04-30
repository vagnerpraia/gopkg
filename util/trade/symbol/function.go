package gpsymbol

import (
	"strings"

	gpenum "github.com/vagnerpraia/gopkg/enum"
)

func Normalize(symbol string) string {

	if strings.HasPrefix(symbol, gpenum.AssetWDO.String()) {
		return gpenum.AssetWDO.String()
	}

	if strings.HasPrefix(symbol, gpenum.AssetDOL.String()) {
		return gpenum.AssetDOL.String()
	}

	if strings.HasPrefix(symbol, gpenum.AssetWIN.String()) {
		return gpenum.AssetWIN.String()
	}

	if strings.HasPrefix(symbol, gpenum.AssetIND.String()) {
		return gpenum.AssetIND.String()
	}

	return ""
}
