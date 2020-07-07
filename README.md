# Pies is a reactive framework written for Go. [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mrmiguu/Pies/blob/master/LICENSE)

Why Pies
-
Question: when `a` changes, how do you update `b`?
```go
    a = 2
    b = a + 2
```

Engineers still use observers, listeners, events, signals, and polling to keep variables dependent on other variables and data dependent of data up to date.

### Pies keeps all your variables and all your data fresh all the time. No matter how many dependencies.

How Pies works
-
It all starts from the `pie.Mount`. Here, Pies calls the `myApp` function.
```go
func main() {
    pie.Mount(myApp)
}
```
The `myApp` function calls everything. And everything calls everything else. And so on.
```go
func myApp() {
    println("myApp!")

    myComponent()
    myComponent2()
    myComponent3()
    myComponent4()
    ...
    myComponentN()
}
```
In Pies, `pie.Var` represents a variable, and `pie.Do` is for calling code that you only want to run once or when dependencies change.
```go
func myComponent() {
    count, setCount := pie.IntVar(1)

    pie.Do(func() {
        setCount(count + 99)
    })

    pie.Do(func() {
        println("count updated to", count) // 1, 100
    }, count)
}
```
Setting a variable triggers a cascading effect; all functions from the root downwards are re-called, allowing all variable and data changes to naturally propagate throughout your tree of Pies.

In a nutshell
-
Pies is an elegant inversion on the classic iterative programming paradigm. It's a way of focusing on the data itself and always being up to date, always being structured in a way to allow for reacting to those changes smoothly.

Pies serves as a long term stable frame to work within, for many years and engineers to come.
