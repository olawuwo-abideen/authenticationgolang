## Golang Authentication

## Built With:

- Golang
- Gin
- Dotenv
- Validator
- jwt-go
- mongo-driver
- crypto

## Installation

- clone the repository

```sh
git clone git@github.com:olawuwo-abideen/authenticationgolang.git

```

- navigate to the folder

```sh
cd authenticationgolang.git
```

## Run the app in development mode

Open a terminal window session, or the equivalent on your machine, and enter the following command to install all the
Node modules needed to run the app:

```sh
go mod tidy 

```

After running `go mod tidy` navigate to main.go file and enter the following `go run main.go` command:

```sh

npm go run main.go

```

This will start the app and set it up to listen for incoming connections on port 3000. Open up your browser of choice
and go to the url

```sh

http://localhost:3000

```

to start using the app.

## API Endpoints

The following API endpoints are available:

- https://localhost:3000/

- `POST https://localhost:3000/user/signup` - User signup
- `POST https://localhost:3000/user/login` - User login
- `GET https://localhost:3000/user/` - Get all users
- `GET https://localhost:3000/user/:id` - Get user by id

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

Feel free to check the [issues page](https://github.com/olawuwo-abideen authenticationgolang/issues).

## Authors

👤 **Olawuwo Abideen**

- GitHub: [@Olawuwo Abideen](https://github.com/olawuwo-abideen)
- Twitter: [@Olawuwo Abideen](https://twitter.com/olawuwo_abideen)
- LinkedIn: [@Olawuwo Abideen](https://www.linkedin.com/in/olawuwo-abideen/)