# GO TodoList Microservice
This is a very simple boilerplate Go "Microservice" for the Peak Activity "Todo List Challenge". It has everything you need (except for a MongoDB driver) to create the Todolist Microservice needed to complete the "Todo List Challenge".

### Setup
Create an `.env` file with the PORT you would like the service to run on.
```
PORT=8888
```

### Requirements
1. You must use MongoDB as a Database for this microservice. You can get a free MongoDB database @ [mlab.com](http://mlab.com).

### What needs to be done
Below are the tasks that need to be completed in order to finish this Microservice.

#### 1. Endpoint for Adding Items
- Must add item passed through http.Request into a MongoDB collection.
- If the item was added to the MongoDB collection successfully, you must respond with the item and HTTP Status 200.
- If any errors occur, you must respond with a error message and HTTP Status 401.
#### 2. Endpoint for Editing Items
- Must update the existing item in the MongoDB collection with the item passed through http.Request.
- If the item was updated in the MongoDB collection successfully, you must respond with the item and HTTP Status 200.
- If any errors occur, you must respond with a error message and HTTP Status 401.
#### 3. Endpoint for Deleting Items
- Must delete the existing item in the MongoDB collection with the item identifier passed through http.Request.
- If the item was deleted from the MongoDB collection successfully, you must respond with HTTP Status 200.
- If any errors occur, you must respond with a error message and HTTP Status 401.
#### 4. Endpoint for Listing Items
- Must get all of the existing item in the MongoDB collection.
- If the items were retrieved successfully, you must respond with the items and HTTP Status 200.
- If any errors occur, you must respond with a error message and HTTP Status 401.

### What needs to be done after completion
After you complete the above, you can query this microservice by completing the [Todo List React App](https://github.com/PeakActivity/react-todolist-challenge) 

### Submission
Please commit the microservice to a new branch, once completed. After doing so, we will review the microservice and get back to you shortly.