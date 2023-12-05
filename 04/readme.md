## Day 
Thought process was basically that it seemed to be approachable same way as previous day; using regex. I realized fairly fast I could split up the strings quite easily. And then I extracted numbers from both sides and matched in 1 direction.

I like the idea of treating the input as data instead of just a big pile of text. So instead of iterating through it and looking forward I like to put the numbers into slices and then deal with them from there.

### Part 1
Decided to disregard any card id because it wasnt useful here. Also was wondering if there was some fancy binary stuff I could do to, but ended up just making the incrementing function instead

### Part 2
Had to re-do a bit because I needed to structure my loop so that I could re-run it when necessary. I was contemplating if it is possible to some kind of stack or queue but it would require me to keep storing the card IDs and I didnt want to do that. Just wanted to flow through it. My program in the end took maybe a minute or 2 to run so probably not very optimized
