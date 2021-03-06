suppressMessages(library(PivotalR))
gp <- db.connect(port=16526, dbname='madlib', host='127.0.0.1', user='gpadmin', verbose=F, quick=T) # 1, Greenplum database
x <- db.data.frame('madlibtestdata.dt_abalone', verbose=F)

f <- madlib.glm(rings ~ . - id, data = x, family = poisson(log))
AIC(f)
