package main

import (
    "fmt"
    "database/sql"
    "log"
    "net/http"
    "text/template"

    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    AcNo int
    Name string
    Mobile int
    Email string
    Balance int
	Password string
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "pwd"
    dbName := "Kartheek"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}
var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM user ORDER BY acNo")
    if err != nil {
        panic(err.Error())
    }
    user := User{}
    for selDB.Next() {
        var acNo, balance, mobile int
        var name, email, password string
        err = selDB.Scan(&acNo, &name, &mobile, &email, &balance, &password)
        if err != nil {
            panic(err.Error())
        }
        user.AcNo = acNo
        user.Name = name
        user.Email = email
        user.Mobile = mobile
        user.Balance = balance
        user.Password = password
    }
    tmpl.ExecuteTemplate(w, "Index", user)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM user WHERE acNo=?", nId)
    if err != nil {
        panic(err.Error())
    }
    user := User{}
    for selDB.Next() {
        var acNo, mobile, balance int
        var name, email, password string
        err = selDB.Scan(&acNo, &name, &mobile, &email, &balance, &password)
        if err != nil {
            panic(err.Error())
        }
        user.AcNo = acNo
        user.Name = name
        user.Email = email
        user.Mobile = mobile
        user.Balance = balance
        user.Password = password
    }
    tmpl.ExecuteTemplate(w, "Show", user)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        acNo := r.FormValue("acNo") 
        name := r.FormValue("name") 
        //fmt.Println(r.FormValue("id"))
        email := r.FormValue("email")
        mobile := r.FormValue("mobile")
        balance := r.FormValue("balance")
        password := r.FormValue("password")
        
        insForm, err := db.Prepare("INSERT INTO user(acNo, name, mobile, email, balance, password) VALUES(?,?,?,?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(acNo, name, mobile, email, balance, password)
        log.Println("INSERT: Name: " + name)
    }
    defer db.Close()
    http.Redirect(w, r, "/index", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "Edit", nil)
}
func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        mobile := r.FormValue("mobile")
        email := r.FormValue("email")        
        acNo := r.FormValue("acNo")
        insForm, err := db.Prepare("UPDATE user SET name=?, mobile=?, email=? WHERE acNo=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, mobile, email, acNo)
        log.Println("UPDATE: Name: " + name)
    }
    defer db.Close()
    http.Redirect(w, r, "/index", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    user := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM user WHERE acNo=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(user)
    log.Println("DELETE")
    http.Redirect(w, r, "/index", 301)
    defer db.Close()
}

func Login(w http.ResponseWriter, r *http.Request){
    tmpl.ExecuteTemplate(w, "Login", nil)
    fmt.Println(r.Method)
}
func Checkout(w http.ResponseWriter, r *http.Request){
    db := dbConn()
    if r.Method == "POST" {
        acNo := r.FormValue("acNo") 
        password := r.FormValue("password")
        var isAuthenticated bool
        err := db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM user WHERE acNo=? AND password=?",acNo,password).Scan(&isAuthenticated)
        if err != nil {
            log.Fatal(err)
        } 
        if isAuthenticated{
            http.Redirect(w, r, "/index", 301)
        }else{
            //alert("Incorrect!")
            fmt.Fprintf(w,"<html><h4>Incorrect Credentials! Please <a href='/login'>Try Again!</a></h4></html>")
            //http.Redirect(w, r, "/login", 301)
            //tmpl.Execute(w, "<h4><br /><br /><br />Incorrect credentials! </h4>")
        }
    }
}
func Signup(w http.ResponseWriter, r *http.Request){
    tmpl.ExecuteTemplate(w, "Signup", nil)
}
func Checkin(w http.ResponseWriter, r *http.Request){
    db := dbConn()
    if r.Method == "POST" {
        acNo := r.FormValue("acNo")
        name := r.FormValue("name")
        mobile := r.FormValue("mobile")
        email := r.FormValue("email")
        balance := 0 
        password := r.FormValue("password")
        confirmPassword := r.FormValue("confirmPassword")
        var isAuthenticated bool
        err := db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM user WHERE acNo=?",acNo).Scan(&isAuthenticated)
        if err != nil {
            log.Fatal(err)
        } 
        if isAuthenticated{
            fmt.Fprintf(w,"<html><h4>Account Number Already Exists! Please <a href='/signup'>Try Again!</a></h4></html>")
            //http.Redirect(w, r, "/index", 301)
        }else if password == confirmPassword{
            insForm,_ := db.Prepare("INSERT INTO user(acNo, name, mobile, email, balance, password) VALUES(?,?,?,?,?,?)")
            insForm.Exec(acNo, name, mobile, email, balance, password)
            http.Redirect(w, r, "/login", 301)
            defer db.Close()  
        }else{
            fmt.Fprintf(w,"<html><h4>Passwords Didn't Match! Please <a href='/signup'>Try Again!</a></h4></html>")
            //http.Redirect(w, r, "/signup", 301)
        }
    }
}
func Home(w http.ResponseWriter, r *http.Request){
    tmpl.ExecuteTemplate(w, "Home", nil)
}

func main() {
    log.Println("Server started on: http://localhost:8080")
    http.HandleFunc("/", Home)
    http.HandleFunc("/signup", Signup)
    http.HandleFunc("/checkin", Checkin)
    http.HandleFunc("/login", Login)
    http.HandleFunc("/checkout", Checkout)
    http.HandleFunc("/index", Index)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":8080", nil)
}