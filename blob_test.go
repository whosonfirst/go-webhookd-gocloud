package gocloud

import (
	_ "gocloud.dev/blob/fileblob"
)

import (
	"context"
	"github.com/whosonfirst/go-webhookd/v3/dispatcher"
	"testing"
)

func TestDispatcher(t *testing.T) {

	uris := []string{
		"file:///usr/local/data/webhookd/",
		"file:///usr/local/data/webhookd/?metadata=skip",
		"file:///usr/local/data/webhookd/?metadata=skip&dispatch_prefix=_ts_",
	}

	ctx := context.Background()

	for _, u := range uris {

		_, err := dispatcher.NewDispatcher(ctx, u)

		if err != nil {
			t.Fatalf("Failed to create dispatcher for '%s', %v", u, err)
		}
	}
}
