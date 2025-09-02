package search

type ResultType string

const (
	ResultTypeBegin   ResultType = "begin"
	ResultTypeMatch   ResultType = "match"
	ResultTypeEnd     ResultType = "end"
	ResultTypeSummary ResultType = "summary"
)


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
	Type string `json:"type"`
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
