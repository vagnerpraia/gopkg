package gpenum

import (
	"errors"
	"strings"
)

type Asset string

const (
	assetINDAlias = "IND"
	assetDOLAlias = "DOL"
	assetWINAlias = "WIN"
	assetWDOAlias = "WDO"

	AssetIND Asset = assetINDAlias
	AssetDOL Asset = assetDOLAlias
	AssetWIN Asset = assetWINAlias
	AssetWDO Asset = assetWDOAlias
)

func NewAsset(name string) (Asset, error) {

	name = strings.ToUpper(name)

	switch name {
	case assetINDAlias:
		return AssetIND, nil

	case assetDOLAlias:
		return AssetDOL, nil

	case assetWINAlias:
		return AssetWIN, nil

	case assetWDOAlias:
		return AssetWDO, nil

	default:
		return "", errors.New("asset not found")
	}
}

func (a Asset) String() string {

	switch a {
	case AssetIND:
		return "IND (Índice Bovespa)"

	case AssetDOL:
		return "DOL (Dólar Comercial Futuro)"

	case AssetWIN:
		return "WIN (Ibovespa Mini)"

	case AssetWDO:
		return "WDO (Dólar Mini)"

	default:
		return "Unknown Asset"
	}
}

func (a Asset) Symbol() string {

	switch a {
	case AssetIND:
		return "IND"

	case AssetDOL:
		return "DOL"

	case AssetWIN:
		return "WIN"

	case AssetWDO:
		return "WDO"

	default:
		return "Unknown Asset"
	}
}
