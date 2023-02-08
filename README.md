# spots-query-and-endpoint
Simple SQL query and GoLang endpoint

To run GoLang endpoints open powershell or terminal and type 'go run spotsEndpoint.go Spots.go'
use Postman to access endpoints.
##### returns all spots
- GET: 'localhost:8080/spots'
##### returs spot by ID
- GET: 'localhost:8080/spots/2887b105-bdc5-4ead-a705-17ff3f046405'

use 'SQL/create table 'MY_TABLE' and insert data.sql' to create table and add data.</br>
use 'SQL/CREATE FUNCTION Parse_For_Domain_Name.sql' to add filter Domain_Name function.</br>
use 'SQL/spots have the same domain.sql' to search detailed information about spots.
returns:

<table>
  <tr>
    <th>Domain</th>
    <th>SpotCount</th>
  </tr>
  <tr>
    <td>co.uk</td>
    <td>3583</td>
  </tr>
  <tr>
    <td>facebook.com</td>
    <td>88</td>
  </tr>
  <tr>
    <td>com.au</td>
    <td>80</td>
  </tr>
</table>

Table above shows that 3583 spots have the same domain named 'co.uk'
