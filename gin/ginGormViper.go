package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/spf13/viper"
)

var DB dbdata

type DEPARTMENT struct {
	DEPT_ID   int    `gorm:"column:DEPT_ID;"`
	DEPT_NAME string `gorm:"column:DEPT_NAME;"`
	HEAD_ID   int    `gorm:"column:HEAD_ID;"`
	DEPT_HEAD string `gorm:"column:DEPT_HEAD;"`
} //Department

type dbdata struct {
	db *gorm.DB
} //config password

func connectdb() (string, int, string, string, string) {
	viper.SetConfigFile("config2.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fetal error config file: %s", err))
	}

	server1 := viper.GetString("mssql.server")
	port1 := viper.GetInt("mssql.port")
	user1 := viper.GetString("mssql.user")
	password1 := viper.GetString("mssql.password")
	database1 := viper.GetString("mssql.database")

	return server1, port1, user1, password1, database1
} //connectdatabase

func main() {
	r := gin.Default()
	r.POST("/member", database)
	r.POST("/insert", insertdb)
	r.POST("/update/:DEPT_ID/:DEPT_NAME/:HEAD_ID/:DEPT_HEAD", updatedb)
	r.POST("/delete", deletedb)
	server, port, user, password, database := connectdb()
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)
	db, err := gorm.Open("mssql", connectionString)
	if err != nil {
		panic("failed to connect database")
	}
	DB.db = db
	defer db.Close()

	r.Run(":8011") //run port
} //Main

func database(c *gin.Context) {
	var deptInfo []DEPARTMENT
	err := DB.db.Select("DEPT_ID,DEPT_NAME,HEAD_ID,DEPT_HEAD").Table("DEPARTMENT").Find(&deptInfo).Error
	if err != nil {
		fmt.Print(err)
		panic("failed to connect database")
	}
	fmt.Printf("deptInfo = %+v", deptInfo)
	c.JSON(200, gin.H{
		"Status": deptInfo,
	})
} //database

func insertdb(c *gin.Context) {
	var deptInfo []DEPARTMENT
	err := DB.db.Table("DEPARTMENT").Exec("INSERT INTO DEPARTMENT (DEPT_ID ,DEPT_NAME,HEAD_ID,DEPT_HEAD) VALUES (8,'backend',350,'mike')").Find(&deptInfo).Error
	if err != nil {
		fmt.Print(err)
		panic("failed to connect database")
	}
	fmt.Printf("deptInfo = %+v", deptInfo)

	c.JSON(200, gin.H{
		"Status": deptInfo,
	})
} //insert

func updatedb(c *gin.Context) {
	var deptInfo DEPARTMENT
	e := c.BindJSON(&deptInfo)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("deptInfo = %+v", deptInfo)
	fmt.Println(deptInfo.DEPT_ID)
	err := DB.db.Table("DEPARTMENT").
		Where("DEPT_ID = ? ", deptInfo.DEPT_ID).
		Update(&deptInfo)
	if err == nil {
		fmt.Println(err)
		panic("failed to insert database")
	}


	c.JSON(200, gin.H{
		"Status": deptInfo,
	})
} //update

func deletedb(c *gin.Context) {
	var deptInfo []DEPARTMENT
	err := DB.db.Exec("DELETE FROM DEPARTMENT WHERE DEPT_ID = '8'").Error
	err = DB.db.Select("DEPT_ID,DEPT_NAME,HEAD_ID,DEPT_HEAD").Table("DEPARTMENT").Find(&deptInfo).Error
	if err != nil {
		fmt.Print(err)
		panic("failed to connect database")
	}

	fmt.Printf("deptInfo = %+v", deptInfo)

	c.JSON(200, gin.H{
		"Status": deptInfo,
	})
} //delete
