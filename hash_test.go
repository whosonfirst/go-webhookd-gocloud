package gocloud

import (
	"context"
	"testing"
)

func TestHashBody(t *testing.T) {

	tests := map[string]string{
		"hello world": "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9",
	}

	ctx := context.Background()

	for k, expected_v := range tests {

		v, err := HashBody(ctx, []byte(k))

		if err != nil {
			t.Fatalf("Failed to hash '%s', %v", k, err)
		}

		if v != expected_v {
			t.Fatalf("Unexpected value for '%s': %s", k, v)
		}
	}

}
