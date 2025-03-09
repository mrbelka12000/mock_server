# Mock Server

A simple mock server for handling requests and responses based on predefined routes and cases.

## Features

- Create and manage services
- Assign handlers with cases
- Define request and response bodies
- Configure request and response headers
- View registered services

## Installation
   
```shell
  docker volume create mock_server
  docker run -d -p 5552:5552 -v mock_server:/db \
  -e PATH_TO_DB=/db/data.db \
  -e HTTP_PORT=5552 \
  mrbelka12000/mock_server
```

#### Mock Server is now running on http://0.0.0.0:5552.

## Usage

1. **Create a Service**
    - Start by creating a service with a name.

2. **Configure Handlers**
    - Click on the created service.
    - Add handlers to define the routes the mock server should check.
    - Each handler can have multiple cases, which determine how requests are matched and responses are returned.

3. **Define Cases**
    - Cases specify the matching criteria and response behavior.
    - Each case includes:
        - **Tag:**
            - `1 = Default` → Matches requests based only on route equality.
            - `2 = On Equal` → Matches requests based on route, request body, and request headers.
        - **Request Headers**
        - **Request Body**
        - **Response Headers**
        - **Response Body**
    - If `Tag = 2`, the **Request Headers** and **Request Body** will be compared with the incoming request data before returning a response.

4. **Send Requests to the Mock Server**
    - Once all necessary data is set up, you can send an HTTP request to:
      ```
      http://$host/api/$service_name/$route
      ```
    - The mock server will match the incoming request against the defined handlers and cases to return the appropriate response.  
