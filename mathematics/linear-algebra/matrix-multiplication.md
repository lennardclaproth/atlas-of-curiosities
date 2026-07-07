# Matrix Multiplication as Composition of Transformations

The central idea of this lesson is that **matrix multiplication represents applying one linear transformation after another**.

Understanding this geometric meaning makes matrix multiplication far more intuitive than treating it as a formula to memorize.

## Quick Recap: Matrices Represent Transformations

A linear transformation moves space while preserving two properties:

1. Straight grid lines remain straight, parallel, and evenly spaced.
2. The origin stays fixed.

Every transformation is completely determined by where it sends the basis vectors:

- î = (1, 0)
- ĵ = (0, 1)

The coordinates of their destinations become the columns of a matrix.

Because every vector is a linear combination of the basis vectors, knowing where the basis vectors go tells us where every vector goes.

## Applying Multiple Transformations

Often, we want to perform one transformation and then another.

For example:

1. Rotate the plane.
2. Apply a shear.

The overall result is itself a new linear transformation called the **composition** of the two transformations.

Just like any other transformation, it can be represented by a single matrix.

## What Matrix Multiplication Really Means

Suppose:

- Matrix A represents one transformation.
- Matrix B represents another.

Applying B first and then A means:

1. Transform the vector using B.
2. Transform the result using A.

The matrix that captures this entire process is the product:

A × B

This gives matrix multiplication its geometric interpretation:

> Matrix multiplication is the composition of linear transformations.

## Why the Order Matters

One surprising feature is that matrix multiplication is read from right to left.

In:

A × B × v

the vector is transformed by B first, then by A.

This follows the same logic as ordinary function composition.

Because transformations can affect space differently, changing the order usually changes the result.

For example:

- Shear → Rotate
- Rotate → Shear

produce completely different final transformations.

This means matrix multiplication is generally **not commutative**:

A × B ≠ B × A

## How Matrix Multiplication Is Computed

To build the matrix representing a composition:

1. Follow where î ends up after both transformations.
2. Follow where ĵ ends up after both transformations.
3. Use their final positions as the columns of the new matrix.

This leads directly to the familiar matrix multiplication formula.

Rather than memorizing the formula, it helps to remember:

> The columns of the product matrix are simply the transformed versions of the columns of the original matrix.

The algebra follows naturally from tracking the basis vectors.

## Why the Formula Exists

Traditional courses often present matrix multiplication as a numerical procedure.

But the numbers are not the main idea.

The formula exists because:

- Columns represent transformed basis vectors.
- Applying another transformation means transforming those columns.
- The resulting columns describe the composed transformation.

The computation is just bookkeeping for a geometric process.

## Associativity Becomes Obvious

One famous property of matrix multiplication is associativity:

(A × B) × C = A × (B × C)

Algebraically, proving this can be tedious.

Geometrically, it is almost trivial.

Both sides simply mean:

1. Apply C.
2. Apply B.
3. Apply A.

The order of transformations never changes, only the grouping.

Since the same sequence of transformations occurs in both cases, the results must be identical.

This provides both a proof and an intuitive explanation for why matrix multiplication is associative.

## The Big Idea

Matrices are not merely arrays of numbers.

They represent transformations of space.

Once that viewpoint is adopted:

- Matrix–vector multiplication means applying a transformation.
- Matrix multiplication means composing transformations.
- Non-commutativity reflects the importance of order.
- Associativity follows naturally from performing the same sequence of transformations.

This geometric perspective turns matrix multiplication from a mechanical calculation into a meaningful description of how transformations combine.

# Resources
1. [Matrix multiplication as composition | Chapter 4, Essence of linear algebra - 3Blue1Brown](https://youtu.be/XkY2DOUCWMU)