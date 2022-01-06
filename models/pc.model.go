package models

import (
	"server_pc/db"
	"time"
)

type Pc struct {
	Id         int       `json:"id"`
	Nama       string    `json:"nama"`
	LastActive time.Time `json:"last_active"`
	Perintah   string    `json:"perintah"`
	Lokasi     string    `json:"lokasi"`
}

func FetchAllPc() (Response, error) {
	var obj Pc
	var arrObj []Pc
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM pc"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&obj.Id, &obj.Nama, &obj.LastActive, &obj.Perintah, &obj.Lokasi)
		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = 200
	res.Message = "Sukses"
	res.Data = arrObj

	return res, nil

}

func InsertPc(nama string, lokasi string) (Response, error) {
	var res Response
	// var last_active time.Time
	con := db.CreateCon()

	// last_active = time.Now()
	sqlStatement := "INSERT pc (nama, lokasi) VALUES (?, ?)"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(nama, lokasi)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = 200
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdatePc(id int, nama string, lokasi string) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "UPDATE pc SET nama=?, lokasi=? WHERE id=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(nama, lokasi, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = 200
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func UpdateLastActive(id int) (Response, error) {

	var res Response
	var perintah string
	con := db.CreateCon()

	last_update := time.Now()

	// select perintah pc
	sqlStatement := "SELECT perintah FROM pc WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&perintah)
	if err != nil {
		return res, nil
	}

	// update pc
	sqlStatement = "UPDATE pc SET last_active = ?, perintah = '' WHERE id = ?"
	stmt, err = con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(last_update, id)
	if err != nil {
		return res, nil
	}

	res.Status = 200
	res.Message = "Sukses"
	res.Data = map[string]interface{}{
		"id":          id,
		"last_active": last_update,
		"perintah":    perintah,
	}

	return res, nil
}

func DeletePc(id int) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "DELETE FROM pc WHERE id=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = 200
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func RestartAllPc() (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "UPDATE pc SET perintah = 'restart'"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	defer stmt.Close()

	result, err := stmt.Exec()
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = 200
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func ShutdownAllPc() (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "UPDATE pc SET perintah = 'shutdown'"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	defer stmt.Close()

	result, err := stmt.Exec()
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = 200
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
