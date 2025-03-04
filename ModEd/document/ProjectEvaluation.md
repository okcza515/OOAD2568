# Criteria Project Evaluation

## Basic of OOP

The implementation of the project must apply Object Oriented concept of Go
programming language e.g. Struct, Composition, Interface, and other. It should
be emphasized that the procedural implementation should be avoided. 

## Application of Design Pattern and SOLID

The more design patterns and SOLID applied in the project, the more will be scored.
Moreover, student must be able to explain the reason of using each design pattern.

## Direction of Dependency

Try to reduce the implementation on the base library/module using SOLID and
Go composition and interface. Each request of an extra implementation on callee
side, the caller will get -1.0 point of penalty.

## Dependency

By conflict of dependency, the group of callee will get -1.0 point of penalty score
for each point and caller -0.5 point. Hence, the callee side should be very generic
and no consideration of the caller side.

## Refactoring

After getting comment from instructor, the student is able to refactor the code
to improve the characteristic.

## Code Characteristic

The implemented code must have a good characteristic according to design pattern
e.g. no god-object, code length, dependency management, code reusable, etc..

## Code Convention

The code must be compiled with the coding convention so that the difference between
the code from different developer cannot be seen.

## Code Testability

Each piece of code should be testable in different manners: unit testing, functional testing
or others. It is recommended to use Go programming language built-in test.