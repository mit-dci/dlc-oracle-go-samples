# Simple command-line example oracle

This simple command-line oracle will generate and store a new password protected private key. It will then print out the Oracle's public key, and then enter a loop:

* Generate new one-time signing scalar
* Print out the equivalent public key (R-point)
* Ask for a numeric value to sign
* Sign the value with the scalar and print out the signature

You can use this oracle to easily test Discreet Log Contracts. You can form a contract based on the public key printed out at the start, and the R-point printed out in the loop. Then using the value and signature you can settle the contract.

For examples of running a Discreet Log Contract using LIT, see...