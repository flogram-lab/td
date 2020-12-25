package telegram_test

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"os"
	"testing"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/transport"
)

func TestMTProxy(t *testing.T) {
	addr, ok := os.LookupEnv("GOTD_MTPROXY_ADDR")
	if !ok {
		t.Skip("Skipped. Set GOTD_MTPROXY_ADDR to enable external e2e mtproxy test.")
	}

	secret, err := hex.DecodeString(os.Getenv("GOTD_MTPROXY_SECRET"))
	if err != nil {
		t.Fatal("secret parsing failed", err)
	}

	trp, err := transport.MTProxy(nil, 2, secret)
	if err != nil {
		t.Fatal("secret invalid", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	client := telegram.NewClient(telegram.TestAppID, telegram.TestAppHash, telegram.Options{
		Addr:      addr,
		Transport: trp,
	})

	if err := client.Connect(ctx); err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = client.Close()
	}()

	if err := telegram.NewAuth(
		telegram.TestAuth(rand.Reader, 2),
		telegram.SendCodeOptions{},
	).Run(ctx, client); err != nil {
		t.Fatal(err)
	}

	if _, err := client.Self(ctx); err != nil {
		t.Fatal(err)
	}
}
