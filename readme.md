# Payment-API

### How To Run Application
1. Run the sql query in "config/database" init.sql file to our mysql.
######
2. Run the program by typing command "go run ." in the terminal.
######
3. Open our postman and import file "ASSESTMENT_MNC.postman_collection.json".
######
4. After importing the file, click the collection "Assestment_MNC".
######
5. After that, the first part we click on is the register request, choose Body, fill in the json data and click the send button.
######
6. After the data is registered, we can select Login request, choose Body, fill in the username and password data, and then select the send button. After that copy the "token" response.
######
7. After the login and token are copied, select the Customer folder and click "Get All Customers". Then select "Authorization" and fill in the token field with the token copied earlier and click the send button. Then a response will appear, copy the customer "id" data as the sender and receiver in the response.
######
8. After that, select the Transaction folder and click "Create New Transaction". Then select "Authorization" and fill in the token field with the token copied earlier. Then select "Body" also fill in the "id" copied earlier in "SenderId" as the sender and "RecieverId" as the reciever, then also fill in "Amount". Then select the send button.
######
9. And to log out, we only have to wait 30 minutes for our token to expire.