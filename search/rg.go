package search

import (
	"hsh/utils"
	"encoding/json"
	"errors"
	"fmt"
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
	Data struct {
		Path struct {
			Text string `json:"text"`
		} `json:"path"`
	} `json:"data"`
}
type RGDataMatch struct {
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
type RGDataEnd struct {
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
type RGDataSummary struct {
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

	parameters := append([]string{rg.Query, rg.Path, "--json"}, rg.Options...)

	result, err := utils.ExecCommand(COMMAND, parameters)
	if err != nil {
		fmt.Println("err here?")
		return nil, err
	}

	lineResult := strings.Split(result, "\n")

	searchResults := make([]SearchResult, 0)

	// last item is an empty new line
	for _, line := range lineResult[:len(lineResult)-1] {
		bytesLine := []byte(line)

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
			newSearchRes := SearchResult{
				Path:       begin.Data.Path.Text,
				MatchLines: make([]MatchLine, 0),
			}
			searchResults = append(searchResults, newSearchRes)
		case "match":
			var match RGDataMatch
			err = json.Unmarshal(bytesLine, &match)
			if err != nil {
				return nil, err
			}
			lastSearchRes := &searchResults[len(searchResults)-1]
			newMatch := MatchLine{
				LineContent: match.Data.Lines.Text,
				LineNumber:  match.Data.LineNumber,
			}
			lastSearchRes.MatchLines = append(lastSearchRes.MatchLines, newMatch)
		case "end":
			var end RGDataEnd
			err = json.Unmarshal(bytesLine, &end)
			if err != nil {
				return nil, err
			}
			lastSearchRes := &searchResults[len(searchResults)-1]
			lastSearchRes.NumResults = end.Data.Stats.Matches
		case "summary":
			// for now, we ignore summary
		}
	}

	return searchResults, nil
}
