## Day 
this was a bit nasty. I decided I didnt want to do what I would usually do and iterate over the whole shit. So I went for some regex and extracted all the numbers and symbols from that.

### Part 1
I didnt want to split the lines into slices either so I just calculated the 
length of a row and worked from that. I really dislike these kind of puzzles which seem quite similar every aoc, where you just have to check every adjacent piece. Just ends up with a bunch of for loops.

Anyway, I had fun figuring out how to do it without iterating everything

### Part 2
Used the same logic from the first part, but iterating from the perspective of the symbols instead of numbers, and then adding the matched numbers to a map to avoid duplicates and then multiplying and summing up
