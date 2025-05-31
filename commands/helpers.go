package commands

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/zigzagalex/gator/internal/database"
)

func PrettyPrintFeeds(feeds []database.GetFeedsRow) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Headers
	fmt.Fprintln(w, "Created At\tFeed Name\tURL\tOwner")

	// Rows
	for _, f := range feeds {
		owner := "NULL"
		if f.Name_2.Valid {
			owner = f.Name_2.String
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			f.CreatedAt.Format(time.RFC3339),
			f.Name,
			f.Url,
			owner,
		)
	}

	w.Flush()
}
