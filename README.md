# Ocelloids Data Steward CLI

The Ocelloids Data Steward CLI fetches aggregated data in JSON format from the Ocelloids data steward agent and outputs it to `stdout`, allowing you to easily use it with other command-line tools like `jq`.

```shell
Ocelloids Data Steward Agent command line interface

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

---

NOTES

to clean up :)

```
jq --slurp . < assets.jsonl > assets.json
```

go run main.go download --http-url http://127.0.0.1:3000 | jq --slurp . > assets.json