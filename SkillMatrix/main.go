package main
import (
    "fmt"
    "time"
    //"strconv"
    "database/sql"
    "log"
    "net/http"
    "text/template"
    //"encoding/gob"

    //"github.com/gorilla/sessions"
    //"github.com/gorilla/securecookie"
    _ "github.com/go-sql-driver/mysql"
)
// User holds a users account information
type Session map[string] string
var S = Session{}
var tmpl = template.Must(template.ParseGlob("form/*"))
type Result struct{
    emp []Employee
    skill []Skill
}
type Skill struct{
    Id int64
    Name string
}
type Employee struct {
    Id    int64
    Name  string
    Mobile int64
    Email string
    YearsOfExperience int64
	ProjectName string
	Designation string
	Skillset string
	CompletedTrainings string
	ProjectAquiredSkills string
	Achievements string
    EmployeeStatus string
    Github string
    Linkedin string 
    Techgig string
    HackerRank string
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
func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
        selDB, err := db.Query("SELECT * FROM employee ORDER BY id")
        if err != nil {
            panic(err.Error())
        }
        emp := Employee{}
        res := []Employee{}
        for selDB.Next() {
            var id, yearsOfExperience, mobile int64
            var name, email, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus, github, linkedin, techgig, hackerRank string
            err = selDB.Scan(&id, &name, &mobile, &email, &yearsOfExperience, &projectName, &designation, &skillset, &completedTrainings, &projectAquiredSkills, &achievements, &employeeStatus, &github, &linkedin, &techgig, &hackerRank)
            if err != nil {
                panic(err.Error())
            }
            emp.Id = id
            emp.Name = name
            emp.Mobile = mobile
            emp.Email = email
            emp.YearsOfExperience = yearsOfExperience
            emp.ProjectName = projectName
            emp.Designation = designation
            emp.Skillset = skillset
            emp.CompletedTrainings = completedTrainings
            emp.ProjectAquiredSkills = projectAquiredSkills
            emp.Achievements = achievements
            emp.EmployeeStatus = employeeStatus
            emp.Github = github
            emp.Linkedin = linkedin
            emp.Techgig = techgig
            emp.HackerRank = hackerRank
            res = append(res, emp)
        }
        tmpl.ExecuteTemplate(w, "Index", res)
        defer db.Close()
}
func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.FormValue("id")

        selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
        if err != nil {
            panic(err.Error())
        }
        emp := Employee{}
        for selDB.Next() {
            var id, yearsOfExperience, mobile int64
            var name, email, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus, github, linkedin, techgig, hackerRank string
            err = selDB.Scan(&id, &name, &mobile, &email, &yearsOfExperience, &projectName, &designation, &skillset, &completedTrainings, &projectAquiredSkills, &achievements, &employeeStatus, &github, &linkedin, &techgig, &hackerRank)
            if err != nil {
                panic(err.Error())
            }
            emp.Id = id
            emp.Name = name
            emp.Mobile = mobile
            emp.Email = email
            emp.YearsOfExperience = yearsOfExperience
            emp.ProjectName = projectName
            emp.Designation = designation
            emp.Skillset = skillset
            emp.CompletedTrainings = completedTrainings
            emp.ProjectAquiredSkills = projectAquiredSkills
            emp.Achievements = achievements
            emp.EmployeeStatus = employeeStatus
            emp.Github = github
            emp.Linkedin = linkedin
            emp.Techgig = techgig
            emp.HackerRank = hackerRank
        }
        tmpl.ExecuteTemplate(w, "Show", emp)
        defer db.Close()
}
func New(w http.ResponseWriter, r *http.Request) {

        tmpl.ExecuteTemplate(w, "New", nil)
}
func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

        selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
        if err != nil {
            panic(err.Error())
        }
        emp := Employee{}
        for selDB.Next() {
            var id, yearsOfExperience, mobile int64
            var name, email, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus, github, linkedin, techgig, hackerRank string
            err = selDB.Scan(&id, &name, &mobile, &email, &yearsOfExperience, &projectName, &designation, &skillset, &completedTrainings, &projectAquiredSkills, &achievements, &employeeStatus, &github, &linkedin, &techgig, &hackerRank)
            if err != nil {
                panic(err.Error())
            }
            emp.Id = id
            emp.Name = name
            emp.Mobile = mobile
            emp.Email = email
            emp.YearsOfExperience = yearsOfExperience
            emp.ProjectName = projectName
            emp.Designation = designation
            emp.Skillset = skillset
            emp.CompletedTrainings = completedTrainings
            emp.ProjectAquiredSkills = projectAquiredSkills
            emp.Achievements = achievements
            emp.EmployeeStatus = employeeStatus
            emp.Github = github
            emp.Linkedin = linkedin
            emp.Techgig = techgig
            emp.HackerRank = hackerRank
        }
        tmpl.ExecuteTemplate(w, "Edit", emp)
        defer db.Close()
}
func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

        if r.Method == "POST" {
            name := r.FormValue("name") 
            fmt.Println(r.FormValue("id"))
            mobile := r.FormValue("mobile") 
            email := r.FormValue("email")
            yearsOfExperience := r.FormValue("yearsOfExperience")
            projectName := r.FormValue("projectName")
            designation := r.FormValue("designation")
            skillset := r.FormValue("skillset")
            completedTrainings := r.FormValue("completedTrainings")
            projectAquiredSkills := r.FormValue("projectAquiredSkills")
            achievements := r.FormValue("achievements")
            employeeStatus := r.FormValue("employeeStatus")
            github := r.FormValue("github")
            linkedin := r.FormValue("linkedin")
            techgig := r.FormValue("techgig")
            hackerRank := r.FormValue("hackerRank")
            insForm, err := db.Prepare("INSERT INTO employee(name, mobile, email, yearsOfExperience, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus, github, linkedin, techgig, hackerRank) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
            if err != nil {
                panic(err.Error())
            }
            insForm.Exec(name, mobile, email, yearsOfExperience, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus, github, linkedin, techgig, hackerRank)
            log.Println("INSERT: Name: " + name)
        }
        defer db.Close()
        http.Redirect(w, r, "/index", 301)

}
func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

        if r.Method == "POST" {
            name := r.FormValue("name") 
            fmt.Println(r.FormValue("id"))
            mobile := r.FormValue("mobile") 
            email := r.FormValue("email")
            yearsOfExperience := r.FormValue("yearsOfExperience")
            projectName := r.FormValue("projectName")
            designation := r.FormValue("designation")
            skillset := r.FormValue("skillset")
            completedTrainings := r.FormValue("completedTrainings")
            projectAquiredSkills := r.FormValue("projectAquiredSkills")
            achievements := r.FormValue("achievements")
            employeeStatus := r.FormValue("employeeStatus")
            github := r.FormValue("github")
            linkedin := r.FormValue("linkedin")
            techgig := r.FormValue("techgig")
            hackerRank := r.FormValue("hackerRank")
            id := r.FormValue("uid")
            insForm, err := db.Prepare("UPDATE employee SET name=?, mobile=?, email=?, yearsOfExperience=?, projectName=?, designation=?, skillset=?, completedTrainings=?, projectAquiredSkills=?, achievements=?, employeeStatus=?, github=?, linkedin=?, techgig=?, hackerRank=? WHERE id=?")
            if err != nil {
                panic(err.Error())
            }
            insForm.Exec(name, mobile, email, yearsOfExperience, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus, github, linkedin, techgig, hackerRank, id)
            log.Println("UPDATE: Name: " + name)
        }
        defer db.Close()
        http.Redirect(w, r, "/index", 301)

}
func Delete(w http.ResponseWriter, r *http.Request) {
    
    db := dbConn()
    emp := r.URL.Query().Get("id")

        delForm, err := db.Prepare("DELETE FROM employee WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        delForm.Exec(emp)
        log.Println("DELETE")
        http.Redirect(w, r, "/index", 301)
        defer db.Close()

}
func Checkout(w http.ResponseWriter, r *http.Request){
    db := dbConn()
    if r.Method == "POST" {
        id := r.FormValue("id") 
        password := r.FormValue("password")
        var isAuthenticated bool
        err := db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM admin WHERE id=? AND password=?",id,password).Scan(&isAuthenticated)
        if err != nil {
            log.Fatal(err)
        } 
        if isAuthenticated{
            fmt.Println(S)
            flag := 0
            for _, v := range S{
                fmt.Println(v, "hgfygd", flag)
                if v == id{
                    flag = 1
                }
            }
            if flag == 0{
                t := time.Now()
                session := id + " : " + t.String()
                S[session] = id
                fmt.Println(session, S)
            }
            if id == "1"{
                selDB,_ := db.Query("SELECT id, name, email FROM employee ORDER BY id")
                selDB2,_ := db.Query("SELECT id, name FROM skills ORDER BY id")
                emp := Employee{}
                res := []Employee{}
                for selDB.Next() {
                    var id int64
                    var name, email string
                    err = selDB.Scan(&id, &name, &email)
                    if err != nil {
                        panic(err.Error())
                    }
                    emp.Id = id
                    emp.Name = name
                    emp.Email = email
                    
                    res = append(res, emp)
                }
                skill := Skill{}
                res2 := []Skill{}
                for selDB2.Next() {
                    var id int64
                    var name string
                    err = selDB2.Scan(&id, &name)
                    if err != nil {
                        panic(err.Error())
                    }
                    skill.Id = id
                    skill.Name = name
                    res2 = append(res2, skill)   
                }
                //fmt.Println(result)
                var data = struct {
                    Skills []Skill
                    Employees []Employee
                }{
                    Skills: res2, Employees: res,
                }
                tmpl.ExecuteTemplate(w, "Admin", data)
                }else{
                http.Redirect(w, r, "/index", 301)
            }
        }else{
            //alert("Incorrect!")
            tmpl.ExecuteTemplate(w, "Login", 301)
            fmt.Fprintf(w,"<html><h6 align='center'>Incorrect Credentials! Please Try Again!</h6></html>")
        }
    }
}
func Signup(w http.ResponseWriter, r *http.Request){
    tmpl.ExecuteTemplate(w, "Signup", nil)
}
func Home(w http.ResponseWriter, r *http.Request){
    tmpl.ExecuteTemplate(w, "Home", nil)
}
func Search(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

        if r.Method=="POST"{
            keyword := "%"+r.FormValue("keyword")+"%"
            fmt.Println(keyword)
        selDB, err := db.Query("SELECT * FROM employee WHERE name like ? OR email like ? OR projectName like ? OR designation like ? OR  achievements like ? OR employeeStatus like ? OR skillset like ? OR completedTrainings like ? OR projectAquiredSkills like ? OR github like ? OR linkedin like ? OR techgig like ? OR hackerRank like ? ORDER BY id", keyword, keyword, keyword, keyword, keyword, keyword, keyword, keyword, keyword, keyword, keyword, keyword, keyword)
        if err != nil {
            panic(err.Error())
        }
        emp := Employee{}
        res := []Employee{}
        for selDB.Next() {
            var id, yearsOfExperience, mobile int64
            var name, email, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus, github, linkedin, techgig, hackerRank string
            err = selDB.Scan(&id, &name, &mobile, &email, &yearsOfExperience, &projectName, &designation, &skillset, &completedTrainings, &projectAquiredSkills, &achievements, &employeeStatus, &github, &linkedin, &techgig, &hackerRank)
            if err != nil {
                panic(err.Error())
            }
            emp.Id = id
            emp.Name = name
            emp.Mobile = mobile
            emp.Email = email
            emp.YearsOfExperience = yearsOfExperience
            emp.ProjectName = projectName
            emp.Designation = designation
            emp.Skillset = skillset
            emp.CompletedTrainings = completedTrainings
            emp.ProjectAquiredSkills = projectAquiredSkills
            emp.Achievements = achievements
            emp.EmployeeStatus = employeeStatus
            emp.Github = github
            emp.Linkedin = linkedin
            emp.Techgig = techgig
            emp.HackerRank = hackerRank

            res = append(res, emp)
        }
        tmpl.ExecuteTemplate(w, "Index", res)
    }
    defer db.Close()

}
func main() {
    log.Println("Server started on: http://localhost:8080")
    http.HandleFunc("/", Home)
    http.HandleFunc("/search", Search)
    http.HandleFunc("/signup", Signup)
    http.HandleFunc("/checkin", func (w http.ResponseWriter, r *http.Request){
        db := dbConn()
        if r.Method == "POST" {
            id := r.FormValue("id") 
            password := r.FormValue("password")
            confirmPassword := r.FormValue("confirmPassword")
            var isAuthenticated bool
            err := db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM admin WHERE id=?",id).Scan(&isAuthenticated)
            if err != nil {
                log.Fatal(err)
            } 
            if isAuthenticated{
                tmpl.ExecuteTemplate(w, "Signup", 301)
                fmt.Fprintf(w,"<html><h6 align='center'>User ID Already Exists! Please Try Again!</h6></html>")
            }else if password == confirmPassword{
                insForm,_ := db.Prepare("INSERT INTO admin(id, password) VALUES(?,?)")
                insForm.Exec(id, password)
                http.Redirect(w, r, "/login", 301)
                defer db.Close()  
            }else{
                tmpl.ExecuteTemplate(w, "Signup", 301)
                fmt.Fprintf(w,"<html><h6 align='center'>Passwords Didn't Match! PleaseTry Again!</h6></html>")
            }
        }        
    })
    http.HandleFunc("/login", func (w http.ResponseWriter, r *http.Request){
        tmpl.ExecuteTemplate(w, "Login", nil)
        fmt.Println(r.Method)
    })
    http.HandleFunc("/checkout", Checkout)
    http.HandleFunc("/index", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    //http.HandleFunc("/skills", Skills)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":8080", nil)
}
