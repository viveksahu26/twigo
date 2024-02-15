package options

import "github.com/spf13/cobra"

type SendOptions struct {
	Type string
	To   string
	From string
	Body string
}

type SendEmailOptions struct {
	To               string
	From             string
	SenderUser       string
	RecieverUser     string
	Subject          string
	PlainTextContent string
	HtmlContent      string
}

// var _ Interface = (*ShortOptions)(nil)

func (o *SendOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.Type, "type", "", "type of chatting")
	cmd.Flags().StringVar(&o.To, "to", "", "mobile number of reciever")
	cmd.Flags().StringVar(&o.From, "from", "", "mobile number of sender")
	cmd.Flags().StringVar(&o.Body, "body", "", "message")
}

func (o *SendEmailOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.From, "from", "", "senders email address")
	cmd.Flags().StringVar(&o.To, "to", "", "recievers email address")
	cmd.Flags().StringVar(&o.Subject, "subject", "", "subject of an email")
	cmd.Flags().StringVar(&o.SenderUser, "su", "", "sender user name")
	cmd.Flags().StringVar(&o.RecieverUser, "ru", "", "reciever user name")
	cmd.Flags().StringVar(&o.HtmlContent, "hc", "", "html content")
	cmd.Flags().StringVar(&o.PlainTextContent, "ptc", "", "plain text content")
}
