Arash plans to implement a heavy web service, but he knows that from the very beginning, the number of user requests to endpoints with high processing load will be very high. So he decided to use rate limiter. Since he doesn't know how to implement this rate limiter, he asked you to implement it for him.

Project details
  The structure of the project files is as follows:

rate-limit
├── db
│ └── connection.go
├── handlers
│ └── root.go
├── limiters
│ ├── by_app_key.go
│ └── by_ip.go
├── test
│ └── ratelimit_sample_test.go
├── go.mod
├── go.sum
└── main.go

The time/rate dependency is defined in the project and you must use it to implement the rate limiter.



Database
PostgreSQL database is used in this project. The program contains a table named app_keys, which contains two columns: id of BIGSERIAL type and key of VARCHAR(255) type.

In the db package, a function called GetConnection is defined, which you should use to get the database connection. You can change the database connection information in this file.

Roots
Only one route with address / is defined in the program, which can apply different rate limiters on its handler.

ByIp function
This function includes the following three parameters:

next is of type http.Handler: it is the main handler of the root.
refillRate of type rate.Limit: the amount of increase in the number of user's allowed requests per second (for example, if its value is equal to 3, 3 units will be added to the number of user's allowed requests per second, provided that the current number of allowed requests is smaller than tokenBucketSize.)
tokenBucketSize of type int: the limit of the number of consecutive user requests
Implement this function to limit requests based on IP. If the user has not faced any restrictions, the request should be processed by next. Otherwise, the response code should be changed to 429, the Content-Type header value should be set to application/json, and the response body should look like this:

{"error": "too many requests"}
JSON
ByAppKey function
The signature of this function is similar to the ByIp function.

Implement this function to limit requests based on the value of the X-App-Key header in the request. If the value of this header does not exist in the app_keys table, the restriction should not be applied. Otherwise, the restriction should apply. If the user encountered a restriction, the response should be similar to the response expected for the ByIp function.

The X-App-Key header is guaranteed to be present in the request when using this function.

Combination of rate limiters
We may want to combine rate limiters. If the ByIp and ByAppKey functions are used together, the restriction must be applied both by IP and by the value of the X-App-Key header. For example, if the value of tokenBucketSize is equal to 5 and we send 5 requests in a row with different IPs, but with the same X-App-Key (so that it is present in the app_keys table), we should face the limitation. Also, if we send these requests with the same IP, but with different X-App-Keys (so that they are in the app_keys table), we still have to face the limitation.

Hints
There is no need for persistent information about rate limiters (the data about rate limiters can be lost when the program is terminated).
You are only allowed to make changes in the limiters folder.
If needed, you can define new files in the limiters folder.
What you need to upload
  In the newsletter