package main

import (
	"context"
	"github.com/heroiclabs/nakama-common/runtime"
)

func saveDataToDb(ctx context.Context, logger runtime.Logger, out []byte, err error, module runtime.NakamaModule) error {
	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok || userID == "" {
		logger.Error("No user ID found in context")
		// return unauthenticated error
		return runtime.NewError("No user ID in context", 16)
	}
	writes := []*runtime.StorageWrite{
		{
			Collection:      "user_data",
			Key:             "server_info",
			UserID:          userID,
			Value:           string(out),
			PermissionRead:  1, // only owner read
			PermissionWrite: 1, // owner write permission
		},
	}
	_, err = module.StorageWrite(ctx, writes)
	if err != nil {
		logger.Error("Error while saving to DB : %v", err)
		// return internal error
		return runtime.NewError("Error while fetching content", 13)
	}
	return nil
}
