# Mini-Project
Golang Mini Project at Celerates

### Create New Rest API
- Used Gin, GORM, and PostgreSQL
- Used "jobs" as domain

### SQL
```
CREATE TABLE "users" (
"user_id" SERIAL primary key not null,
  "username" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "role" varchar(20) NOT NULL,
  "company_id" int DEFAULT NULL
);

CREATE TABLE "jobs" (
  "job_id" SERIAL PRIMARY KEY NOT NULL,
  "title" varchar(20) NOT NULL,
  "city" varchar(20) NOT NULL,
  "status" varchar(20) NOT null,
  "company_id" int DEFAULT NULL
);

INSERT INTO "jobs" VALUES
  (DEFAULT,'Back-End Developer','Jakarta', 'open', 101),
  (DEFAULT,'Front-End Developer','Bandung', 'open', 101),
  (DEFAULT,'Quality Assurance','Malang', 'open', 102),
  (DEFAULT,'Full-Stack Developer','Surabaya', 'open', 102);
```

### REST API
Routes below are endpoints used in this project

#### Jobs

- `GET` All Jobs (implemented pagination) `/jobs`

With pagination (show page 1 and limit by 1 data) `/jobs?page=1&limit=1`

- `GET` Job by ID (get job with id = 1) `/jobs/1`

- `POST` Add a data to table `/jobs`

Insert data to body using JSON format. Below is the example
```
{
    "title": "Web Developer",
    "city": "Jakarta",
    "status": "Open",
    "company_id": 101
}
```
- `PUT` Update existed data (update data with id = 1) `/jobs/1`
- `DELETE` Delete existed data (delete data with id = 1) `/jobs/1`

#### Users
- `POST` Register User `/users`

Insert data to body
```
{
    "username": "abc",
    "password": "abc",
    "role": "user",
    "company_id": 101
}
```

- `POST` Login `/login`

Insert data to body
```
{
    "username": "abc",
    "password": "abc"
}
```

Example of response after logged in
```
"id": 1,
    "username": "abc",
    "company_id": 101
    "role": "user",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjF9.06lbO3Sb1wyS45SCYsUxwrUyon5u6l1bnCbzwp83wbI"
```
