package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/abdelmoubine/IT-Tools/ui"
	"github.com/abdelmoubine/IT-Tools/engine"
)

func main() {
	// CLI flags for quick tasks
	exportFlag := flag.Bool("export", false, "Export a sample CSV/TXT report")
	scanNet := flag.String("scan", "", "Quick scan CIDR or comma list of IPs")
	guiFlag := flag.Bool("gui", true, "Start GUI (default true)")
	flag.Parse()

	// Ensure runtime directories
	os.MkdirAll("exports", 0755)
	os.MkdirAll("tmp", 0700)

	// If CLI quick tasks requested
	if *exportFlag {
		fn, err := engine.ExportSampleCSV()
		if err != nil {
			fmt.Println("Export error:", err)
			os.Exit(1)
		}
		fmt.Println("Exported sample CSV:", filepath.Clean(fn))
		return
	}
	if *scanNet != "" {
		results := engine.QuickScan(*scanNet)
		fmt.Printf("Quick scan results: %+v\n", results)
		return
	}

	// Default: start GUI
	if *guiFlag {
		err := ui.RunGUI()
		if err != nil {
			fmt.Println("GUI error:", err)
			os.Exit(1)
		}
		return
	}
}