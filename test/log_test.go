package test

import (
	"context"
	"github.com/marsli9945/mars-go/marsLog"
	"os"
	"testing"
)

type model struct {
	CookieCode string `json:"cookie_code"`
	FeedId     string `json:"feed_id"`
}

func TestLog(t *testing.T) {
	err := os.Setenv("DEBUG_ENABLE", "true")
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()

	marsLog.Logger().Debug("Debug")
	marsLog.Logger().Info("Info")
	marsLog.Logger().Warn("Warn")
	marsLog.Logger().Error("Error")
	marsLog.Logger().Fatal("Debug")

	marsLog.Logger().DebugF("F: %s", "DebugF")
	marsLog.Logger().InfoF("F: %s", "InfoF")
	marsLog.Logger().WarnF("F: %s", "WarnF")
	marsLog.Logger().ErrorF("F: %s", "ErrorF")
	marsLog.Logger().FatalF("F: %s", "FatalF")

	marsLog.Logger().DebugFX(ctx, "FX: %s", "DebugFX")
	marsLog.Logger().InfoFX(ctx, "FX: %s", "InfoFX")
	marsLog.Logger().WarnFX(ctx, "FX: %s", "WarnFX")
	marsLog.Logger().ErrorFX(ctx, "FX: %s", "ErrorFX")
	marsLog.Logger().FatalFX(ctx, "FX: %s", "FatalFX")

	split := model{
		CookieCode: "123",
		FeedId:     "456",
	}
	marsLog.Logger().Json(split)
	marsLog.Logger().JsonFormat(split)
}
