library(PivotalR)

f <- lm(rings ~ . - id, data = abalone)

summary(f)

pdf("test01.pdf")
plot(abalone$rings, predict(f))
dev.off()
