# Linear Transformations and Matrices

Among all topics in linear algebra, one idea helps everything else start to make sense:

**Matrices describe linear transformations of space.**

Understanding this connection turns matrix multiplication from a memorized procedure into something intuitive and visual.

## What Is a Linear Transformation?

A transformation is simply a function that takes a vector as input and produces another vector as output.

Instead of thinking about individual vectors moving, it is often easier to imagine every point in space moving at once. The result looks like space itself is being stretched, rotated, squished, or sheared.

A transformation is **linear** if it satisfies two visual rules:

1. Straight lines remain straight.
2. The origin stays fixed.

Linear transformations preserve the structure of space, keeping grid lines parallel and evenly spaced.

## The Key Insight

To describe an entire linear transformation, you do not need to know what happens to every vector.

You only need to know what happens to the two basis vectors:

- **î** = (1, 0)
- **ĵ** = (0, 1)

Every vector can be written as a linear combination of these basis vectors.

For example:

v = x·î + y·ĵ

Because linear transformations preserve linear combinations, the transformed vector becomes:

v' = x·T(î) + y·T(ĵ)

This means that once you know where î and ĵ land, you automatically know where every vector lands.

## Matrices as Encoded Transformations

A matrix is simply a compact way to record where the basis vectors move.

For a matrix:

[A B]
[C D]

- The first column (A, C) tells where î lands.
- The second column (B, D) tells where ĵ lands.

The columns are not arbitrary numbers—they are transformed basis vectors.

This is the geometric meaning of a matrix.

## Matrix–Vector Multiplication

When a matrix acts on a vector (x, y), the result is found by:

1. Scaling the first column by x.
2. Scaling the second column by y.
3. Adding the results.

This produces:

(Ax + By, Cx + Dy)

Viewed geometrically, matrix multiplication is simply constructing the correct linear combination of the transformed basis vectors.

## Examples of Transformations

### Rotation

A 90° counterclockwise rotation moves:

- î → (0, 1)
- ĵ → (-1, 0)

The matrix columns are exactly these transformed basis vectors.

### Shear

In a shear transformation:

- î remains fixed.
- ĵ shifts sideways.

The resulting matrix records these new positions, and multiplication tells us how every vector is affected.

## Reading a Matrix Visually

You can also work backwards.

Given a matrix:

1. Move î to the location described by the first column.
2. Move ĵ to the location described by the second column.
3. Imagine the rest of space adjusting while preserving straight lines, parallelism, and spacing.

This reconstructs the entire transformation.

## When Space Collapses

If the two matrix columns point in the same direction, they are linearly dependent.

In this case:

- The transformation loses a dimension.
- The entire plane is squished onto a single line.
- Every transformed vector lies within the span of those two dependent columns.

This shows how matrices connect directly to span and linear dependence.

## The Big Idea

A matrix is not just a table of numbers.

It is a description of how space is transformed.

- Columns represent where the basis vectors land.
- Matrix multiplication computes the resulting linear combination.
- Every matrix corresponds to a geometric transformation.

This perspective becomes the foundation for nearly every major topic in linear algebra, including matrix multiplication, determinants, change of basis, and eigenvalues. Once matrices are viewed as transformations of space rather than collections of numbers, the subject becomes far more intuitive.

# Resources
1. [Linear transformations and matrices | Chapter 3, Essence of linear algebra - 3Blue1Brown](https://youtu.be/kYB8IZa5AuE)