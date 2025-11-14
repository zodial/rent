# Research: saas-rent-management (Phase 0)

**Created**: 2025-11-14

## Purpose
Resolve open questions from the spec and capture decisions that affect design and implementation.

## Resolved Questions

- Q1: 多租户隔离策略（每租户独立 DB vs 单实例多租户）
  - Decision: 初始实现采用单实例多租户（单 SQLite 文件）并通过 `tenant_id` 在应用层隔离数据。
  - Rationale: 简化部署、测试与 CI，满足 MVP 需求。便于快速开发和 TDD 流程。
  - Alternatives: per-tenant DB 或 schema-per-tenant，适用于有较强隔离/合规要求的场景，作为未来迁移路径。

- Q2: 收入分成精度与四舍五入规则
  - Decision: 在内部以分（cent）为单位（整数）存储货币，并在计算时使用银行家舍入（banker's rounding）。记录每次分配的原始值与四舍五入调整。
  - Rationale: 降低累计误差并符合财务实践；方便实现准确的分成。
  - Alternatives: 引入高精度十进制库（shopspring/decimal）以支持更高精度或多币种场景。

- Q3: 支付/结算通道集成
  - Decision: 初始版本不直接集成支付网关；系统仅生成分成报表与交易记录，供外部或手动结算。
  - Rationale: 减少合规与对接风险，优先实现准确账务逻辑。
  - Alternatives: 后续可集成 Stripe/Payoneer 等并实现自动转账。

## Impact on Design
- 单 DB 设计要求所有核心表含 `tenant_id` 字段并在 service 层强制筛选。
- 会计模块需要使用整数金钱模型与集中结算任务（单元测试覆盖）。
- 报表模块应可导出 CSV/JSON 并核对 transactions 与 allocations。

## Next Steps (Phase 1 prerequisites)
- Generate `data-model.md` (基于 spec entities + tenant_id 设计，列出字段与索引)。
- Produce OpenAPI contracts for core flows (properties, rooms, contracts, tenants, settlements).
- Create `quickstart.md` with local dev steps (how to run backend + frontend with SQLite).
