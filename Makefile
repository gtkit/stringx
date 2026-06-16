.PHONY: tool check



LINT_TARGETS ?= ./...
GOLANGCI_LINT_MIN_MAJOR ?= 2

tool: ## Lint Go code with the installed golangci-lint
	@ echo "▶️ golangci-lint run"
	@ version=$$(golangci-lint --version | awk '{print $$4}'); \
	case "$$version" in \
		$(GOLANGCI_LINT_MIN_MAJOR).*) ;; \
		*) echo "golangci-lint $$version is unsupported; install v$(GOLANGCI_LINT_MIN_MAJOR).x for .golangci.yml" >&2; exit 1 ;; \
	esac
	golangci-lint config verify -c .golangci.yml
	golangci-lint run $(LINT_TARGETS)
	gofumpt -l -w .
	@ echo "✅ golangci-lint run"

## govulncheck 检查漏洞 go install golang.org/x/vuln/cmd/govulncheck@latest
check:
	govulncheck ./...
	gosec -exclude=G103,G104,G115,G304,G404 ./...


tag:
	@current=$$(grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+' version.go | head -n1 | tr -d 'v'); \
	if [ -z "$$current" ]; then echo "version not found in version.go"; exit 1; fi; \
	maj=$$(echo $$current | cut -d. -f1); \
	min=$$(echo $$current | cut -d. -f2); \
	patch=$$(echo $$current | cut -d. -f3); \
	newpatch=$$(expr $$patch + 1); \
	new="v$$maj.$$min.$$newpatch"; \
	printf "Bump: v%s -> %s\n" "$$current" "$$new"; \
	sed -E -i.bak 's/(const Version = ")([^"]+)(")/\1'"$$new"'\3/' version.go; \
	git add version.go; \
	git commit -m "chore(release): $$new"; \
	printf "Release: %s\n" "$$new"; \
	git push gtkit HEAD; \
	git tag -a "$$new" -m "release $$new"; \
	printf "Tag: %s\n" "$$new"; \
	git push gtkit "$$new"; \
	printf "Done\n"
	rm -f version.go.bak

gittag:
	git tag --sort=-version:refname | grep -E '^v(0|1)\.' | head -1