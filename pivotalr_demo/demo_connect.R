library(PivotalR)
gp <- db.connect(port = 16526, dbname = 'madlib', host='127.0.0.1', user='gpadmin', quick = TRUE) # Greenplum database
pg <- db.connect(port = 5333, dbname = 'madlib', host='127.0.0.1', user='gpadmin', quick = TRUE) # Postgres database
hq <- db.connect(port = 18526, dbname = 'madlib', host='127.0.0.1', user='gpadmin', quick = TRUE) # HAWQ database
db.list()
