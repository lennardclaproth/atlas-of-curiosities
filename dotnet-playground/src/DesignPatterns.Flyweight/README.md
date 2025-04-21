# FlyWeight
Use sharing to support large numbers of fine-grained objects efficiently.

## Applicability
The flyweight pattern's effectiveness depends heavily on how and where it's used. Apply the Flyweight pattern when aal of the following are true:
- An application uses a large number of objects.
- Storage costs are high because of the sheer quantity of objects.
- Most object state can be made extrinsic.
- Many groups of objects may be replaced by relatively few shared objects once extrinsic state is removed.

## Participants
- Flyweight
  - Declares an interface through which flyweights can receive and act on extrinsic state
- ConcreteFlyweight
  - Implements the Flyweight interface and adds storage for intrinsic state, if any. A ConcreteFlyweight object must be shareable. Any state it stores must be intrinsic; that is, it must be independent of the ConcreteFlyweight object's context
- UnsharedConcreteFlyweight
  - Not all Flyweight subclasses need to be shared. The Flyweight interface enables sharing; it doesn't enforce it. It's common for UnsharedConcreteFlyweight objects to have ConcreteFlyweight objects as children at some level in the flyweight object structure.
- FlyweightFactory
  - Creates and manages flyweight objects.
  - Ensures that flyweights are shared properly. When a client requests a flyweight, the FlyweightFactory object supplies an existing instance or creates one if none exists.
- Client
  - Maintains a reference to flyweight(s).
  - Computes or stores the extrinsic state of flyweight(s)