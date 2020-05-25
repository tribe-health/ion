package proto

import "encoding/json"

// ProcessInfo describes an a/v process
type ProcessInfo struct {
	MID          string          `json:"mid"`
	RID          string          `json:"rid"`
	Processor    string          `json:"processor"`
	ClientConfig json.RawMessage `json:"config"`
}
