# Rent Constitution

## Core Principles

### Stack & Tooling Constraints
后端必须使用 Golang。前端必须使用 Vue 3 与 TypeScript，使用原生 Web API 进行网络与 DOM 操作。项目应避免非必要的第三方依赖库；任何新增依赖必须在 PR 中明确理由与替代方案。该原则为项目实现一致性、可维护性与轻量化的基础许可政策。

### Test-Driven Development (NON-NEGOTIABLE)
MUST: 所有新功能、修复与重构均须遵循 TDD 流程：先写失败测试、实现、再重构。MUST: 单元测试与集成测试覆盖率不得低于 90%。MUST: CI 必须在合并前验证测试覆盖率与测试通过。任何未达到覆盖率的变更必须附带明确的风险说明与后续补全计划。

### Performance & Responsiveness
MUST: 前端 UI 响应时间（首交互或关键路径响应）目标为 <= 100ms。SHOULD: 后端端到端响应应有明确 SLO/目标并通过基准与负载测试验证。性能退化必须在 PR 中提供基准对比数据与回退或优化计划。所有性能断言须纳入自动化性能测试。

### Simplicity & Observability
代码设计应优先简洁、可理解与可测试（YAGNI 原则）。MUST: 必要的结构化日志、指标与链路追踪必须就位以支持故障定位。SHOULD: 新模块自带基本健康检查与可测性入口（可测试的边界、副作用隔离）。

### Versioning & Stability
MUST: 遵循语义化版本控制。对外兼容性破坏的变更需标记为 MAJOR，并附带迁移说明与兼容层或 deprecation 计划。次要功能添加或扩展为 MINOR，文档/修正为 PATCH。发布需包含变更日志与回滚策略。

## Additional Constraints

- 技术栈要求：后端使用 Go modules；前端使用原生 ESM + Vite 或等价浅依赖构建（仅当必要）。网络协议优先使用 REST/JSON（或明确说明的替代方案）。
- 安全与合规：敏感数据必须加密传输与存储；任何第三方服务需通过安全评估并写入依赖清单与风险登记表。
- 构建与部署：构建产物必须可重复生成；应包含可自动化的 CI/CD 流程定义与回滚步骤。

## Development Workflow

- 分支策略：主分支为 `main`。所有功能通过 feature branch + PR 合并（feature branch 命名建议以三位序号开头，例如 `001-feature-name`）。
- 代码审查：所有 PR 至少一位维护者审批并通过自动化测试/性能检查。重大设计变更须附设计文档与兼容性评估。
- 质量门：CI 必须在合并前验证：静态检查（lint）、测试（覆盖率阈值）、安全扫描（如有）、性能基准（关键路径）。
- Triage 与 发布：紧急回归应有隔离分支与热修复流程；常规发布需伴随发布说明与变更日志。

## Governance

本宪章优先于其他非正式实践；任何修订须经 PR、维护者批准并在变更说明中列出影响范围与迁移步骤。若修订引入不兼容规则，版本号须做 MAJOR 升级并提供迁移计划。合规审查由维护者轮值小组负责，关键合规问题必须记录并在下一次发布前解决。

**Version**: 1.0.0 | **Ratified**: TODO(RATIFICATION_DATE) | **Last Amended**: 2025-11-14
