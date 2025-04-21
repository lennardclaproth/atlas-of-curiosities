
# Propositions

A proposition is a declarative statement that are either true or false.

You can construct a compound proposition by combining multiple propositions using *connectives*.

connectives:

| Type          | Operator | Action         |
| ------------- | -------- | -------------- |
| Negation      | ¬        | Not            |
| Conjunction   | ∧        | And            |
| Disjunction   | ∨        | Or             |
| Implication   | ⇾        | If, then       |
| Biconditional | ⇿        | If and only If |

## Negations

The negation of proposition

$
 p \text{ is } ¬p \text{ (not p)}
$

Therefore we could make an easy negation based on the following proposition.

$
p=\text{My dog is the cutest dog}\\
¬p=\text{My dog is \underline{not} the cutest dog}
$

Important to note is that if we have a statemen like the following:

$
p=\text{The door is not open}
$

The negtion of that becomes:

$
¬p=\text{The door is open}
$

We can comprise the following truth table from the information gathered so far:

| p   | ¬p  |
| --- | --- |
| T   | F   |
| F   | T   |

It is important to note that if you make a truth table you can follow the rule:

$
size = \text{ncol}^{\text{nrow}}
$

## Conjunctions

The conjunction of propositions is denoted p ∧ q. For a conjunction to be true both propositions must be true.

Say we have the following propositions:

$
p = \text{It is raining}
$

$
q = \text{I am home}
$

Based on the propositions above we can create the following conjunction:

$
p ∧ q
$

With the following truth table:

| p   | q   | p ∧ q |
| --- | --- | ----- |
| T   | T   | T     |
| T   | F   | F     |
| F   | T   | F     |
| F   | F   | F     |

## Disjunction

The disjunction of propositions p and q is denoted by p ∨ q and reads p or q. For a disjunction to be true, *either* proposition must be true.

| p   | q   | p ∨ q |
| --- | --- | ----- |
| T   | T   | T     |
| T   | F   | T     |
| F   | T   | T     |
| F   | F   | F     |

There is also a different type of disjunction, the exclusive or. Base don this we now know two disjunctions:

- Inclusive or -> p ∨ q -> The prerequisite for MA420 is MA315 or MA335
- Exclusive or -> p ⊕ q -> You get soup or salad with your entree

This will be then the complete truth table:

| p   | q   | p ∨ q | p ⊕ q |
| --- | --- | ----- | ----- |
| T   | T   | T     | F     |
| T   | F   | T     | T     |
| F   | T   | T     | T     |
| F   | F   | F     | F     |

## Implications

The implication of propositions p and q is denoted as p ⇾ q and reads, *if p then q* or *p implies q*.

Based on this we can conclude the following:

> When the hypothesis is TRUE, the conclusion must be TRUE for the implication to be true. When the hypothesis is FALSE, the conclusion is TRUE

This will give us the following truth table:

| p   | q   | p ⇾ q |
| --- | --- | ----- |
| T   | T   | T     |
| T   | F   | F     |
| F   | T   | T     |
| F   | F   | T     |

Based on the truth table above we see an "anomaly" as the third value however this is to be expected considering that this implication is one directional, meaning that q says nothing about p. So we can conclude that on line three if q is true it does not mean anything for the value of p.

### Converse, inverse, contrapositive

From the implication mentioned above we can construct 3 new conditional statements:

| name           | notation | action            |
| -------------- | -------- | ----------------- |
| Converse       | q ⇾ p    | switch order      |
| Inverse        | ¬p ⇾ ¬q  | negate            |
| Contrapositive | ¬q ⇾ ¬p  | switch and negate |

So now if we state the following *It is raining is a sufficient condition for me not going to town.*

I want to construct the IF THEN form.

IF *it is raining*, THEN *I won't go to town*.

we can conclude that:

$
p=\text{It is raining} \\
q=\text{I won't go to town} \\
p⇾q
$

So now if I would want the converse what would it look like?
IF *I won't go to town* THEN *it is raining*

For the inverse it would then become:
IF *It is not raining* THEN *I won't not go to town*

And lastly the contrapositive: IF *I won't not go to town* THEN *It is not raining*

Besides the aforementioned implication we also have a second implication which is a Biconditional. The biconditional of p and q is denoted as p ⇿ q which reads *p if and only if q*

>**Note: A biconditional can also be written as a compound proposition:**
>
>$
>p ⇿ q=(p ⇾ q) ∧ (q ⇾ p)
>$

## Logical Equivalences of Disjunctions and Conjunctions [1, Discrete Maths - 1.3.2]

### Identity Laws

These laws show that disjunction with false and conjunction with true don't change the value of the proposition.

- $p \lor F \equiv p$ → Adding false to a statement with OR doesn't change its truth value.
- $p \land T \equiv p$ → AND-ing a statement with true also keeps it the same.

---

### Domination Laws

These laws demonstrate that disjunction with true always results in true, and conjunction with false always results in false.

- $p \lor T \equiv T$ → Anything OR true is always true.
- $p \land F \equiv F$ → Anything AND false is always false.

---

### Idempotent Laws

Repeating the same proposition in a disjunction or conjunction doesn’t change the result.

- $p \lor p \equiv p$ → OR-ing a proposition with itself is just itself.
- $p \land p \equiv p$ → AND-ing a proposition with itself is also just itself.

---

### Double Negation Law

Negating a negation brings you back to the original proposition.

- $\neg (\neg p) \equiv p$ → Double negatives cancel out.

---

### Absorption Laws

These laws eliminate redundancy by showing how combining repeated elements simplifies the expression.

- $p \lor (p \land q) \equiv p$ → If $p$ is true, the whole disjunction is true, regardless of $q$.
- $p \land (p \lor q) \equiv p$ → If $p$ is false, the whole conjunction is false, regardless of $q$.

---

### Negation Laws

These capture key logical principles like the law of the excluded middle and contradiction.  
[4] Refers to classical logic axioms.

- $p \lor \neg p \equiv T$ → Law of the Excluded Middle [4]: a statement is either true or its negation is.
- $p \land \neg p \equiv F$ → Contradiction [4]: a statement can't be both true and false at once.

---

### Commutative Laws

The order of propositions doesn't matter in OR and AND operations.

- $p \lor q \equiv q \lor p$
- $p \land q \equiv q \land p$

---

### Associative Laws

When using only ORs or only ANDs, grouping doesn’t matter.

- $(p \lor q) \lor r \equiv p \lor (q \lor r)$
- $(p \land q) \land r \equiv p \land (q \land r)$

---

### Distributive Laws

These show how to distribute OR over AND and vice versa.

- $p \lor (q \land r) \equiv (p \lor q) \land (p \lor r)$
- $p \land (q \lor r) \equiv (p \land q) \lor (p \land r)$

---

### De Morgan’s Laws

These laws transform negations of compound statements into equivalent forms using the opposite operator.

- $\neg(p \land q) \equiv \neg p \lor \neg q$
- $\neg(p \lor q) \equiv \neg p \land \neg q$

Now knowing all this let's proof the following equation:

$$
\neg ( p\lor (\neg p \land q))\equiv \neg p \land \neg q
$$

Let's start with the left hand side, we can see De Morgan's second law there:
$$
\neg ( A \lor B) \\
\text{where } A = p\\
\text{where } B = (\neg p \land q)
$$

therefore we can write the following:
$$
\neg p \land \neg (\neg p \land q) \equiv \neg ( p\lor (\neg p \land q))
$$

We can further simplify this by applying the first law of De Morgan:
$$
\neg (A \land B) = \neg A \lor \neg B\\
\text{Where } A = \neg p \\
\text{Where } B = q \\
\text{which will give us }(\neg \neg p \lor \neg q)
$$

Now having this we can combine the results:
$$
\neg p \land (p \lor \neg q) \equiv \neg ( p\lor (\neg p \land q))
$$

If we now apply the law of distribution on this which looks as follow:
$p∧(q∨r)≡(p∧q)∨(p∧r)$ we get the following:

$$
(\neg p \land p) \lor (\neg p \land \neg q) \equiv \neg p \land (p \lor \neg q)
$$

This statement however contains a contradiction $(\neg p \land p)$ which we can immediately convert to False.

Which will give us:
$$
F \lor (\neg p \land \neg q)
$$

which is the first Identity law of disjunctions, *"Adding false to a statement with OR doesn't change its truth value."* so we can conclude:
$$
\neg p \land \neg q \equiv F \lor (\neg p \land \neg q)
$$

Which brings us back to:
$$
\neg p \land \neg q \equiv \neg p \land \neg q
$$

### Tool table (other helpful laws)

| Tool                                                               | Formal Name                     | Also Known As               | Category                          |
| ------------------------------------------------------------------ | ------------------------------- | --------------------------- | --------------------------------- |
| $p \rightarrow q \equiv \neg p \lor$                               | Material Implication            | Implication-as-disjunction  | Logical Equivalence               |
| $p \rightarrow q \equiv \neg q \rightarrow \neg p$                 | Contrapositive Law              | Contraposition              | Logical Equivalence               |
| $\neg(p \land q) \equiv \neg p \lor \neg q$                        | De Morgan’s Law                 | —                           | Logical Equivalence               |
| $p \land \neg q \equiv \neg(p \rightarrow q)$                      | Implication Negation            | Proof by contradiction tool | Logical Equivalence               |
| $(p \land q) \rightarrow r \equiv p \rightarrow (q \rightarrow r)$ | Currying (in logic/computation) | Nested Implication          | Higher-Order Logic / Curry–Howard |
| $p \lor \neg p \equiv T$                                          | Law of the Excluded Middle      | —                           | Classical Logic                   |
| $p \land \neg p \equiv F$                                         | Law of Non-Contradiction        | —                           | Classical Logic                   |

these fall under the formal categories:

- Propositional Logic [5]
- First-order Logic [6]
- Natural Deduction [7]

<!-- 

THIS IS WRONG STILL INTERESTING THOUGH

### Three laws of thought

There are three laws of thought,

- The law of identity *"Whatever is is."* can be denoted as $p \equiv p$
- The law of non-contradiction *"Nothing can both be and not be."* can be denoted as $p∨F=p \text{ or } \neg (p \land \neg p)$
- The law of excluded middle *"Everything must either be or not be."* can be denoted as  p \lor \neg p

-->

## Resources

1. [Discrete maths by Kimberly Brehm 1.x.x](https://youtube.com/playlist?list=PLl-gb0E4MII28GykmtuBXNUNoej-vY5Rz&feature=shared)
2. [De Morgan's laws](https://en.wikipedia.org/wiki/De_Morgan%27s_laws)
3. [Law of identity](https://en.wikipedia.org/wiki/Law_of_identity)
4. [Law of thought](https://en.wikipedia.org/wiki/Law_of_thought#The_three_traditional_laws)
5. [Propositional logic](https://en.wikipedia.org/wiki/Propositional_calculus)
6. [First order logic](https://en.wikipedia.org/wiki/First-order_logic)
7. [Natural deduction](https://en.wikipedia.org/wiki/Natural_deduction)
