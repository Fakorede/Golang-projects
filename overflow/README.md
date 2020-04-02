# Overflow

## Explanation

This program dives deep into how Go behaves when a value goes beyond the range of its declared type i.e. an Overflow

> But Go cannot catch Overflow errors in runtime

Instead, the values _wraps around_ to the minimum negative or maximum positive depending on the type.

**Signed Types**
|-------------------|--------------|
| Minimum value - 1 | Max Positive |  
| Maximum value + 1 | Min Negative |

Example for **int8(-128 to 127)**
|----------|-------|
| -128 - 1 | 127 |
| 127 + 1 | -128 |

**Unsigned Types**
|-------------------|--------------|
| Minimum value - 1 | Max Positive |
| Maximum value + 1 | Zero |

Example for **uint8(0 to 255)**
|----------|-------|
| 0 - 1 | 255 |
| 255 + 1 | 0 |
