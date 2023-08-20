# action

GitHub Action to ...

## Inputs

| Name | Description | Required |
|------|-------------|----------|
| key  | description | true     |

## Outputs

| Name   | Description |
|--------|-------------|
| result | result      |

## Example

`.github/workflows/main.yml`

```yml
name: main
on:
  push:
    branches:
      - main
jobs:
  main:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3

      <...>

      - name: ...
        uses: ...
        with:
          key: ...
```
