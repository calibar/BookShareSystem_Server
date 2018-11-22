package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type BookTransaction struct {
	Id               int       `orm:"column(book_id);auto"`
	BookName         string    `orm:"column(book_name);size(64)"`
	BookAuthor       string    `orm:"column(book_author);size(64);null"`
	BookDescription  string    `orm:"column(book_description);null"`
	BookCover        string    `orm:"column(book_cover);null"`
	BookOwner        string    `orm:"column(book_owner);size(32);null"`
	BookBorrower     string    `orm:"column(book_borrower);size(32);null"`
	Campus           string    `orm:"column(campus);size(32)"`
	PostExpiration   time.Time `orm:"column(post_expiration);type(date)"`
	ExpectReturnTime time.Time `orm:"column(expect_return_time);type(date)"`
	ActualReturnTime time.Time `orm:"column(actual_return_time);type(date);null"`
	PostDate         time.Time `orm:"column(post_date);type(timestamp);auto_now"`
	OwnerRating      int       `orm:"column(owner_rating);null"`
	BorrowerRating   int       `orm:"column(borrower_rating);null"`
	OwnerComment     string    `orm:"column(owner_comment);null"`
	BorrowerComment  string    `orm:"column(borrower_comment);null"`
	BookStatus       string    `orm:"column(book_status);size(32)"`
}

func (t *BookTransaction) TableName() string {
	return "book_transaction"
}

func init() {
	orm.RegisterModel(new(BookTransaction))
}

// AddBookTransaction insert a new BookTransaction into database and returns
// last inserted Id on success.
func AddBookTransaction(m *BookTransaction) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBookTransactionById retrieves BookTransaction by Id. Returns error if
// Id doesn't exist
func GetBookTransactionById(id int) (v *BookTransaction, err error) {
	o := orm.NewOrm()
	v = &BookTransaction{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBookTransaction retrieves all BookTransaction matches certain condition. Returns empty list if
// no records exist
func GetAllBookTransaction(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BookTransaction))
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

	var l []BookTransaction
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

// UpdateBookTransaction updates BookTransaction by Id and returns error if
// the record to be updated doesn't exist
func UpdateBookTransactionById(m *BookTransaction) (err error) {
	o := orm.NewOrm()
	v := BookTransaction{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBookTransaction deletes BookTransaction by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBookTransaction(id int) (err error) {
	o := orm.NewOrm()
	v := BookTransaction{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BookTransaction{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
