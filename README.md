# Ocelloids Data Steward CLI

The Ocelloids Data Steward CLI fetches aggregated data in JSON format from the Ocelloids data steward agent and outputs it to `stdout`, allowing you to easily use it with other command-line tools like `jq`.

```
Ocelloids Data Steward Agent command-line interface

Usage:
  steward [command]

Available Commands:
  fetch       Streams either assets or chains data to stdout
  help        Help about any command

Flags:
  -k, --api-key string    Ocelloids API key
      --config string     config file (default is $HOME/.stw.yaml)
  -h, --help              help for steward
  -u, --http-url string   HTTP API base URL (default "https://dev-api.ocelloids.net")

Use "steward [command] --help" for more information about a command.
```

## Usage

### 1. Download Assets Data

Fetch assets data and save it to a JSON Lines (JSONL) file:

```bash
steward fetch assets > assets.jsonl
```

To convert the JSONL file into a JSON array, use `jq`:

```bash
jq --slurp . < assets.jsonl > assets.json
```

### 2. Download Chains Data

Fetch chains data and save it to a JSONL file:

```bash
steward fetch chains > chains.jsonl
```

---

Have fun!