package engine

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// ExportSampleCSV creates an example CSV with metadata and sample rows
func ExportSampleCSV() (string, error) {
	fn := fmt.Sprintf("exports/scan_%d.csv", time.Now().Unix())
	f, err := os.Create(fn)
	if err != nil {
		return "", err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	// metadata
	w.Write([]string{"file_version", "app_version", "scan_type", "start_time", "end_time", "operator", "language"})
	w.Write([]string{"1.0", "1.0", "quick_scan", time.Now().Add(-time.Minute).Format(time.RFC3339), time.Now().Format(time.RFC3339), "operator", "en"})
	w.Write([]string{})
	// headers
	w.Write([]string{"scan_id", "ip", "hostname", "port", "protocol", "state", "service", "banner", "timestamp"})
	// sample rows
	w.Write([]string{"scan-001", "192.168.1.12", "printer-office", "80", "tcp", "open", "http", "Apache/2.4.41", time.Now().Format(time.RFC3339)})
	w.Write([]string{"scan-001", "192.168.1.45", "workstation-23", "3389", "tcp", "filtered", "rdp", "-", time.Now().Format(time.RFC3339)})
	return fn, nil
}