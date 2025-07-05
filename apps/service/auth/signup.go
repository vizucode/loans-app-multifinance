package auth

import (
	"context"
	"fmt"
	"math/rand"
	"multifinancetest/apps/domain"
	"multifinancetest/apps/models"
	"multifinancetest/helpers/constants/rpcstd"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
	"golang.org/x/crypto/bcrypt"
)

func (s *auth) SignUp(ctx context.Context, req domain.RequestCustomer) (err error) {

	err = s.validator.StructCtx(ctx, &req)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	dateBirth, err := time.Parse("2006-01-02", req.DateBirth)
	if err != nil {
		return err
	}

	resultUser, err := s.db.FirstCustomerByEmail(ctx, req.Email)
	if err != nil {
		return errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "email already used")
	}

	if resultUser.Email != "" {
		return errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "email already used")
	}

	// process national identity image
	nationalIdentityFile, err := s.processImage(ctx, req.NationalIdentityImageURL, "100x100")
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}
	req.NationalIdentityImageURL = fmt.Sprintf("%s/%d.png", "/profile-pictures", time.Now().Unix())
	err = s.fs.UploadFile(ctx, req.NationalIdentityImageURL, nationalIdentityFile)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	// process selfie image
	selfieFile, err := s.processImage(ctx, req.SelfieImageURL, "")
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}
	req.SelfieImageURL = fmt.Sprintf("%s/%d.png", "/profile-pictures", time.Now().Unix())
	err = s.fs.UploadFile(ctx, req.SelfieImageURL, selfieFile)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	userId := uuid.New().String()
	err = s.db.CreateCustomer(ctx, models.Customer{
		ID:                       userId,
		Email:                    req.Email,
		Password:                 string(hashedPassword),
		FullName:                 req.FullName,
		LegalName:                req.LegalName,
		DateBirth:                dateBirth,
		BornAt:                   req.BornAt,
		Salary:                   req.Salary,
		NationalIdentityNumber:   req.NationalIdentityNumber,
		NationalIdentityImageUrl: req.NationalIdentityImageURL,
		SelfieImageUrl:           req.SelfieImageURL,
	})
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	// insert customer tenor
	resultTenor, err := s.db.GetAllTenor(ctx)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	for _, tenor := range resultTenor {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		min := 100
		max := 500_000
		randomAmount := r.Intn(max-min+1) + min
		err = s.db.CreateCustomerTenor(ctx, models.CustomerTenors{
			ID:              uuid.New().String(),
			CustomerID:      userId,
			TenorID:         tenor.ID,
			LimitLoanAmount: float64(randomAmount),
		})
		if err != nil {
			logger.Log.Error(ctx, err)
		}
	}

	return nil
}
