---
phase: 01-tokenization-counting
plan: 02
subsystem: testing
tags: [go, exact-token, punctuation, case-sensitivity, tdd]
requires:
  - phase: 01-tokenization-counting
    provides: whitespace tokenization contract and stdlib Top10 pipeline
provides:
  - "README-based exact-token regression tests for punctuation, case, and standalone hyphen semantics"
  - "Explicit raw-token counting contract preserved in Top10"
affects: [top.go, top_test.go, Phase 02]
tech-stack:
  added: []
  patterns: ["README example as executable regression", "exact-token map keys without normalization"]
key-files:
  created: []
  modified: [top.go, top_test.go]
key-decisions:
  - "Зафиксировал README example дословно в тесте, чтобы future changes не уводили базовую версию в starred-task semantics."
  - "Сохранил подсчет по raw token string и ограничился поясняющим комментарием в top.go, потому что реализация уже соответствовала контракту."
patterns-established:
  - "README bullet points превращаются в отдельные subtests внутри TestTop10."
  - "Пунктуация, регистр и standalone '-' считаются разными точными токенами в base task."
requirements-completed: [PARS-03, RANK-01]
duration: 1 min
completed: 2026-05-17
---

# Phase 01 Plan 02: Tokenization counting Summary

**README-semantics exact-token regressions now lock punctuation, case, and standalone hyphen handling while `Top10` keeps counting raw tokens unchanged.**

## Performance

- **Duration:** 1 min
- **Started:** 2026-05-17T04:21:39Z
- **Completed:** 2026-05-17T04:22:44Z
- **Tasks:** 2
- **Files modified:** 2

## Accomplishments
- Добавлены regression subtests для README-примера с точным ожидаемым порядком результатов.
- Зафиксирована exact-token семантика для `Нога`, `нога`, `нога,` и standalone `-`.
- В `top.go` явно задокументировано, что базовая версия считает raw tokens без нормализации.

## Task Commits

Each task was committed atomically:

1. **Task 1: Add failing exact-token regression tests from README semantics** - `ec3d51a` (test)
2. **Task 2: Preserve raw-token counting and prove the whole homework package still passes** - `a979e0d` (refactor)

**Plan metadata:** pending

## Files Created/Modified
- `top_test.go` - Добавлены README-based exact-token regression tests для punctuation/case/hyphen semantics.
- `top.go` - Добавлен комментарий, фиксирующий raw-token counting contract без изменения поведения.

## Decisions Made
- README использован как канонический источник истины и перенесен в прямые regression assertions.
- Реализация `Top10` не менялась по существу, потому что уже соблюдала exact-token contract; добавлено только пояснение для читаемости.

## Deviations from Plan

None - plan executed exactly as written.

## Issues Encountered
- RED-фаза Task 1 не дала ожидаемого падения: после добавления новых тестов `Top10` сразу прошел их, потому что реализация из `01-01` уже соответствовала README exact-token semantics.
- Обязательная команда `go test -v -count=1 -race -timeout=1m ./...` не выполнилась в текущем окружении: сначала `-race` потребовал `CGO_ENABLED=1`, затем при `CGO_ENABLED=1` сборка уперлась в отсутствие `gcc` в `PATH`.
- В качестве доступного доказательства корректности успешно выполнен `go test -v -count=1 ./...`, включая новые regression tests.

## User Setup Required

None - no external service configuration required.

## Next Phase Readiness
- Phase 1 закрывает базовую exact-token семантику и готов передать работу в Phase 2 ranking tasks.
- Для полного соответствия репозиторной команде верификации в окружении по-прежнему нужен доступный `gcc` для `go test -race`.

## Self-Check: PASSED
- FOUND: `.planning/phases/01-tokenization-counting/01-02-SUMMARY.md`
- FOUND: `ec3d51a`
- FOUND: `a979e0d`

---
*Phase: 01-tokenization-counting*
*Completed: 2026-05-17*
