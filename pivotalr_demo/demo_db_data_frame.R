library(PivotalR)
pg <- db.connect(port = 16526, dbname = 'madlib', host='127.0.0.1', user='gpadmin', verbose = FALSE, quick = TRUE)
dat <- db.data.frame('madlibtestdata.dt_abalone')
