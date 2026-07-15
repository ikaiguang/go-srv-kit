# Agent Instructions

## Working Rules

- Read the target files and adjacent implementations before changing code.
- Follow existing package patterns and keep changes narrowly scoped.
- Run the smallest relevant tests after changes, then expand verification when risk warrants it.
- Do not edit generated Proto or Wire output directly; change the source and run the repository generator.
- Preserve public API compatibility unless the user explicitly approves a breaking change.

## Skills

- Use `.agents/skills/my-project` for repository-specific Go, module, Proto, generation, test, and debugging context.
- Use `.agents/skills/code-audit-repair` only for explicit whole-repository or multi-module audits.
- Use matching global Superpowers skills for general design, planning, debugging, TDD, review, and verification workflows.
- Use `docs/superpowers/specs/` and `docs/superpowers/plans/` as the only design and implementation-plan document locations.
- User instructions and repository facts take precedence over generic skill defaults.

## Repository

- This is a multi-module Go repository. Locate the nearest `go.mod` before choosing commands or package paths.
- Run Go tests from the owning module directory; start with the affected package.
- Treat root and module Makefiles as the source of truth for Proto and other generation commands.

## Context7

- Use Context7 for current documentation when the user asks about a library, framework, SDK, API, CLI tool, or cloud service.
- Start with `resolve-library-id`, then call `query-docs` with the selected `/org/project` ID and the user's full concept-specific question.
- Use separate documentation queries for distinct concepts.
- Do not use Context7 for business-logic debugging, refactoring, code review, scripts written from scratch, or general programming concepts.
