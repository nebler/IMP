imp project: https://sulzmann.github.io/ModelBasedSW/imp.html#(2)

We first finished pretty and eval from the sketch. <br>
We change the eval of Or and And so it doesn't short-circuit evaluate. <br>
We created a new struct to describe variables since before they were only
the name but to implement nested looping they also need to have the scope they are defined in.
This also means that every state now has a dedicated name.

```
type ValName [2]string

type ValState struct {
    name string
    vals map[ValName]Val
}
```

We give the first state the name of global. <br>

Inside of declare the name of the current state as well as the variable name will be saved. This will be used after an if or while block to determine if any of the global variables have been altered. (update method)
Inside of every if or while we also create a new state which is a copy of the current one just with an additional prefix. <br>

The parser was mostly just iterating over the current characters and creating edge cases to fit the syntax. The only complicating part was the grouping character: ().
We solved this issue by writing an enitre new fucntion to just solve operations inside the grouping which will be called after solving the first part of the expression. We also wrote a method to check when a number ends by going over the string and checking if the current character is a number. We did the same to determine when the name of a variable ends.
