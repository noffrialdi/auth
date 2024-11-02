# auth
Sign up &amp; Login

#========= requipment =================#
1. Golang minimum v1.21
2. Mariadb



#========= Setup ========#
1. first setup mariadb in localhost
2. create database auth
3. create table user with ddl in folder sql/user.sql
4. create config yaml in config/files/main.development.yaml (copy value main.development.example)

#========= running ========#
1. go main.go serve-http
2. import postman Auth.postman_collection
3. and running with postman

