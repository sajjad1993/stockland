package openai

import (
	"fmt"
	"testing"
)

func TestProcessor_ProcessRequest(t *testing.T) {
	client := NewOpenAIProcessor("xxx")

	fmt.Println(client.ProcessRequest("hello"))

}
