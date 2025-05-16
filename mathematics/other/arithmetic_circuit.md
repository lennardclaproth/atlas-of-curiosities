# Arithmetic circuit

An arithmetic circuit $C$ is a way of computing a polynomial. it does this as follows:

- An input gate computes the polynomial it is labeled by
- A sum gate ($+$) $v$ computes the sum of the polynomials computed by its children (a gate $u$ is a *child* of $v$ if the directed edge $(v,u)$ is in the graph)
- A product gate ($\times$) gate computes the product of the polyonmials computed by its children.

![alt text](./img/ArithmeticCircuit.svg "figure: Directed acyclic graph")

The image above represents a Directed acyclic graph. If we follow the steps explained in the beginning and compute the graph from left to right on the input gates we will get the following:

The sum gates will compute as follows:

$x_1+x_2+1$

the product gates will compute as:

$(x_1+x_2)x_2(x_2+1)$

Now that we grasp the basic concept of an Arithmetic circuit it is important to emphasize some of the boundaries in which it functions.

An arithmetic circuit can only be defined over the field $\mathbb{F}$ and the set of variables $x_1,...,x_n$[2]

A circuit has two complexity measres associated with it: size and depth. the *size* f a circuit is the number of gates in it, the *depth* of a circuit is the length of the longest directed path in it. For exmple, the circuit in the figure has size six and depth two.

## Resources

1. [Arithmetic circuit complexity](https://en.wikipedia.org/wiki/Arithmetic_circuit_complexity)
2. [Field (mathematics)](https://en.wikipedia.org/wiki/Field_(mathematics))
3. [Directed acyclic graph](https://en.wikipedia.org/wiki/Directed_acyclic_graph)
