package domain

import "context"

type AIProcessor interface {
	ProcessRequest(ctx context.Context, prompt string, sub string) (Offer, error)
	AnalyzeImage(ctx context.Context, imagePath string) ([]string, error)
}
