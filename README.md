# Description

Imagine a dartboard, where there are 2 possible scores each time you throw the dart, 12 or 7.

![Dartboard](https://raw.githubusercontent.com/justmathstuff/dart-unreachable-score-problem/master/IMG.jpg)

What is the highest score that is impossible for you to reach?

Hint: it's below 100

# Solution

The problem may seem hard at first, but once the pattern is discovered, it becomes very easy.

The first key point, is to realize that we can represent every single score as:

```
x = number of 7's
y = number of 12's
7x + 12y
```

Now from here, it is tempting to think about which numbers can't be represented with this expression, which I did at first, but it just led me down a rabbit hole.

The second key point is to change up our perspective. Instead of focusing on what numbers **can't** be represented, let's think about how we can use the expression to get all the numbers that **can** be represented.

First off, trying to list all the numbers at once is very time consuming and prone to errors, so a smarter way might be to find a pattern that we can use to generate these numbers.
If we look at the expression `7x + 12y`, we see that there are two variables or `knobs` that we can turn to get numbers that fit the expression, `x` and `y`. Now, if we start at a number that fit this expression, such as `0`, which can be represented with `7(0) + 12(0)`, how can we change `x` and `y` to get more numbers like this? Well, we can do either `x + 1`, `y + 1`, or both!

So what we can do to generate numbers that can be represented with the expression within the range 0 to 100, is we can start at 0, and use the method that we found to get 3 more numbers, and repeat the whole thing with those 3 numbers instead of 0, and so on and so forth, and we are good! Now, we just need to ignore all those representable numbers, and pick out the largest number that we haven't touched yet!

## Alternatively, a method that's more efficient

So the original method was to take 0, and branch off to 3 different numbers, and then to branch off again from all 3 of those to 9 more numbers. But keeping track of which few branches you need to extend is not efficient, so a better way might be to just branch off from the first number in the list which is representable, and haven't been used to branch off before (because it would be meaningless to repeat a branch off).

Example:

![Example Problem with 3 & 5](https://raw.githubusercontent.com/justmathstuff/dart-unreachable-score-problem/master/IMG1.jpg)

## Tree Diagram

![Tree Diagram](https://raw.githubusercontent.com/justmathstuff/dart-unreachable-score-problem/master/IMG2.jpg)

# Software

The `interactive/solution` binary is a solution to this problem written in the Go programming language.

Usage:

`interactive/solution --A=7 --B=12 --Range=100`