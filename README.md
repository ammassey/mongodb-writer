# mongodb-writer
This is a microservice written by Alan Massey for Andronik Mkrtytchev
to provide write functionality for users who would like enter a 
recipe manually in the event they did not receive a desired recipe by
the api Andronik's program integrates with.

Communication Contract:
Communication will be established with the microservice via rest API. I agree to be available in the event that Andronik is unable to integrate this
microservice into his main program. My availablity it 5PM-Midnight PST and open 
availability on weekends. Once hosted, I will provide the correct endpoint for Andronik
and troubleshoot any issue.

Posting Data:
Posting data to MongoDB by why of this microservice requires the data to be formatted as follows:
{
    "title": "",
    "ingredients": "",
    "servings": "",
    "instructions": "",
    "Password": "",
    "user": ""
}

There is no header required an example call would be as follows:
requests.post(url, data)

Response Data: 
The response that the user will receive from the server is as the text as follows:
{"_id":"","title":"","":"":"","instructions":"","password":"","user":""}

A 201 status code indicates that the write to mongoDB was successful:

response = requests.post(url, data)
print(response.status_code)
print(response.text)

