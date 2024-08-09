# Ocelloids Data Steward CLI

The Ocelloids Data Steward CLI fetches aggregated data in JSON format from the Ocelloids data steward agent and outputs it to `stdout`, allowing you to easily use it with other command-line tools like `jq`.

```
Ocelloids Data Steward Agent command-line interface

Usage:
  steward [command]

Available Commands:
  fetch       Prints assets or chains data to stdout
  help        Help about any command

Flags:
  -k, --api-key string    Ocelloids API key
      --config string     config file (default is $HOME/.stw.yaml)
  -h, --help              help for steward
  -u, --http-url string   HTTP API base URL (default "https://dev-api.ocelloids.net")

Use "steward [command] --help" for more information about a command.
```

## Install

To install the Ocelloids Data Steward CLI, download the appropriate binary for your platform from the [GitHub Releases page](https://github.com/sodazone/go-steward/releases).

<details>
<summary>Steps to Install</summary>

1. **Visit the Releases Page**: Go to the [Ocelloids Data Steward CLI Releases](https://github.com/sodazone/go-steward/releases).

2. **Choose Your Platform**: Select the binary that matches your operating system (e.g., Linux, macOS, Windows) and architecture (e.g., x86, x64, ARM).

3. **Download the Binary**: Click on the appropriate binary to start the download.

4. **Make the Binary Executable**: Once downloaded, ensure the binary is executable (only necessary on Unix-based systems). Use the following command:

   ```bash
   chmod +x steward
   ```
5. **Move the Binary to a Directory in Your PATH**: Move the binary to a directory included in your system's PATH environment variable for easy access:
   ```bash
   mv steward /usr/local/bin/
   ```
6. **Verify Installation**: Confirm the installation by running the following command:
   ```bash
   steward help
   ```
</details>

## Usage

### Download Assets Data

Fetch assets data and save it to a JSON Lines (JSONL) file:

```bash
steward fetch assets > assets.jsonl
```

To convert the JSONL file into a JSON array, use `jq`:

```bash
jq --slurp . < assets.jsonl > assets.json
```

### Download Chains Data

Fetch chains data and save it to a JSONL file:

```bash
steward fetch chains > chains.jsonl
```

---

Have fun!