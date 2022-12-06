package gocloud

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/whosonfirst/go-webhookd/v3"
	"github.com/whosonfirst/go-webhookd/v3/dispatcher"
	"gocloud.dev/pubsub"
	"net/url"
)

func init() {

	ctx := context.Background()

	for _, scheme := range pubsub.DefaultURLMux().TopicSchemes() {

		err := dispatcher.RegisterDispatcher(ctx, scheme, NewPubSubDispatcher)

		if err != nil {
			panic(err)
		}
	}
}

type PubSubDispatcher struct {
	webhookd.WebhookDispatcher
	topic *pubsub.Topic
	mode  string
}

func NewPubSubDispatcher(ctx context.Context, uri string) (webhookd.WebhookDispatcher, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	mode := "lines"

	q_mode := q.Get("mode")
	switch q_mode {
	case "all", "lines":
		mode = q_mode
	default:
		return nil, fmt.Errorf("Invalid or unsupported mode, %s", q_mode)
	}

	// Because the gocloud packages are often fussy about unknown query parameters

	q.Del("mode")
	u.RawQuery = q.Encode()
	uri = u.String()

	t, err := pubsub.OpenTopic(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create topic, %w", err)
	}

	d := &PubSubDispatcher{
		topic: t,
		mode:  mode,
	}

	return d, nil
}

func (d *PubSubDispatcher) Dispatch(ctx context.Context, body []byte) *webhookd.WebhookError {

	var err error

	switch d.mode {
	case "all":
		err = d.sendMessage(ctx, body)
	default: // lines
		err = d.dispatchLines(ctx, body)

	}

	if err != nil {
		return &webhookd.WebhookError{Code: 999, Message: err.Error()}
	}

	return nil
}

func (d *PubSubDispatcher) dispatchLines(ctx context.Context, body []byte) error {

	br := bytes.NewReader(body)
	scanner := bufio.NewScanner(br)

	for scanner.Scan() {

		err := d.sendMessage(ctx, scanner.Bytes())

		if err != nil {
			return err
		}
	}

	err := scanner.Err()

	if err != nil {
		return fmt.Errorf("Scanner reported an error, %w", err)
	}

	return nil
}

func (d *PubSubDispatcher) sendMessage(ctx context.Context, body []byte) error {

	msg := &pubsub.Message{
		Body: body,
	}

	err := d.topic.Send(ctx, msg)

	if err != nil {
		return fmt.Errorf("Failed to send message, %w", err)
	}

	return nil
}
