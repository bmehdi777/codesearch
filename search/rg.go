package search

import (
	"codesearch/utils"
	"encoding/json"
	"errors"
	"strings"
)

const COMMAND = "rg"

type Ripgrep struct {
	Query   string
	Path    string
	Options []string
}

type ResultType string

const (
	ResultTypeBegin   ResultType = "begin"
	ResultTypeMatch   ResultType = "match"
	ResultTypeEnd     ResultType = "end"
	ResultTypeSummary ResultType = "summary"
)

type ResultSearch interface{}

type ResultSearchBegin struct {
	Type ResultType `json:"type"`
	Data struct {
		Path struct {
			Text string `json:"text"`
		} `json:"path"`
	} `json:"data"`
}

type ResultSearchMatch struct {
	Type ResultType `json:"type"`
	Data struct {
		Path struct {
			Text string `json:"text"`
		} `json:"path"`
		Lines struct {
			Text string `json:"text"`
		} `json:"lines"`
		LineNumber     int `json:"line_number"`
		AbsoluteOffset int `json:"absolute_offset"`
		Submatches     []struct {
			Match struct {
				Text string `json:"text"`
			} `json:"match"`
			Start int `json:"start"`
			End   int `json:"end"`
		} `json:"submatches"`
	} `json:"data"`
}

type ResultSearchEnd struct {
	Type ResultType `json:"type"`
	Data struct {
		Path struct {
			Text string `json:"text"`
		} `json:"path"`
		BinaryOffset any `json:"binary_offset"`
		Stats        struct {
			Elapsed struct {
				Secs  int    `json:"secs"`
				Nanos int    `json:"nanos"`
				Human string `json:"human"`
			} `json:"elapsed"`
			Searches          int `json:"searches"`
			SearchesWithMatch int `json:"searches_with_match"`
			BytesSearched     int `json:"bytes_searched"`
			BytesPrinted      int `json:"bytes_printed"`
			MatchedLines      int `json:"matched_lines"`
			Matches           int `json:"matches"`
		} `json:"stats"`
	} `json:"data"`
}

type ResultSearchSummary struct {
	Type ResultType `json:"type"`
	Data struct {
		ElapsedTotal struct {
			Human string `json:"human"`
			Nanos int    `json:"nanos"`
			Secs  int    `json:"secs"`
		} `json:"elapsed_total"`
		Stats struct {
			BytesPrinted  int `json:"bytes_printed"`
			BytesSearched int `json:"bytes_searched"`
			Elapsed       struct {
				Human string `json:"human"`
				Nanos int    `json:"nanos"`
				Secs  int    `json:"secs"`
			} `json:"elapsed"`
			MatchedLines      int `json:"matched_lines"`
			Matches           int `json:"matches"`
			Searches          int `json:"searches"`
			SearchesWithMatch int `json:"searches_with_match"`
		} `json:"stats"`
	} `json:"data"`
}

func (rg *Ripgrep) Search() error {
	exist, err := utils.VerifyPathExist(rg.Path)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("Path doesn't exist")
	}

	parameters := append([]string{rg.Query, rg.Path}, rg.Options...)

	result, err := utils.ExecCommand(COMMAND, parameters)
	if err != nil {
		return err
	}

	lineResult := strings.Split(result, "\n")

	for _, line := range lineResult {
		byteLine := []byte(line)
		rg.MapToStruct(byteLine)
		// convert these struct into a local search struct
	}

	return nil
}

func (rg *Ripgrep) MapToStruct(line []byte) (ResultSearch, error) {
	var m map[string]interface{}
	err := json.Unmarshal(line, &m)
	if err != nil {
		return nil, err
	}

	var resultSearch ResultSearch

	switch m["type"].(string) {
	case "begin":
		resultSearch = &ResultSearchBegin{}
	case "match":
		resultSearch = &ResultSearchMatch{}
	case "end":
		resultSearch = &ResultSearchEnd{}
	case "summary":
		resultSearch = &ResultSearchSummary{}
	}

	err = json.Unmarshal(line, resultSearch)

	return resultSearch, err
}
