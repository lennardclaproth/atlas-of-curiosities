# What is a snark

A snark is a succinct proof that a certain statemen is tru.

Example statement: "I know an $m$ such that $\text{SHA256}(m) = 0$"

A **SNARK** is a *proof* that is short and *fast* to verify

It's actual definition is: **Zero-knowledge Succinct Non-Interactive Argument of Knowledge**

So say for example we have $m$ which is 1GB in size. To verify that the statement above is true we can use the trivial way:

- send the 1GB data
- Hash it with SHA256
- proof that it is zero.

The issue with this however is that this will take more time, we want to reduce this.

a **zk-SNARK** is a proof the "reveals nothing" about $m$

The advantage of this is that you can use this to create private transaction on a blockchain. This is because it can be verified yet the actual data is not exposed.

Another application of a zk-SNARK is Compliance. i.e. a private proof of solvency and compliance (verifying that a customer complies with a non public transaction of less than 10000 dollars withou actually exposing the amount).

## Cryptography

timestamps of video:

- Arithmetic circuits 6min
- Argument systems 9min

### Arithmetic circuit

*from [Arithmetic circuit](../mathematics/other/arithmetic_circuit.md)*

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

#### Directed Acyclic Graph

An arithmetic circuit is represented as a DAG:

- Vertices (nodes): inputs, gates, and output nodes.
- Edges: represent the flow of value from inputs through gates, to outputs
- Acyclic: there are no cycles, computation flows strictly forwards.

Gates can be multiple different operators,

an example of a graph

```md
    x      y
     \    /
      + (s = x + y)
       \
        * (f = s * s)
```

this will result in $f=(x+y)(x+y)$ which in turn results in the polynomial:

$$
(x+y)^2=x^2+2xy+y^2
$$

We come to this via:
$$
(x+y)^2=(x+y)\times(x+y)
$$
Apply the distributive property (FOIL method if you prefer):
We multiply each term in the first parenthesis by each term in the second:

$x \times x = x^2$

$x \times y = xy$

$y \times x = yx$ (but $yx = xy$ by commutativity)

$y \times y = y^2$

Which will give us:
$$
    x^2+xy+xy+y^2
$$

And results into:
$$
x^2+2xy+y^2
$$

## Argument system

- prover, verifier
- non interactive argument system

## Resources

1. [ZK Whiteboard Sessions - Module One: What is a SNARK? by Dan Boneh](https://youtu.be/h-94UhJLeck)
2. [Arithmetic circuit complexity](https://en.wikipedia.org/wiki/Arithmetic_circuit_complexity)
3. [Field (mathematics)](https://en.wikipedia.org/wiki/Field_(mathematics))
4. [Proofs, Arguments and Zero-Knowledge]
