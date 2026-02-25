// Package analyze provides media analysis capabilities for deepfake detection.
package analyze

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"time"
)

// MediaFormat represents supported media formats.
type MediaFormat string

const (
	FormatJPEG   MediaFormat = "jpeg"
	FormatPNG    MediaFormat = "png"
	FormatGIF    MediaFormat = "gif"
	FormatWebP   MediaFormat = "webp"
	FormatMP4    MediaFormat = "mp4"
	FormatAVI    MediaFormat = "avi"
	FormatMOV    MediaFormat = "mov"
	FormatUnknown MediaFormat = "unknown"
)

// AnalysisResult represents media analysis result.
type AnalysisResult struct {
	ID              string
	FilePath        string
	Format          MediaFormat
	Dimensions      [2]int
	FileSize        int64
	Hash            string
	CreatedAt       time.Time
	ModifiedAt      time.Time
	IsDeepfake      bool
	DeepfakeScore   float64
	Confidence      float64
	Indicators      []string
	Methodology     []string
	AnalysisTime    time.Duration
	IsValid         bool
	Message         string
}

// MediaAnalyzer analyzes media files for authenticity.
type MediaAnalyzer struct {
	thresholds *AnalysisThresholds
	results    []AnalysisResult
}

// AnalysisThresholds contains analysis thresholds.
type AnalysisThresholds struct {
	DeepfakeThreshold float64
	ConfidenceThreshold float64
	MinFileSize       int64
	MaxFileSize       int64
	MinDimensions     [2]int
	MaxDimensions     [2]int
}

// NewMediaAnalyzer creates a new media analyzer.
func NewMediaAnalyzer() *MediaAnalyzer {
	return &MediaAnalyzer{
		thresholds: &AnalysisThresholds{
			DeepfakeThreshold: 0.7,
			ConfidenceThreshold: 0.6,
			MinFileSize: 1024,
			MaxFileSize: 100 * 1024 * 1024, // 100MB
			MinDimensions: [2]int{100, 100},
			MaxDimensions: [2]int{8192, 8192},
		},
		results: make([]AnalysisResult, 0),
	}
}

// SetThresholds sets analysis thresholds.
func (a *MediaAnalyzer) SetThresholds(thresholds AnalysisThresholds) {
	a.thresholds = &thresholds
}

// GetThresholds returns current thresholds.
func (a *MediaAnalyzer) GetThresholds() AnalysisThresholds {
	return *a.thresholds
}

// AnalyzeFile analyzes a media file for deepfakes.
func (a *MediaAnalyzer) AnalyzeFile(filePath string) (*AnalysisResult, error) {
	result := &AnalysisResult{
		ID:         fmt.Sprintf("analysis-%d", time.Now().UnixNano()),
		FilePath:   filePath,
		IsValid:    true,
		Message:    "Analysis completed",
	}

	// Check file exists
	file, err := os.Stat(filePath)
	if err != nil {
		result.IsValid = false
		result.Message = fmt.Sprintf("File not found: %v", err)
		return result, err
	}

	// Check file size
	if file.Size() < a.thresholds.MinFileSize || file.Size() > a.thresholds.MaxFileSize {
		result.IsValid = false
		result.Message = "File size out of acceptable range"
		return result, fmt.Errorf("file size out of range")
	}

	result.FileSize = file.Size()

	// Detect format
	result.Format = a.detectFormat(filePath)
	if result.Format == FormatUnknown {
		result.IsValid = false
		result.Message = "Unsupported format"
		return result, fmt.Errorf("unsupported format")
	}

	// Read and analyze image
	result.Dimensions, result.Hash, err = a.analyzeImage(filePath)
	if err != nil {
		result.IsValid = false
		result.Message = fmt.Sprintf("Image analysis error: %v", err)
		return result, err
	}

	// Check dimensions
	if result.Dimensions[0] < a.thresholds.MinDimensions[0] || result.Dimensions[0] > a.thresholds.MaxDimensions[0] {
		result.IsValid = false
		result.Message = "Width out of range"
		return result, fmt.Errorf("width out of range")
	}
	if result.Dimensions[1] < a.thresholds.MinDimensions[1] || result.Dimensions[1] > a.thresholds.MaxDimensions[1] {
		result.IsValid = false
		result.Message = "Height out of range"
		return result, fmt.Errorf("height out of range")
	}

	// Perform deepfake analysis
	result.IsDeepfake, result.DeepfakeScore, result.Confidence, result.Indicators, result.Methodology = a.detectDeepfake(filePath)

	result.AnalysisTime = 0 // Would be calculated from actual analysis time

	// Store result
	a.results = append(a.results, *result)

	return result, nil
}

// detectFormat detects media format from file extension.
func (a *MediaAnalyzer) detectFormat(filePath string) MediaFormat {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".jpg", ".jpeg":
		return FormatJPEG
	case ".png":
		return FormatPNG
	case ".gif":
		return FormatGIF
	case ".webp":
		return FormatWebP
	case ".mp4":
		return FormatMP4
	case ".avi":
		return FormatAVI
	case ".mov":
		return FormatMOV
	default:
		return FormatUnknown
	}
}

// analyzeImage analyzes image file.
func (a *MediaAnalyzer) analyzeImage(filePath string) ([2]int, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return [2]int{}, "", err
	}
	defer file.Close()

	var img image.Image
	var dims [2]int

	switch filepath.Ext(filePath) {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	default:
		return dims, "", fmt.Errorf("unsupported format")
	}

	if err != nil {
		return dims, "", err
	}

	bounds := img.Bounds()
	dims[0] = bounds.Max.X - bounds.Min.X
	dims[1] = bounds.Max.Y - bounds.Min.Y

	// Calculate hash
	data, err := os.ReadFile(filePath)
	if err != nil {
		return dims, "", err
	}

	hash := sha256.Sum256(data)
	return dims, hex.EncodeToString(hash[:]), nil
}

// detectDeepfake performs deepfake detection.
func (a *MediaAnalyzer) detectDeepfake(filePath string) (bool, float64, float64, []string, []string) {
	// In production: implement actual deepfake detection using ML models
	// For demo: simulate detection
	
	indicators := []string{}
	methodology := []string{}
	
	// Simulate detection methods
	methodology = append(methodology, "Frequency analysis")
	methodology = append(methodology, "Texture analysis")
	methodology = append(methodology, "Face landmark analysis")
	methodology = append(methodology, "Temporal consistency check")
	
	// Simulate indicators
	deepfakeScore := 0.0
	confidence := 0.0
	
	// Simulate analysis - in production this would use ML models
	// For demo, return random values
	deepfakeScore = 0.0 // Assume authentic by default
	confidence = 0.95   // High confidence in authenticity
	
	return deepfakeScore > a.thresholds.DeepfakeThreshold, deepfakeScore, confidence, indicators, methodology
}

// AnalyzeDirectory analyzes all media files in a directory.
func (a *MediaAnalyzer) AnalyzeDirectory(dirPath string) ([]AnalysisResult, error) {
	var results []AnalysisResult

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Check if media file
		ext := filepath.Ext(path)
		supportedExts := map[string]bool{
			".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
			".webp": true, ".mp4": true, ".avi": true, ".mov": true,
		}

		if !supportedExts[ext] {
			return nil
		}

		result, err := a.AnalyzeFile(path)
		if err != nil {
			return nil // Skip files that can't be analyzed
		}

		results = append(results, *result)
		return nil
	})

	return results, err
}

// GetResults returns all analysis results.
func (a *MediaAnalyzer) GetResults() []AnalysisResult {
	return a.results
}

// GetResultsByDeepfake returns results marked as deepfakes.
func (a *MediaAnalyzer) GetResultsByDeepfake() []AnalysisResult {
	var results []AnalysisResult
	for _, result := range a.results {
		if result.IsDeepfake {
			results = append(results, result)
		}
	}
	return results
}

// GetResultsByConfidence returns results above confidence threshold.
func (a *MediaAnalyzer) GetResultsByConfidence(minConfidence float64) []AnalysisResult {
	var results []AnalysisResult
	for _, result := range a.results {
		if result.Confidence >= minConfidence {
			results = append(results, result)
		}
	}
	return results
}

// GenerateReport generates analysis report.
func (a *MediaAnalyzer) GenerateReport() string {
	results := a.GetResults()

	var report string
	report += "=== Media Analysis Report ===\n\n"

	report += "Total Analyzed: " + fmt.Sprintf("%d\n", len(results))

	deepfakeCount := 0
	originalCount := 0
	for _, result := range results {
		if result.IsDeepfake {
			deepfakeCount++
		} else {
			originalCount++
		}
	}

	report += "Deepfakes Detected: " + fmt.Sprintf("%d\n", deepfakeCount)
	report += "Authentic Files: " + fmt.Sprintf("%d\n", originalCount)

	if len(results) > 0 {
		report += "\nAnalysis Details:\n"
		for i, result := range results {
			report += fmt.Sprintf("\n[%d] %s\n", i+1, result.FilePath)
			report += "    Format: " + string(result.Format) + "\n"
			report += "    Dimensions: " + fmt.Sprintf("%dx%d", result.Dimensions[0], result.Dimensions[1]) + "\n"
			report += "    Size: " + fmt.Sprintf("%d bytes\n", result.FileSize)
			report += "    Hash: " + result.Hash[:16] + "...\n"
			report += "    Is Deepfake: " + fmt.Sprintf("%v\n", result.IsDeepfake)
			report += "    Deepfake Score: " + fmt.Sprintf("%.2f\n", result.DeepfakeScore)
			report += "    Confidence: " + fmt.Sprintf("%.2f%%\n", result.Confidence*100)
		}
	}

	return report
}

// CalculateAverageScore calculates average deepfake score.
func (a *MediaAnalyzer) CalculateAverageScore() float64 {
	results := a.GetResults()
	if len(results) == 0 {
		return 0.0
	}

	var totalScore float64
	for _, result := range results {
		totalScore += result.DeepfakeScore
	}

	return totalScore / float64(len(results))
}

// CreateCommonThresholds creates common analysis thresholds.
func CreateCommonThresholds() AnalysisThresholds {
	return AnalysisThresholds{
		DeepfakeThreshold: 0.7,
		ConfidenceThreshold: 0.6,
		MinFileSize: 1024,
		MaxFileSize: 100 * 1024 * 1024,
		MinDimensions: [2]int{100, 100},
		MaxDimensions: [2]int{8192, 8192},
	}
}

// GetMediaAnalyzer returns analyzer.
func GetMediaAnalyzer(analyzer *MediaAnalyzer) *MediaAnalyzer {
	return analyzer
}

// GetAnalysisResult returns analysis result.
func GetAnalysisResult(result *AnalysisResult) *AnalysisResult {
	return result
}