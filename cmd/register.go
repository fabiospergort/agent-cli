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
nor the server can sign transactions alone.

The wallet starts in "unclaimed" status until a human owner claims it.

Naming your wallet (--name is required):
  Use your agent's name for a general-purpose wallet, or describe its purpose
  for a dedicated one. The name helps the human owner recognize it.

  General:   "<AgentName>'s Wallet"    e.g. "Orion's Wallet"
  Purpose:   "<Purpose> Wallet"        e.g. "Research Budget", "API Allowance"`,
	Example: `  # General-purpose wallet (use your agent's name)
  botwallet register --name "Orion's Wallet" --owner human@example.com

  # Purpose-specific wallet
  botwallet register --name "Research Budget" --owner sarah@example.com

  # Minimal (owner can be added later)
  botwallet register --name "My Wallet"`,
	Run: runWalletCreate,
}

func init() {
	registerCmd.Flags().StringVarP(&walletCreateName, "name", "n", "", "Name for your wallet (required)")
	registerCmd.Flags().StringVarP(&walletCreateAgentModel, "model", "m", "", "Agent model (e.g., 'gpt-4', 'claude-3')")
	registerCmd.Flags().StringVar(&walletCreateOwner, "owner", "", "Owner's email (wallet appears in their portal)")
}
