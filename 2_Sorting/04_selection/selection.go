package main

// Proposition A. Selection sort uses ~N2/2 compares and N exchanges to sort an array of length N.
// Proof: You can prove this fact by examining the trace, which is an N-by-N table in which unshaded letters correspond to compares. About one-half of the entries in the table are unshaded—those on and above the diagonal. The entries on the diagonal each correspond to an exchange. More precisely, examination of the code reveals that, for each i from 0 to N − 1, there is one exchange and N − 1 − i compares, so the totals are N exchanges and (N − 1) + (N − 2) + . . . + 2 + 1+ 0 = N(N − 1)/2 ~ N2/2 compares.

