// main.go
package main
 
import (
	"auth-api/database"
	"auth-api/routes"
 
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
 
func main() {
	// GORMセット
	database.Connect()
 
	app := fiber.New()

	//CROSの設定
	app.Use(cors.New(cors.Config{
		// https://docs.gofiber.io/api/middleware/cors#config
		AllowCredentials: true,
	}))

	routes.Setup(app)
 
	app.Listen(":90")
}


// package main
 
// import (
// 	"fmt"
// 	"os"
 
// 	"github.com/gofiber/fiber/v2"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )
 
// var (
// 	schema = "%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local"
// 	// docker-compose.ymlに設定した環境変数を取得
// 	username       = os.Getenv("MYSQL_USER")
// 	password       = os.Getenv("MYSQL_PASSWORD")
// 	dbName         = os.Getenv("MYSQL_DATABASE")
// 	datasourceName = fmt.Sprintf(schema, username, password, dbName)
// )
 
// func main() {
// 	// GORMセット
// 	db, err := gorm.Open(mysql.Open(datasourceName), &gorm.Config{})
// 	if err != nil {
// 		panic("Could not connect to the database")
// 	}
 
// 	fmt.Println(db)
// }
