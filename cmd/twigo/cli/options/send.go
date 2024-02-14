package options

import "github.com/spf13/cobra"

type SendOptions struct {
	Type string
	To   string
	From string
	Body string
}

// var _ Interface = (*ShortOptions)(nil)

func (o *SendOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.Type, "type", "", "type of chatting")
	cmd.Flags().StringVar(&o.To, "to", "", "source mobile number")
	cmd.Flags().StringVar(&o.From, "from", "", "destination mobile number")
	cmd.Flags().StringVar(&o.Body, "body", "", "message")
}
