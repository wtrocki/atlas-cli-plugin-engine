#  Atlas CLI plugin generator

> NOTE: This project is just an experiment. Sill WIP.

Generate Atlas CLI plugins from OpenAPI file.

## How it works

1. API updates engine was extended to detect new changes thru oasdiff.
Diff can be triggered manually by running: 
2. Based on diff file we generating or updating atlas CLI plugin:
https://github.com/wtrocki/atlas-cli-plugin-engine/blob/main/openapi/command.diff
3. Plugin gets automerged and released. 
4. Plugin engine can discover the new extension and autoinstall it.



