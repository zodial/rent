# Feature Specification: [FEATURE NAME]

**Feature Branch**: `[###-feature-name]`  
**Created**: [DATE]  
**Status**: Draft  
**Input**: User description: "$ARGUMENTS"

## Purpose (用户价值) *(mandatory)*

[简短说明：本功能为何对用户/业务有价值。用一句话概括核心用户价值。]

## Scope & Out-of-Scope

- Includes: [列出当前交付范围]  
- Excludes: [列出明确不包含的内容]

## Prioritized User Scenarios & Acceptance Tests *(mandatory & test-first friendly)*

<!--
  IMPORTANT: 每个场景必须是一个可独立交付和测试的用户旅程。
  优先级使用 P1 (必须)、P2 (重要)、P3 (可选)。
-->

### User Story 1 - [Brief Title] (Priority: P1)

[用简短非技术语言描述用户旅程和预期价值]

**Why this priority**: [说明业务价值]

**Independent Test / TDD guidance**: 
- 测试类型：Unit / Integration / E2E
- 必须先写失败测试：列出至少1个关键断言（例如：DB 中存在行；API 返回 201；权限为 403）

**Acceptance Scenarios (Given / When / Then)**:

1. **Given** [初始状态], **When** [动作], **Then** [预期结果: 可断言的输出或 DB 状态]
2. **Given** [初始状态], **When** [动作], **Then** [预期结果]

**Performance / Non-functional checks** (if applicable):
- 关键路径响应阈值（例如：首交互 ≤ 100ms）或 CI 性能断言

---

### User Story 2 - [Brief Title] (Priority: P2)

[描述...]

**Why this priority**: [说明价值]

**Independent Test / TDD guidance**: [列出应该先编写的测试类型与关键断言]

**Acceptance Scenarios (Given / When / Then)**:

1. **Given** ..., **When** ..., **Then** ...

---

### User Story 3 - [Brief Title] (Priority: P3)

[描述...]

**Acceptance Scenarios**:

1. **Given** ..., **When** ..., **Then** ...

---

## Edge Cases *(must enumerate and make testable)*

- [边界情况 1] -> 预期行为与失败断言（例如：重复提交应返回 409）
- [边界情况 2] -> 预期行为与失败断言
- [错误场景] -> 例如：权限不足、网络超时、DB 约束冲突

## Requirements *(mandatory & testable)*

### Functional Requirements (每项需可被自动化断言)

- **FR-001**: System MUST [清晰行为描述，可直接转化为测试断言]
- **FR-002**: System MUST [清晰行为描述]
- **FR-003**: System MUST [清晰行为描述]

*If a requirement is unclear mark it as [NEEDS CLARIFICATION] with the specific question.*

### Non-Functional Requirements

- Performance: [例如：关键页面首交互 ≤ 100ms]
- Security: [例如：敏感字段需在 DB 中加密存储]
- Observability: [例如：关键写入必须产生日志和审计记录]
- Test Coverage: [例如：Unit + Integration coverage >= 90%]

## Data Model / Key Entities *(high-level, no implementation details)*

- **EntityName**: { key_attr, key_attr }
- **EntityName**: { relation: OtherEntity }

## Success Criteria *(mandatory & measurable)*

- **SC-001**: [业务可衡量指标，例如：用户能在移动端在 2 分钟内完成主要任务]
- **SC-002**: [例如：API 在并发 N 个请求下响应不超过 X ms（如适用）]
- **SC-003**: [测试覆盖率 >= 90% 并且 CI 在合并前强制校验]
- **SC-004**: [如有性能目标，需在 PR 中附基准对比数据]

## Testing Guidance & CI Gates *(enforce TDD & automated checks)*

- TDD rule: 开发时必须先添加失败测试（列出至少 1 个单元测试与 1 个集成测试）。
- CI gates (must pass before merge): lint -> unit tests (coverage) -> integration tests -> performance checks (if provided).
- 所有 PR 必须包含测试结果与覆盖率报告并在描述中引用关键断言。

## API / UX Contracts *(optional but encouraged to include examples)*

- API: [HTTP verb] /path -> 请求体 / 响应示例 / 错误情形
- UX: 简短说明关键界面与交互的必备元素（尤其针对移动端）

## Implementation Notes & Constraints

- Technology constraints (if any): [例如：Backend Go, Frontend Vue3+TS, SQLite]
- 儿童条款：尽量避免引入非必要第三方库；任何新增依赖需在 PR 中说明理由。

## Open Questions

- Q1: [填入开放问题并标注影响范围]
- Q2: [填入开放问题]

## Change Log

- [DATE] - Created by [AUTHOR]
- [DATE] - Updated: [summary of changes]

<!-- End of spec template -->
