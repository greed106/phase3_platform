package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

// PostgresConfig 表示PostgreSQL数据库的配置
type PostgresConfig struct {
	Dialect         string `yaml:"dialect"`
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	DBName          string `yaml:"dbname"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
	// 可以为 SSLMode 添加更多的配置选项，如下所示
	// SSLMode         string `yaml:"sslmode"`
}

var log = logrus.New()
var DB *gorm.DB

func InitializePostgresConnection(confFilePath string) error {
	var conf PostgresConfig
	// 读入配置文件
	file, err := os.ReadFile(confFilePath)
	if err != nil {
		log.Errorf("Error reading config file: %v", err)
		return err
	}
	// 解析配置文件，将解析结果存入conf中
	if err = yaml.Unmarshal(file, &conf); err != nil {
		log.Errorf("Error unmarshalling config file: %v", err)
		return err
	}

	// 构造连接字符串
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.DBName,
	)

	// 使用 gorm 连接到数据库
	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Errorf("Error connecting to the database: %v", err)
		return err
	}

	// 设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Errorf("Error getting underlying sql.DB: %v", err)
		return err
	}

	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetime) * time.Second)

	log.Info("PostgreSQL connected successfully with gorm!")
	return nil
}
