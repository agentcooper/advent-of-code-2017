### Part 1

Quite easy!

### Part 2

Spent too much time on this!

I solved part 2 by automatically translating input assembly into Go program with `goto` statements, see `GoProgram` function. Output of that can be found in `generated-raw/generated-raw.go`. After that I manually optimized it by translating `goto` into `if` and `for` statements, see `generated-optimized/generated-optimized.go`.

After optimization the gist of the code becomes clear and it is easy to come up with a new optimized program. I was too lazy to do that, so my final program runs for about 10 minutes.
