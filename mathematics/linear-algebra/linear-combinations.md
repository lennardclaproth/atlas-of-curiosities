# Span, Basis, and Linear Independence

The previous discussion introduced vectors as coordinates, but there is a deeper way to think about those coordinates. Instead of seeing a coordinate pair like (3, -2) as just numbers, think of each number as a **scalar** that scales a **basis vector**.

In the standard 2D coordinate system, the basis vectors are:
- **î**: one unit in the x-direction
- **ĵ**: one unit in the y-direction

A vector can be built by scaling these basis vectors and adding them together. For example, (3, -2) means:

- Scale î by 3
- Scale ĵ by -2
- Add the results

This process is called a **linear combination**.

## Basis Vectors

A basis is the set of vectors that coordinates refer to when they act as scalars. Importantly, the standard basis is not the only choice. Any pair of vectors that can generate every vector in the plane can serve as a basis.

This highlights an important idea:

> Coordinates only make sense relative to a chosen basis.

Different bases produce different coordinate systems, even though they describe the same space.

## Span

The collection of all vectors that can be reached through linear combinations of a set of vectors is called their **span**.

In two dimensions:

- Most pairs of vectors span the entire plane.
- If the vectors point in the same direction, their span collapses to a single line.
- If both vectors are zero, the span is only the origin.

Span answers a central question in linear algebra:

> What vectors can be reached using only vector addition and scalar multiplication?

## Vectors as Points

When working with many vectors at once, it is often easier to think of each vector as the point at its tip rather than as an arrow.

This allows us to visualize spans more easily:

- A line represents all vectors in a one-dimensional span.
- The entire plane represents all vectors in a two-dimensional span.

## Span in Three Dimensions

The idea becomes more interesting in 3D.

Two non-parallel vectors span a flat plane passing through the origin. Every vector on that plane can be formed by scaling and adding those two vectors.

Adding a third vector creates two possibilities:

1. If the third vector already lies on the plane, the span does not grow.
2. If the third vector points in a new direction, the span expands to fill all of three-dimensional space.

Each genuinely new direction increases the dimensionality of the span.

## Linear Dependence and Independence

Sometimes a vector adds no new capability because it can already be created from the others.

When one vector can be expressed as a linear combination of the others, the vectors are called **linearly dependent**.

When every vector contributes a new dimension to the span, the vectors are **linearly independent**.

In simple terms:

- **Dependent:** at least one vector is redundant.
- **Independent:** every vector adds something new.

## The Technical Definition of a Basis

A basis is:

> A set of linearly independent vectors that span a space.

This definition captures two essential requirements:

1. The vectors can reach every point in the space (they span it).
2. None of the vectors are redundant (they are linearly independent).

Together, these properties make a basis the minimal set of building blocks needed to describe an entire space.

These ideas—linear combinations, span, basis, and linear independence—form the foundation for understanding matrices and linear transformations, which build on them in the next stage of linear algebra.

# Resources
1. [Linear combinations, span, and basis vectors | Chapter 2, Essence of linear algebra - 3Blue1Brown ](https://youtu.be/k7RM-ot2NWY)