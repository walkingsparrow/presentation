library(PivotalR)
pg <- db.connect(port = 16526, dbname = 'madlib', verbose = FALSE, quick = TRUE)
dat <- db.data.frame('madlibtestdata.dt_abalone', verbose = FALSE)
dat
