# 项目说明

本项目包含三个主要脚本文件：`realize_go.py`、`rectify.py` 和 `build_protocol.bat`。这些脚本用于处理和生成协议文件，并将协议注册及处理函数写入相应的 Go 文件中。

## 文件说明

### `config.json`

该文件用于存储配置信息，包括 `.proto` 文件目录、要遍历的目录、要过滤的 `.proto` 文件列表以及节点和 `handle_msg.go` 文件路径的映射关系。

#### 示例

```json
{
    "proto_dir": "../internal/protocol",
    "gofile_directory": "../internal/protocol/gofile",
    "filtered_files": ["login.proto", "base_type.proto", "baseinfo.proto"],
    "node_to_go_file": {
        "game": "../nodes/game/module/player/handle_msg.go",
        "user": "../nodes/user/module/player/handle_msg.go"
    }
}
```

### `realize_go.py`

该脚本用于提取 `.proto` 文件中以 `Req` 结尾的消息，并生成相应的注册代码和处理函数，写入对应的 `handle_msg.go` 文件中。

#### 功能

1. **提取消息**：从指定目录中提取所有未被注释且以 `Req` 结尾的消息及其注释。
2. **生成注册代码**：生成注册代码行，不包含注释。
3. **生成处理函数**：生成处理函数代码，包含注释。
4. **更新 Go 文件**：将生成的注册代码和处理函数写入对应的 `handle_msg.go` 文件中。
5. **注释指定行**：注释 `gofile` 目录下所有 `.go` 文件中包含指定 proto 语句的行。

#### 使用方法

1. 修改 `config.json` 文件，设置您的 `.proto` 文件目录、要遍历的目录、要过滤的 `.proto` 文件列表以及节点和 `handle_msg.go` 文件路径的映射关系。
2. 运行脚本：`python realize_go.py`

### `rectify.py`

该脚本用于调整 `.proto` 文件的编码格式，并将字段名转换为驼峰命名。

#### 功能

1. **列出文件**：列出指定目录下的所有 `.proto` 文件。
2. **处理字段**：将字段名转换为驼峰命名，并将自定义类型的首字母大写。
3. **调整协议编码**：调整协议编码，并将包名本地化。

#### 使用方法

1. 修改 `fileDir` 变量，设置您的 `.proto` 文件目录。
2. 运行脚本：`python rectify.py`

### `build_protocol.bat`

该批处理文件用于执行 `rectify.py` 和 `realize_go.py` 脚本，并生成 Go 和 JavaScript 协议文件。

#### 功能

1. **执行 `rectify.py`**：调整 `.proto` 文件的编码格式。
2. **生成协议文件**：使用 `protoc` 生成 Go 和 JavaScript 协议文件。
3. **执行 `realize_go.py`**：生成协议注册及处理函数，并写入相应的 Go 文件中。

#### 使用方法

1. 确保 `rectify.py` 和 `realize_go.py` 脚本在同一目录下。
2. 运行批处理文件：`build_protocol.bat`

## 示例

假设您的项目目录结构如下：

```
project/
│
├── internal