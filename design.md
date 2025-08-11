1. Purpose
The goal of this application is to prevent excessive API calls. It will be done by creating a middle-ware that would store callers info and amount of requests to prevent excessive usage.
2. Not purpose
Creation of fully functional HTTP server with various functionalities.
3. Main Components
    a. Server: contains a simple HTTP server that would receive and respond to GET requests ONLY.
    b. In-Memory Map: a map that would store API keys as well as the amount of requests.
    c. Interception: component that would receive the request, use in-memory map to see if the rate-limit is not violated, allowing the HTTP requests to pass. When limit is hit, the response returned will be 429.
    d. Cleanup: will be handled with goroutines that will delete expired data entries.
    e. Threat safety: ensure that requests coming at once will be safe from being read and written at the same time.
4. Testing
    a. A set of API keys will be used to send requests, with some of them exceed the limit. Expected: the API keys that exceeded the limit must be blocked for X amount of time.
    b. Expired data entries that exceeded limit will be added to the map, then calls will be made using API keys existing in these entries. Expected: goroutines must clean up the data so the API keys limit is reset.
    

Basic config:
1. Initial request limit: 15 per minute