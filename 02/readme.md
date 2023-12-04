## Day 
Pretty nice task also. Always a blast with Regex, use regexr.io to try it out though. Also had to look up some of the strings. functions, eg cutPrefix.

### Part 1
Not much to say, went with the first solution that popped in my mind. I decided to just calculate the gameID instead of reading it, that way I could discard the first part of the lines right away. Also added a map here, maps make stuff easier to do fast.

### Part 2
This was a fairly easy expansion of the first part. Instead of looking if any of them crossed the threshold I made sure to store the highest one and then multiply them as I was iterating through.
