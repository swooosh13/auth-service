db.auth("mongoadm", "mongoadm")

authdb = db.getSiblingDB("auth")
authdb.createUser({
  "user": "atuser",
  "pwd" : "atuser",
  "roles": [
    { "role" : "readWrite", "db" : "auth"}
  ],
  "mechanisms": [ "SCRAM-SHA-1" ],
  "passwordDigestor": "client"
})
authdb.auth("atuser", "atuser")