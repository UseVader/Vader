# Vader

**Vader** is the command-line tool that allows developers to execute diagnostic scripts fetched from a Vader instance (self-hosted or public). It’s built for speed, trust, and simplicity, enabling seamless debugging and validation of developer environments.

While the VaderCore Platform is where scripts are authored and managed, the CLI is what developers actually use to run those scripts and generate structured outputs.

Video demo here: [Vader Demo](https://www.loom.com/share/bbe81b692a2640a8b09f5991efbb1460?sid=77df57e0-ed35-452d-9027-4d0d2bae70cb)

---

## Features

- Run diagnostic scripts with a single command.
- Preview the exact commands before execution.
- Export structured logs or generate shareable links.
- Connect to multiple Vader instances (via alias).
- Security-first: Trust is built with preview + strict execution patterns.

---

## Installation

### Requirements

- Go version 1.20+

### Steps

```bash
git clone https://github.com/UseVader/vader.git
cd vader
go build -o vader .
```

Make the binary globally accessible:

```bash
mv vader /usr/local/bin/
```

---

## Basic CLI Usage

- `vader run <script-id>` – Run a script with preview.
- `vader run <script-id> --no-preview` – Run a script directly, skipping the preview.
- `vader run <script-id> --export` – Generate a shareable report or export output.

For advanced commands and features, see [Advanced CLI Usage](./docs/advanced-cli.md).

---

## Contributing

Contributions are welcome!

We’d love your help adding new features, refining the CLI experience, improving trust and security mechanisms, or contributing more useful flags.

If you're interested in contributing, please check out the issues tab and open a PR.

---

## License

Coming Soon.
