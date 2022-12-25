package dao

import "test/program/2/model"

func Crate(name string) (err error) {
	_, err = DB.Exec("create table ? (goods_name varchar(20) not null, goods_num int not null, id bigint auto_increment primary key)", name)
	return err
}

func StoreManage_in(name string, s model.Storehouse) (err error) {
	_, err = DB.Exec("insert into ? (goods_name,goods_num) values(?,?)", name, s.GoodsName, s.GoodsNum)
	return err
}

func StoreManage_out(storeName string, name string) (err error) {
	_, err = DB.Exec("delete from ? where goods_name=?", storeName, name)
	return err
}
