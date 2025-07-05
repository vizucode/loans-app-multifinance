package domain

type RequestCustomer struct {
	Email                    string  `json:"email" validate:"required,email"`                     // Email harus ada dan formatnya email
	Password                 string  `json:"password" validate:"required,min=8"`                  // Password harus ada dan minimal 8 karakter
	FullName                 string  `json:"full_name" validate:"required"`                       // Nama lengkap harus ada
	LegalName                string  `json:"legal_name" validate:"required"`                      // Pointer string untuk nullable
	DateBirth                string  `json:"date_birth" validate:"required,datetime=2006-01-02"`  // Tanggal lahir harus ada dan format YYYY-MM-DD
	BornAt                   string  `json:"born_at" validate:"required"`                         // Pointer string untuk nullable
	Salary                   float64 `json:"salary" validate:"required"`                          // Pointer float64 untuk nullable DECIMAL
	NationalIdentityNumber   string  `json:"national_identity_number" validate:"required,len=16"` // NIK harus ada dan panjang 16
	NationalIdentityImageURL string  `json:"national_identity_image_url" validate:"required"`     // Pointer string untuk nullable URL
	SelfieImageURL           string  `json:"selfie_image_url" validate:"required"`                // Pointer string untuk nullable URL
}

type ResponseCustomer struct {
	Email                    string  `json:"email" validate:"required,email"`                     // Email harus ada dan formatnya email
	Password                 string  `json:"password" validate:"required,min=8"`                  // Password harus ada dan minimal 8 karakter
	FullName                 string  `json:"full_name" validate:"required"`                       // Nama lengkap harus ada
	LegalName                string  `json:"legal_name" validate:"required"`                      // Pointer string untuk nullable
	DateBirth                string  `json:"date_birth" validate:"required,datetime=2006-01-02"`  // Tanggal lahir harus ada dan format YYYY-MM-DD
	BornAt                   string  `json:"born_at" validate:"required"`                         // Pointer string untuk nullable
	Salary                   float64 `json:"salary" validate:"required"`                          // Pointer float64 untuk nullable DECIMAL
	NationalIdentityNumber   string  `json:"national_identity_number" validate:"required,len=16"` // NIK harus ada dan panjang 16
	NationalIdentityImageURL string  `json:"national_identity_image_url" validate:"required"`     // Pointer string untuk nullable URL
	SelfieImageURL           string  `json:"selfie_image_url" validate:"required"`
}

type RequestSignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResponseSignIn struct {
	AccessToken        string `json:"access_token"`
	RefreshToken       string `json:"refresh_token"`
	AccessTokenExpired string `json:"access_token_expired"`
}

type RequestRefreshToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
