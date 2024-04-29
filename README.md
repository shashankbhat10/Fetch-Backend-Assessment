# Fetch-Backend-Assessment

Backend assessment for Fetch Rewards

# How to run

To run the project, go to root folder of project and run the following command: </br>

<pre>go run ./main.go</pre>

# Error scenario

Since begaviour for incorrect scenario isnt mentioned, following are the cases where error response is returned:

<ul>
<li>No request body for POST /receipts</li>
<li>Incorrect JSON in request body for POST /receipts</li>
<li>ID parameter is blank for GET /receipts/{id}/points</li>
<li>Incorrect ID format for GET /receipts/{id}/points</li>
<li>ID for receipt not found in GET /receipts/{id}/points</li>
<li>Price has incorrect format in GET /receipts/{id}/points</li>
<li>Invalid date or time format in GET /receipts/{id}/points</li>
</ul>

# Assumptions

1. Time is in HH:MM format and 12-hour clock
2. Date is in YYYY-MM-DD format
3. In-memory map is sufficient for this assessment

# Author

<b>Name:</b> Shashank Bhat </br>
<b>Email:</b> bhatshashank94@gmail.com
