# 博客系统后端

作者: xibeiwind 创建时间: 2025-11-10 最后更新: 2025-11-10 版本: v1.0.0

基于 Go + Gin + GORM 开发的博客系统后端，提供完整的文章管理、用户认证和评论功能。

## 项目特性

1. **完整的文章管理功能** - 包括创建、编辑、删除、查看文章列表、文章详情等
2. **用户认证功能** - 包括用户注册、登录，密码的加密存储
3. **评论功能** - 包括创建、编辑、删除、查看评论列表、评论详情等
4. **基于JWT的Token认证** - 安全的身份验证机制
5. **SQLite数据库** - 轻量级数据库，便于部署和开发

## 技术栈

- **后端框架**: Gin
- **数据库**: SQLite + GORM
- **认证**: JWT (JSON Web Tokens)
- **密码加密**: bcrypt
- **开发语言**: Go 1.25.1

## 项目结构

```
task_four/
├── main.go                 # 应用入口，路由配置
├── middlewares.go          # 中间件（数据库、JWT认证）
├── go.mod                  # Go模块依赖
├── go.sum                  # 依赖校验
├── task4.db                # SQLite数据库文件
└── blogs/                  # 业务逻辑模块
    ├── models.go           # 数据模型定义
    ├── auth.go             # 用户认证逻辑
    ├── articles.go         # 文章管理逻辑
    ├── comments.go         # 评论管理逻辑
    └── errors.go           # 错误处理
```

## 数据模型

### User (用户表)
- `ID` - 主键
- `Username` - 用户名（唯一，非空）
- `Password` - 密码（加密存储）
- `Email` - 邮箱（唯一，非空）

### Post (文章表)
- `ID` - 主键
- `Title` - 文章标题
- `Content` - 文章内容
- `UserID` - 作者ID（外键）
- `User` - 作者信息
- `Comments` - 评论列表
- `CreatedAt` - 创建时间
- `UpdatedAt` - 更新时间

### Comment (评论表)
- `ID` - 主键
- `Content` - 评论内容
- `UserID` - 评论者ID（外键）
- `PostID` - 文章ID（外键）
- `User` - 评论者信息
- `Post` - 文章信息
- `CreatedAt` - 创建时间

## API 文档

### 认证相关接口

#### 用户注册
- **URL**: `POST /register`
- **认证**: 不需要
- **参数**:
  ```json
  {
    "username": "string",
    "password": "string", 
    "email": "string"
  }
  ```
- **响应**:
  ```json
  {
    "code": 200,
    "msg": "注册成功",
    "data": null
  }
  ```

#### 用户登录
- **URL**: `POST /login`
- **认证**: 不需要
- **参数**:
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **响应**:
  ```json
  {
    "code": 200,
    "msg": "登录成功",
      "token": "jwt_token_string"
    
  }
  ```

### 文章相关接口

#### 获取文章列表
- **URL**: `GET /articles`
- **认证**: 不需要
- **响应**: 文章列表

#### 获取文章详情
- **URL**: `GET /articles/:id`
- **认证**: 不需要
- **响应**: 文章详情

#### 创建文章
- **URL**: `POST /articles`
- **认证**: 需要 (Bearer Token)
- **参数**:
  ```json
  {
    "title": "string",
    "content": "string"
  }
  ```

#### 更新文章
- **URL**: `PUT /articles/:id`
- **认证**: 需要 (Bearer Token)
- **参数**: 同创建文章

#### 删除文章
- **URL**: `DELETE /articles/:id`
- **认证**: 需要 (Bearer Token)

### 评论相关接口

#### 获取文章评论
- **URL**: `GET /comments/:post_id`
- **认证**: 不需要
- **响应**: 评论列表

#### 创建评论
- **URL**: `POST /comments/:post_id`
- **认证**: 需要 (Bearer Token)
- **参数**:
  ```json
  {
    "content": "string"
  }
  ```
### 更新评论
- **URL**: `PUT /comments/:comment_id`
- **认证**: 需要 (Bearer Token)
- **参数**:
  ```json
  {
    "content": "string"
  }
  ```
#### 删除评论
- **URL**: `DELETE /comments/:comment_id`
- **认证**: 需要 (Bearer Token)

## 安装和运行

### 环境要求
- Go 1.25.1 或更高版本

### 安装步骤

1. **克隆项目**
   ```bash
   git clone https://github.com/xibeiwind/study_task.git
   cd study_task/task_four
   ```

2. **安装依赖**
   ```bash
   go mod tidy
   ```

3. **运行项目**
   ```bash
   go run main.go
   ```

4. **访问API**
   服务器将在 `http://localhost:5055` 启动

### 构建可执行文件
```bash
go build -o blog-api
./blog-api
```

## 开发说明

### 数据库配置
- 使用 SQLite 数据库，数据库文件为 `task4.db`
- 自动迁移：启动时自动创建表结构
- 开发环境使用明文密钥，生产环境建议使用环境变量

### JWT 配置
- 密钥: `your-secret-key` (开发环境)
- Token 格式: Bearer Token
- 需要在请求头中添加: `Authorization: Bearer <token>`

### 错误处理
- 统一的响应格式:
  ```json
  {
    "code": 状态码,
    "msg": "消息",
    "data": 数据或null
  }
  ```

## 注意事项

1. **生产环境部署**:
   - 修改 JWT 密钥为强密码
   - 使用环境变量管理敏感信息
   - 考虑使用 PostgreSQL 或 MySQL 替代 SQLite

2. **安全考虑**:
   - 密码使用 bcrypt 加密存储
   - JWT Token 有过期时间
   - 输入验证和参数绑定

3. **性能优化**:
   - 可添加 Redis 缓存
   - 数据库连接池配置
   - API 限流和熔断

## 许可证

MIT License
