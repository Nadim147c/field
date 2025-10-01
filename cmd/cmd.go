package cmd

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"unicode"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var (
	delimiter   = "space"
	ignoreEmpty = false
	limit       = math.MaxInt
)

func init() {
	flags := Command.Flags()
	flags.StringVarP(&delimiter, "delimiter", "d", delimiter, "delimiter for field separation")
	flags.BoolVarP(&ignoreEmpty, "ignore-empty", "i", ignoreEmpty, "ignore empty lines")
	flags.IntVarP(&limit, "limit", "n", limit, "number of field to separate")

	if slices.Contains(os.Args, "_carapace") {
		carapace.Gen(Command)
	}
}

// Command is root command
var Command = &cobra.Command{
	Use:   "field [--flags] ...<range>",
	Short: "Extract and print selected fields from each input line",
	Example: `
# Extract the second field from ps output (PID) and kill those processes
ps aux | grep bad-process | field 2 | xargs kill

# Print only the usernames (first field) from /etc/passwd and ignore empty lines
cat /etc/passwd | field -i -d: 1

# Show just the command (limit number of field 11) from ps output
ps aux | field -n11 11

# Extract multiple fields (user and PID) and print them
ps aux | field 1 2
`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ranges := make([]*Range, len(args))
		for i, a := range args {
			r, err := ParseRange(a, false)
			if err != nil {
				return err
			}
			ranges[i] = r
		}

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			var fields []string
			if cmd.Flags().Changed("delimiter") {
				fields = FieldN(line, delimiter, limit)
			} else {
				fields = FieldNPred(line, unicode.IsSpace, limit)
			}

			selected := make([]string, 0, 10)
			for r := range slices.Values(ranges) {
				selected = append(selected, r.Select(fields)...)
			}

			if ignoreEmpty && len(selected) == 0 {
				continue
			}

			fmt.Println(strings.Join(selected, " "))
		}

		return scanner.Err()
	},
}
