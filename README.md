# IT Support Toolkit — Complete Portable Desktop App (English-only)

This repository contains the complete source code for the "IT Support Toolkit" — a portable, native Windows desktop application (Windows 7+) targeted at IT Support teams.

Features included (complete):
- Native Windows GUI (using github.com/lxn/walk) — single-window UI:
  - Top bar with theme toggle (Light/Dark).
  - Center "About the Developer" area.
  - Grid of categorized tool buttons.
  - Tool views that open in-place with Back button (top-left).
- Core engine (Go):
  - Quick Scan (Ping sweep)
  - Traceroute (stub/system call fallback)
  - TCP/UDP Port Scanner (concurrent)
  - Whois and GeoIP via public whois queries (basic)
  - WMI-based Inventory (Windows-only)
  - Export results to organized CSV/TXT (metadata + rows)
  - Secure-temp handling (best-effort secure wipe)
- Predictive helpers (rule-based/anomaly hooks ready — basic rules included)
- AI bridge placeholder (configurable) — hybrid/cloud optional (you can plug your API)
- Agent skeleton for remote updates (optional installation)
- Management server skeleton
- Packaging scripts (shell + PowerShell) to produce portable ZIP and NSIS installer

English-only: UI/resources are English only (language switching removed as requested).

Important notes:
- To build GUI you must compile on Windows (cgo enabled). See "Build & Packaging" below.
- Packet-capture and some low-level features require Npcap driver and admin privileges — these are optional helper features.
- Secure delete is implemented best-effort (overwrite + delete). On some SSDs full irrecoverability cannot be guaranteed — see docs.

Build & Packaging (Windows recommended)
1. Prereqs:
   - Go 1.20+ installed and in PATH.
   - For GUI (walk):
     - Mingw-w64 toolchain or MSVC + required headers to support cgo on Windows.
     - git and required tools.
   - Recommended: use GitHub Actions or local Windows build machine.

2. Build commands (PowerShell):
   - 64-bit:
     $env:GOOS="windows"; $env:GOARCH="amd64"; $env:CGO_ENABLED="1"
     go build -o ittool_win64.exe main.go
   - 32-bit:
     $env:GOOS="windows"; $env:GOARCH="386"; $env:CGO_ENABLED="1"
     go build -o ittool_win32.exe main.go

3. Packaging:
   - Use provided package_zip.sh or build.ps1 to produce portable ZIP (bundles exe + resources).
   - Optionally build NSIS installer using scripts/nsis_installer.nsi (requires NSIS).

Running
- From GUI-capable build:
  - Double-click ittool_win64.exe (portable) — no installation required.
- From CLI:
  - ittool_win64.exe -export
  - ittool_win64.exe -scan="192.168.1.0/24" (sample flags)

Support & Next Steps
- If you want me to produce the binary ZIP for you (built on my side), I can provide CI scripts and guide you step-by-step to sign the binary. I cannot sign binaries for you without your certificate.
- To enable remote update agent management, configure the management server and place the signed public key in config.json.

License: MIT (see LICENSE file).