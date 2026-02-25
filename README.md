# deepscan - Deepfake Detection & Media Authentication

[![Go](https://img.shields.io/badge/Go-1.21-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

**Detect deepfakes and verify media authenticity with forensic analysis.**

Advanced deepfake detection and media authentication tool for verifying digital content.

## 🚀 Features

- **Deepfake Detection**: Identify AI-generated fake media
- **Media Authentication**: Verify media authenticity
- **Forensic Analysis**: Comprehensive forensic examination
- **Multiple Analysis Methods**: Frequency, texture, landmark analysis
- **Metadata Verification**: Validate EXIF and other metadata
- **Blockchain Verification**: Verify media via blockchain registry

## 📦 Installation

### Build from Source

```bash
git clone https://github.com/hallucinaut/deepscan.git
cd deepscan
go build -o deepscan ./cmd/deepscan
sudo mv deepscan /usr/local/bin/
```

### Install via Go

```bash
go install github.com/hallucinaut/deepscan/cmd/deepscan@latest
```

## 🎯 Usage

### Analyze Media

```bash
# Analyze media for deepfakes
deepscan analyze /path/to/image.jpg
```

### Verify Authenticity

```bash
# Verify media authenticity
deepscan verify /path/to/image.jpg
```

### Forensic Analysis

```bash
# Perform forensic analysis
deepscan forensic /path/to/image.jpg
```

### Compare Media

```bash
# Compare two media files
deepscan compare image1.jpg image2.jpg
```

### Generate Report

```bash
# Generate analysis report
deepscan report
```

### Programmatic Usage

```go
package main

import (
    "fmt"
    "github.com/hallucinaut/deepscan/pkg/analyze"
    "github.com/hallucinaut/deepscan/pkg/verify"
    "github.com/hallucinaut/deepscan/pkg/forensic"
)

func main() {
    // Create analyzer
    analyzer := analyze.NewMediaAnalyzer()
    
    // Analyze file
    result, err := analyzer.AnalyzeFile("/path/to/image.jpg")
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Is Deepfake: %v\n", result.IsDeepfake)
    fmt.Printf("Confidence: %.2f%%\n", result.Confidence*100)
    
    // Verify media
    verifyResult, err := verify.VerifyFile("/path/to/image.jpg")
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Verification: %v\n", verifyResult.IsValid)
    
    // Forensic analysis
    forensicAnalyzer := forensic.NewForensicAnalyzer()
    forensicResult, err := forensicAnalyzer.AnalyzeFile("/path/to/image.jpg")
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Authentic: %v\n", forensicResult.IsAuthentic)
}
```

## 📚 Detection Methods

### Frequency Analysis
- Analyze frequency domain artifacts
- Detect unnatural frequency patterns
- Identify AI-generated frequencies

### Texture Analysis
- Examine texture consistency
- Detect synthetic textures
- Identify smoothing artifacts

### Face Landmark Analysis
- Analyze facial feature consistency
- Detect landmark irregularities
- Identify manipulation artifacts

### Temporal Consistency
- Check frame-to-frame consistency
- Detect temporal anomalies
- Identify synthetic motion

## 🧪 Analysis Methods

| Method | Description |
|--------|-------------|
| Frequency | Analyze frequency domain |
| Texture | Analyze texture patterns |
| Landmark | Analyze facial landmarks |
| Metadata | Verify EXIF metadata |
| Forensic | Comprehensive forensic analysis |

## 🏗️ Architecture

```
deepscan/
├── cmd/
│   └── deepscan/
│       └── main.go          # CLI entry point
├── pkg/
│   ├── analyze/
│   │   ├── analyze.go      # Media analysis
│   │   └── analyze_test.go # Unit tests
│   ├── verify/
│   │   ├── verify.go       # Media verification
│   │   └── verify_test.go  # Unit tests
│   └── forensic/
│       ├── forensic.go     # Forensic analysis
│       └── forensic_test.go # Unit tests
└── README.md
```

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test -v ./pkg/analyze -run TestAnalyzeFile
```

## 📋 Example Output

```
$ deepscan analyze /path/to/image.jpg

Media Analysis for Deepfake Detection
======================================

Analysis Result:
  File: /path/to/image.jpg
  Format: JPEG
  Dimensions: 1920x1080
  Size: 245678 bytes
  Hash: a1b2c3d4e5f67890...
  Is Deepfake: false
  Deepfake Score: 5.00%
  Confidence: 95.00%
  Status: Valid
```

## 🔒 Security Use Cases

- **Deepfake Detection**: Detect AI-generated fake media
- **Media Authentication**: Verify media authenticity
- **Forensic Investigation**: Investigate media authenticity
- **Content Verification**: Verify social media content
- **Legal Evidence**: Authenticate digital evidence

## 🛡️ Best Practices

1. **Use multiple methods**: Combine detection methods
2. **Verify metadata**: Check EXIF and other metadata
3. **Compare with originals**: If available
4. **Check blockchain**: Verify via blockchain registry
5. **Maintain chain of custody**: For legal evidence
6. **Regular updates**: Keep detection models updated

## 📄 License

MIT License

## 🙏 Acknowledgments

- Deepfake research community
- Media forensics researchers
- AI security community

## 🔗 Resources

- [Deepfake Detection Challenge](https://www.kaggle.com/c/deepfake-detection-challenge)
- [FakeCatcher](https://www.researchgate.net/publication/347574893_FakeCatcher)
- [Forensic Image Analysis](https://www.ieice.org/~is&f/forensics.html)

---

**build with GPU by [hallucinaut](https://github.com/hallucinaut)**