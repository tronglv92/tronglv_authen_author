package auth

import (
	"context"
	"fmt"
	"github/tronglv_authen_author/helper/define"
)

func GetAuthData(ctx context.Context) (AuthData, error) {
	data, ok := ctx.Value(define.AuthDataContextKey).(AuthData)
	if !ok {
		return nil, fmt.Errorf("can't get info data from context")
	}
	return data, nil
}
