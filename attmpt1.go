package main

	import (
		"net/http"
		"github.com/julienschmidt/httprouter"
		"fmt"
		"github.com/ttacon/chalk"
		_ "github.com/mattn/go-sqlite3"
		"log"
		"database/sql"
		"strconv"
		"html/template"

	)
	// data structure for trend

	

	type User struct{

		
		Roll_number int
		Reg_number int
		Air int
		Name string
		Sex string
		Parent_name string
		Nationality string
		Category string
		P_addr string
		C_addr string
		P_city string
		C_city string
		P_pincode int 
		C_pincode int
		P_landline int 
		C_landline int
		P_mobile int
		C_mobile int 
		Email string

		}

	func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){

		http.ServeFile(w,r,"index.html")
       
       ip  := r.RemoteAddr

		
		fmt.Println(chalk.Yellow,ip," requested first-year page...",chalk.Reset)

	}

	func ServeHTMl(w http.ResponseWriter, r *http.Request, _ httprouter.Params){

			ip  := r.RemoteAddr

			filepath:= r.URL.Path

			filename:=filepath[1:]

			if len(filename)==0 {
				http.ServeFile(w,r,"home.html")

				fmt.Println(chalk.Yellow,ip," requested home page...",chalk.Reset)	

			} else{  
				http.ServeFile(w,r,filename+".html")

				fmt.Println(chalk.Yellow,ip," requested",filename,"page...",chalk.Reset)	
			}
		
	}

	func Submit(w http.ResponseWriter, r *http.Request, p httprouter.Params){


			rollnumber,_ := strconv.Atoi(r.FormValue("rollnumber"))
			air,_ := strconv.Atoi(r.FormValue("air"))
			regnumber,_ :=strconv.Atoi(r.FormValue("regnumber"))
			P_pincode,_  := strconv.Atoi(r.FormValue("pincode1"))
			C_pincode,_:= strconv.Atoi(r.FormValue("pincode2"))
			P_landline,_ := strconv.Atoi(r.FormValue("phone1"))
			C_landline,_ := strconv.Atoi(r.FormValue("phone2"))
			P_mobile,_ := strconv.Atoi(r.FormValue("mobile1"))
			C_mobile,_ := strconv.Atoi(r.FormValue("mobile2"))
				
		student := User{
			
			rollnumber,
			regnumber,
			air,
			r.FormValue("name"),
			r.FormValue("sex"),
			r.FormValue("parentname"),
			r.FormValue("nationality"),
			r.FormValue("catgy"),
			r.FormValue("addr"),
			r.FormValue("commaddr"),
			r.FormValue("city1"),
			r.FormValue("city2"),
			P_pincode,
			C_pincode,
			P_landline,
			C_landline,
			P_mobile,
			C_mobile,
			r.FormValue("email")  }

		output:= "You have been registered as roll number : "+r.FormValue("rollnumber")

		fmt.Println(output)

		

		t,_ := template.ParseFiles("success.html")

		t.Execute(w,student)

		db, err := sql.Open("sqlite3", "./test.db")
			riperr(err)    	
    	defer db.Close()

    	query1,err := db.Prepare("insert into student_1 values(?,?,?,?,?)" )

    	query1.Exec(student.Roll_number,student.Reg_number,student.Air,student.Name,student.Sex)
    	

    	query2,err := db.Prepare("insert into student_2 values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

    	query2.Exec(student.Name,student.Parent_name,student.Nationality,student.Category,student.P_addr,student.C_addr,student.P_city,student.C_city,student.P_pincode,student.C_pincode,student.P_landline,student.C_landline,student.P_mobile,student.C_mobile,student.Email)
	}


	func GoBack(w http.ResponseWriter, r *http.Request, p httprouter.Params){

		http.Redirect(w, r, "/", 301)
	}

	func main(){
		

		

		newfeed:= chalk.Red.NewStyle().WithBackground(chalk.Black)

		


		Server:= httprouter.New()

		
		Server.GET("/firstyear",ServeHTMl)

		Server.POST("/submit",Submit)

		Server.GET("/submit",GoBack)

        Server.ServeFiles("/resources/*filepath", http.Dir("resources"))

        Server.GET("/",ServeHTMl)

		fmt.Println(newfeed,"waiting at :3000",chalk.Reset);

		http.ListenAndServe(":3000",Server)

	}
 

 func riperr(err error){
 	if err!= nil{

    		log.Fatal(err)

    	}
 }

