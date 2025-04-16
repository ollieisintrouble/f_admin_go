package models

type GetAssetRequest struct {
	AssetId   int64  `json:"assetId"`
	FindMany  bool   `json:"findMany"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
