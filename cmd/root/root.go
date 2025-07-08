package root

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/henrywhitakercommify/cfimport/internal/dns"
	"github.com/spf13/cobra"
)

var (
	token          string
	to             string
	importTemplate = `import {
	from = "{{ .ID }}"
	to = {{ .To }}
}`
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "cfimport [zone]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := dns.New(token)

			records, err := client.Records(cmd.Context(), args[0])
			if err != nil {
				return err
			}

			tmpl, err := template.New("import").Parse(importTemplate)
			if err != nil {
				return err
			}

			out := strings.Builder{}

			for _, r := range records {
				data := map[string]any{
					"ID": r.ID,
					"To": to,
				}
				if err := tmpl.Execute(&out, data); err != nil {
					return fmt.Errorf("execute import template: %w", err)
				}
				out.WriteString("\n")
			}

			fmt.Println(out.String())

			return nil
		},
	}

	cmd.Flags().StringVar(&token, "token", "", "The api token to authenticate with cloudflare")
	cmd.Flags().StringVar(&to, "to", "", "The path to import the record to")

	return cmd
}
