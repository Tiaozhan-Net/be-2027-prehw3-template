# be-2027-prehw3-template：Gin 基础与打卡接口

本项目为西安交通大学挑战网社团技术部后端第三次培训课前预习作业模板。核心目标为跑通基础 Gin Web 服务，掌握路由注册、JSON 参数绑定（`ShouldBindJSON`）与结构体字段规则校验。

---

## 一、接口规范与需求

> **说明**：本次作业仅需在 `main.go` 中编写代码，其余配置请勿改动。

在 `main.go` 中实现打卡接口：

1. **接口路径**：`POST /api/v1/checkin`
2. **请求格式**：`application/json`
3. **字段要求**：
   - `name`（string，必填）：姓名
   - `student_id`（string，必填）：学号，长度为 10 ~ 12 位
4. **响应规范**：
   - **成功（HTTP 200 OK）**：返回打卡成功信息与回显数据

     ```json
     {
       "code": 200,
       "message": "打卡成功",
       "data": {
         "name": "张三",
         "student_id": "3122101234"
       }
     }
     ```

   - **失败（HTTP 400 Bad Request）**：参数缺失或学号长度非法时返回具体错误信息

     ```json
     {
       "code": 400,
       "message": "参数校验失败: ..."
     }
     ```

> **说明**：学号字段建议在结构体中定义为 `StudentID string` 类型，配合 `binding:"required,min=10,max=12"` 校验字符长度。若声明为数值类型（如 int），min/max 标签将判定数值大小而非字符位数。

---

## 二、本地运行与验证

```bash
# 1. 整理与下载依赖
go mod tidy

# 2. 启动服务（默认监听 8080 端口）
go run main.go
```

启动后访问 `http://localhost:8080/ping`，收到 `{"message":"pong"}` 即代表本地 Go 与 Gin 环境正常。

---

## 三、接口调试示例

使用 Apifox 或任一 HTTP 工具发送请求：

- **Method**：`POST`
- **URL**：`http://localhost:8080/api/v1/checkin`
- **Header**：`Content-Type: application/json`
- **Body**：

  ```json
  {
    "name": "张三",
    "student_id": "3122101234"
  }
  ```

需分别验证正常入参（返回 200）以及缺少字段或学号长度非法（返回 400）两种场景。

---

## 四、提交与验收

1. **仓库命名**：通过模板建仓（建仓方式与往常一致），建议统一命名为 `be-2027-prehw3-姓名拼音`（例如 `be-2027-prehw3-zhangsan`）；
2. **CI 验证**：代码 push 后 GitHub Actions 自动化测试通过（绿勾 ✅）；
3. **抓包截图**：将 Apifox 截图存放在 `docs/images/` 并更新至第五节。

---

## 五、抓包截图展示

截图存放规则：

1. 截图文件存放于本工程 `docs/images/` 目录；
2. 必须使用 Markdown 相对路径引用，不得使用操作系统本地绝对路径；
3. 截图内容应完整展示请求路径、入参、响应状态码与响应体。

### 1. 打卡成功（HTTP 200）

*对应图片路径：`docs/images/checkin_200.png`*

![打卡成功截图](docs/images/checkin_200.png)

### 2. 参数校验失败（HTTP 400）

*对应图片路径：`docs/images/checkin_400.png`*

![参数校验失败截图](docs/images/checkin_400.png)
