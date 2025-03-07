# mock_server


## Mock external APIs using mock_server. 

### Info:

#### Service defines some object or domain logic of your needs

#### Handler defines implementation that need to be mocked in srv.

#### Handler's option:

1) ServiceID 
2) Route
3) Cases
   4) Tag (equal | default)
   5) Request body (in case of equal, need to set real request body that will be sent)
   6) Response body (random string that you need)



## Usage:

1) Create srv
2) Assign some handlers to the srv, mock_server will be automatically update routes for the next use
3) Try to send request on your handler
4) Finished