# State (p. 305 Design Patterns)
Allow an object to alter its behavior when its internal state changes. The object will appear to change its class

## When to use
Use the State pattern in either of the following cases:
- An object's behavior depends on its state, and it must change its behavior at run-time depending on that state
- Operations have large, multipart conditional statements that depend on the object's state. This state is usually represented by one or more enumerated constants. Often, several operations will contain this same condition structure. The state pattern puts each branch of the conditional in a separate class. This lets you treat the object's state as an object in its own right that can vary independently of other objects.

## Creating and destroying state objects
A common implementation trade-off worth considering is whether (1) to create State objects only when they are needed and destroy them thereafter versus (2) creating them ahead of time and never destroying them.

The first choice is often times preferable when the states that will be entered aren't known at runtime, and contexts change state infrequently. This approach avoids creating objects that won't be used, which is important if the state objects store a lot of information.

The second approach is better when state changes occur rapidly, in which case you want to avoid destroying states, because they may be needed again shortly. Instantiation costs are paid once upfront, and there are no destruction costs at all. This approach might be inconvenient, though, because the Context must keep references to all states that might be entered.