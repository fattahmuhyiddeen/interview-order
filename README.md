# Deployment in Heroku
1. Port number will be always changed. Therefore, always read port number from env `PORT`. Do not set this explicitly in Heroku
1. Database credential always changed. Therefore, always read from env `DATABASE_URL` if in heroku server, or get it from Heroku CLI command `heroku config:get DATABASE_URL -a emeltrack. Do not set this explicitly in Heroku


#env
PORT=1323
DATABASE_URL=postgres://USERNAME:PASSWORD@localhost:5432/DATABASE?sslmode=disable
ENV=local
