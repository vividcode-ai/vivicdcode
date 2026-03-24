# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

VividCode is a Go-based terminal AI assistant for software development. It provides an interactive TUI built with Bubble Tea for interacting with various AI models (OpenAI, Anthropic Claude, Google Gemini, AWS Bedrock, Azure, Groq, OpenRouter). Features include session management, file tracking, MCP tool integration, and LSP support for code intelligence.

## Common Commands

```bash
# Build the application
go build -o vividcode

# Run the application
./vividcode

# Run tests
go test ./...

# Run a single test
go test -run TestName ./path/to/package

# Run with debug logging
./vividcode -d
```

## Architecture

The codebase follows a modular architecture with clear separation of concerns:

- **cmd/**: Cobra CLI commands and application entry point
- **internal/app/**: Core application services (sessions, messages, permissions, agent)
- **internal/config/**: Configuration loading from JSON files and environment variables
- **internal/db/**: SQLite database with goose migrations; provides sessions, messages, and files tables
- **internal/llm/**: LLM integration layer
  - **agent/**: Agent service that orchestrates AI interactions with tools
  - **models/**: Model definitions and capabilities for each provider
  - **provider/**: API clients for each LLM provider
  - **prompt/**: Prompt templates for different agent tasks
  - **tools/**: Tool definitions the AI can use (bash, glob, grep, edit, view, write, etc.)
- **internal/lsp/**: Language Server Protocol client implementation
- **internal/tui/**: Bubble Tea TUI components and layouts
- **internal/logging/**: Structured logging with pub/sub event system

## Key Patterns

- **Services**: Core business logic is organized into services (Sessions, Messages, Permissions, CoderAgent) that communicate via pub/sub
- **Database**: Uses sqlc-generated Go code from SQL files in internal/db/sql/
- **Tools**: AI tools are defined as Go interfaces in internal/llm/tools/ and registered with the agent
- **Config**: Configuration supports cascading from .vividcode.json files and environment variables

## Database Migrations

Migrations are stored in internal/db/migrations/ and use goose. Run migrations via db.Connect() which auto-runs pending migrations.

## Configuration

Config is loaded from (in order of priority):
1. `./.vividcode.json` (local)
2. `$XDG_CONFIG_HOME/vividcode/.vividcode.json`
3. `$HOME/.vividcode.json`

Environment variables like `ANTHROPIC_API_KEY`, `OPENAI_API_KEY`, `GEMINI_API_KEY`, etc. can also be used.
