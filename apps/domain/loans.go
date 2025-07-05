package domain

type RequestLoans struct {
	Otr         float64 `json:"otr"`
	AssetName   string  `json:"asset_name"`
	PickedTenor int     `json:"picked_tenor"`
}

type ResponseLoans struct {
	Otr                     float64 `json:"otr"`
	AssetName               string  `json:"asset_name"`
	PickedTenor             int     `json:"picked_tenor"`
	InterestFee             float64 `json:"interest_fee"`
	AdminFee                float64 `json:"admin_fee"`
	InterestRate            float64 `json:"interest_rate"`
	TotalMonthlyInstalement float64 `json:"total_monthly_instalement"`
	TotalInstallmentAmount  float64 `json:"total_installment_amount"`
}
