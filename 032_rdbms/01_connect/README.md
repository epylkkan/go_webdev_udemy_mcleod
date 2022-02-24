Using MySQL
Install MySQL
Download MySQL Community Server
We will need a MySQL driver
go get github.com/go-sql-driver/mysql
read the documentation
see all SQL drivers

Astaxie's book
Include the driver in your imports
_ "github.com/go-sql-driver/mysql"
Read the documentation

AWS: 
Determine the Data Source Name
user:password@tcp(localhost:5555)/dbname?charset=utf8
Read the documentation
Open a connection
db, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")

================

go get github.com/go-sql-driver/mysql


====================

https://medium.com/@DazWilkin/google-cloud-sql-6-ways-golang-a4aa497f3c67


PROJECT=ep-appengine-v2    // gcp-project-id
REGION=europe-north1
ROOT=gotraining
INSTANCE=${PROJECT}:${REGION}:${ROOT}

DBNAME=test
DBUSER=root
DBPASS=Purpletree462


===================

https://cloud.google.com/sql/docs/mysql/connect-app-engine-standard#go

GCP:

var (
        dbUser                 = mustGetenv("DB_USER")                  // e.g. 'my-db-user'
        dbPwd                  = mustGetenv("DB_PASS")                  // e.g. 'my-db-password'
        instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
        dbName                 = mustGetenv("DB_NAME")                  // e.g. 'my-database'
)

socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
if !isSet {
        socketDir = "/cloudsql"
}

dbURI := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

// dbPool is the pool of database connections.
dbPool, err := sql.Open("mysql", dbURI)
if err != nil {
        return nil, fmt.Errorf("sql.Open: %v", err)
}

// ...

return dbPool, nil
