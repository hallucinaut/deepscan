package forensic

import (
	"testing"
	"time"
)

func TestNewForensicAnalyzer(t *testing.T) {
	analyzer := NewForensicAnalyzer()
	if analyzer == nil {
		t.Fatal("Expected analyzer to be created")
	}
}

func TestAddIndicator(t *testing.T) {
	analyzer := NewForensicAnalyzer()
	indicator := ForensicIndicator{
		Name:        "Test Indicator",
		Description: "Test description",
		Severity:    "low",
		Confidence:  0.8,
	}

	analyzer.AddIndicator(indicator)
	indicators := analyzer.GetIndicators()

	if len(indicators) != 1 {
		t.Errorf("Expected 1 indicator, got %d", len(indicators))
	}
	if indicators[0].Name != "Test Indicator" {
		t.Errorf("Expected indicator name 'Test Indicator', got '%s'", indicators[0].Name)
	}
}

func TestGetIndicators(t *testing.T) {
	analyzer := NewForensicAnalyzer()
	analyzer.AddIndicator(ForensicIndicator{
		Name:        "Indicator 1",
		Description: "Description 1",
		Severity:    "medium",
	})

	indicators := analyzer.GetIndicators()
	if len(indicators) != 1 {
		t.Errorf("Expected 1 indicator, got %d", len(indicators))
	}
}

func TestGetForensicAnalyzer(t *testing.T) {
	analyzer := NewForensicAnalyzer()
	retrieved := GetForensicAnalyzer(analyzer)

	if retrieved != analyzer {
		t.Error("Expected analyzer to be the same instance")
	}
}

func TestGetForensicAnalysisResult(t *testing.T) {
	result := &ForensicAnalysisResult{
		ID:                "forensic-001",
		FilePath:          "/path/to/image.jpg",
		IsAuthentic:       true,
		AuthenticityScore: 0.95,
		Message:           "Analysis completed",
	}

	retrieved := GetForensicAnalysisResult(result)
	if retrieved.ID != "forensic-001" {
		t.Errorf("Expected ID 'forensic-001', got '%s'", retrieved.ID)
	}
	if !retrieved.IsAuthentic {
		t.Error("Expected IsAuthentic to be true")
	}
}

func TestForensicAnalysisResult(t *testing.T) {
	result := ForensicAnalysisResult{
		ID:                "forensic-001",
		FilePath:          "/path/to/image.jpg",
		Timestamp:         time.Now(),
		IsAuthentic:       true,
		AuthenticityScore: 0.95,
		Methodology:       []string{"Header analysis", "EXIF analysis"},
		Message:           "Analysis completed",
	}

	if result.AuthenticityScore != 0.95 {
		t.Errorf("Expected AuthenticityScore 0.95, got %f", result.AuthenticityScore)
	}
	if !result.IsAuthentic {
		t.Error("Expected IsAuthentic to be true")
	}
	if len(result.Methodology) != 2 {
		t.Errorf("Expected 2 methodology items, got %d", len(result.Methodology))
	}
}

func TestForensicIndicator(t *testing.T) {
	indicator := ForensicIndicator{
		Name:        "Test Indicator",
		Description: "Test description",
		Severity:    "high",
		Confidence:  0.85,
		Evidence:    []string{"Evidence 1", "Evidence 2"},
	}

	if indicator.Severity != "high" {
		t.Errorf("Expected severity 'high', got '%s'", indicator.Severity)
	}
	if indicator.Confidence != 0.85 {
		t.Errorf("Expected confidence 0.85, got %f", indicator.Confidence)
	}
	if len(indicator.Evidence) != 2 {
		t.Errorf("Expected 2 evidence items, got %d", len(indicator.Evidence))
	}
}