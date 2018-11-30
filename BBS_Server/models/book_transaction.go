package models

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
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
	PostExpiration   string `orm:"column(post_expiration);type(date)"`
	ExpectReturnTime string `orm:"column(expect_return_time);type(date)"`
	ActualReturnTime string `orm:"column(actual_return_time);type(date);null"`
	PostDate         time.Time `orm:"column(post_date);type(timestamp);auto_now"`
	OwnerRating      int       `orm:"column(owner_rating);null"`
	BorrowerRating   int       `orm:"column(borrower_rating);null"`
	OwnerComment     string    `orm:"column(owner_comment);null"`
	BorrowerComment  string    `orm:"column(borrower_comment);null"`
	BookStatus       string    `orm:"column(book_status);size(32)"`
}
type BookTransHasApplicants struct {
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
	Applicants       []ApplicantList
}
type days struct {
	days string
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
	if m.BookStatus=="request"{
		var requester UserProfile
		err=o.Raw("select * from user_profile where username=?",m.BookBorrower).QueryRow(&requester)
		requester.RequestCount+=1
		fmt.Println(requester.RequestCount)
		if num, err := o.Update(&requester); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}else if m.BookStatus=="post"{
		var poster UserProfile
		err=o.Raw("select * from user_profile where username=?",m.BookOwner).QueryRow(&poster)
		poster.PostCount+=1
		if num, err := o.Update(&poster); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
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
		if m.ExpectReturnTime!=""{
			v.ExpectReturnTime=m.ExpectReturnTime
		}
		if m.BookBorrower!=""{
			v.BookBorrower=m.BookBorrower
		}
		if m.ActualReturnTime!=""{
			v.ActualReturnTime=m.BookBorrower
		}
		if m.BookAuthor!=""{
			v.BookAuthor=m.BookAuthor
		}
		if m.BookCover!="" {
			v.BookCover=m.BookCover
		}
		if m.BookDescription!=""{
			v.BookDescription=m.BookDescription
		}
		if m.BookName!="" {
			v.BookName=m.BookName
		}
		if m.BookOwner!="" {
			v.BookOwner=m.BookOwner
		}
		if m.BorrowerComment!="" {
			v.BorrowerComment=m.BorrowerComment
			Rescore(v.BookBorrower,2)
		}
		if m.BorrowerRating!=0 {
			v.BorrowerRating=m.BorrowerRating
			Rescore(v.BookBorrower,1)
		}
		if m.OwnerComment!="" {
			v.OwnerComment=m.OwnerComment
			Rescore(v.BookOwner,2)
		}
		if m.OwnerRating!=0 {
			v.OwnerRating=m.OwnerRating
			Rescore(v.BookOwner,1)
		}
		if m.PostExpiration!="" {
			v.PostExpiration=m.PostExpiration
		}
		if m.Campus!=""{
			v.Campus=m.Campus
		}
		if m.BookStatus!=""{
			if m.BookStatus!=v.BookStatus{
				if m.BookStatus=="borrowed"{
					v.BookStatus=m.BookStatus
					if num, err = o.Update(&v); err == nil {
						fmt.Println("Number of records updated in database:", num)
					}
					var borrower UserProfile
					err=o.Raw("select * from user_profile where username=?",v.BookBorrower).QueryRow(&borrower)
					borrower.BorrowCount+=1
					fmt.Println(borrower)
					if num, err = o.Update(&borrower); err == nil {
						fmt.Println("Number of records updated in database:", num)
					}
					/*var ownerUser UserProfile
					err=o.Raw("select * from user_profile where username=?",v.BookOwner).QueryRow(&ownerUser)
					ownerUser.Score=ownerUser.Score+5;
					if _,err :=o.Update(&ownerUser);err==nil{
						fmt.Println(ownerUser.Score)
					}*/
					RescoreLendRecore(v.BookOwner,8,v.BookName)
				}else if m.BookStatus=="returned"{
					v.BookStatus=m.BookStatus
					if num, err = o.Update(&v); err == nil {
						fmt.Println("Number of records updated in database:", num)
					}
					/*var userPunish UserProfile
					err=o.Raw("select * from user_profile where username=?",v.BookBorrower).QueryRow(&userPunish)
					userPunish.Score=userPunish.Score+1;
					if _,err :=o.Update(&userPunish);err==nil{
						fmt.Println(userPunish.Score)
					}*/
					Rescore(v.BookBorrower,1)
				}else{
					v.BookStatus=m.BookStatus
					if num, err = o.Update(&v); err == nil {
						fmt.Println("Number of records updated in database:", num)
					}
				}

			}
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
// Check if expect return time has already reached
func CheckExpiration()  {
	for true {
		o :=orm.NewOrm()
		var PostsExpired []BookTransaction
		var ReturnExpired []BookTransaction
		current_time := time.Now().UTC()
		current_time_str:=current_time.Format("2006-01-02")
		num,err:=o.Raw("select * from book_transaction where post_expiration < ?",current_time_str ).QueryRows(&PostsExpired)
		if err==nil{
			if num!=0{
				fmt.Println(num)
				fmt.Println(PostsExpired)
				for _,v :=range PostsExpired{
					var postUser UserProfile
					err=o.Raw("select * from user_profile where username=?",v.BookBorrower).QueryRow(&postUser)
					postUser.Score=postUser.Score+1;
					if _,err :=o.Update(&postUser);err==nil{
						fmt.Println(postUser.Score)
					}

					DeleteBookTransaction(v.Id)
				}
			}else{
				fmt.Println("no post expired")
			}

		}else{
			fmt.Println(err)
		}
		num1,err:=o.Raw("select * from book_transaction where book_status='borrowed' and expect_return_time < ?",current_time_str).QueryRows(&ReturnExpired)
		if err==nil{
			if num1!=0{
				for _,v :=range ReturnExpired {
					/*var userPunish UserProfile
					err=o.Raw("select * from user_profile where username=?",v.BookBorrower).QueryRow(&userPunish)
					userPunish.Score=userPunish.Score-1;
					if _,err :=o.Update(&userPunish);err==nil{
						fmt.Println(userPunish.Score)
					}*/
					layout := "2006-01-02"
					expectReturnTime,err:=time.Parse(layout,v.ExpectReturnTime)
					if err != nil {
						fmt.Println(err)
					}else {
						fmt.Println(expectReturnTime)
						subdays := current_time.Sub(expectReturnTime).Hours()/24
						subsocre:=math.Floor(subdays)
						fmt.Println(subsocre)
						Rescore(v.BookBorrower,-int(subsocre))
					}

				}

			}else {
				fmt.Println("no return expired")
			}

		}else{
			fmt.Println(err)
		}
		time.Sleep(24*time.Hour)
	}


}
func Rescore(username string, scoreChange int)  {
	o :=orm.NewOrm()
	var user UserProfile
	err:=o.Raw("select * from user_profile where username=?",username).QueryRow(&user)
	usernum,err:=o.Raw("select * from user_profile").QueryRows()
	if usernum >10{
		user.Score=user.Score+int(float64(scoreChange)/math.Log10(float64(usernum)))
	}else {
		user.Score=user.Score+scoreChange
	}

	if _,err =o.Update(&user);err==nil{
		fmt.Println(user.Username+"score is"+strconv.Itoa(user.Score))
	}else {
		fmt.Println(err)
	}
}
func RescoreLendRecore(username string, scoreChange int, bookname string)  {
	o :=orm.NewOrm()
	var user UserProfile
	err:=o.Raw("select * from user_profile where username=?",username).QueryRow(&user)
	var users []UserProfile
	usernum,err:=o.Raw("select * from user_profile").QueryRows(&users)
	var demands []BookTransaction
	demand,err:=o.Raw("select * from book_transaction where book_name=? and book_status=?",bookname,"request").QueryRows(&demands)
	if usernum >10{
		if demand>2{
			user.Score=user.Score+int(float64(scoreChange)*math.Log2(float64(demand))/math.Log10(float64(usernum)))
		}else{
			user.Score=user.Score+int(float64(scoreChange)/math.Log10(float64(usernum)))
		}
	}else {
		if demand>2{
			user.Score=user.Score+int(float64(scoreChange)*math.Log2(float64(demand)))
		}else {
			user.Score=user.Score+scoreChange
		}

	}

	if _,err =o.Update(&user);err==nil{
		fmt.Println(user.Username+"score is"+strconv.Itoa(user.Score))
	}else {
		fmt.Println(err)
	}
}