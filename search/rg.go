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

type RGResultSearch struct {
	Type ResultType      `json:"type"`
	Data json.RawMessage `json:"data"`
}
type RGDataBegin struct {
	Path struct {
		Text string `json:"text"`
	} `json:"path"`
}
type RGDataMatch struct {
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
}
type RGDataEnd struct {
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
}
type RGDataSummary struct {
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

	for i := 0; i < len(lineResult); i++ {
		bytesLine := []byte(lineResult[i])

		var rgResult RGResultSearch
		err := json.Unmarshal(bytesLine, &rgResult)
		if err != nil {
			return nil, err
		}

		switch rgResult.Type {
		case "begin":
			var begin RGDataBegin
			err = json.Unmarshal(bytesLine, &begin)
			if err != nil {
				return nil, err
			}
			newSearchRes := SearchResult{}
			searchResults = append(searchResults, newSearchRes)
		case "match":
			var match RGDataMatch
			err = json.Unmarshal(bytesLine, &match)
			if err != nil {
				return nil, err
			}
			lastSearchRes := searchResults[len(searchResults)-1]
			matchLine := MatchLine{
				LineContent: match.Lines.Text,
				LineNumber: match.LineNumber,
			}
			lastSearchRes.MatchLines = append(lastSearchRes.MatchLines, matchLine)
		case "end":
			var end RGDataEnd
			err = json.Unmarshal(bytesLine, &end)
			if err != nil {
				return nil, err
			}
			//lastSearchRes := searchResults[len(searchResults)-1]


		case "summary":
			// for now, we ignore summary
		}
	}

	return searchResults, nil
}
