# One Siam Restaurant

`List of components`

Go: [https://golang.org/](https://golang.org/)

Gin: [https://gin-gonic.com/](https://gin-gonic.com/)

Docker: [https://www.docker.com/](https://www.docker.com/)

GoDotEnv: [https://github.com/joho/godotenv](https://github.com/joho/godotenv)


## Getting Started

To get started with the One Siam Restaurant application, You can run locally using following command

```bash
    go run main.go
```

Or you'll need Docker installed on your machine and following steps below to run the application:

1. **Clone the repository:**

```bash
git clone https://github.com/yourgithubusername/onesiamrestaurant.git
cd onesiamrestaurant
```

2. Build the Docker container:
Run the container:

```bash
docker build -t onesiamrestaurant .
```

3. Run the container:

```bash
docker run -d -p 8080:8080 onesiamrestaurant
```

After running these commands, the application should be accessible at http://localhost:8080.


License
This project is licensed under the MIT License.