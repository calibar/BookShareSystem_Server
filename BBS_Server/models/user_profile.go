package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type UserProfile struct {
	Id           int       `orm:"column(pid);auto"`
	Username     string    `orm:"column(username);size(32)"`
	Nickname     string    `orm:"column(nickname);size(32);null"`
	Email        string    `orm:"column(email);size(32);null"`
	Campus       string    `orm:"column(campus);size(32);null"`
	StudentId    string    `orm:"column(student_id);size(16);null"`
	Avatar       string    `orm:"column(avatar);null"`
	LendCount    int       `orm:"column(lend_count)"`
	BorrowCount  int       `orm:"column(borrow_count)"`
	PostCount    int       `orm:"column(post_count)"`
	RequestCount int       `orm:"column(request_count)"`
	Score        int       `orm:"column(score)"`
	SignupDate   time.Time `orm:"column(signup_date);type(timestamp);auto_now"`
	Badge        string    `orm:"column(badge);size(32);null"`
}
type UPrank struct {
	Username     string    `orm:"column(username);size(32)"`
	Score        int       `orm:"column(score)"`
	Rank int				`orm:"column(rank)"`
}
type UserProfileHasRank struct {
	Id           int       `orm:"column(pid);auto"`
	Username     string    `orm:"column(username);size(32)"`
	Nickname     string    `orm:"column(nickname);size(32);null"`
	Email        string    `orm:"column(email);size(32);null"`
	Campus       string    `orm:"column(campus);size(32);null"`
	StudentId    string    `orm:"column(student_id);size(16);null"`
	Avatar       string    `orm:"column(avatar);null"`
	LendCount    int       `orm:"column(lend_count)"`
	BorrowCount  int       `orm:"column(borrow_count)"`
	PostCount    int       `orm:"column(post_count)"`
	RequestCount int       `orm:"column(request_count)"`
	Score        int       `orm:"column(score)"`
	SignupDate   time.Time `orm:"column(signup_date);type(timestamp);auto_now"`
	Badge        string    `orm:"column(badge);size(32);null"`
	Rank int				`orm:"column(rank)"`
}

func (t *UserProfile) TableName() string {
	return "user_profile"
}

func init() {
	orm.RegisterModel(new(UserProfile))
}

// AddUserProfile insert a new UserProfile into database and returns
// last inserted Id on success.
func AddUserProfile(m *UserProfile) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserProfileById retrieves UserProfile by Id. Returns error if
// Id doesn't exist
func GetUserProfileById(id int) (v *UserProfile, err error) {
	o := orm.NewOrm()
	v = &UserProfile{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetUserProfileByUsername(username string) (v *UserProfile, err error) {
	o :=orm.NewOrm()
	userProfile:=UserProfile{Username:username}
	err=o.Read(&userProfile,"Username")
	if err==orm.ErrNoRows{
		return v,nil
	}else{
		fmt.Println(err)
	}
	return nil,err
}
// GetAllUserProfile retrieves all UserProfile matches certain condition. Returns empty list if
// no records exist
func GetAllUserProfile(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserProfile))
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

	var l []UserProfile
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
// Get single user profile
func GetSingleUserProfileByUsername(username string) (up UserProfileHasRank,err error) {
	o:=orm.NewOrm()
	var q orm.RawSeter
	q=o.Raw("select * from (" +
		"select *, case " +
		"when @prev_value=score then @rank_count " +
		"when @prev_value:=score then @rank_count:=@rank_count+1 " +
		"end as rank " +
		"from user_profile as up, (select @prev_value:=null,@rank_count:=0) as r " +
		"order by score desc) as s " +
		"where username= ?",username)
	err=q.QueryRow(&up)
	if err==nil{
		fmt.Println(up)
	}else {
		fmt.Println(err)
	}
	return up,err
}
// UpdateUserProfile updates UserProfile by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserProfileById(m *UserProfile) (err error) {
	o := orm.NewOrm()
	v := UserProfile{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUserProfile deletes UserProfile by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserProfile(id int) (err error) {
	o := orm.NewOrm()
	v := UserProfile{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UserProfile{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
