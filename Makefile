# Makefile (GNU make). Works on macOS, Linux, and Windows (GNU make 4.x).
.DEFAULT_GOAL := help

# -------- OS detection & suffix --------
ifeq ($(OS),Windows_NT)
  EXEEXT := .exe
  DETECTED_OS := windows
else
  UNAME_S := $(shell uname -s)
  ifeq ($(UNAME_S),Darwin)
    DETECTED_OS := mac
  else
    DETECTED_OS := linux
  endif
  EXEEXT :=
endif

# -------- Project config --------
BIN_NAME ?= pulse
OUT_DIR  ?= ./bin
BIN      := $(OUT_DIR)/$(BIN_NAME)$(EXEEXT)
PKG_MAIN ?= ./cmd

# -------- Tools --------
GOLANGCI_LINT := golangci-lint

# -------- Platform helpers --------
ifeq ($(DETECTED_OS),windows)
  MKDIR_P = powershell.exe -NoLogo -NoProfile -Command "New-Item -ItemType Directory -Force -Path '$(1)' | Out-Null"
  RM_RF = powershell.exe -NoLogo -NoProfile -Command "if (Test-Path -LiteralPath '$(1)') { Remove-Item -LiteralPath '$(1)' -Recurse -Force }"
  TOUCH = powershell.exe -NoLogo -NoProfile -Command "New-Item -ItemType File -Path '$(1)' -Force | Out-Null"
  FMT_CHECK = powershell.exe -NoLogo -NoProfile -Command "Set-StrictMode -Version 2; $$ErrorActionPreference = 'Stop'; $$files = gofmt -l .; if ($$files) { Write-Host \"The following files need 'gofmt -s -w':\"; $$files | ForEach-Object { Write-Host $$_ }; exit 1 }"
  ENSURE_GO_MOD_CLEAN = powershell.exe -NoLogo -NoProfile -Command "Set-StrictMode -Version 2; $$ErrorActionPreference = 'Stop'; if (Test-Path '.git') { git diff --quiet -- go.mod go.sum; if ($$LASTEXITCODE -ne 0) { git diff -- go.mod go.sum; Write-Host \"go.mod/go.sum changed; commit the result of 'go mod tidy'\"; exit 1 } }"
  CHECK_TOOL = powershell.exe -NoLogo -NoProfile -Command "Set-StrictMode -Version 2; if (-not (Get-Command $(1) -ErrorAction SilentlyContinue)) { Write-Error 'Missing required tool: $(1). Install it and ensure it is on PATH.'; exit 1 }"
else
  MKDIR_P = mkdir -p $(1)
  RM_RF = rm -rf $(1)
  TOUCH = touch $(1)
  FMT_CHECK = sh -c 'set -eu; files="$$(gofmt -l .)"; if [ -n "$$files" ]; then printf "%s\n" "The following files need '\''gofmt -s -w'\'':"; printf "%s\n" "$$files"; exit 1; fi'
  ENSURE_GO_MOD_CLEAN = sh -c 'set -eu; if [ -d .git ]; then if ! git diff --quiet -- go.mod go.sum; then git diff -- go.mod go.sum; printf "%s\n" "go.mod/go.sum changed; commit the result of '\''go mod tidy'\''"; exit 1; fi; fi'
  CHECK_TOOL = sh -c 'set -eu; if ! command -v $(1) >/dev/null 2>&1; then printf "%s\n" "Missing required tool: $(1). Install it and ensure it is on PATH."; exit 1; fi'
endif

# -------- Phony targets --------
.PHONY: help build test run lint lint-fast fmt fmt-check vet tidy clean ci tools lint-docker

help:
	@echo Detected OS: $(DETECTED_OS)
	@echo make setup		- download go modules
	@echo make build        - build $(BIN)
	@echo make test         - run unit tests
	@echo make run          - run $(PKG_MAIN)
	@echo make lint         - run golangci-lint (requires manual install)
	@echo make lint-fast    - fast checks only (requires manual install)
	@echo make fmt          - gofmt (write changes)
	@echo make fmt-check    - verify formatting
	@echo make vet          - go vet
	@echo make tidy         - go mod tidy (assert clean)
	@echo make clean        - remove build artifacts
	@echo make tools        - verify required dev tools
	@echo make ci	        - run lint and test commands

# -------- Core --------
setup:
	go mod download

build:
	@$(call MKDIR_P,$(OUT_DIR))
	go build -o "$(BIN)" $(PKG_MAIN)

test:
	go test -race -cover ./...

run:
	go run $(PKG_MAIN)

# -------- Lint / format / vet --------
lint:
	@$(call CHECK_TOOL,$(GOLANGCI_LINT))
	$(GOLANGCI_LINT) run ./...

lint-fast:
	@$(call CHECK_TOOL,$(GOLANGCI_LINT))
	$(GOLANGCI_LINT) run --fast ./...

fmt:
	gofmt -s -w .

fmt-check:
	@$(FMT_CHECK)

vet:
	go vet ./...

tidy:
	go mod tidy
	@# Fail CI if go.mod/go.sum changed
	@$(ENSURE_GO_MOD_CLEAN)

# -------- Tool bootstrap --------
tools:
	@$(call CHECK_TOOL,$(GOLANGCI_LINT))
	@$(GOLANGCI_LINT) version

# -------- Cleanup --------
clean:
	@$(call RM_RF,$(OUT_DIR))

# Convenience
ci: setup lint test build
