# Binary
There are different base systems, base-10 is the most commonly used system by us humans. Base-10 is the digits represented by 0 through 9 (also called decimal). However in the computer science world we mostly use Base-2 which is a 0 or a 1. To calculate the potential options within a base you can do 

The maximum number of unique options you can represent with $n$ digits in base-$b$ is:

$$
\text{Options} = b^n
$$

For example, in base-10 with 3 digits:

$$
10^3 = 1000 \text{ options (from 000 to 999)}
$$

in the binary system to represent a byte we can thus conclude
$$
2^8=256 \text{ (a byte consists of 8 bits)}
$$

Counting binary goes from right to left. A couple examples below:
$$
000=0 \\
001=1 \\
010=2 \\
011=3 \\
100=4 \\
101=5 => 4+1\\
110=6 => 4+2\\
111=7 => 4+2+1\\
\text {And so on}
$$

So now knowing this we can take a short dive into ASCII, ASCII are characters that are assigned a number which van be represented in binary. However ASCII has only 8 bits which is a byte which have 256 options (2^8). For english this is fine but suppose we have to support other languages like chinese or the cyrillic script or sanskrit we run into some issues considering that 256 options are simply not enough.

This introduced unicode which rests on 32 bits which is also backwards compatible with ASCII.

We understand how text works but how does it then work with color? Let's take for example RGB. Every color can be mapped to a set of 3 bytes (171, 255, 111), where (R, G, B), this will give you a color (for the sake of readability the values are written in base-10)