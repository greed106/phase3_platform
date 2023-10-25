testProject
|-- app
|   |-- ai
|   |   |-- handler.go
|   |   |-- router.go
|   |   |-- models
|   |       |-- ai_models.go
|   |-- sensor
|   |   |-- handler.go
|   |   |-- router.go
|   |   |-- models
|   |       |-- sensor_models.go
|-- models  // 公共模型
|   |-- field.go
|-- routers
|   |-- router.go  // 主路由文件，组合ai和sensor的路由
|-- middlewares
|   |-- auth_middleware.go
|-- utils
|   |-- date_utils.go
|-- config
|   |-- config.yml
|   |-- loader.go
|-- go.mod
|-- go.sum
|-- main.go
