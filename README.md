Example of Using curl with Real IDs
For GET request, you don't need any ID. Just hit the endpoint:

`curl -X GET http://localhost:8080/companies`

------------

For POST request, create a new company:

`curl -X POST http://localhost:8080/companies \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Tech Solutions",
        "address": "123 Silicon Valley, CA",
        "website": "http://techsolutions.com"
    }'`
	
------------

After adding a company, you can use the PUT or PATCH requests to update it:

`curl -X PUT http://localhost:8080/companies/12345 \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Tech Solutions Inc.",
        "address": "456 New Silicon Valley, CA",
        "website": "http://techsolutionsinc.com"
    }'`
	
------------

Finally, delete the company using its ID:

`curl -X DELETE http://localhost:8080/companies/12345`

---

```mermaid
flowchart TD
    A[GET /companies] --> B[POST /companies]
    B --> C[PUT /companies/:id]
    C --> D[DELETE /companies/:id]


flowchart TD
    A[Client Browser] --> B[Web Server]
    B --> C[Application Server]
    C --> D[Database]