package main

/*
生成sql对应的model结构数据 至sql/model目录
*/
import (
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
)

const sqlModelDir = "nodes/leaf/jettengame/sql/model/"
const modelPkgPath = "model"

// 请清空 [sqlmodel/model]目录下的文件,再执行
func main() {
	user := "root"
	password := "ko8899110"
	address := "127.0.0.1"
	port := "3306"
	dbName := "minigame"
	dataSourceName := user + ":" + password + "@tcp(" + address + ":" + port + ")/" + dbName + "?charset=utf8"
	//注意 注意 此处为移除gorm的日志自定义了相关结构。正式使用时 请放开
	db, err1 := gorm.Open(gmysql.Open(dataSourceName), &gorm.Config{
		//Logger:      DiscardLogger{}.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "t_",   // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true,                              // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "CId"), // use name replacer to change struct/field name before convert it to db name
		},
		PrepareStmt: false,
	})
	if err1 != nil {
		panic(err1)
	}

	// 生成gorm的model文件,如果已经生成则
	GenModel(sqlModelDir, db)
	pDb, _ := db.DB()
	pDb.Close()
}

func GenModel(outDir string, db *gorm.DB) {
	fileInfos, err := os.ReadDir(outDir)
	if err != nil {
		if err = os.Mkdir(outDir, 0755); err != nil {
			log.Printf("生成[%v]目录失败 err%v", outDir, err)
			return
		}
	}
	if 0 < len(fileInfos) {
		log.Printf("[%v]目录下已有model文件. 如需生成,请清空目录下的所有model文件!!!", outDir)
		//return
	}
	// 生成model数据 	从当前数据库的所有表生成结构
	g := gen.NewGenerator(gen.Config{
		OutPath:      outDir,
		ModelPkgPath: modelPkgPath,
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		//FieldNullable:     false, // generate pointer when field is nullable
		//FieldCoverable:    true,  // generate pointer when field has default value
		//FieldWithIndexTag: true,  // generate with gorm index tag
		//FieldWithTypeTag:  true,  // generate with gorm column type tag
	})

	// 选择数据库连接
	g.UseDB(db)
	tables, _ := db.Migrator().GetTables()

	// 过滤管理后台的数据表
	for _, table := range tables {
		if !strings.Contains(table, "sys_") {
			g.GenerateModel(table)
		}
	}

	// 调整部分数据类型
	var dataMap = map[string]func(gorm.ColumnType) (dataType string){
		// int mapping
		"int": func(columnType gorm.ColumnType) (dataType string) {
			if n, ok := columnType.Nullable(); ok && n {
				return "*int32"
			}
			return "int32"
		},

		// bool mapping
		"tinyint": func(columnType gorm.ColumnType) (dataType string) {
			ct, _ := columnType.ColumnType()
			if strings.HasPrefix(ct, "tinyint(1)") {
				return "bool"
			}
			return "byte"
		},
	}
	g.WithDataTypeMap(dataMap)

	// 生成代码
	g.Execute()
}
