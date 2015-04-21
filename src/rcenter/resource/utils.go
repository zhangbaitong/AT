package resource

import (
	"common"
	"database/sql"
	"fmt"
	"time"
)

const (
	SELECT_ALL string = "select * from resource_tab"
	INSERT     string = "insert into resource_tab (res_id,res_name,owner_acid,operator_acid,status,create_time) values (?,?,?,?,?,?)"
	UPDATE     string = "update resource_tab  set res_name=?,owner_acid=?,operator_acid=?,status=?,create_time=? where res_id=?"
	DELETE     string = "delete from resource_tab  where res_id=?"
)

func Query(sqlstr string) (ress []Resource, err error) {
	rows, err := common.GetDB().Query(sqlstr)
	if err != nil {
		fmt.Println(err)
	} else {
		ress = getResWithRows(rows)
	}
	return ress, err
}

func Insert(res_name string, owner_acid int, operator_acid int) {
	tx, err := common.GetDB().Begin()
	if err != nil {
		fmt.Println(err)
	}
	stmt, err := tx.Prepare(INSERT)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(common.GetUID(), res_name, owner_acid, operator_acid, 0, time.Now().Unix())

	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()
}

func Update(uuidstr string, res_name string, owner_acid int, operator_acid int, status int) {
	tx, err := common.GetDB().Begin()
	if err != nil {
		fmt.Println(err)
	}
	stmt, err := tx.Prepare(UPDATE)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(res_name, owner_acid, operator_acid, status, time.Now().Unix(), uuidstr)

	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()
}

func Delete(uuidstr string) {
	tx, err := common.GetDB().Begin()
	if err != nil {
		fmt.Println(err)
	}
	stmt, err := tx.Prepare(DELETE)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(uuidstr)

	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()
}

//=========================================================================================================

func exec(sqlstr string, uuidstr string, res_name string, owner_acid int, operator_acid int, status int, create_time int64) {

}

func getResWithRows(rows *sql.Rows) (ress []Resource) {
	ress = make([]Resource, 0)
	for i := 0; rows.Next(); i++ {
		var res Resource
		rows.Scan(&res.res_id, &res.res_name, &res.owner_acid, &res.operator_acid, &res.status, &res.create_time)
		ress = append(ress, res)
	}
	return ress
}
