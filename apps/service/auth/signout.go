package auth

import (
	"context"
	"net/http"

	"multifinancetest/apps/middlewares/security"
	"multifinancetest/helpers/constants/rpcstd"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (uc *auth) SignOut(ctx context.Context) (err error) {
	currentUser, found := security.ExtractUserContext(ctx)
	if !found {
		return errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "user not found")
	}

	err = uc.db.RevokeAllTokenByAccessTokenByUserId(ctx, currentUser.Id)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return nil
}
