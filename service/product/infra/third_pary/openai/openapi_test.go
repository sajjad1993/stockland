package openai

import (
	"context"
	"fmt"
	"testing"
)

func TestProcessor_ProcessRequest(t *testing.T) {
	client := NewOpenAIProcessor("xxxx")

	fmt.Println(client.ProcessRequest(context.Background(), "hello", "there"))
}
