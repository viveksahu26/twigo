package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/viveksahu26/twigo/cmd/twigo/cli"
	twigoError "github.com/viveksahu26/twigo/cmd/twigo/errors"
)

func main() {
	ctx := context.Background()
	for i, arg := range os.Args {
		if (strings.HasPrefix(arg, "-") && len(arg) == 2) || (strings.HasPrefix(arg, "--") && len(arg) >= 4) {
			continue
		}
		if strings.HasPrefix(arg, "--") && len(arg) == 3 {
			// Handle --o, convert to -o
			newArg := fmt.Sprintf("-%c", arg[2])
			msg := fmt.Sprintf("%s", ctx, "the flag %s is deprecated and will be removed in a future release. Please use the flag %s.", arg, newArg)
			fmt.Fprintf(os.Stderr, "WARNING: %s\n", msg)
			os.Args[i] = newArg
		} else if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			// Handle -output, convert to --output
			newArg := fmt.Sprintf("-%s", arg)
			newArgType := "flag"
			if newArg == "--version" {
				newArg = "version"
				newArgType = "subcommand"
			}
			msg := fmt.Sprintf("%s", ctx, "the %s flag is deprecated and will be removed in a future release. "+
				"Please use the %s %s instead.",
				arg, newArg, newArgType,
			)
			fmt.Fprintf(os.Stderr, "WARNING: %s\n", msg)
			os.Args[i] = newArg
		}
	}
	if err := cli.New().Execute(); err != nil {
		var urlShortnerError *twigoError.TwigoError
		if errors.As(err, &urlShortnerError) {
			log.Printf("error during command execution: %v", err)
			os.Exit(urlShortnerError.ExitCode())
		}

		// we don't call os.Exit as Fatalf does both PrintF and os.Exit(1)
		log.Fatalf("error during command execution: %v", err)
	}
}
