package repository

import "github.com/KowalskiPiotr98/gotabase"

// ScanObjects iterates through the result rows and calls the scan function for each one.
func ScanObjects[T any](rows gotabase.Rows, scan func(row gotabase.Row) (*T, error)) ([]*T, error) {
	results := make([]*T, 0)
	for rows.Next() {
		row, err := scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}
