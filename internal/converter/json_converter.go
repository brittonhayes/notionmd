package converter

import (
	"encoding/json"

	"github.com/dstotijn/go-notion"
)

// ConvertToJSON takes a markdown document and returns Notion blocks as JSON
func ConvertToJSON(markdown string) ([]map[string]any, error) {
	// Use the existing converter to get notion blocks
	blocks, err := Convert(markdown)
	if err != nil {
		return nil, err
	}

	// Convert each block to a map using JSON marshaling
	var jsonBlocks []map[string]any
	for _, block := range blocks {
		if jsonBlock := blockToMap(block); jsonBlock != nil {
			jsonBlocks = append(jsonBlocks, jsonBlock)
		}
	}

	return jsonBlocks, nil
}

// blockToMap converts a notion.Block to a map[string]any
func blockToMap(block notion.Block) map[string]any {
	// Use JSON marshaling to convert the block to a map
	jsonData, err := json.Marshal(block)
	if err != nil {
		return nil
	}

	var result map[string]any
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil
	}

	return result
}
