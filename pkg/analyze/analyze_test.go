package analyze

import (
	"testing"
)

func TestNewMediaAnalyzer(t *testing.T) {
	analyzer := NewMediaAnalyzer()
	if analyzer == nil {
		t.Fatal("Expected analyzer to be created")
	}
	if analyzer.thresholds == nil {
		t.Error("Expected thresholds to be initialized")
	}
}

func TestSetThresholds(t *testing.T) {
	analyzer := NewMediaAnalyzer()
	thresholds := AnalysisThresholds{
		DeepfakeThreshold:   0.8,
		ConfidenceThreshold: 0.7,
		MinFileSize:         2048,
	}

	analyzer.SetThresholds(thresholds)
	currentThresholds := analyzer.GetThresholds()

	if currentThresholds.DeepfakeThreshold != 0.8 {
		t.Errorf("Expected DeepfakeThreshold 0.8, got %f", currentThresholds.DeepfakeThreshold)
	}
	if currentThresholds.MinFileSize != 2048 {
		t.Errorf("Expected MinFileSize 2048, got %d", currentThresholds.MinFileSize)
	}
}

func TestDetectFormat(t *testing.T) {
	analyzer := NewMediaAnalyzer()

	testCases := []struct {
		path     string
		expected MediaFormat
	}{
		{"/path/to/image.jpg", FormatJPEG},
		{"/path/to/image.jpeg", FormatJPEG},
		{"/path/to/image.png", FormatPNG},
		{"/path/to/image.gif", FormatGIF},
		{"/path/to/image.webp", FormatWebP},
		{"/path/to/file.mp4", FormatMP4},
		{"/path/to/file.unknown", FormatUnknown},
	}

	for _, tc := range testCases {
		result := analyzer.detectFormat(tc.path)
		if result != tc.expected {
			t.Errorf("detectFormat(%s) = %s, want %s", tc.path, result, tc.expected)
		}
	}
}

func TestCreateCommonThresholds(t *testing.T) {
	thresholds := CreateCommonThresholds()

	if thresholds.DeepfakeThreshold == 0 {
		t.Error("Expected non-zero DeepfakeThreshold")
	}
	if thresholds.ConfidenceThreshold == 0 {
		t.Error("Expected non-zero ConfidenceThreshold")
	}
}

func TestGetAnalysisResult(t *testing.T) {
	result := &AnalysisResult{
		ID:             "analysis-001",
		FilePath:       "/path/to/image.jpg",
		Format:         FormatJPEG,
		IsDeepfake:     false,
		DeepfakeScore:  0.05,
		Confidence:     0.95,
		IsValid:        true,
	}

	retrieved := GetAnalysisResult(result)
	if retrieved.ID != "analysis-001" {
		t.Errorf("Expected ID 'analysis-001', got '%s'", retrieved.ID)
	}
	if retrieved.IsDeepfake != false {
		t.Error("Expected IsDeepfake to be false")
	}
}

func TestGetMediaAnalyzer(t *testing.T) {
	analyzer := NewMediaAnalyzer()
	retrieved := GetMediaAnalyzer(analyzer)

	if retrieved != analyzer {
		t.Error("Expected analyzer to be the same instance")
	}
}

func TestAnalysisResult_Metrics(t *testing.T) {
	result := AnalysisResult{
		ID:            "analysis-001",
		FilePath:      "/path/to/image.jpg",
		Format:        FormatJPEG,
		Dimensions:    [2]int{1920, 1080},
		FileSize:      245678,
		Hash:          "a1b2c3d4e5f6789012345678901234567890abcd",
		IsDeepfake:    false,
		DeepfakeScore: 0.05,
		Confidence:    0.95,
		IsValid:       true,
		Message:       "Analysis completed",
	}

	if result.FileSize != 245678 {
		t.Errorf("Expected FileSize 245678, got %d", result.FileSize)
	}
	if result.Confidence != 0.95 {
		t.Errorf("Expected Confidence 0.95, got %f", result.Confidence)
	}
	if !result.IsValid {
		t.Error("Expected IsValid to be true")
	}
}