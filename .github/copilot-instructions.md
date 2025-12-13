# Copilot Instructions (Agent-Focused)

This file is for automated coding agents (e.g. Copilot) performing changes in the repo. It deliberately avoids duplicating human-facing documentation in `README.md`, `CONTRIBUTING.md`, or `DESIGN.md` and focuses on operational constraints, invariants, and safe-change heuristics.

## 1. Purpose

Enable precise, low-risk edits that preserve project guarantees: thin wrapper design, 100% statement coverage, Conventional Commits, and predictable exit code semantics.

## 2. Core Invariants (Do Not Break)

- Thin wrapper: do not reimplement chezmoi logic; only orchestrate invocation.
- Exit codes: 0 (success); propagated non-zero from child; -1 execution failure (command not found / could not start); map non-nil error from `executor.ExecuteChezmoi` to process exit 1 after writing error to stderr.
- Coverage: MUST remain 100% statement coverage after any change.
- POSIX-only: no Windows support scaffolding.
- Destination is always `/`; do not introduce flags to override source/dest.

## 3. Safe Change Checklist

Before applying a patch:

1. Add or adjust tests FIRST (red → green cycle).
2. Keep diff surgical: modify only what the task requires.
3. Do not add external dependencies without explicit task instruction.
4. Maintain existing public function signatures unless a failing test proves necessity.
5. Preserve logging format (lines starting with `info:`) unless tasked to change it.

## 4. Action Boundaries

Allowed without extra approval:

- Adding table-driven tests.
- Refactoring internals for clarity with unchanged behavior + coverage.
- Minor doc clarifications (non-duplicative).

Defer unless explicitly requested:

- Introducing new CLI flags / breaking user-facing interface.
- Modifying release automation workflows.
- Adding new build tags or platform conditionals.

Forbidden:

- Generating Windows compatibility code.
- Downgrading coverage or removing tests without replacement.

## 5. Implementation Preconditions (Spec vs Implemented)

Some features are referenced in design but not yet implemented. Mark new code with comments if it fulfills a SPEC item.

| Feature                                    | Status | Preconditions to Implement                                 |
| ------------------------------------------ | ------ | ---------------------------------------------------------- |
| Command classification (sudo vs non-sudo)  | SPEC   | Add classification table + exhaustive tests                |
| Ownership fix after sudo commands          | SPEC   | Introduce FS abstraction + tests simulating modified files |
| Env var mapping (CHEZROOT*\* → CHEZMOI*\*) | SPEC   | Implement builder + precedence tests (global vs profile)   |
| Profile management commands                | SPEC   | Define path resolver + collision & validation tests        |

## 6. Test Requirements

Every new branch or conditional must be covered. For execution logic:

- Success path (exit 0).
- Non-zero exit from `chezmoi` (propagation).
- Execution failure (exit -1) via nonexistent command test.
- Error mapping in `run()` (stderr message + exit 1).
  Use `bytes.Buffer` for stdout/stderr capture; prefer table-driven tests for classification once implemented.

## 7. Failure Handling & Retry

If a change causes tests or lint to fail, attempt up to 3 focused fixes. On persistent failure, stop and summarize root cause, failing output, and next options—do not leave broken state unexamined.

## 8. Logging & Messages

Do not change existing log prefixes (`info:`) or stderr formatting unless the task explicitly requires it. New messages should be deterministic (no timestamps) and testable.

## 9. Roadmap Hooks

Unimplemented planned features should be clearly marked with `// SPEC:` comments rather than partial stubs. Do not simulate unsupported behavior.

## 10. Non-Goals

- Windows / UAC model integration.
- Arbitrary destination directory overrides.
- Silent auto-fixes of user configuration beyond documented behavior.
- Adding heavy abstraction layers without documented need.

## 11. Conventional Commit Hints

Choose types based on change intent:

- `feat`: Implement a SPEC feature (only when fully realized + tested).
- `fix`: Correct logic, test reveals regression or missing edge case.
- `refactor`: Internal structural improvement, no behavior change.
- `test`: Add or improve tests only.
- `docs`: Documentation-only change.
- `ci` / `build`: Pipeline or build system modification.

Subject line: imperative, lowercase, no trailing period. Keep scope when clarifying area (e.g., `executor`, `runner`, `docs`).

## 12. Escalation Rules

If a requested change conflicts with an invariant (e.g., reduces coverage, alters exit semantics), halt and output a concise explanation + safer alternatives instead of proceeding.

## 13. Assumptions Allowed

- `chezmoi` and `sudo` binaries exist on PATH during tests.
- POSIX environment with bash available (tests may use `/bin/sh` fallback if needed).

## 14. Assumptions Forbidden

- Assuming profile directories already populated.
- Assuming environment variables beyond standard OS + test setup.

## 15. Patch Style

- Use minimal context in diffs; avoid unrelated whitespace churn.
- Keep imports sorted naturally; do not introduce blank identifier imports except tooling file.

## 16. ADRs (Optional Future)

If a change alters a core invariant, first add an ADR in `docs/adr/` (format: `NNNN-title.md`) summarizing Context / Decision / Consequences. Only implement after ADR merged.

Refer to:

- User usage & profiles: `README.md`
- Contributor workflow & coverage commands: `CONTRIBUTING.md`
- Architectural rationale & SPEC features: `DESIGN.md` (add when created)

End of agent-focused instructions.
