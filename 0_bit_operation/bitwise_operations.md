# Bitwise operations
- [ ] [Bits cheat sheet](bits-cheat-sheet.pdf) - you should know many of the powers of 2 from (2^1 to 2^16 and 2^32)
- [ ] Get a really good understanding of manipulating bits with: &, |, ^, ~, >>, <<
    - [ ] [words](https://en.wikipedia.org/wiki/Word_(computer_architecture))
    - [ ] [Binary To Decimal](https://www.rapidtables.com/convert/number/how-binary-to-decimal.html)
      - Shifting all of a number's bits to the left by 1 bit is equivalent to multiplying the number by 2. Thus, all of a number's bits to the left by n bits is equivalent to multiplying that number by 2^n.
      - Right shifts are equivalent to dividing a number by 2 
    - [ ] Good intro:
      [Bit Manipulation (video)](https://www.youtube.com/watch?v=7jkIUgLC29I)
        - CHECK IF EVEN
            ```python
              (x & 1) == 0
            ```
        - CHECK IF POWER OF TWO
            ```python
              (x & x-1) == 0
            ```
    - [ ] [C Programming Tutorial 2-10: Bitwise Operators (video)](https://www.youtube.com/watch?v=d0AwjSpNXR0)
    - [ ] [Bit Manipulation](https://en.wikipedia.org/wiki/Bit_manipulation)
    - [ ] [Bitwise Operation](https://en.wikipedia.org/wiki/Bitwise_operation)
    - [ ] [Bithacks](https://graphics.stanford.edu/~seander/bithacks.html)
    - [ ] [The Bit Twiddler](https://bits.stephan-brumme.com/)
    - [ ] [The Bit Twiddler Interactive](https://bits.stephan-brumme.com/interactive.html)
    - [ ] [Bit Hacks (video)](https://www.youtube.com/watch?v=ZusiKXcz_ac)
    - [ ] [Practice Operations](https://pconrad.github.io/old_pconrad_cs16/topics/bitOps/)
- [ ] 2s and 1s complement
    - [Binary: Plusses & Minuses (Why We Use Two's Complement) (video)](https://www.youtube.com/watch?v=lKTsv6iVxV4)
    - [1s Complement](https://en.wikipedia.org/wiki/Ones%27_complement)
    - [2s Complement](https://en.wikipedia.org/wiki/Two%27s_complement)
- [ ] Count set bits
    - [4 ways to count bits in a byte (video)](https://youtu.be/Hzuzo9NJrlc)
    - [Count Bits](https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan)
    - [How To Count The Number Of Set Bits In a 32 Bit Integer](http://stackoverflow.com/questions/109023/how-to-count-the-number-of-set-bits-in-a-32-bit-integer)
- [ ] Swap values:
    - [Swap](https://bits.stephan-brumme.com/swap.html)
- [ ] Absolute value:
    - [Absolute Integer](https://bits.stephan-brumme.com/absInteger.html)
