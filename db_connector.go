package main

import (
	"context"
	"errors"
	"github.com/heroiclabs/nakama-common/runtime"
)

func saveDataToDb(ctx context.Context, logger runtime.Logger, out []byte, module runtime.NakamaModule, userID string) error {

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
	_, err := module.StorageWrite(ctx, writes)
	if err != nil {
		logger.Error("Error while saving to DB : %v", err)
		// return internal error
		return errors.New("error while saving data")
	}
	return nil
}
