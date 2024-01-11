<h1 align="center">Welcome to Simple Bank 👋</h1>

> 😋  This repository hosts my Backend Project for A Simple Bank, a Golang-based web service to create and manage money transfers between available accounts.  😋

## About
Through this project, I've developed a robust backend system for a banking application. 

Key features include account management, transaction records, and inter-account money transfers. 

The project spans several critical areas of backend development, including database design with DBML, transaction handling, API development using the Gin framework, user authentication with JWT and PASETO, and implementing robust testing strategies. Additionally, the repository demonstrates proficiency in deploying applications using Docker and Kubernetes on AWS, complete with domain registration and traffic routing. 

In this project, I've delved deeply into backend web service development, showcasing my practical abilities and grasp of key tools and technologies. I want to use this project as a platform to demonstrate my dedication and eager to learn backend development and my ability to contribute meaningfully in a professional context.

## Key Technologies and Concepts

This project incorporates a diverse set of technologies and concepts essential for backend development:

- **Programming Language**: Golang
- **Web Framework**: Gin
- **API Documentation**: Swagger
- **Database Design and Interaction**: DBML, SQL, Database Isolation Levels
- **Security and Authentication**: JWT, PASETO, HTTPS, TLS, Let's Encrypt
- **Containerization and Orchestration**: Docker, Kubernetes
- **Cloud Computing**: AWS, EKS
- **Continuous Integration/Deployment**: GitHub Actions

## Details
### 1. Working with database (Postgres + SQLC)
<details open>
  <summary> I designed DB schema and generated SQL code with dbdiagram.io </summary>
  <img src="https://github.com/ceciliachenguo/simplebank/assets/121702916/4b928eca-332a-4c90-a582-f2aa5c5948a7" alt="Simple Bank Schema" width="800">
</details>

<details open>
  <summary> Docker + Postgres + TablePlus </summary>
  <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 
    I used <a href="https://www.docker.com/" target="_blank">Docker</a> to run containers and chose Postgres 12 as my relational database. For easier look up to the actual data in the local database, I used the <a href="https://tableplus.com/" target="_blank">TablePlus</a> GUI. When using Docker, I use terminal command extentively, such as docker ps -a, docker start, and docker exec.
  </p>
</details>

<details open>
  <summary> Database Migration</summary>
  <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 
    When working with a database, schema migration is often necessary to adapt to new business requirements. I ran and managed these database schema migrations using the <a href="https://github.com/golang-migrate/migrate" target="_blank">Golang Migrate</a> library. This library offers various customized commands for migrating schemas up and down. The SQL code for schema migration is stored in the <a href="https://github.com/ceciliachenguo/simplebank/tree/main/db/migration" target="_blank">'db/migration'</a> folder. 
  </p>
</details>

<details open>
  <summary> Makefile and .Phony list</summary>
  <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 
    Remembering and entering various lengthy terminal commands can be exhausting, so I defined those commonly used commands in my <a href="https://github.com/ceciliachenguo/simplebank/blob/main/Makefile" target="_blank">Makefile</a> and listed them under .PHONY targets for easy execution.
  </p>
</details>

<details open>
  <summary> Generate CRUD Golang code from SQL</summary>
  <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 
    After comparing various libraries for converting SQL queries into type-safe Go code, such as database/sql, Gorm, sqlx, and sqlc, I decided to use <a href="https://github.com/sqlc-dev/sqlc/tree/v1.4.0" target="_blank">sqlc</a> for interacting (Create, Read, Update, and Delete operations) with my database. After executing sqlc, it automatically generates struct definitions for models, function definitions with parameters, and the dbtx interface(this allows me to use either a database or a transaction to execute a query).
  </p>
</details>

<details open>
  <summary> Unit Tests for database CRUD</summary>
  <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 
    Using Go's testing package and <a href="https://github.com/stretchr/testify" target="_blank">Testify</a> library(require package), I wrote various unit tests for the database CRUD operations. The files are named ending with "_test.go".
  </p>
</details>

<details open>
  <summary> DB Transaction lock, and handle Deadlock </summary>
  <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 
    In order to show deadlock, I used TDD (test driven development) to create multiple go routine to execute transfer transactions concurrrently, then iterate through the list of results to check the created transfer and entry objects, and finally check the balances of those accounts accordingly. <br> 
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 
    A deadlock occurs when multiple processes try to access or modify the same data, with each of them waiting for the others to release the data while continuing to hold onto it themselves. This leads to a situation where none of the processes can actually control the data. A simple solution to prevent deadlocks is to include the 'FOR NO KEY UPDATE' clause in the SQL code. This informs PostgreSQL that the current operation will not modify the foreign key (account id), even though the main purpose of the SQL statement is to change the balance. As a result, the transaction lock does not hold onto the accounts table (where the account id is the primary key), thus reducing the risk of deadlocks. Deadlocks can also occur due to the order in which transactions update shared resources. For example, if Transaction T1 locks Account A first and then tries to lock Account B, while simultaneously Transaction T2 locks Account B and then tries to lock Account A, a deadlock ensues. To resolve this, I implemented a solution that enforces a consistent ordering: always locking the account with the smaller ID before locking the account with a larger ID.
  </p>
</details>

<details open>
  <summary> ACID Property </summary>
  <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 
    The default isolation level for my PostgreSQL database is 'Repeatable Read'.
  </p>
</details>




## Others
 ✨ [See my iOS Projects](https://github.com/ceciliachenguo/iOSAppPortfolio_Cecilia_in_Marlo)

## Author

👤 **Cecilia Chen**

## Thanks!!!

* LinkedIn: https://www.linkedin.com/in/ceciliaguochen/
* Github: [@ceciliachenguo](https://github.com/ceciliachenguo)
