# Roadmap: Домашнее задание №3: Частотный анализ

## Overview

Путь к готовому решению короткий и линейный: сначала зафиксировать точную семантику токенизации и подсчета, затем добавить детерминированное ранжирование top 10, после чего закрыть риски через явные edge-case тесты.

## Phases

**Phase Numbering:**

- Integer phases (1, 2, 3): Planned milestone work
- Decimal phases are reserved for urgent insertions if scope changes later

- [ ] **Phase 1: Tokenization & Counting** - Реализовать корректное выделение слов и подсчет частот по базовым правилам.
- [ ] **Phase 2: Deterministic Top-10 Ranking** - Добавить стабильную сортировку и ограничение результата до 10 слов.
- [ ] **Phase 3: Verification & Edge Coverage** - Зафиксировать спецификацию расширенными тестами и финальной проверкой.

## Phase Details

### Phase 1: Tokenization & Counting

**Goal**: As a homework author, I want to implement Top10 correctly and lock behavior with tests, so that tests pass reliably.
**Mode:** mvp
**Depends on**: Nothing (first phase)
**Requirements**: [PARS-01, PARS-02, PARS-03, RANK-01, QUAL-01]
**Success Criteria** (what must be TRUE):

  1. Вызывающий код получает пустой результат для пустой строки без паники.
  2. Слова, разделенные пробелами, табами и переводами строк, выделяются одинаково корректно.
  3. Регистр и пунктуация сохраняются как часть точного токена при подсчете частот.

**Plans**: 2 plans

Plans:
**Wave 1**

- [x] 01-01-PLAN.md — Зафиксировать blank-input и whitespace tokenization slice через `Top10`

**Wave 2** *(blocked on Wave 1 completion)*

- [ ] 01-02-PLAN.md — Зафиксировать exact-token counting slice по README semantics

### Phase 2: Deterministic Top-10 Ranking

**Goal**: `Top10` возвращает детерминированный top 10 по частоте и слову.
**Mode:** mvp
**Depends on**: Phase 1
**Requirements**: [RANK-02, RANK-03, RANK-04, RANK-05]
**Success Criteria** (what must be TRUE):

  1. Результат всегда отсортирован по убыванию частоты.
  2. При равной частоте слова идут в лексикографическом порядке.
  3. Функция возвращает не более 10 элементов даже при большом числе уникальных слов.

**Plans**: 2 plans

Plans:

- [ ] 02-01: Добавить сортируемое представление частот и компаратор для tie-break
- [ ] 02-02: Реализовать безопасное ограничение результата до 10 слов и проверить tie-case

### Phase 3: Verification & Edge Coverage

**Goal**: Полностью подтвердить контракт README набором целевых тестов и итоговой верификацией.
**Mode:** mvp
**Depends on**: Phase 2
**Requirements**: [QUAL-02]
**Success Criteria** (what must be TRUE):

  1. В тестах явно покрыты happy-path, whitespace edge cases, пунктуация и tie-break.
  2. Полная команда `go test -v -count=1 -race -timeout=1m ./...` проходит в каталоге задания.
  3. Проверяющему достаточно прочитать тесты, чтобы увидеть отражение ключевых правил README.

**Plans**: 1 plan

Plans:

- [ ] 03-01: Расширить тесты под требования README и прогнать итоговую верификацию

## Progress

**Execution Order:**
Phases execute in numeric order: 1 → 2 → 3

| Phase | Plans Complete | Status | Completed |
|-------|----------------|--------|-----------|
| 1. Tokenization & Counting | 1/2 | In Progress|  |
| 2. Deterministic Top-10 Ranking | 0/2 | Not started | - |
| 3. Verification & Edge Coverage | 0/1 | Not started | - |
