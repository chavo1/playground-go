int checker could be used as a storage for bits. Every bit in integer value can be treated as a flag, so eventually int is an array of bits (flag). Each bit in your code states whether the character with bit's index was found in string or not. You could use bit vector for the same reason instead of int. There are two differences between them:

  - Size. int has fixed size, usually 4 bytes which means 8*4=32 bits (flags). Bit vector usually can be of different size or you should specify the size in constructor.

  - API. With bit vectors you will have easier to read code, probably something like this:
```
    vector.SetFlag(4, true); // set flag at index 4 as true

    for int you will have lower-level bit logic code:

    checker |= (1 << 5); // set flag at index 5 to true
```
- Also probably int may be a little bit faster, because operations with bits are very low level and can be executed as-is by CPU. BitVector allows writing a little bit less cryptic code instead plus it can store more flags.