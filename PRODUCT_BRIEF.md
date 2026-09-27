# Qy API Change Toolkit

## 一句话定位

**Detect changes before they break your workflow.**

中文：**在变化影响工作流之前，发现它。**

## 产品类型

一次性购买的开发者工具包，不是订阅制 SaaS。

包含四个共享报告引擎的工具：

```text
openapi-change   OpenAPI 破坏性变化检查
schema-change    JSON Schema 兼容性检查
sitemap-change   Sitemap 新增、删除、修改检查
feed-change      RSS/JSON Feed 内容变化检查
```

用户在本地或 CI 中运行工具，不需要注册产品账号，也不需要把 API 文件上传到你的服务器。

## 目标用户

- 使用第三方 API 的开发团队
- 维护 OpenAPI/Swagger 文档的后端团队
- 需要检查供应商 API、网站结构和 Feed 变化的技术负责人
- 希望让 AI Agent 自动识别 API 变化的开发者

## 核心问题

外部 API、Schema、Sitemap 或 Feed 更新后，用户通常只能在流程中断后才发现。需要提前知道：

- 哪些接口被删除或改名
- 哪些请求参数变成必填
- 哪些返回字段被删除或改类型
- 哪些认证 Scope 发生变化
- 哪些状态码和错误结构发生变化
- 哪些接口已经废弃
- 哪些 URL 或内容发生新增、删除、修改

## 产品输出

```text
旧文件/URL + 新文件/URL
        |
        v
Qy API Change Toolkit
        |
        +-- Breaking Changes
        +-- Compatible Changes
        +-- Deprecations
        +-- Authentication Changes
        +-- URL / Content Changes
        +-- JSON Event
        +-- Markdown / HTML / SARIF Report
```

## 免费版

```text
本地 CLI
OpenAPI 3.x diff
Markdown 报告
基础 Breaking / Compatible 分类
```

## Pro 一次性版

首发价格：¥79

```text
JSON Event 输出
SARIF 输出
GitHub Action
HTML 报告
四个工具全部解锁
多个文件和 URL 比较
长期版本更新
```

## 工具模块

```text
openapi-change   API 接口与 Schema 变化
schema-change    JSON Schema 兼容性变化
sitemap-change   网站 URL 结构变化
feed-change      RSS/JSON Feed 内容变化
```

## 产品简介

### 中文

Qy API Change Toolkit 是一个面向开发者和 AI Agent 的变化分析工具包，覆盖 OpenAPI、JSON Schema、Sitemap 和 RSS/JSON Feed。它比较两个版本，自动识别破坏性变化、兼容变化、废弃接口、字段变化、URL 变化和内容变化，并生成可读报告与机器可读事件。工具支持本地 CLI 和 GitHub Actions，所有分析都可以在本地完成。

### English

Qy API Change Toolkit is a change analysis toolkit for developers and AI agents, covering OpenAPI, JSON Schema, sitemaps, and RSS/JSON feeds. It compares two versions and identifies breaking changes, compatible changes, deprecations, schema changes, URL changes, and content changes. Generate human-readable reports and machine-readable events from the CLI or GitHub Actions.

## 首页标题

```text
Catch every change before it breaks your workflow.
```

## 首页副标题

```text
Compare APIs, schemas, sitemaps, and feeds. Get a clear report your team or AI agent can act on.
```

## 行动按钮

```text
Run the free diff
Download Pro
View documentation
```

## 不提供的服务

```text
人工安装
定制开发
远程调试
企业 SLA
电话支持
客户专属规则
用户数据托管
```

## 自动化交付

```text
支付
→ 7天自动退款
→ 下载链接
→ License
→ GitHub Release
→ 发票
→ 固定退款政策
```

## 获客入口

```text
GitHub Marketplace 免费 Action
GitHub README
npm 免费 CLI
开发者工具目录
OpenAPI 相关搜索词
```

## 个人卖家准备材料

```text
个人邮箱
GitHub 账户
Paddle 个人卖家账户
身份证件
地址
税务信息
收款账户
产品邮箱：qingy5461@outlook.com
产品链接
支持邮箱
```

## 验收目标

```text
产品可本地运行
两个 OpenAPI 文件可生成报告
免费版和 Pro 版边界清晰
一次购买可自动交付
不需要用户账号
不需要云端数据库
不需要人工安装
连续30天无需人工维护
```
