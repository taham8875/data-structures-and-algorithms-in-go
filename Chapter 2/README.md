# Algorithms Analysis

# Algorithm

An algorithm is a set of instructions to solve a problem or perform a task. It is a sequence of steps that are followed to solve a problem.

The most important property of an algorithm is **correctness**. An algorithm should always produce the correct result and **efficiency**.

Efficiency is measured in two parameters, **time** and **space**. Time is the number of steps required to solve a problem and space is the amount of memory required to solve a problem.

# Asymptotic Analysis

Asymptotic analysis is a way of analyzing the efficiency of an algorithm independently of any particular dataset of programming language.

We are interested in the order of growth of the algorithm not the exact time or space required to execute the algorithm.

# Big-O Notation

Big-O notation is a way of expressing the efficiency of an algorithm. It is the upper bound of the function that represents the efficiency of an algorithm.

Common Big-O notations are:

- O(1) - Constant Time
- O(log n) - Logarithmic Time
- O(n) - Linear Time
- O(n log n) - Linearithmic Time
- O(n^2) - Quadratic Time
- O(n^3) - Cubic Time
- O(n^c) - Polynomial Time
- O(c^n) - Exponential Time
- O(n!) - Factorial Time (or O(n^n))

## Constant Time - O(1)

An algorithm is said to have constant time if the time required to execute the algorithm is the same regardless of the size of the input data.

example:

- Accessing the nth element of an array
- Push and pop operations on a stack
- Accessing an element in a hash table

## Logarithmic Time - O(log n)

An algorithm is said to have logarithmic time if the time required to execute the algorithm is proportional to the **logarithm** of the input size.

So each step of an algorithm reduces the size of the input by a significant amount.

example:

- Binary Search

## Linear Time - O(n)

An algorithm is said to have linear time if the time required to execute the algorithm is proportional to the input size.

example:

- Sequential Search

## Linearithmic Time - O(n log n)

An algorithm is said to have linearithmic time if the time required to execute the algorithm is proportional to the **product** of the input size and the logarithm of the input size.

example:

- Merge Sort
- Heap Sort
- Quick Sort (average case)

## Quadratic Time - O(n^2)

An algorithm is said to have quadratic time if the time required to execute the algorithm is proportional to the **square** of the input size.

example:

- Bubble Sort
- Insertion Sort
- Selection Sort

# Deriving the Runtime Function of an Algorithm

## Constants

Constants are ignored in Big-O notation.

for example, if an algorithm takes 3n + 2 steps to execute, then the runtime function is O(n).

If each step takes a constant time, them the time complexity of the algorithm is O(1).

## Loops

The time complexity of a loop is the length of the loop times the time complexity of the statements inside the loop.

for example, if a loop executes n times and each step takes a constant time, then the time complexity of the loop is O(n).

## Nested Loops

The time complexity of nested loops is the product of the time complexities of all the loops.

for example, if a loop executes n times and inside it another loop executes n times, then the time complexity of the nested loops is O(n^2).

## Consecutive Statements

The time complexity of consecutive statements is the maximum of the time complexities of the statements.

for example, if two statements take O(n) and O(n^2) time respectively, then the time complexity of the two statements is O(n^2).

## If-Else Statements

The time complexity of an if-else statement is the maximum of the time complexities of the if and else blocks. (remember we measure the worst case)

for example, if the if block takes O(n) time and the else block takes O(n^2) time, then the time complexity of the if-else statement is O(n^2).

## Logarithmic Time

If in each iteration tje input size is reduced by a constant factor, then the time complexity of the algorithm is O(log n).
