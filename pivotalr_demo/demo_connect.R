library(PivotalR)
pg <- db.connect(port = 5333, dbname = 'madlib', quick = TRUE) # Postgres database
gp <- db.connect(port = 16526, dbname = 'madlib', quick = TRUE) # Greenplum database
hq <- db.connect(port = 18526, dbname = 'madlib', quick = TRUE) # HAWQ database
db.list()
