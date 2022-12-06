package gocloud

import (
	_ "gocloud.dev/pubsub/mempubsub"
)

import (
	"context"
	"github.com/whosonfirst/go-webhookd/v3/dispatcher"
	"testing"
)

func TestPubSubDispatcher(t *testing.T) {

	uris := []string{
		"mem://testing",
		"mem://testing?mode=all",
		"mem://testing?mode=lines",
	}

	messages := [][]byte{
		[]byte("hello world"),
		[]byte("hello\nworld"),
	}

	ctx := context.Background()

	for _, u := range uris {

		d, err := dispatcher.NewDispatcher(ctx, u)

		if err != nil {
			t.Fatalf("Failed to create dispatcher for '%s', %v", u, err)
		}

		for _, m := range messages {

			err2 := d.Dispatch(ctx, m)

			if err2 != nil {
				t.Fatalf("Failed to dispatch message, %v", err2)
			}
		}
	}
}
