package handlers

import (
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"
	"github.com/gin-gonic/gin"
)

// GetAssetsChartHistory returns
func GetAssetsChartHistory(c *gin.Context) {
	result := make([]models.AssetChartHistory, 0)

	limit := int(24)

	for _, assetName := range models.ChosenAssetNames {
		asset, err := s.client.GetAsset(assetName)
		if err != nil {
			s.l.Printf("failed to get asset detail information: %s\n", err)
		}

		charts, err := s.db.QueryAssetChartHistory(assetName, limit)
		if err != nil {
			s.l.Printf("failed to query asset chart history: %s", err)
		}

		prices := make([]models.Prices, 0)

		for _, chart := range charts {
			tempPrices := &models.Prices{
				Price:     chart.Price,
				Timestamp: chart.Timestamp,
			}

			prices = append(prices, *tempPrices)
		}

		tempResult := &models.AssetChartHistory{
			Name:         asset.Name,
			Asset:        asset.Asset,
			MappedAsset:  asset.MappedAsset,
			CurrentPrice: asset.Price,
			QuoteUnit:    asset.QuoteUnit,
			ChangeRange:  asset.ChangeRange,
			Supply:       asset.Supply,
			Marketcap:    asset.Price * asset.Supply,
			AssetImage:   asset.AssetImg,
			Prices:       prices,
		}

		result = append(result, *tempResult)
	}

	models.Respond(c.Writer, result)
	return
}
