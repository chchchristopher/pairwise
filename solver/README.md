# pairwise_solver

## What it does

This takes in a (currently hard-coded) set of arrays, and permutes them such
that every element appears in a column exactly once.

## Notes

Uncomment `Enumerate all results` and wait for 3 billion years.

## Todo

- Reduce the solution space by (order+1)! when enumerating all results, each time a solution is found by eliminating all possibilties of the form of shifting entire columns of the set. 
- Move to a Go setting to avoid writing and reading
- In the meantime, if the `solver` flag is `true`:
- - PPmaker should output to a text file
- - The python solver should read input from a text file, perform solving, and
    export to a text file
- - `pairwise` should then pick up and do things to it (send it to `hexmaker`
    or whatever future card/tile creation
- - Probably all organized by a shell script/TUI/GUI

## Credits

thanks to a random redditor for this, and also to the googles, I wrote almost
none of it