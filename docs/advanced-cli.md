# Vader CLI Commands

Below is a list of current and planned commands for the Vader CLI:

| Command                               | Description                                                  | Status |
| ------------------------------------- | ------------------------------------------------------------ | ------ |
| `vader run <script-id>`               | Executes a script with a preview of commands                 | ✅     |
| `vader run <script-id> --no-preview`  | Runs a script directly without showing the preview           | ⏳     |
| `vader run <script-id> --export`      | Generates and exports a report/log of the execution          | ⏳     |
| `vader run <script-id> --dry-run`     | Simulates script execution without running any command       | ⏳     |
| `vader inspect <script-id>`           | Displays a detailed breakdown of the script’s commands       | ⏳     |
| `vader history`                       | Lists recently executed scripts                              | ⏳     |
| `vader bookmark <script-id>`          | Adds a script to bookmarks for quick access                  | ⏳     |
| `vader bookmarks`                     | Lists all bookmarked scripts                                 | ⏳     |
| `vader search <keyword>`              | Searches for scripts by name, tags, or metadata              | ⏳     |
| `vader add org <alias> <backend-url>` | Adds a new org alias to point to a self-hosted Vader backend | ⏳     |
| `vader run <org-alias> <script-id>`   | Executes a script from a specific org                        | ⏳     |
| `vader set-default-org <alias>`       | Sets a default org to avoid specifying it every time         | ⏳     |
| `vader config set-backend <url>`      | Manually sets the backend URL                                | ⏳     |
| `vader config show`                   | Displays the current CLI configuration                       | ⏳     |
