package engine

import (
	"fmt"

	"github.com/StackExchange/wmi"
)

type Win32_OperatingSystem struct {
	Caption     *string
	OSArchitecture *string
	InstallDate *string
	CSName     *string
}

type InventoryItem struct {
	ComputerName string
	OS           string
	Arch         string
	InstallDate  string
}

func CollectWMIInventory() ([]InventoryItem, error) {
	var dst []Win32_OperatingSystem
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		return nil, err
	}
	items := make([]InventoryItem, 0)
	for _, d := range dst {
		item := InventoryItem{
			ComputerName: safeString(d.CSName),
			OS:           safeString(d.Caption),
			Arch:         safeString(d.OSArchitecture),
			InstallDate:  safeString(d.InstallDate),
		}
		items = append(items, item)
	}
	return items, nil
}

func safeString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}