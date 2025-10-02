package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"math"
	"os"
	"slices"
	"strings"
	"unicode"

	"github.com/carapace-sh/carapace"
	shlex "github.com/carapace-sh/carapace-shlex"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/valyala/fasttemplate"
)

var (
	delimiter   = "space"
	format      = "none"
	ignoreEmpty = false
	shell       = false
)

var limit limitValue = math.MaxInt

func init() {
	flags := Command.Flags()
	flags.BoolVarP(&ignoreEmpty, "ignore-empty", "i", ignoreEmpty, "ignore empty lines")
	flags.BoolVarP(&shell, "shlex", "s", ignoreEmpty, "spilt qoute like shells")
	flags.StringVarP(&delimiter, "delimiter", "d", delimiter, "delimiter for field separation")
	flags.StringVarP(&format, "format", "f", format, "field printing format")
	flags.VarP(&limit, "limit", "n", "number of field to separate")

	if slices.Contains(os.Args, "_carapace") {
		carapace.Gen(Command)
	} else {
		handler := slog.New(log.New(os.Stderr))
		slog.SetDefault(handler)
	}
}

// MinimumNArgs returns an error if there is not at least N args.
func MinimumNArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if !cmd.Flags().Changed("format") && len(args) < n {
			return fmt.Errorf("requires at least %d arg(s), only received %d", n, len(args))
		}
		return nil
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
ps aux | field -n 11 11

# Shows the PID and the command from ps command with <PID>:<COMMAND> format
ps aux | field -n 11 -f "{2}:{11}"

# Extract multiple fields (user and PID) and print them
ps aux | field 1 2

# Extract a directory and get all the deleted files
rm -vrf bad-directory | field -s -- -1
`,
	Args: MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var template *fasttemplate.Template
		var writter *bufio.Writer
		if cmd.Flags().Changed("format") {
			t, err := fasttemplate.NewTemplate(format, "{", "}")
			if err != nil {
				return err
			}
			template = t
			writter = bufio.NewWriter(os.Stdout)
		}

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
			if shell {
				s, err := shlex.Split(line)
				if err != nil {
					slog.Error("Failed to parse qouted field", "error", err)
					continue
				}
				fields = s.Strings()
			} else if cmd.Flags().Changed("delimiter") {
				fields = FieldN(line, delimiter, limit.Int())
			} else {
				fields = FieldNPred(line, unicode.IsSpace, limit.Int())
			}

			selected := make([]string, 0, 10)
			for r := range slices.Values(ranges) {
				selected = append(selected, r.Select(fields)...)
			}

			if template == nil {
				if ignoreEmpty && len(selected) == 0 {
					continue
				}
				fmt.Println(strings.Join(selected, " "))
				continue
			}

			selector := func(w io.Writer, tag string) (int, error) {
				r, err := ParseRange(tag, false)
				if err != nil {
					return 0, err
				}
				return fmt.Fprint(w, strings.Join(r.Select(fields), " "))
			}

			if _, err := template.ExecuteFunc(writter, selector); err != nil {
				slog.Error("Failed execute format template", "error", err)
				continue
			}

			if _, err := writter.WriteString("\n"); err != nil {
				slog.Error("Failed execute format template", "error", err)
				continue
			}

			if err := writter.Flush(); err != nil {
				return err
			}
		}

		return scanner.Err()
	},
}
