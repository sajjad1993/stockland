package domain

type AIProcessor interface {
	ProcessRequest(prompt string) (string, error)
	AnalyzeImage(imagePath string) ([]string, error)
}
