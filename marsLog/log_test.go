package marsLog

import (
	"context"
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

	Logger().Debug("Debug")
	Logger().Info("Info")
	Logger().Warn("Warn")
	Logger().Error("Error")
	Logger().Fatal("Debug")

	Logger().DebugF("F: %s", "DebugF")
	Logger().InfoF("F: %s", "InfoF")
	Logger().WarnF("F: %s", "WarnF")
	Logger().ErrorF("F: %s", "ErrorF")
	Logger().FatalF("F: %s", "FatalF")

	Logger().DebugFX(ctx, "FX: %s", "DebugFX")
	Logger().InfoFX(ctx, "FX: %s", "InfoFX")
	Logger().WarnFX(ctx, "FX: %s", "WarnFX")
	Logger().ErrorFX(ctx, "FX: %s", "ErrorFX")
	Logger().FatalFX(ctx, "FX: %s", "FatalFX")

	split := model{
		CookieCode: "123",
		FeedId:     "456",
	}
	Logger().Json(split)
	Logger().JsonFormat(split)
}
