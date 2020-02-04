# webapptester
Command line tool that generates boilerplate unit tests for your http handlers. Get a head start in writing your unit tests with these auto generated (Table Driven) tests. Now this doesn't completely unit test your code but you can easily add more test cases and variables to your test structure. This even parses your file for Mux router variables, and sets them up in the table for you! Each handler is tested separately from the router, so the unit is truely being tested on its own.  

## motivation
At one of my previous internships, I built a lot of endpoints in Go. So naturally this meant I wrote a lot of unit tests for these endpoints. What I realized when writing these unit tests was that there is a lot of repetitive actions in building these tests. Tasks like creating a http request for your test case, setting up all the necessary variables to run your test, and just structuring the test cases in general.. seemed very repetitive.  

## how-to 
This project has no outside dependencies (other than Go)

`go get github.com/yaoalex/webapptester`
`go install`
