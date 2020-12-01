# Day 1 

### Inital Thoughts - a
> *Disclaimer: Never done AoC so my thought process is probably wrong and all over the place :joy:

Go is cool but takes work. 

I need to parse the .txt file to get all the numbers out. The parse will return a 'slice' type. Since it won't limit the size - which might not be useful for day 01 but I could reuse it on another day. 

I wonder if its worth ordering the the numbers. After this there are two options: 
* Brute force: Get a number -> add to a different number -> check if 2020 -> if no then try a different number (go through all of the slice and then switch the 'first' number) -> if yes return the two numbers
* Smart way: No idea yet 

Then the rest is easy. 
