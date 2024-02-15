# twigo CLI

twigo is a cli to send sms/whatapps in golang

## How to use it

### Send SMS

Suppose you want to send SMS to your friend:

```bash
#   twigo chat --type sms --to "+9XXXXXXXXXXX6" --from "+19519774919" --body "hey buddy !!" 
  go run ./cmd/twigo/main.go send sms --to "+919XXXXXXXX6" --from "+19519774919" --body "hey buddy..."

  twigo send sms --to "+919XXXXXXXX6" --from "+19519774919" --body "hey buddy !!"

# send join <dry-canal> from whatapps and you will get your virtual whatapps number:
  go run ./cmd/twigo/main.go send whatsup --to "+9197XXXXXXX76" --from "+14155238886" --body "hey buddy..."

  twigo send voice --to "+9197XXXXXX76" --from "+19519774919" --body "hey buddy !!"

  twigo send sms --to "+9197XXXXXXX76" --from "+19519774919" --body "hey buddy !!"


```

Understand the above commad:

- twigo: root command
- chat: sub-command
- type: flag-  <sms/whatapps/voice/etc>
- to: flag - destination mobile no.
- from: flag - source mobile no.
- body: flag - message

### Send whatapps msg

Suppose you want to send whatapps to your friend:

` twigo chat --type whatapps `

### Send voice note

Suppose you want to send voice to your friend:

` twigo chat --type voice `

### Send verification/2FA msg

Suppose you want to send verify to your friend:

` twigo chat --type voice `
