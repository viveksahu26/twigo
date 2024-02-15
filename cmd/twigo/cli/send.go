package cli

import (
	"github.com/spf13/cobra"
	"github.com/viveksahu26/twigo/cmd/twigo/cli/options"
	"github.com/viveksahu26/twigo/cmd/twigo/cli/send"
)

func Send() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send",
		Short: "send message",
	}
	cmd.AddCommand(
		sms(),
		whatsup(),
		voice(),
		email(),
	)
	return cmd
}

// sub-command of send
func sms() *cobra.Command {
	o := &options.SendOptions{}

	cmd := &cobra.Command{
		Use:   "sms",
		Short: "Send sms directly to mobile phone",
		Example: `
		twigo send sms --to "+919XXXXXXXX6" --from "+19519774919" --body "hey buddy !!"
		`,
		// Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return send.SendMsg(cmd.Context(), o.To, o.From, o.Body, args)
		},
	}
	o.AddFlags(cmd)

	return cmd
}

// sub-command of send
func whatsup() *cobra.Command {
	o := &options.SendOptions{}

	cmd := &cobra.Command{
		Use:   "whatsup",
		Short: "Send whatsup messag",
		Example: `
		twigo send whatsup --to "+919XXXXXXXX6" --from "+19519774919" --body "hey buddy !!"
		`,
		// Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return send.SendWhatsUP(cmd.Context(), o.To, o.From, o.Body, args)
		},
	}
	o.AddFlags(cmd)

	return cmd
}

// sub-command of send
func voice() *cobra.Command {
	o := &options.SendOptions{}

	cmd := &cobra.Command{
		Use:   "voice",
		Short: "Send voice messag",
		Example: `
		twigo send voice --to "+91XXXXXXXXX6" --from "+19519774919" --body "hey buddy !!"
		`,
		// Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return send.SendVoiceMsg(cmd.Context(), o.To, o.From, o.Body, args)
		},
	}
	o.AddFlags(cmd)

	return cmd
}

// sub-command of send
func email() *cobra.Command {
	o := &options.SendEmailOptions{}

	cmd := &cobra.Command{
		Use:   "email",
		Short: "Send email messag",
		Example: `
		twigo send email --su "<sender user name>" --from "<sender-email-address>"    --ru "reciever user name" --to "reciever-email-address" --subject "Twilio 101 Workshop"  --ptc "and easy to do anywhere, even with Go"  --hc "<strong>and easy to do anywhere, even with Go</strong>"`,
		// Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return send.SendEmailMsg(cmd.Context(), o.SenderUser, o.From, o.RecieverUser, o.To, o.Subject, o.HtmlContent, o.PlainTextContent, args)
		},
	}
	o.AddFlags(cmd)

	return cmd
}
