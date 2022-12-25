package dao

func Crate(name string) (err error) {
	_, err = DB.Exec("create table " + name + "(goods_name varchar(20) not null, goods_num int not null, id bigint auto_increment primary key)")
	return err
}
