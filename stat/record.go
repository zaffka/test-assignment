package stat

import (
	"os"
	"sort"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// Record represents a single record for any "name-score" evaluation.
type Record struct {
	Name  string
	Score int
}

// Records represents a slice of stat.Record items.
type Records []*Record

// SortedByScore is sorting a stat.Records slice by score.
func (rs Records) SortedByScore() Records {
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].Score > rs[j].Score
	})

	return rs
}

// PrintTable prints a stat.Recors slice as a ASCII table to os.Stdout.
func (rs Records) PrintTable(statType string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{statType, "Score"})

	for _, rec := range rs.SortedByScore()[:10] {
		table.Append([]string{rec.Name, strconv.Itoa(rec.Score)})
	}
	table.Render()
}
