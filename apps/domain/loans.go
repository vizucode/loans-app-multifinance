package domain

type RequestLoans struct {
	Otr         float64 `json:"otr" validate:"required"`
	AssetName   string  `json:"asset_name" validate:"required"`
	PickedTenor int     `json:"picked_tenor" validate:"required"`
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

type ResponseLimitLoans struct {
	LoanId              string  `json:"loan_id"`
	LoanMonth           int     `json:"loan_month"`
	TotalLoanAmount     float64 `json:"total_loan_amount"`
	RemainingLoanAmount float64 `json:"remaining_loan_amount"`
}
