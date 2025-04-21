# Predicate Logic

Often, statements contain **variables**, such as $x + 2 = 5$. Based on this information, we can conclude that the statement is **neither true nor false** as long as the variable is not assigned a value — it is an **open statement**.

This expression consists of two parts:

- The first part, the variable $x$, is the **subject**.
- The second part, “**plus 2 equals 5**”, is the **predicate**, which describes a property or condition that may or may not hold for $x$.

We can represent this using **predicate notation** as:

$$
P(x) = x + 2 = 5
$$

Here, $P(x)$ is a **predicate** (also called a propositional function). It becomes a **statement** that is either **true** or **false** only once a specific value for $x$ is substituted.

Connectives from propositional logic carry over to predicate logic. therefore we can compound them and use them in statements.

$$
P(3) \lor P(-1) \text{ where } P= x \gt 0
$$

The solution to the above mentioned will be *True or False = True*

## Quantifiers

Using quantifiers we can turn propositional function/ predicates into a proposition.

| Type                   | Symbol     | Meaning                         | Truth Condition                       |
| ---------------------- | ---------- | ------------------------------- | ------------------------------------- |
| Universal Quantifier   | $\forall$  | For all values in the domain    | Must be true for every possible value |
| Existential Quantifier | $\exists$  | There exists at least one value | True if true for some value           |
| Unique Existential     | $\exists!$ | There exists exactly one        | True if true for one and only one     |

### Universal Quantifier

Suppose we have the open statement $x > 0$.  
Using the universal quantifier, we can express the proposition:

$$
\forall x\, P(x)
$$

This means that the predicate $P(x)$ (in this case, "$x > 0$") must be **true for every value of $x$** in the specified **domain** (also called the *universe of discourse*).

Let us consider two different domains:

1. $U = \mathbb{Z}$ → All integers  
2. $U = \mathbb{Z}^+$ → All **positive** integers

We want to evaluate the truth of:

$$
\forall x \in U,\, x > 0
$$

We ask:

> Is it true that **all integers are greater than 0**?

We can immediately find a counterexample for the first case:

- $-3 \in \mathbb{Z}$
- $P(-3) = -3 > 0 \equiv \text{False}$

Therefore:

$$
\forall x \in \mathbb{Z},\, x > 0 \quad \text{is false}
$$

Now the domain is only **positive integers**.

- Every $x \in \mathbb{Z}^+$ is by definition $> 0$
- So $P(x)$ is true for all $x$

Therefore:

$$
\forall x \in \mathbb{Z}^+,\, x > 0 \quad \text{is true}
$$

we can finally conlude that the **truth of a universally quantified statement** depends entirely on the **domain of discourse**. A statement like $\forall x\, P(x)$ is:

- **True** if $P(x)$ holds for **every element** in the domain
- **False** if **any counterexample exists**

### Existential Quantifier

Suppose we have the same open statement as before: $x > 0$.  
Using the **existential quantifier**, we express the proposition:

$$
\exists x\, P(x)
$$

This means that the predicate $P(x)$ must be **true for at least one** value of $x$ in the domain (also known as the *universe of discourse*).

Consider the following two domains:

1. $U = \mathbb{Z}$ → All integers  
2. $U = \mathbb{Z}^-$ → All **negative** integers

We evaluate the truth of the statement:

$$
\exists x \in U,\ x > 0
$$

We are asking:  
> “Is there at least one integer such that $x > 0$?”

For the first domain we can already quickly conclude that it is the case, for example, $x = 1$ satisfies the condition.  
So the proposition is **true**.

Howeber the second domain includes only **negative integers**.

We are asking:  
> “Is there a negative integer such that $x > 0$?”

Here we can conclude that such value does not exist in the domain, so the predicate is **false for every** $x$ in $\mathbb{Z}^-$.  
Thus, the existential statement is **false**.

We now know that the existential quantifier $\,\exists x\,P(x)$ asserts that **at least one value** in the domain makes the predicate **true**.  

- It is **true** if there exists *any* satisfying value.
- It is **false** if **none** exist.

## Negating quantifiers

Let $P(x)$ be the statement *"x has taken a course in programming"* for the domain students in the class.

$\neg \forall xP(x)$ which can be read as *"There is a student in the class that has **not** taken a course in programming"* Which can in turn be written as $\exists x \neg P(x)$

$\neg \exists x P(x)$ Which can be read as *"All students in the class have not taken a course in programming"* $\forall x \neg P(x)$

The two examples mentioned above are the ***De Morgan's laws*** for quantifiers.

The first law is as follows:
$$
\neg \exists x P(x) \equiv \forall x \neg P(x)
$$

- This is ***TRUE*** when $P(x)$ is ***FALSE*** for every x
- This is ***FALSE*** when there is an x for which $P(x)$ is ***TRUE***

The second law is as follows
$$
\neg \forall x P(x) \equiv \exists \neg P(x)
$$

- This is ***TRUE*** when there is an x for which $P(x)$ is ***FALSE***
- This is ***FALSE*** when $P(x)$ is ***TRUE*** for every x

## Resources

1. [Discrete maths by Kimberly Brehm 1.x.x](https://youtube.com/playlist?list=PLl-gb0E4MII28GykmtuBXNUNoej-vY5Rz&feature=shared)
2. [Geeks for geeks quantifiers](https://www.geeksforgeeks.org/mathematic-logic-predicates-quantifiers/)