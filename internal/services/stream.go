package services

import (
	"context"
	"time"

	"github.com/GetStream/getstream-go"
	"github.com/danilovict2/go-interview-RTC/models"
	"github.com/google/uuid"
)

func CreateNewVideoCall(client *getstream.Stream, authUser models.User, startTime time.Time) (string, error) {
	callID := uuid.NewString()
	call := client.Video().Call("default", callID)
	callRequest := getstream.GetOrCreateCallRequest{
		Data: &getstream.CallRequest{
			CreatedByID: getstream.PtrTo(authUser.UUID.String()),
			StartsAt:    &getstream.Timestamp{Time: &startTime},
			SettingsOverride: &getstream.CallSettingsRequest{
				Recording: &getstream.RecordSettingsRequest{
					Mode:    "auto-on",
					Quality: getstream.PtrTo("720p"),
				},
			},
		},
	}

	resp, err := call.GetOrCreate(context.TODO(), &callRequest)
	if err != nil {
		return "", err
	}

	return resp.Data.Call.ID, nil
}
