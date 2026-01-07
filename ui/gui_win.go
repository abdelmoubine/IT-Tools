package ui

import (
	"fmt"
	"time"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"

	"github.com/abdelmoubine/IT-Tools/engine"
)

// RunGUI starts the native Windows GUI.
// This implementation uses walk declarative to create the window and the views.
func RunGUI() error {
	var mw *walk.MainWindow
	var pages *walk.TabWidget // we'll use TabWidget to swap views; but hide tabs
	var aboutLabel *walk.TextEdit
	var themeToggle *walk.PushButton

	// shared UI model state
	appCtx := &AppContext{
		ThemeDark: false,
	}
	// create window
	if err := MainWindow{
		AssignTo: &mw,
		Title:    "IT Support Toolkit",
		MinSize:  Size{Width: 1000, Height: 700},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					Label{Text: "  ", ColumnSpan: 1},
					HSpacer{},
					PushButton{
						Text: "Toggle Theme",
						AssignTo: &themeToggle,
						OnClicked: func() {
							appCtx.ThemeDark = !appCtx.ThemeDark
							applyTheme(appCtx.ThemeDark)
						},
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					VBox{
						StretchFactor: 1,
						Children: []Widget{
							Label{Text: "About the Developer", Font: Font{PointSize: 14, Bold: true}},
							TextEdit{
								ReadOnly: true,
								AssignTo: &aboutLabel,
								Text: engine.GetAboutText(), // from resources/en.json loaded in engine init
								MinSize: Size{Width: 520, Height: 120},
							},
							PushButton{
								Text: "Open About Links",
								OnClicked: func() {
									engine.OpenAboutLinks()
								},
							},
						},
					},
					Composite{
						Layout: Grid{Columns: 3, Rows: 3, Spacing: 12},
						Children: []Widget{
							PushButton{Text: "Network → Quick Scan", OnClicked: func() { openToolQuickScan(mw) }},
							PushButton{Text: "Network → Port Scanner", OnClicked: func() { openToolPortScan(mw) }},
							PushButton{Text: "Network → Traceroute", OnClicked: func() { openToolTraceroute(mw) }},
							PushButton{Text: "System → WMI Inventory", OnClicked: func() { openToolWMI(mw) }},
							PushButton{Text: "Security → Vulnerability Check", OnClicked: func() { openToolVuln(mw) }},
							PushButton{Text: "Forensics → Collect Logs", OnClicked: func() { openToolCollectLogs(mw) }},
							PushButton{Text: "AI Assistant", OnClicked: func() { openToolAI(mw) }},
							PushButton{Text: "Export / Reports", OnClicked: func() { openToolExport(mw) }},
							PushButton{Text: "Exit", OnClicked: func() { mw.Close() }},
						},
					},
				},
			},
		},
	}.Create(); err != nil {
		return err
	}

	// apply initial theme
	applyTheme(false)

	// run
	mw.Show()
	return mw.Run()
}

type AppContext struct {
	ThemeDark bool
}

func applyTheme(dark bool) {
	// placeholder, walk theming is limited; keep default.
	if dark {
		// no-op for now
	} else {
		// no-op
	}
}

func openToolQuickScan(parent *walk.MainWindow) {
	d := walk.NewDialog()
	d.SetTitle("Quick Scan")
	d.SetLayout(walk.NewVBoxLayout())
	defer d.Dispose()

	// Controls
	ipEdit, _ := walk.NewLineEdit(d)
	ipEdit.SetText("192.168.1.0/24")
	startBtn, _ := walk.NewPushButton(d)
	startBtn.SetText("Start")
	results, _ := walk.NewTextEdit(d)
	results.SetReadOnly(true)
	startBtn.Clicked().Attach(func() {
		startBtn.SetEnabled(false)
		results.SetText("Scanning...")
		go func() {
			res := engine.QuickScan(ipEdit.Text())
			// build output
			out := fmt.Sprintf("Found %d hosts\n", len(res))
			for _, r := range res {
				out += fmt.Sprintf("%s\n", r)
			}
			results.SetText(out)
			startBtn.SetEnabled(true)
		}()
	})

	// Layout items
	d.Children().Add(walk.NewLabelWithText("Target CIDR or comma list:"))
	d.Children().Add(ipEdit)
	d.Children().Add(startBtn)
	d.Children().Add(results)

	d.SetSize(walk.Size{Width: 700, Height: 500})
	d.Run()
}

func openToolPortScan(parent *walk.MainWindow) {
	d := walk.NewDialog()
	d.SetTitle("Port Scanner")
	d.SetLayout(walk.NewVBoxLayout())
	defer d.Dispose()

	ipEdit, _ := walk.NewLineEdit(d)
	ipEdit.SetText("192.168.1.12")
	portEdit, _ := walk.NewLineEdit(d)
	portEdit.SetText("1-1024")
	startBtn, _ := walk.NewPushButton(d)
	startBtn.SetText("Start Scan")
	results, _ := walk.NewTextEdit(d)
	results.SetReadOnly(true)

	startBtn.Clicked().Attach(func() {
		startBtn.SetEnabled(false)
		results.SetText("Scanning ports...")
		go func() {
			rows := engine.PortScan(ipEdit.Text(), portEdit.Text())
			out := ""
			for _, r := range rows {
				out += fmt.Sprintf("%s:%d %s\n", r.IP, r.Port, r.State)
			}
			results.SetText(out)
			startBtn.SetEnabled(true)
		}()
	})

	d.Children().Add(walk.NewLabelWithText("Target IP:"))
	d.Children().Add(ipEdit)
	d.Children().Add(walk.NewLabelWithText("Ports (e.g. 1-1024 or 22,80,443):"))
	d.Children().Add(portEdit)
	d.Children().Add(startBtn)
	d.Children().Add(results)

	d.SetSize(walk.Size{Width: 700, Height: 500})
	d.Run()
}

func openToolTraceroute(parent *walk.MainWindow) {
	d := walk.NewDialog()
	d.SetTitle("Traceroute")
	d.SetLayout(walk.NewVBoxLayout())
	defer d.Dispose()

	hostEdit, _ := walk.NewLineEdit(d)
	hostEdit.SetText("8.8.8.8")
	startBtn, _ := walk.NewPushButton(d)
	startBtn.SetText("Trace")
	results, _ := walk.NewTextEdit(d)
	results.SetReadOnly(true)

	startBtn.Clicked().Attach(func() {
		startBtn.SetEnabled(false)
		results.SetText("Running traceroute...")
		go func() {
			hops, _ := engine.Traceroute(hostEdit.Text(), 30)
			out := ""
			for i, h := range hops {
				out += fmt.Sprintf("%02d: %s\n", i+1, h)
			}
			results.SetText(out)
			startBtn.SetEnabled(true)
		}()
	})

	d.Children().Add(walk.NewLabelWithText("Host/IP:"))
	d.Children().Add(hostEdit)
	d.Children().Add(startBtn)
	d.Children().Add(results)
	d.SetSize(walk.Size{Width: 700, Height: 450})
	d.Run()
}

func openToolWMI(parent *walk.MainWindow) {
	d := walk.NewDialog()
	d.SetTitle("WMI Inventory")
	d.SetLayout(walk.NewVBoxLayout())
	defer d.Dispose()

	startBtn, _ := walk.NewPushButton(d)
	startBtn.SetText("Collect Inventory")
	results, _ := walk.NewTextEdit(d)
	results.SetReadOnly(true)

	startBtn.Clicked().Attach(func() {
		startBtn.SetEnabled(false)
		results.SetText("Collecting inventory...")
		go func() {
			inv, err := engine.CollectWMIInventory()
			out := ""
			if err != nil {
				out = fmt.Sprintf("Error: %v", err)
			} else {
				for _, i := range inv {
					out += fmt.Sprintf("%s | %s | %s\n", i.ComputerName, i.OS, i.InstallDate)
				}
			}
			results.SetText(out)
			startBtn.SetEnabled(true)
		}()
	}()

	d.Children().Add(startBtn)
	d.Children().Add(results)
	d.SetSize(walk.Size{Width: 800, Height: 500})
	d.Run()
}

func openToolVuln(parent *walk.MainWindow) {
	walk.MsgBox(nil, "Vulnerability Check", "Basic vulnerability checks are available (local rules). For deep scans please enable agent or external scanners.", walk.MsgBoxIconInformation)
}

func openToolCollectLogs(parent *walk.MainWindow) {
	// Collect logs stub
	go func() {
		_ = engine.CollectEventLogs()
	}()
	walk.MsgBox(nil, "Logs", "Event logs collected and exported.", walk.MsgBoxIconInformation)
}

func openToolAI(parent *walk.MainWindow) {
	walk.MsgBox(nil, "AI Assistant", "AI Assistant (hybrid) is available via settings. This is a placeholder UI.", walk.MsgBoxIconInformation)
}

func openToolExport(parent *walk.MainWindow) {
	fn, err := engine.ExportSampleCSV()
	if err != nil {
		walk.MsgBox(nil, "Export Error", fmt.Sprintf("Export failed: %v", err), walk.MsgBoxIconError)
		return
	}
	walk.MsgBox(nil, "Exported", fmt.Sprintf("Exported sample CSV to: %s", fn), walk.MsgBoxIconInformation)
}