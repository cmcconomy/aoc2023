# Day 03

This was a fun one - I like the strategy I came up with, which was to store for the engine:
1. the plain engine "diagram",
2. a "mask" storing adjacent digits, and
3. a map which maps a mask-id to its parsed integer value

With this information I was able to build a series of functions that went from parsing the engine input, to finding symbols, to finding adjacent masks, to finding adjacent numbers.

Part 2 was gratifying because it required me to identify only a certain 'symbol', which required a very minor behaviour flag in the original "find symbols" function.

Since I had broken this into parts, I could use the intermediary "find symbols" function to interrogate the number of adjacent masks/ints to the special symbols, and only keep the ones that were requested.

Overall another fun one.