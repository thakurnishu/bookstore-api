# BookStore API

The Bookstore API is a sample REST application that provides a simple way to register, list, and find books. It is designed to help developers learn and practice building RESTful APIs using Golang. The API is documented using Apiary and provides a clear and concise overview of its functionality. The API includes endpoints for registering books, listing all books, and finding books by title. The API requires Docker, Golang, and Go-Migrate CLI to be installed on the local machine. The API is open-source and can be cloned from the GitHub repository

## Perquisites
- Docker Should be [Installed](https://docs.docker.com/engine/install/)
- Golang Should be [Installed](https://go.dev/doc/install)

## Run API On Local Machine
- Clone repo
     -
       git clone https://github.com/thakurnishu/bookstore-api.git
       cd bookstore-api
       go mod download #Downlad all dependencies
- Choose which Database you want to use `postgres` or `mysql`
- In `main.go` file uncomment store which you want and comment other store
- Like you want to go with `postgres`, uncomment postgres store and comment mysql store
     - ```go
       go
       func main() {

	      listenAdrr := flag.String("listenAdrr", "3000", "Port where server will listen")
	      flag.Parse()

	      store, err := storage.NewPostgresStore()
	      // store, err := storage.NewMySQLStore()
	      if err != nil {
		     log.Fatalln(err)
	      }
	      if err = store.Init(); err != nil {
		     log.Fatalln(err)
	      }

	      server := api.NewServer(*listenAdrr, store)
	      server.Run()

       }
       ```
- To Start Postgres Database [Docker Container]
     -
       make postgres-start
- To Start Mysql Database [Docker Container]
     -
       make mysql-start
- Start API Server
     - 
       make run
- After that **register** some books and try **listing** and **finding** book. 

#### Register 
- Register Book At `/book/register`
- Added Book use `POST`
- In JSON Send
```json
   {
     "title": "Book Title",
     "publication": "Publication Name",
     "author": "Author Name",
     "isbn": 1234567293215, # Unique Value (isbn-13)
     "available": 5                                                  
   }
```
#### List All Books
- Get All Books At `/book`
- To list Books use `GET`
- Response will look like
```json
[
  {
    "title": "Title 1",
    "author": "Author 1",
    "publication": "Publication 1",
    "isbn": 1234567893415,
    "available": 2,
    "added_at": "2024-01-16T20:12:06.551696Z"
  },
  {
    "title": "Title 2",
    "author": "Author 2",
    "publication": "Publication 2"
    "isbn": 1234567893215,
    "available": 5,
    "added_at": "2024-01-16T20:12:24.740663Z"
  },
  {
    "title": "Title 3",
    "author": "Author 3",
    "publication": "Publication 4"
    "isbn": 1234567293215,
    "available": 10,
    "added_at": "2024-01-16T20:12:35.421522Z"
  }
]
```
#### Find Book
- Find Book At `/book` by using Query Parameter with `parameter` = `title` and `value`
- example if book name **First Book** val to query parameter `first-book`
- To list Book use `GET`
- Response will look like
```json
{
  "title": "First Book",
  "author": "Author",
  "publication": "Publication",
  "isbn": 1234567893415,
  "available": 5,
  "added_at": "2024-01-16T20:12:06.551696Z"
}
```
### CleanUp
- Stop Server using `ctrl+c`
- Stop and Delete Postgres Container
     -
       make db-stop
