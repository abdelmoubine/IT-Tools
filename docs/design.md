# Design - IT Support Toolkit (English-only)

Overview
- Single binary portable desktop app for Windows 7+.
- Written in Go with native UI using walk.
- Offline-first, with hybrid optional cloud AI and optional agent.

Components
- main.go: bootstraps CLI or GUI.
- ui/: GUI implementation and dialogs for each tool.
- engine/: core functionality (network, ports, whois, WMI, export).
- agent/: optional background service for update and management.
- server/: management server skeleton for pushing signed resources.

Security
- All exported artifacts include metadata and can be encrypted.
- Temporary files stored in tmp/ and are overwritten and deleted on request (best-effort).
- Remote updates require resources to be signed and server trusted.

Packaging
- build scripts produce portable EXE for x86/x64.
- package_zip.sh produces portable ZIP including resources.

Extensibility
- Plugins support: future work includes WASM-based plugins loaded and sandboxed.
- ML: ONNX inference integration can be added for predictive maintenance.