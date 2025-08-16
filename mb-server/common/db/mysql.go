package db

import (
	"fmt"
	"mb-server/common/config"
	"mb-server/common/logger"

	"gorm.io/gorm/clause"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MysqlDB *gorm.DB
)

func connectMysql(c *config.Mysql) error {
	newDb, err := ConnectMysqlSlf(c.Username, c.Password, c.Host, c.Port, c.DatabaseName)
	if err != nil {
		return err
	}
	logger.Infof("init mysql success: %v:%v", c.Host, c.Port)
	MysqlDB = newDb
	return nil
}
func ConnectMysqlSlf(userName string, password string, host string, port string, databaseName string) (*gorm.DB, error) {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	logger.Debugf("databaseName: %s", databaseName)
	dbConn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", userName, password, host, port, databaseName)
	newDb, err := gorm.Open(mysql.Open(dbConn), &gorm.Config{})
	if err != nil {
		logger.Errorf("init mysql err " + err.Error())
		return nil, err
	}
	return newDb, err
}

func MysqlGet(table string, key string, value interface{}, dst interface{}) (int64, error) {
	return MysqlGetSlf(MysqlDB, table, key, value, dst)
}
func MysqlGetSlf(slf *gorm.DB, table string, key string, value interface{}, dst interface{}) (int64, error) {
	query := slf.Table(table).Where(key+" = ?", value)
	var cnt int64
	err := query.Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	if cnt <= 0 {
		return 0, nil
	}

	err = query.Scan(dst).Error
	if err != nil {
		return 0, err
	}

	return cnt, nil
}

func MysqlGetWhereForPage(table string, page_count int, page_size int, dst interface{}, query interface{}, args ...interface{}) (int64, error) {
	return MysqlGetWhereForPageSlf(MysqlDB, table, page_count, page_size, dst, query, args...)
}
func MysqlGetWhereForPageSlf(slf *gorm.DB, table string, page_count int, page_size int, dst interface{}, query interface{}, args ...interface{}) (int64, error) {
	var cnt int64
	err := slf.Table(table).Where(query, args...).Count(&cnt).Scopes(Paginate(page_count, page_size)).Find(dst).Error
	if err != nil {
		return 0, err
	}
	if cnt <= 0 {
		return 0, nil
	}

	return cnt, nil
}

func MysqlGetWhereForPageOrder(table string, order string, page_count int, page_size int, dst interface{}, query interface{}, args ...interface{}) (int64, error) {
	return MysqlGetWhereForPageOrderSlf(MysqlDB, table, order, page_count, page_size, dst, query, args...)
}
func MysqlGetWhereForPageOrderSlf(slf *gorm.DB, table string, order string, page_count int, page_size int, dst interface{}, query interface{}, args ...interface{}) (int64, error) {
	var cnt int64
	err := slf.Table(table).Where(query, args...).Count(&cnt).Order(order).Scopes(Paginate(page_count, page_size)).Find(dst).Error
	if err != nil {
		return 0, err
	}
	if cnt <= 0 {
		return 0, nil
	}

	return cnt, nil
}

func MysqlGetCountByTwoOp(table string, key string, value interface{}, key1 string, value1 interface{}) (int64, error) {
	return MysqlGetCountByTwoOpSlf(MysqlDB, table, key, value, key1, value1)
}
func MysqlGetCountByTwoOpSlf(slf *gorm.DB, table string, key string, value interface{}, key1 string, value1 interface{}) (int64, error) {
	var cnt int64
	err := slf.Table(table).Where(key+" = ? and "+key1+" = ?", value, value1).Count(&cnt).Error
	if err != nil {
		return 0, err
	}

	return cnt, nil
}

func MysqlGetTable(table string, dst interface{}) (int64, error) {
	return MysqlGetTableSlf(MysqlDB, table, dst)
}
func MysqlGetTableSlf(slf *gorm.DB, table string, dst interface{}) (int64, error) {
	query := slf.Table(table)
	var cnt int64
	err := query.Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	if cnt <= 0 {
		return 0, nil
	}

	err = query.Scan(dst).Error
	if err != nil {
		return 0, err
	}

	return cnt, nil
}

func MysqlGetForPage(table string, key string, value interface{}, page_count int, page_size int, dst interface{}, sort string) (int64, error) {
	return MysqlGetForPageSlf(MysqlDB, table, key, value, page_count, page_size, dst, sort)
}
func MysqlGetForPageSlf(slf *gorm.DB, table string, key string, value interface{}, page_count int, page_size int, dst interface{}, sort string) (int64, error) {
	var cnt int64
	err := slf.Table(table).Where(key+" = ?", value).Count(&cnt).Order(sort + " desc").Scopes(Paginate(page_count, page_size)).Find(dst).Error
	if err != nil {
		return 0, err
	}
	if cnt <= 0 {
		return 0, nil
	}

	return cnt, nil
}

func MysqlGetForPageByTwoOp(table string, key string, value interface{}, key1 string, value1 interface{}, page_count int, page_size int, dst interface{}, sort string) (int64, error) {
	return MysqlGetForPageByTwoOpSlf(MysqlDB, table, key, value, key1, value1, page_count, page_size, dst, sort)
}
func MysqlGetForPageByTwoOpSlf(slf *gorm.DB, table string, key string, value interface{}, key1 string, value1 interface{}, page_count int, page_size int, dst interface{}, sort string) (int64, error) {
	var cnt int64
	err := slf.Table(table).Where(key+" = ? and "+key1+" = ?", value, value1).Count(&cnt).Order(sort + " desc").Scopes(Paginate(page_count, page_size)).Find(dst).Error
	if err != nil {
		return 0, err
	}
	if cnt <= 0 {
		return 0, nil
	}

	return cnt, nil
}

func MysqlCreate(table string, data interface{}) error {
	return MysqlCreateSlf(MysqlDB, table, data)
}
func MysqlCreateSlf(slf *gorm.DB, table string, data interface{}) error {
	return slf.Table(table).Create(data).Error
}

func MysqlSet(table string, data interface{}) error {
	return MysqlSetSlf(MysqlDB, table, data)
}
func MysqlSetSlf(slf *gorm.DB, table string, data interface{}) error {
	return slf.Table(table).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(data).Error
}

func MysqlSetCol(table string, key string, keyValue interface{}, col string, colValue interface{}) error {
	return MysqlSetColSlf(MysqlDB, table, key, keyValue, col, colValue)
}
func MysqlSetColSlf(slf *gorm.DB, table string, key string, keyValue interface{}, col string, colValue interface{}) error {
	return MysqlDB.Table(table).Where(key+" = ?", keyValue).UpdateColumn(col, colValue).Error
}

func MysqlUpdates(tabel string, key string, keyValue any, data any) error {
	return MysqlUpdatesSlf(MysqlDB, tabel, key, keyValue, data)
}
func MysqlUpdatesSlf(slf *gorm.DB, tabel string, key string, keyValue any, data any) error {
	return slf.Table(tabel).Where(key+"=?", keyValue).Updates(data).Error
}

func MysqlIncCol(table string, key string, keyValue interface{}, col string, incValue int) error {
	return MysqlIncColSlf(MysqlDB, table, key, keyValue, col, incValue)
}
func MysqlIncColSlf(slf *gorm.DB, table string, key string, keyValue interface{}, col string, incValue int) error {
	return slf.Table(table).Where(key+" = ?", keyValue).Update(col, gorm.Expr(col+"+?", incValue)).Error
}

func MysqlExec(sql string, values ...interface{}) error {
	return MysqlExecSlf(MysqlDB, sql, values...)
}
func MysqlExecSlf(slf *gorm.DB, sql string, values ...interface{}) error {
	return slf.Exec(sql, values...).Error
}

func MysqlExecFind(dst interface{}, sql string, values ...interface{}) error {
	return MysqlExecFindSlf(MysqlDB, dst, sql, values...)
}
func MysqlExecFindSlf(slf *gorm.DB, dst interface{}, sql string, values ...interface{}) error {
	return slf.Raw(sql, values...).Scan(dst).Error
}

func Paginate(page int, page_size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case page_size > 100:
			page_size = 100
		case page_size <= 0:
			page_size = 10
		}
		offset := (page - 1) * page_size
		return db.Offset(offset).Limit(page_size)
	}
}
