package cmd

import (
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Create a new Botwallet (same as 'wallet create')",
	Long: `Create a new Botwallet for your AI agent.

This is an alias for 'wallet create' — both commands do exactly the same thing.

What happens:
  1. Generates a threshold key pair (your key share stays on your machine)
  2. Registers with Botwallet (server holds its own separate key share)
  3. Saves credentials to ~/.botwallet/

The wallet is secured with 2-of-2 threshold signing — neither the agent
nor the server can sign transactions alone. The wallet starts "unclaimed"
until a human owner claims it via a claim link.

Naming (--name is required):
  This is your wallet — just pick a name and proceed. No need to ask your
  human; it keeps the experience seamless. Use your own name, a creative name,
  or describe the wallet's purpose. Your human will see it when they claim.`,
	Example: `  botwallet register --name "Assistant's Wallet"
  botwallet register --name "Research Budget" --owner human@example.com
  botwallet register --name "x402 APIs Allowance"`,
	Run: runWalletCreate,
}

func init() {
	registerCmd.Flags().StringVarP(&walletCreateName, "name", "n", "", "Name for your wallet (required)")
	registerCmd.Flags().StringVarP(&walletCreateAgentModel, "model", "m", "", "Agent model (e.g., 'gpt-4', 'claude-3')")
	registerCmd.Flags().StringVar(&walletCreateOwner, "owner", "", "Owner's email (wallet appears in their portal)")
}
