package report

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/BShaT/qy-api-change-toolkit/internal/diff"
)

func JSON(changes []diff.Change) ([]byte, error) {
	return json.MarshalIndent(changes, "", "  ")
}

func Markdown(changes []diff.Change) []byte {
	var buffer bytes.Buffer
	buffer.WriteString("# Qy API Change Report\n\n")
	if len(changes) == 0 {
		buffer.WriteString("No changes detected.\n")
		return buffer.Bytes()
	}
	buffer.WriteString("| Severity | Type | Path | Message |\n")
	buffer.WriteString("|---|---|---|---|\n")
	for _, change := range changes {
		buffer.WriteString(fmt.Sprintf("| %s | %s | `%s` | %s |\n",
			change.Severity,
			change.Type,
			escapeMarkdown(change.Path),
			escapeMarkdown(change.Message),
		))
	}
	return buffer.Bytes()
}

func escapeMarkdown(value string) string {
	return strings.ReplaceAll(value, "|", "\\|")
}
