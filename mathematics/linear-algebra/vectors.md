# The Story of Vectors: The Foundation of Linear Algebra

Linear algebra begins with one fundamental idea: **the vector**.

Different fields view vectors differently:
- **Physicists** see vectors as arrows with a length and direction.
- **Computer scientists** see vectors as ordered lists of numbers.
- **Mathematicians** define vectors more generally as objects that can be added together and scaled by numbers.

For intuition, it is useful to think of a vector as an arrow starting at the **origin** of a coordinate system. Its coordinates describe how to reach its tip. In 2D, this requires an x-value and a y-value; in 3D, a z-value is added. Every vector corresponds to exactly one set of coordinates, and every set of coordinates corresponds to exactly one vector.

Two operations form the core of linear algebra:

## Vector Addition
Vector addition combines movements. If one vector represents a step and another represents a second step, their sum represents the overall movement. Geometrically, this is visualized by placing the tail of the second vector at the tip of the first. Numerically, it means adding corresponding coordinates.

## Scalar Multiplication
Scalar multiplication changes a vector's size. Positive scalars stretch or shrink a vector, while negative scalars also reverse its direction. Numerically, each coordinate is multiplied by the scalar.

## Why It Matters
The power of linear algebra comes from translating between two perspectives:

- **Geometry:** vectors as arrows in space.
- **Algebra:** vectors as lists of numbers.

This translation allows us to visualize data, describe physical space, create computer graphics, and perform computations efficiently.

Everything else in linear algebra—such as span, bases, linear dependence, and linear transformations—builds on these two fundamental operations: **vector addition** and **scalar multiplication**.

## Resources
1. [Vectors | Chapter 1, Essence of linear algebra - 3Blue1Brown](https://youtu.be/fNk_zzaMoSs) 