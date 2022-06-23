package gocloud

import (
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
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

	d, err := dispatcher.NewDispatcher(ctx, "mem://")

	if err != nil {
		t.Fatalf("Failed to create new mem:// dispatcher, %v", err)
	}

	err2 := d.Dispatch(ctx, []byte("hello world"))

	if err2 != nil {
		t.Fatalf("Failed to dispatch message, %v", err2)
	}

}
