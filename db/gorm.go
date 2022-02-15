package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db *gorm.DB

func IsRecordNotFound(s string) bool {
	var product Product
	if err := db.Where("code = ?", s).First(&product).Error; err != nil {
		return true
	}
	return false
}

func WorkflowGorm() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "[go-utils] ", log.Lshortfile|log.Ldate|log.Lmicroseconds),
			logger.Config{
				SlowThreshold:             time.Second,   // 慢 SQL 阈值
				LogLevel:                  logger.Silent, // 日志级别
				IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,         // 禁用彩色打印
			},
		),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	if IsRecordNotFound("D42") && IsRecordNotFound("D43") {
		db.Create(&[]Product{{Code: "D42", Price: 100}, {Code: "D43", Price: 101}})
	}

	// Read
	var product Product
	//db.First(&product, 1)
	//log.Println(product)
	//db.First(&product, "code = ?", "D43") // 查找 code 字段值为 D43 的记录
	//log.Println(product)

	// 获取第一条匹配的记录
	result := db.Where("code = ?", "D49").First(&product)
	if result.Error == gorm.ErrRecordNotFound {
		log.Println("db.Where(\"code = ?\", \"D49\").First(&product) not found")
	}

	var products []Product
	result = db.Find(&products)
	log.Println("products.Count:", result.RowsAffected)
	for _, p := range products {
		log.Println(p)
	}

	// Update - 将 product 的 price 更新为 200
	db.Where("code = ?", "D43").First(&product)
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&product, 1)
	db.Unscoped().Delete(&product, 1)
}
