# rest_api
A test of interacting with a server RESTfully

POST: a 'block' of data up to 64MiB gets POSTed to the server and returns a SHA1 hash

GET: given a SHA1 hash, returns block previously POSTed

Assumption is that the hashing will happen on the server side, and only return previously posted, or applicable error states

Error states should include:
* Disk is full?
* Data not found 404
* Data corrupt on disk?
* Data already exists?

Unsure how to handle 'blocks', replacing with json. 

Returning with SHA1 didn't work as expected, not enough time to find a fix