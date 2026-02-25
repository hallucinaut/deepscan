package verify

import (
	"testing"
	"time"
)

func TestNewMetadataCollector(t *testing.T) {
	collector := NewMetadataCollector()
	if collector == nil {
		t.Fatal("Expected collector to be created")
	}
}

func TestCollect(t *testing.T) {
	collector := NewMetadataCollector()
	// Test with non-existent file - should return empty map without error
	metadata, err := collector.Collect("/nonexistent/file.jpg")
	// Expected to fail or return empty
	if err == nil && len(metadata) == 0 {
		// This is acceptable for non-existent file
		return
	}
}

func TestVerifyMetadata(t *testing.T) {
	metadata := map[string]string{
		"filename": "test.jpg",
		"size":     "12345",
		"created":  "2024-01-01T00:00:00Z",
		"modified": "2024-01-01T00:00:00Z",
	}

	issues := VerifyMetadata(metadata)
	// Should have no issues for complete metadata
	if len(issues) > 0 {
		t.Errorf("Expected no issues, got %d: %v", len(issues), issues)
	}
}

func TestNewHashVerifier(t *testing.T) {
	expectedHash := "a1b2c3d4e5f6789012345678901234567890abcd"
	verifier := NewHashVerifier(expectedHash)

	if verifier == nil {
		t.Fatal("Expected verifier to be created")
	}
	if verifier.expectedHash != expectedHash {
		t.Errorf("Expected hash '%s', got '%s'", expectedHash, verifier.expectedHash)
	}
}

func TestNewBlockchainVerifier(t *testing.T) {
	verifier := NewBlockchainVerifier("https://registry.example.com")

	if verifier == nil {
		t.Fatal("Expected verifier to be created")
	}
}

func TestNewWatermarkVerifier(t *testing.T) {
	verifier := NewWatermarkVerifier()

	if verifier == nil {
		t.Fatal("Expected verifier to be created")
	}
}

func TestNewForensicAnalyzer(t *testing.T) {
	analyzer := NewForensicAnalyzer()

	if analyzer == nil {
		t.Fatal("Expected analyzer to be created")
	}
}

func TestVerifyFile(t *testing.T) {
	// Test with non-existent file
	result, err := VerifyFile("/nonexistent/file.jpg")
	// Expected to fail
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
	if result == nil {
		t.Error("Expected result to not be nil")
	}
}

func TestGenerateVerificationReport(t *testing.T) {
	results := []VerificationResult{
		{
			FilePath:  "/path/to/image1.jpg",
			Method:    MethodHash,
			IsValid:   true,
			Confidence: 0.95,
			Message:   "Verification successful",
			Hash:      "a1b2c3d4e5f6789012345678901234567890abcd",
		},
		{
			FilePath:  "/path/to/image2.jpg",
			Method:    MethodMetadata,
			IsValid:   false,
			Confidence: 0.6,
			Message:   "Verification failed",
			Hash:      "b2c3d4e5f6789012345678901234567890abcdef",
		},
	}

	report := GenerateVerificationReport(results)

	if report == "" {
		t.Error("Expected report to not be empty")
	}
	if len(report) < 50 {
		t.Errorf("Expected report to be at least 50 characters, got %d", len(report))
	}
}

func TestCreateCommonVerifications(t *testing.T) {
	methods := CreateCommonVerifications()

	if len(methods) == 0 {
		t.Error("Expected at least one verification method")
	}

	// Check that methods have valid values
	for i, method := range methods {
		if string(method) == "" {
			t.Errorf("Method %d has empty value", i)
		}
	}
}

func TestGetVerificationResult(t *testing.T) {
	result := &VerificationResult{
		ID:         "verify-001",
		FilePath:   "/path/to/image.jpg",
		Method:     MethodHash,
		IsValid:    true,
		Confidence: 0.95,
	}

	retrieved := GetVerificationResult(result)
	if retrieved.ID != "verify-001" {
		t.Errorf("Expected ID 'verify-001', got '%s'", retrieved.ID)
	}
	if !retrieved.IsValid {
		t.Error("Expected IsValid to be true")
	}
}

func TestGetMetadataEntry(t *testing.T) {
	entry := &MetadataEntry{
		Key:   "camera_model",
		Value: "Test Camera",
	}

	retrieved := GetMetadataEntry(entry)
	if retrieved.Key != "camera_model" {
		t.Errorf("Expected key 'camera_model', got '%s'", retrieved.Key)
	}
}

func TestVerificationResult_Structure(t *testing.T) {
	result := VerificationResult{
		ID:             "verify-001",
		FilePath:       "/path/to/image.jpg",
		Method:         MethodHash,
		IsValid:        true,
		Confidence:     0.95,
		Message:        "Verification successful",
		Timestamp:      time.Now(),
		Hash:           "a1b2c3d4e5f6789012345678901234567890abcd",
		Metadata:       map[string]string{"key": "value"},
		ForensicResult: map[string]interface{}{"test": true},
	}

	if result.Confidence != 0.95 {
		t.Errorf("Expected Confidence 0.95, got %f", result.Confidence)
	}
	if len(result.Metadata) != 1 {
		t.Errorf("Expected 1 metadata entry, got %d", len(result.Metadata))
	}
}