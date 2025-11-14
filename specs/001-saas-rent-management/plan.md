# Implementation Plan: saas-rent-management

**Branch**: `001-saas-rent-management` | **Date**: 2025-11-14 | **Spec**: specs/1-saas-rent-management/spec.md
**Input**: Feature specification from `specs/1-saas-rent-management/spec.md`

## Summary
构建一个移动优先的 SAAS 出租屋管理 Web 平台。后端使用 Golang（Gin + GORM）并以 SQLite 作为开发/初期存储；前端使用 Vue 3 + TypeScript，采用轻量级移动优先后台管理 UI（自定义或轻量组件库），遵循 TDD 与性能限制（关键路径首交互 ≤ 100ms）。

## Technical Context

**Language/Version**: Go 1.20+ (建议)  
**Primary Dependencies**: Gin (web framework), GORM (ORM), testify (testing), go-sqlite3 (driver)  
**Storage**: SQLite (single-file DB for initial deployment; design for future migration)  
**Testing**: Go's testing package + testify for asserts; integration tests use in-memory or temp SQLite files; Frontend uses Vitest + @testing-library/vue  
**Target Platform**: Linux/macOS server for backend; modern mobile browsers for frontend  
**Project Type**: Web application (backend API + frontend SPA)  
**Performance Goals**: Backend: p95 < 200ms for typical API calls; Frontend key-path first interaction median ≤ 100ms  
**Constraints**: Minimal third-party JS libs; TDD enforced; test coverage >= 90%  
**Scale/Scope**: MVP supports thousands of properties/rooms in single-DB mode (SQLite limits to be assessed if scale increases)

## Constitution Check

GATE: Verify plan vs constitution (must pass or list violations):
- Backend language: Go — OK
- Frontend: Vue3+TS — OK
- TDD & 90% coverage — OK (must be enforced in CI)
- Performance constraint (<=100ms) — OK (requires performance tests in CI)
- Avoid non-essential third-party libs — OK (must justify any additions)

No violations detected. Plan may proceed to Phase 0.

## Project Structure

Selected layout: Web application split into `backend/` and `frontend/`.

backend/
├── cmd/server/main.go        # server entry
├── internal/
│   ├── api/                  # HTTP handlers (gin)
│   ├── service/              # business logic
│   ├── models/               # GORM models
│   └── db/                   # migrations & sqlite helper
└── tests/                    # integration tests

frontend/
├── src/
│   ├── components/
│   ├── pages/
│   └── services/             # API clients (fetch) and utilities
└── tests/                    # vitest + e2e hooks (if needed)

**Structure Decision**: Separate backend/frontend provides clear separation of concerns, allows independent testing and CI gates for each part, and maps to the spec's API-driven workflow.

## Phase 0: Research & Unknowns

The spec raised three open questions. We'll resolve them now with decisions and rationale.

- Q1: 多租户隔离策略（每租户独立 DB vs 单实例多租户）
  - Decision: 初始实现采用单实例多租户（单 SQLite 文件）并通过 `tenant_id` 在应用层隔离数据。原因：SQLite 管理多个文件会增加复杂度，单 DB 更便于部署和测试，且满足 MVP。未来若需更强隔离或并发伸缩，计划迁移到 PostgreSQL 并支持 per-tenant DB。
  - Rationale: 需求中并未明确高并发或严格隔离需求；选择单 DB 简化测试 (TDD) 与 CI，并降低早期运维复杂度。
  - Alternatives considered: per-tenant DB（更强隔离，但增加迁移和运维成本）；schema-per-tenant（需 DB 支持与复杂迁移策略）。

- Q2: 收入分成精度与四舍五入规则
  - Decision: 使用以分为单位（两位小数，货币以整数分存储）的会计原则，结算与分配在分（cent）级别进行，最终显示为保留两位小数。分配时采用 "银行家舍入"（bankers rounding）或明确四舍六入五成双策略以减少累计误差。记录分配的原始分配值与四舍五入调整。
  - Rationale: 使用整数表示货币并采用银行家舍入可降低累计偏差，并符合多数财务实践。
  - Alternatives: 使用高精度十进制库（如 shopspring/decimal）以避免浮点误差；如果后续需支持多币种或更高精度可引入 decimal 类型.

- Q3: 内置支付/结算通道需求
  - Decision: 初始迭代不直接集成第三方支付网关；系统仅生成分成报表与交易记录，供手动或外部系统结算。支付集成列为扩展任务。
  - Rationale: 支付集成属于风险较高、合规和对接复杂性大的功能，建议在稳定的报表与结算逻辑后再引入。
  - Alternatives: 集成 Stripe/Adyen 等（需额外审计和合规流程）。

**Phase 0 output**: The above decisions will be written to `specs/1-saas-rent-management/research.md` and used to guide Phase 1.

## Next Steps (Phase 1 preview)
- Generate `data-model.md` from spec entities and decisions above.
- Produce OpenAPI contracts for core flows (auth omitted here if external), focusing on properties, rooms, contracts, tenants, allocations.
- Create `quickstart.md` with instructions for running backend+frontend locally using SQLite.

## Artifacts to be created
- `specs/1-saas-rent-management/research.md`
- `specs/1-saas-rent-management/data-model.md`
- `specs/1-saas-rent-management/contracts/*`
- `specs/1-saas-rent-management/quickstart.md`
