You have a real-time fraud-detection system that accepts transactions and returns `true` or `false` if transaction is fraudulent or not.
Transaction class
```js
class Transaction {
  // a UUID of transaction
  string ID;
  // in USD, typically a value betwen 0.01 and 1000 USD. 
  decimal Amount;
  // a 2-letter country code of where the bank is located
  string BankCountryCode;
}
```
You have a method that accepts a collection of transactions

```js
function processTransactions(transactions []Transaction){
  results = []
  foreach transaction in transactions {
    result = {id: transaction.ID}
    result.fraudulent = processTransaction(transaction)
    results.push(result)
  } 
  return results
}
```
The `processTransactions` method is called every second and receives an array of transactions;

Once your code processes all of them, it returns an array of results (id and fraudulent or not)
To verify a specific transaction, you need to call an existing function that will internally call the bank API and verify the transaction using this API.

Calling a bank API can be very fast (1-10ms) or very slow (200-300ms) depending on where the bank is located.

If the bank is somewhere in the US - the API call will take very little time.

If the bank is somewhere in Australia, the call will take more time.

Over time, you start noticing that the number of transactions pushed to your `processTransactions` method every second becomes too big,
and you can no longer process all transactions in one second. 

You need to write a function, that will prioritize the transactions based on their `Amount`. 

The more money you handle, the better, the goal of this exercise is to maximize the total USD value your function processes in 1 second.

You can only process transactions sequentially, i.e. you cannot spawn additional threads or routines. 

**Task:**

Write a `prioritize` function body, which returns a subset of `transactions` that have the maximum total USD in a given `totalTime`.
For example, if the `totalTime` parameter is 1000ms, the total time of the prioritized transactions should be less than or 
equal to 1000ms.

```js
// function should return a subset (or full array)
// that will maximize the USD value and fit the transactions under 1 second
function prioritize(transaction []Transaction, totalTime=1000): []Transaction {
}
```

**Question:**
What is the max USD value that can be processed in 50ms, 60ms, 90ms, 1000ms?

**Instructions:**
Use the api_latencies.json to tell how long each API call takes.

Use the transactions.csv to calculate the total USD values for each question.

Implement the solution as a private GitHub repository and invite `borislobanov92`, `IvanHR21`, `GeorgeFingerprint`, `kimdrip`, and `penyaev` GitHub users to it. Make sure to include answers to all questions in the repo's readme.

Take as much time as you want to get back to us with a solution, but don't spend more than 4 hours on the actual implementation after you started working on it.

You can use any language for the implementation, but our preference is TypeScript or Golang.

**Notes**:
For simplicity, please ignore the time to run the prioritization itself, your goal is to fit the maximum USD transaction amount into the allotted time slots.

**Please briefly describe your algorithm and why you think it's optimal. If you had multiple solutions in mind, why did you choose this particular one?**