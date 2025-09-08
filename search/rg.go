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

type RGResultSearch interface{}

type RGResultSearchBegin struct {
	Type ResultType `json:"type"`
	Data struct {
		Path struct {
			Text string `json:"text"`
		} `json:"path"`
	} `json:"data"`
}

type RGResultSearchMatch struct {
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

type RGResultSearchEnd struct {
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

type RGResultSearchSummary struct {
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

func (rg *Ripgrep) Search() ([]SearchResult, error) {
	exist, err := utils.VerifyPathExist(rg.Path)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New("Path doesn't exist")
	}

	parameters := append([]string{rg.Query, rg.Path}, rg.Options...)

	result, err := utils.ExecCommand(COMMAND, parameters)
	if err != nil {
		return nil, err
	}

	lineResult := strings.Split(result, "\n")

	searchResults := make([]SearchResult, 0)


	// these two blockes certainly need a refacto
	rgResults := make([]RGResultSearch, 0)
	for _, line := range lineResult {
		byteLine := []byte(line)
		res, err := rg.MapJSONToRGStruct(byteLine)
		if err != nil {
			return nil, err
		}

		rgResults = append(rgResults, res)
	}

	return searchResults, nil
}

func (rg *Ripgrep) MapJSONToRGStruct(line []byte) (RGResultSearch, error) {
	var m map[string]interface{}
	err := json.Unmarshal(line, &m)
	if err != nil {
		return nil, err
	}

	var resultSearch RGResultSearch

	switch m["type"].(string) {
	case "begin":
		resultSearch = &RGResultSearchBegin{}
	case "match":
		resultSearch = &RGResultSearchMatch{}
	case "end":
		resultSearch = &RGResultSearchEnd{}
	case "summary":
		resultSearch = &RGResultSearchSummary{}
	}

	err = json.Unmarshal(line, resultSearch)

	return resultSearch, err
}

func (rg *Ripgrep) MapToSearchResult(searchInput RGResultSearch) (SearchResult, error) {

	return SearchResult{}, nil
}
