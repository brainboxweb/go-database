Database connection example
====



Mysql
----

Docker database

        docker run --name brainbox -e MYSQL_DATABASE=brainbox -e  MYSQL_ROOT_PASSWORD=carrot -p 3306:3306  -d mysql:latest 
        

Connection string
----
        
        db, err := sql.Open("mysql", "root:carrot@tcp(192.168.99.100:3306)/brainbox")
