```
go run main.go
```

```
docker-compose up --build
```


## 1. User logs in 
```
curl --request POST \
  --url http://localhost:8000/users/login \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "admin@admin.com",
	"password": "secrets"
}'
```
Response:
```
{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsImVtYWlsIjoiZWExMXJhdGhhdXNAdXNlci5jb20iLCJ0eXBlIjoicmVmcmVzaCIsInJvbGUiOiJ1c2VyIiwic3RhbmRhcmRfY2xhaW1zIjp7ImV4cCI6MTY0OTk0OTk3Nn19.njAbzh6WH-a-40jEPASXe7GktEzaNii78l7KqWNn8s4",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsImVtYWlsIjoiZWExMXJhdGhhdXNAdXNlci5jb20iLCJ0eXBlIjoiYWNjZXNzIiwicm9sZSI6IlVudGVyYWJzY2huaXR0Iiwic3RhbmRhcmRfY2xhaW1zIjp7ImV4cCI6MTYzOTgzMzk0M319.mLYi9LIRhAnbjVpe4FWSXENrGFAQBKI-teF6R8MDgNM",
    "role": "Unterabschnitt"
}
```
Everytime when user sends a request to our API, in the header of the request a value of the field "Authorization" should be "Bearer " + `access_token`.

## 2. Refresh access token
Everytime when user sends the request, frontend side should check whether access token is expired or not. If it is expired, then before the user's request, frontend should send the request /refresh to get a new access_token
Here we do not need Authorization in header of the request. 
Request:
```   curl --request POST \
   --url http://localhost:8000/users/refresh \
   --header 'Content-Type: application/json' \
   --data '{
   "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsImVtYWlsIjoiZWExMXJhdGhhdXNAdXNlci5jb20iLCJ0eXBlIjoicmVmcmVzaCIsInJvbGUiOiJ1c2VyIiwic3RhbmRhcmRfY2xhaW1zIjp7ImV4cCI6MTY0OTk0OTk3Nn19.njAbzh6WH-a-40jEPASXe7GktEzaNii78l7KqWNn8s4"
   }'
```
Response:
```
{
	"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsImVtYWlsIjoiZWExMXJhdGhhdXNAdXNlci5jb20iLCJ0eXBlIjoicmVmcmVzaCIsInJvbGUiOiJ1c2VyIiwic3RhbmRhcmRfY2xhaW1zIjp7ImV4cCI6MTY0OTk0OTk3Nn19.njAbzh6WH-a-40jEPASXe7GktEzaNii78l7KqWNn8s4",
	"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsImVtYWlsIjoiZWExMXJhdGhhdXNAdXNlci5jb20iLCJ0eXBlIjoiYWNjZXNzIiwicm9sZSI6IlVudGVyYWJzY2huaXR0Iiwic3RhbmRhcmRfY2xhaW1zIjp7ImV4cCI6MTY0MDExNjg1NX19.hNB3GhE-N9d-fxX1vETYxTG-7Soj9rOvf_m-ofIlEes"
}
```

## 3. User list
Request:
```   curl --request GET \
   --url http://localhost:8000/users \
   --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiJhZG1pbkBhZG1pbi5jb20iLCJ0eXBlIjoiYWNjZXNzIiwicm9sZSI6ImFkbWluIiwic3RhbmRhcmRfY2xhaW1zIjp7ImV4cCI6MTYzOTU4NTQ1OX19.t13VAdthI9y-Bcq6vGzgnxTcYw-8zArDdptdFrisOsw' \
   --header 'Content-Type: application/json'
```
Response:
```
[
	{
		"id": 1,
		"name": "Einsatzleiter",
		"Password": "",
		"email": "admin@admin.com",
		"branch_name": "Einsatzleiter",
		"branch_id": 2,
		"parent_id": 2,
		"parent_name": "Mollnhof",
		"create_date": "2021-12-15T12:16:44.000415+01:00",
		"update_date": "0001-01-01T00:00:00Z"
	},
    ...
]
```

## 4. Creating a new user.
This action is allowed only to Ensatzleiter.

We put a name of the new branch, its email, set password, phone and parent_id. `parent_id` means what branch has to be its "boss". We choose it from the list that it's possible to get from the request GET /users/.

Request:
```
curl --request POST \
  --url http://localhost:8000/admin/create_user \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiJhZG1pbkBhZG1pbi5jb20iLCJ0eXBlIjoiYWNjZXNzIiwicm9sZSI6ImFkbWluIiwic3RhbmRhcmRfY2xhaW1zIjp7ImV4cCI6MTYzOTU4NTQ1OX19.t13VAdthI9y-Bcq6vGzgnxTcYw-8zArDdptdFrisOsw' \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "EA1.2 - Dom",
	"email": "ea12dom@user.com",
	"password": "secret",
	"phone": "123456",
	"parent_id": 14
}'
```
Response:
```
{
    "err_code": 0,
    "err_message": ""
}
```

## 5. Sending verification email
Admin send the request, that sends the email with OTP to the email of the user
Request:
```
curl --request POST \
  --url http://localhost:8000/admin/email_verification \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiJhZG1pbkBhZG1pbi5jb20iLCJ0eXBlIjoiYWNjZXNzIiwicm9sZSI6IkVpbnNhdHpsZWl0ZXIiLCJzdGFuZGFyZF9jbGFpbXMiOnsiZXhwIjoxNjQwMTE3MjEwfX0.TWAL_riXQ_IjF-CrKrZAGToYasoi3iDIETBIjhnDxBY' \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "ea12dom@user.com"
}'
```

By now, because we did not set up email credentials of admin, from which all these email will be sent to users, we get this response.
Response (by now):
```
{
	"otp": "280928"
}
```
But then we will change it to this:
```
{
    "err_code": 0,
    "err_message": ""
}
```

## 6. Verification of email
The user gets email with this otp. He copies this otp, writes new password and sends this request.
Request:
```
   curl --request POST \
   --url http://localhost:8000/users/activation \
   --header 'Content-Type: application/json' \
   --data '{
   "otp": "290204",
   "password": "secrets"
   }'
```

Response:
```
nil
```

Now he can log in to the system.