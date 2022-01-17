# Getting started

After cloning repo, run:

```
find . -type f \( -name "*.go" \) -exec sed -i '' 's/charmixer\/golang-cli-template/your-repo-name/g' {} +;
go mod init <new repo eg. github.com/username/your-repo-name>
go mod tidy
go run main.go serve
```

To build with supported ldflags use:
```
go build -ldflags="-s -w -X main.version=1.0.0 -X main.commit=qwerty -X main.date=20210101 -X main.tag=v1.0.0 -X main.name=golang-template-api -X main.environment=production" .
```

# What the template gives

## Configuration setup

Reading configuration in following order

1. Read from yaml file by setting `CFG_PATH=/path/to/conf.yaml`
2. Read from environment `PATH_TO_STRUCT_FIELD=override-value`
3. Read from flags `go run main.go serve -f override-value`
4. If none of above use `default:"value"` tag from configuration struct

# Done / TODOs

- [x] Struct-based application config
  - [x] Config from yaml
  - [x] Override config with environment
  - [x] Override environment with flags
  - [x] Default configuration values
- [x] CI pipeline with fmt, vet, staticchecks, build & test
- [x] CD pipeline using goreleaser - triggered by tag push
- [x] Publish to Github Packages (ghcr.io) with GoReleaser - disable by adding skip flag to the docker section in: `.goreleaser.yml` (see GoReleaser docs)
- [x] Setup changelog generator (https://github.com/charmixer/auto-changelog-action) - currently creates pull requests
  - [ ] Should somehow create the changelog before release tag is created, so it gets baked in
- [ ] README.md update with guides
