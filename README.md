<h1 align="center">
  <img src="spike.png" alt="">
  <br>spike<br>
  <p style="margin-top: 10px; font-size: 13px;">Aww, it's purple, Daddy! Can I have one at my house, please?! – 👧</p>
</h1>

Spike is an example app demo'ing how to use Docker, Go, and Aerospike to create a simple REST service. The repo contains a simple application written in Go that contains a single API to return user data.

I've skipped a few steps for the sake of just getting something working. The app loads a single record into the DB, and only accepts requests for that record. I'm also not an expert in Go or Aerospike, so there is definitely room for improvement in the utilization of the language and DB.

As a next step, I would probably throw in request validation and a simple error service.

## Quick start
```
$ docker-compose up -d
```

## Fetching data
```
$ curl http://localhost:8080/user/1?api_key=12345
{
  "api_key":"12345",
  "first_name": "John",
  "last_name": "Doe",
  "company": "Acme"
}
```