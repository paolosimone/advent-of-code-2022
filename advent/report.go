package advent

import (
	"fmt"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/samber/lo"
)

type ReportRow [6]string

type Report struct {
	Header ReportRow
	Rows   []ReportRow
}

func BuildReport(results []DayResult) Report {
	return Report{
		Header: header,
		Rows:   lo.Map(results, func(result DayResult, _ int) ReportRow { return buildReportRow(result) }),
	}
}

func (report Report) Sprint() string {
	tableBuilder := &strings.Builder{}
	table := tablewriter.NewWriter(tableBuilder)
	table.SetAutoFormatHeaders(false)

	table.SetHeader(report.Header[:])
	for _, row := range report.Rows {
		table.Append(row[:])
	}

	table.Render()
	return tableBuilder.String()
}

var header = ReportRow{
	"Day",
	"LoadElapsed",
	"FirstResult",
	"FirstElapsed",
	"SecondResult",
	"SecondElapsed",
}

func buildReportRow(result DayResult) ReportRow {
	return ReportRow{
		fmt.Sprint(result.Day),
		formatDuration(result.LoadElapsed),
		result.FirstResult,
		formatDuration(result.FirstElapsed),
		result.SecondResult,
		formatDuration(result.SecondElapsed),
	}
}

func formatDuration(duration time.Duration) string {
	return fmt.Sprintf("%.6f s", duration.Seconds())
}
