# Qy API Change Toolkit

Detect changes before they break your workflow.

`qy-change` compares OpenAPI, JSON Schema, Sitemap, and RSS/JSON Feed files and reports changes as JSON or Markdown.

## Install

```bash
go install github.com/BShaT/qy-api-change-toolkit/cmd/qy-change@latest
```

## Usage

```bash
qy-change openapi old.yaml new.yaml --format markdown
qy-change schema old.json new.json
qy-change sitemap old.xml new.xml
qy-change feed old.xml new.xml
```

The first release is intentionally local-only. It does not upload files or require a product account.

## License

See `LICENSE`.
