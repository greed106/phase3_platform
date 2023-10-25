## 文件目录（2023/10/25更新）
```plaintext
phase3_platform/                      # 项目的根目录
│
├── app/                             # 主要应用逻辑和功能的目录
│   ├── ai/                          # AI模块
│   │   ├── handler.go               # AI功能的处理程序，如请求处理和响应
│   │   └── router.go                # AI功能的路由配置
│   │
│   ├── models/                      # 数据模型目录，定义了数据结构和与数据库交互的逻辑
│   │   ├── ai_models.go             # AI数据模型
│   │   ├── filed_models.go          # Filed数据模型
│   │   └── sensor_models.go         # 传感器数据模型
│   │
│   ├── repository/                  # 数据存储逻辑，可能包括与数据库的交互
│   │   └── ai_repository.go         # AI数据的存储逻辑
│   │
│   └── sensor/                      # 传感器模块
│       ├── handler.go               # 传感器功能的处理程序
│       └── router.go                # 传感器功能的路由配置
│
├── config/                          # 配置文件目录
│   ├── config.yml                   # 主要的配置文件
│   └── postgresql_config.yml        # PostgreSQL数据库的配置
│
├── database/                        # 数据库连接和配置的目录
│   └── postgresql_connection.go     # PostgreSQL数据库的连接逻辑
│
├── routers/                         # 路由配置目录
│   └── routers.go                   # 主路由配置文件
│
├── go.mod                           # Go模块描述文件
├── main.go                          # 主执行文件，项目的入口点
└── README.md                        # 项目的README，说明和文档
```
