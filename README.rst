===============================
GoBozoCrack
===============================

GoBozoCrack is a depressingly effective and fast MD5 password hash cracker with almost zero CPU/GPU load written in Go. Instead of rainbow tables, dictionaries, or brute force, GoBozoCrack simply *finds* the plaintext password. Specifically, it googles the MD5 hash and hopes the plaintext appears somewhere on the first page of results.

* Free software: MIT license

It works way better than it ever should. Go ahead and try.


How?
----
Basic usage:

   **$ go run main.go -file example.txt**

Or:

    **$ go run main.go -single fcf1eed8596699624167416a1e7e122e**

The input file is expected to contain a single hash per line. GoBozoCrack automatically picks up strings that look like MD5 hashes.


Example with output:

    **$ go run main.go -file -f example.txt**
    
        fcf1eed8596699624167416a1e7e122e:octopus
    
        bed128365216c019988915ed3add75fb:passw0rd
    
        d0763edaa9d9bd2a9516280e9044d885:monkey
    
        dfd8c10c1b9b58c8bf102225ae3be9eb:12081977
    
        ede6b50e7b5826fe48fc1f0fe772c48f:1q2w3e4r5t6y



    **$ go run main.go -single fcf1eed8596699624167416a1e7e122e**

        fcf1eed8596699624167416a1e7e122e:octopus


Why?
----
To show just how bad an idea it is to use plain MD5 as a password hashing mechanism. Honestly, if the passwords can be cracked with *this software*, there are no excuses.


Who?
----
BozoCrack was originally written by Juuso Salonen (http://twitter.com/juusosalonen).

PyBozoCrack and GoBozoCrack were rewritten in Python and Go by Henrique Pereira (http://twitter.com/ikkebr).