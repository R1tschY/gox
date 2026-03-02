# iterx

*Extensions for [iter.Seq](https://pkg.go.dev/iter#Seq)*

[![PkgGoDev](https://pkg.go.dev/badge/github.com/r1tschy/gox/iterx)](https://pkg.go.dev/github.com/r1tschy/gox/iterx)
![GitHub License](https://img.shields.io/github/license/r1tschy/gox)
[![Go Report Card](https://goreportcard.com/badge/github.com/r1tschy/gox)](https://goreportcard.com/report/github.com/r1tschy/gox)

Inspired by [Python's itertools] and [C++'s &lt;algorithm&gt;].

[Python's itertools]: https://docs.python.org/3/library/itertools.html
[C++'s &lt;algorithm&gt;]: https://cppreference.com/w/cpp/algorithm.html

## Generators
| Generator   | Arguments   | Results                 | Example               |
|-------------|-------------|-------------------------|-----------------------|
| [Count]     | start, step | start, start+step, ...  | `10 5 ➞ 10 15 20 ...` |
| [Empty]     |             | (nothing)               | `➞ (empty)`           |
| [Infinite]  |             | (), (), ...             | `➞ () () () ...`      |
| [Range]     | start, end  | start, ..., end-1       | `0 3 ➞ 0 1 2`         |
| [Repeat]    | val, n      | val, val, ... (n times) | `10 3 ➞ 10 10 10`     |
| [SeqOf]     | values      | values                  | `1 2 3 ➞ 1 2 3`       |
| [Singleton] | val         | val                     | `10 ➞ 10`             |

## Transformations
| Iterator       | Arguments   | Results                 | Example                            |
|----------------|-------------|-------------------------|------------------------------------|
| [Accumulate]   | p, fn       | p0, p0+p1, p0+p1+p2...  | `[1,2,3] ➞ 1 3 6`                  |
| [Chain]        | p, q        | p0, ..., pN, q0, ... qN | `[1,2] [3,4] ➞ 1 2 3 4`            |
| [ChainFrom]    | sequence    | p0, ..., pN, q0, ... qN | `[[1,2], [3,4]] ➞ 1 2 3 4`         |
| [Chunk]        | p, n        | (p0, ... pN-1), ...     | `[1,2,3,4,5] 2 ➞ [1,2] [3,4] [5]`  |
| [Cycle]        | p           | p0, ..., pN, p0, ...    | `[1,2] ➞ 1 2 1 2 ...`              |
| [DropWhile]    | p, pred     | pN, pN+1, ...           | `[1,2,3,2,1] <3 ➞ 3 2 1`           |
| [Enumerate]    | p, start    | (start, p0), ...        | `['a','b'] 0 ➞ (0,'a') (1,'b')`    |
| [Filter]       | p, pred     | elements of p if pred   | `[1,2,3] isEven ➞ 2`               |
| [GroupBy]      | p, key      | (k, sub-p), ...         | `[1,2,2,1] ➞ (1,[1]) (2,[2,2])...` |
| [Limit]        | p, n        | p0, p1, ... pN-1        | `[1,2,3,4] 2 ➞ 1 2`                |
| [Map]          | p, fn       | fn(p0), fn(p1), ...     | `[1,2,3] x*2 ➞ 2 4 6`              |
| [Pairwise]     | p           | (p0, p1), (p1, p2), ... | `[1,2,3,4] ➞ (1,2) (2,3) (3,4)`    |
| [Repeat]       | val, n      | val, val, ... (n times) | `10 3 ➞ 10 10 10`                  |
| [Slice]        | p, s, e     | pS, ... pE-1            | `[1,2,3,4] 1 3 ➞ 2 3`              |
| [TakeWhile]    | p, pred     | p0, p1, ... until !pred | `[1,2,3,2,1] <3 ➞ 1 2`             |

## Reductions
| Consumer       | Arguments | Results                    | Example                   |
|----------------|-----------|----------------------------|---------------------------|
| [All]          | p, pred   | true if all pred           | `[2, 4, 6] isEven ➞ true` |
| [Any]          | p, pred   | true if any pred           | `[1, 2, 3] isEven ➞ true` |
| [CollectSlice] | p         | slice of p                 | `1 2 3 ➞ [1, 2, 3]`       |
| [CountIf]      | p, pred   | number of pred             | `[1, 2, 3] isEven ➞ 1`    |
| [First]        | p         | first element              | `1 2 3 ➞ 1`               |
| [Len]          | p         | length of p                | `1 2 3 ➞ 3`               |
| [None]         | p, pred   | true if none pred          | `[1, 3, 5] isEven ➞ true` |
| [Reduce]       | p, i, fn  | fn(...fn(i, p0), pN)       | `[1,2,3] 0 x+y ➞ 6`       |
| [Single]       | p         | single element or not okay | `1 ➞ 1, true`             |


## Combinatoric
| Iterator       | Arguments   | Results                                                          | Example                            |
|----------------|-------------|------------------------------------------------------------------|------------------------------------|
| [Combinations] | p, r        | r-length sequences, in sorted order, no repeated elements        | `[1,2,3] 2 ➞ [1,2] [1,3] [2,3]`    |
| [Permutations] | p, r        | r-length sequences, all possible orderings, no repeated elements | `[1,2,3] 2 ➞ [1,2] [1,3] ...`      |
| [Product]      | p, q        | cartesian product                                                | `[1,2] [3,4] ➞ [1,3] [1,4] ...`    |


[Accumulate]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Accumulate
[Chain]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Chain
[ChainFrom]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#ChainFrom
[Chunk]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Chunk
[Combinations]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Combinations
[Count]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Count
[Cycle]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Cycle
[DropWhile]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#DropWhile
[Empty]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Empty
[Enumerate]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Enumerate
[Filter]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Filter
[GroupBy]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#GroupBy
[Infinite]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Infinite
[Limit]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Limit
[Map]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Map
[Pairwise]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Pairwise
[Permutations]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Permutations
[Product]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Product
[Range]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Range
[Repeat]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Repeat
[SeqOf]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#SeqOf
[Singleton]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Singleton
[Slice]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Slice
[TakeWhile]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#TakeWhile
[All]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#All
[Any]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Any
[CollectSlice]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#CollectSlice
[CountIf]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#CountIf
[First]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#First
[Len]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Len
[None]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#None
[Reduce]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Reduce
[Single]: https://pkg.go.dev/github.com/R1tschy/gox/iterx#Single