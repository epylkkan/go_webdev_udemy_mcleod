locking down your database
create admin super user
use admin
db.createUser(
  {
    user: "jamesbond",
    pwd: "moneypennyrocks007sworld",
    roles: [ { role: "userAdminAnyDatabase", db: "admin" } ]
  }
)
built in user roles

exit mongo & then start again
mongod --auth
mongo -u "jamesbond" -p "moneypennyrocks007sworld" --authenticationDatabase "admin"

