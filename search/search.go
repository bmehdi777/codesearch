package search

// this is a search result for a file only
type SearchResult struct {
	Path       string `json:"path"`
	NumResults int `json:"num_results"`
	MatchLines []MatchLine `json:"match_lines"`
}

type MatchLine struct {
	LineContent string `json:"line_content"`
	LineNumber  int `json:"line_number"`
}

type Searcher interface {
	Search() ([]SearchResult, error)
}
