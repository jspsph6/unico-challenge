package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"log"
)

type Repository struct {
	db *sql.DB
}

type Fair struct {
	ID                int            `json:"id"`                 // id
	ExternalID        int            `json:"external_id"`        // external_id
	Setcens           int            `json:"setcens"`            // setcens
	FairName          string         `json:"fair_name"`          // fair_name
	Register          string         `json:"register"`           // register
	CodeDistrict      int            `json:"code_district"`      // code_district
	CodeSubprefecture int            `json:"code_subprefecture"` // code_subprefecture
	Subprefecture     string         `json:"subprefecture"`      // subprefecture
	Areap             int            `json:"areap"`              // areap
	DistrictName      string         `json:"district_name"`      // district_name
	Region5           string         `json:"region5"`            // region5
	Region8           string         `json:"region8"`            // region8
	Street            string         `json:"street"`             // street
	Number            sql.NullString `json:"number"`             // number
	Neighborhood      string         `json:"neighborhood"`       // neighborhood
	Reference         sql.NullString `json:"reference"`          // reference
	Latitude          int            `json:"latitude"`           // latitude
	Longitude         int            `json:"longitude"`          // longitude
}

type Filter struct {
	DistrictName string
	Region5      string
	FairName     string
	Neighborhood string
}

func NewRepository() (*Repository, error) {
	constr := fmt.Sprintf("%s:%s@%s:%s/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	db, err := sql.Open("mysql", constr)
	if err != nil {
		log.Print(err.Error())
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	defer db.Close()

	return &Repository{db}, nil
}

func (r *Repository) Close() {
	r.db.Close()
}

func (r *Repository) Insert(f *Fair) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	const sqlstr = `INSERT INTO fair.fair (` +
		`external_id, setcens, fair_name, register, code_district, code_subprefecture, subprefecture, areap, district_name, region5, region8, street, number, neighborhood, reference, latitude, longitude` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	stmt, err := r.db.PrepareContext(ctx, sqlstr)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, sqlstr, f.ExternalID, f.Setcens, f.FairName, f.Register, f.CodeDistrict, f.CodeSubprefecture, f.Subprefecture, f.Areap, f.DistrictName, f.Region5, f.Region8, f.Street, f.Number, f.Neighborhood, f.Reference, f.Latitude, f.Longitude)
	return err
}

func (r *Repository) Update(f *Fair) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	const sqlstr = `UPDATE fair.fair SET ` +
		`external_id = ?, setcens = ?, fair_name = ?, register = ?, code_district = ?, code_subprefecture = ?, subprefecture = ?, areap = ?, district_name = ?, region5 = ?, region8 = ?, street = ?, number = ?, neighborhood = ?, reference = ?, latitude = ?, longitude = ? ` +
		`WHERE id = ?`
	stmt, err := r.db.PrepareContext(ctx, sqlstr)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, sqlstr, f.ExternalID, f.Setcens, f.FairName, f.Register, f.CodeDistrict, f.CodeSubprefecture, f.Subprefecture, f.Areap, f.DistrictName, f.Region5, f.Region8, f.Street, f.Number, f.Neighborhood, f.Reference, f.Latitude, f.Longitude, f.ID)
	return err
}

func (r *Repository) Delete(register string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	const sqlstr = `DELETE FROM fair.fair ` +
		`WHERE register = ?`
	stmt, err := r.db.PrepareContext(ctx, sqlstr)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, register)
	return err
}

func (r *Repository) FindByRegister(register string) (*Fair, error) {
	f := Fair{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	const sqlstr = `SELECT ` +
		`id, external_id, setcens, fair_name, register, code_district, code_subprefecture, subprefecture, areap, district_name, region5, region8, street, number, neighborhood, reference, latitude, longitude ` +
		`FROM fair.fair ` +
		`WHERE register = ?`
	err := r.db.QueryRowContext(ctx, sqlstr, register).Scan(&f.ID, &f.ExternalID, &f.Setcens, &f.FairName, &f.Register, &f.CodeDistrict, &f.CodeSubprefecture, &f.Subprefecture, &f.Areap, &f.DistrictName, &f.Region5, &f.Region8, &f.Street, &f.Number, &f.Neighborhood, &f.Reference, &f.Latitude, &f.Longitude)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return &f, nil
}

func (r *Repository) FindByFilter(filter *Filter) ([]*Fair, error) {
	fairs := []*Fair{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sqlstr := `SELECT ` +
		`id, external_id, setcens, fair_name, register, code_district, code_subprefecture, subprefecture, areap, district_name, region5, region8, street, number, neighborhood, reference, latitude, longitude ` +
		`FROM fair.fair WHERE `
	clauses := []string{}
	filters := []interface{}{}

	var clause string
	if filter.DistrictName != "" {
		clause = " district_name = ? "
		clauses = append(clauses, clause)
		filters = append(filters, filter.DistrictName)
	}
	if filter.FairName != "" {
		clause = " fair_name = ? "
		clauses = append(clauses, clause)
		filters = append(filters, filter.FairName)
	}
	if filter.Neighborhood != "" {
		clause = " neighborhood = ? "
		clauses = append(clauses, clause)
		filters = append(filters, filter.Neighborhood)
	}
	if filter.Region5 != "" {
		clause = " region5 = ? "
		clauses = append(clauses, clause)
		filters = append(filters, filter.Region5)
	}
	sqlstr = sqlstr + strings.Join(clauses, "AND")
	rows, err := r.db.QueryContext(ctx, sqlstr, filters)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		f := &Fair{}
		err = rows.Scan(&f.ID, &f.ExternalID, &f.Setcens, &f.FairName, &f.Register, &f.CodeDistrict, &f.CodeSubprefecture, &f.Subprefecture, &f.Areap, &f.DistrictName, &f.Region5, &f.Region8, &f.Street, &f.Number, &f.Neighborhood, &f.Reference, &f.Latitude, &f.Longitude)

		if err != nil {
			log.Print(err.Error())
			return nil, err
		}
		fairs = append(fairs, f)
	}

	return fairs, nil
}
