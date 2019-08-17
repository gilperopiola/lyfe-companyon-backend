package main

import "fmt"

type SearchParameters struct {
	Filter        string
	FilterTagID   int
	FilterImportance int
	SortField     string
	SortDirection string
	Limit         int
	Offset        int
}

func getSearchOrderBy(params *SearchParameters) string {
	orderByString := "id ASC"

	if params.SortField != "" && params.SortDirection != "" {
		orderByString = params.SortField + " " + params.SortDirection
	}

	return orderByString
}

func getSearchURL(params *SearchParameters) string {
	return fmt.Sprintf("?filter=%s&sortField=%s&sortDirection=%s&limit=%d&offset=%d",
		params.Filter, params.SortField, params.SortDirection, params.Limit, params.Offset)
}
