# 用户行为记录API说明

## 概述

本文档描述了游戏分发平台中用户行为记录相关的API接口，包括搜索历史记录、游戏行为跟踪等功能。**用户必须登录后才能使用这些功能**。

## 基础信息

- **基础路径**: `/api/v1/game-engine`
- **认证方式**: 所有接口都需要用户登录认证
- **数据格式**: JSON
- **字符编码**: UTF-8

## 业务逻辑说明

### 用户登录要求
- 用户必须登录后才能玩游戏
- 所有行为记录都与登录用户ID关联
- 确保行为记录的精准性和数据完整性

### 数据隐私保护
- 用户只能访问自己的数据
- 支持用户清空个人搜索历史
- 符合数据保护法规要求

## API接口列表

### 1. 搜索历史管理

#### 1.1 获取搜索历史

**接口描述**: 获取登录用户的搜索历史记录

**请求方式**: `GET`

**接口路径**: `/search-history`

**请求参数**:
```
user_id: int64    // 必填，用户ID
limit: int        // 可选，返回数量限制，默认20
```

**响应示例**:
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "list": [
            {
                "id": 1,
                "search_keyword": "王者荣耀",
                "search_time": "2024-01-15 10:30:00",
                "result_count": 5
            }
        ],
        "total": 1
    }
}
```

#### 1.2 清空搜索历史

**接口描述**: 清空登录用户的所有搜索历史

**请求方式**: `DELETE`

**接口路径**: `/search-history`

**请求参数**:
```json
{
    "user_id": 123    // 必填，用户ID
}
```

**响应示例**:
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "success": true,
        "message": "搜索历史清空成功"
    }
}
```

#### 1.3 获取热门搜索关键词

**接口描述**: 获取最近7天的热门搜索关键词（全站统计）

**请求方式**: `GET`

**接口路径**: `/popular-keywords`

**请求参数**:
```
limit: int    // 可选，返回数量限制，默认10
```

**响应示例**:
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "keywords": ["王者荣耀", "和平精英", "原神"]
    }
}
```

### 2. 游戏行为记录

#### 2.1 记录游戏行为

**接口描述**: 记录用户的游戏相关行为

**请求方式**: `POST`

**接口路径**: `/game-behavior`

**请求参数**:
```json
{
    "user_id": 123,           // 必填，用户ID
    "game_id": 456,           // 必填，游戏ID
    "behavior_type": 1,       // 必填，行为类型(1:查看 2:下载 3:收藏 4:评分)
    "ip_address": "string"    // 可选，IP地址
}
```

**行为类型说明**:
- `1`: 查看游戏详情
- `2`: 下载游戏
- `3`: 收藏游戏
- `4`: 评分游戏

**响应示例**:
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "success": true,
        "message": "游戏行为记录成功"
    }
}
```

#### 2.2 获取游戏历史

**接口描述**: 获取登录用户的游戏行为历史

**请求方式**: `GET`

**接口路径**: `/game-history`

**请求参数**:
```
user_id: int64    // 必填，用户ID
limit: int        // 可选，返回数量限制，默认50
```

**响应示例**:
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "list": [
            {
                "id": 1,
                "game_id": 123,
                "game_name": "王者荣耀",
                "behavior_type": 1,
                "behavior_time": "2024-01-15 10:30:00",
                "ip_address": "192.168.1.1"
            }
        ],
        "total": 1
    }
}
```

## 使用流程

### 基本使用流程

1. **用户登录**: 用户必须先登录获取用户ID
2. **记录行为**: 在用户操作时调用相应的行为记录接口
3. **查询历史**: 使用用户ID查询用户的历史记录
4. **管理数据**: 支持清空搜索历史等数据管理操作

### 集成示例

```javascript
// 假设用户已登录，获取到用户ID
const userId = getCurrentUserId();

// 记录游戏查看行为
async function recordGameView(gameId) {
    const response = await fetch('/api/v1/game-engine/game-behavior', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${getAuthToken()}` // 认证token
        },
        body: JSON.stringify({
            user_id: userId,
            game_id: gameId,
            behavior_type: 1 // 查看
        })
    });
    
    return response.json();
}

// 记录搜索行为
async function recordSearch(keyword, resultCount) {
    // 这里需要调用搜索接口，在搜索成功后记录
    // 具体实现取决于你的搜索接口设计
}

// 获取搜索历史
async function getSearchHistory() {
    const response = await fetch(`/api/v1/game-engine/search-history?user_id=${userId}&limit=20`, {
        headers: {
            'Authorization': `Bearer ${getAuthToken()}`
        }
    });
    
    return response.json();
}

// 清空搜索历史
async function clearSearchHistory() {
    const response = await fetch('/api/v1/game-engine/search-history', {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${getAuthToken()}`
        },
        body: JSON.stringify({
            user_id: userId
        })
    });
    
    return response.json();
}
```

## 注意事项

### 认证要求

1. **登录验证**: 所有接口都需要用户登录认证
2. **用户ID验证**: 确保用户只能访问自己的数据
3. **Token管理**: 前端需要管理认证token

### 数据安全

1. **权限控制**: 用户只能访问自己的行为数据
2. **数据隔离**: 不同用户的数据完全隔离
3. **隐私保护**: 支持用户清空个人数据

### 性能优化

1. **异步记录**: 行为记录采用异步方式，不影响用户体验
2. **批量处理**: 支持批量记录多个行为，减少API调用次数
3. **分页查询**: 历史记录支持分页查询，避免大量数据传输

### 错误处理

1. **认证失败**: 返回401状态码，提示用户重新登录
2. **参数验证**: 返回400状态码，提示参数错误
3. **服务器错误**: 返回500状态码，记录详细错误日志

## 扩展功能

### 智能推荐

基于用户行为历史，提供个性化游戏推荐：
- 相似游戏推荐
- 热门游戏推荐
- 新游戏推荐

### 数据分析

为运营团队提供用户行为分析：
- 用户活跃度统计
- 游戏热度分析
- 搜索趋势分析

### 数据导出

支持用户导出个人数据：
- 搜索历史导出
- 游戏行为记录导出
- 数据格式：CSV、JSON等 