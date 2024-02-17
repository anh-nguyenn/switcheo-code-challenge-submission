/_
Problem 6: Transaction Broadcaster Service
Author: Anh Nguyen
Date: 16/02/2024
_/

Designing a Simple Transaction Broadcaster Service

Overview
This service takes transaction requests, signs them, and sends them to a blockchain. It's built to be reliable, handle requests, and fix problems automatically.

How it Works

- Getting Requests: The API Gateway gets a request, checks it, and gives it to the Processor.
- Working on Requests: The Processor checks, signs, and prepares the transaction.
- Sending Out: The Sender tries to send the transaction. If it fails, it tries again later.
- Keeping Track: The system keeps track of everything and lets admins manage retries.
- Making it Strong and Able to Handle Lots
- Splitting Tasks: Divides work into smaller, manageable pieces for better handling.
- Balancing Load: Shares incoming requests to avoid overload.
- Safety Nets: Stops problems from spreading and fixes them.
- Quick Access Data: Uses memory to quickly get to often-used data.

Main parts (please refer to diagram to see follow of design):

1. API Gateway
   What it does: It's the door for requests. It checks them and lets them in.
   Jobs: Checks requests, limits how many come in, and passes good ones to the Processor.
2. Transaction Processor
   What it does: It's the heart of the service, working on the requests.
   Parts:
   - Checker: Makes sure requests are okay.
   - Signer: Puts a digital signature on the transaction.
   - Sender: Sends the signed transaction to the blockchain.
3. Sender (Broadcast Manager)
   What it does: Manages sending transactions to the blockchain and tries again if there's a problem.
   Parts:
   - RPC Talker: Talks to the blockchain to send transactions.
   - Retry Plan: Tries again if sending fails.
   - Watchdog: Keeps an eye on whether sending worked or not.
4. Data Store
   What it does: Keeps records of transactions and what happened to them.
   Books:
   Log Book: Writes down transaction details.
   History Book: Notes down tries, successes, and fails.
5. Admin Area
   What it does: Lets admins see what's happening and fix problems.
   Tools:
   - Dashboard: Shows transactions and their status.
   - Retry Button: Lets admins try again for failed transactions.
6. Backup Plans
   How it works: Makes sure the service can keep going even after problems.
   Strategies:
   - Queue: Keeps requests in line so none are lost if there's a problem.
   - Remember State: Remembers what it was doing after a restart.
   - Problem Fixers: Handles unexpected errors well.
