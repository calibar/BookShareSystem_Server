package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Campus struct {
	Id         int    `orm:"column(campus_id);auto"`
	CampusName string `orm:"column(campus_name);size(100)"`
	CampusEmailStandard string `orm:"column(campus_email_standard);size(100)"`
	AlterEmailStandard string `orm:"column(alter_email_standard);size(100)"`
}

func (t *Campus) TableName() string {
	return "campus"
}

func init() {
	orm.RegisterModel(new(Campus))
}

// AddCampus insert a new Campus into database and returns
// last inserted Id on success.
func AddCampus(m *Campus) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
func CheckCampusEmail(emailStandard string)bool  {
	o:=orm.NewOrm()
	var campus []Campus
	num,err:=o.Raw("select * from campus where campus_email_standard=? or alter_email_standard=?",emailStandard,emailStandard).QueryRows(&campus)
	fmt.Println(num)
	if err==nil{
		if num>0{
			return true
		}
	}else {
		fmt.Println(err)
	}

	return false
}
// GetCampusById retrieves Campus by Id. Returns error if
// Id doesn't exist
func GetCampusById(id int) (v *Campus, err error) {
	o := orm.NewOrm()
	v = &Campus{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCampus retrieves all Campus matches certain condition. Returns empty list if
// no records exist
func GetAllCampus(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Campus))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Campus
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateCampus updates Campus by Id and returns error if
// the record to be updated doesn't exist
func UpdateCampusById(m *Campus) (err error) {
	o := orm.NewOrm()
	v := Campus{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCampus deletes Campus by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCampus(id int) (err error) {
	o := orm.NewOrm()
	v := Campus{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Campus{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
