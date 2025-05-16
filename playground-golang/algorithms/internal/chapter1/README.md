# Chapter 1 [p1 -p 19]

## Logarithmic scale
What is a logarithmic scale. A logarithmic is the inverse of an exponential. Where 10^2=100 if we want to figure out how many 10s it would take to get to a 100 we write the following: log10 100=2
- 2^3=8 -> log2 8 = 3
- 2^5=32 -> log2 32 = 5
And so on.

## Big O notation
The big O notation is the time complexity notation. Different algorithms grow in running times at different rates. So given the following, we use a simple search to loop over 100 elements, in the big O notation this would be O(n) where n is the number of elements. This is linear, as n grows O grows in a similar fashion. 

Now take the Binary search algorithm, O(logn), as we have seen above we can state that the number of iterations it takes to find the value is the inverse of the power of 2 (log2). therefore we can state that the time complexity of the binary search algorithm is O(logn) where n is the number of elements.

For example: simple search O(n) vs binary search O(logn)

1 million iterations
O(1000000) = 1000000 iterations
O(log2(1000000)) = 20 (rounded upwards)

here we can see the benefit of this.

Some common big O run times are:
- O(log n), also known as log time
- O(n), also known as linear time
- O(n * log n)
- O(n^2)
- O(n!)

Concluding this we can state the following:
- Algorithm sped isnt measured in secods but in growth of the number of operations
- Instead of secods we talk about how quickly the run time of an algorithm increases

## Exercises

### 1.1
[Q] suppose you have a sorted list of 128 names, and you're searching through it using binary search. What's the maximum number of steps it would take.

[A] If you split the list in half every time you search you have the inverse of an exponential therefore we van state the following:

c = 128 => count

p = 2 => parts

log p(c) => log2(128) = 7

### 1.2
[Q] Suppose you double the size of the list. What's the maximum number of steps now

[A] log2(256) = 8

### 1.3
[Q] You have a name and you want to find the persons phone number in the phonebook

[A] O(logn)

### 1.4
[Q] You have a pohnenumber and you want to find the persons name in the phone book (hint you'll have to search the whole phone book)

[A] O(n)


