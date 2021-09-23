## Concurrency Patterns
There are lots of different patterns we can create with goroutines and channels. Two interesting patterns are resource pooling and concurrent searching.

## Notes

* The work code provides a pattern for giving work to a set number of goroutines without losing the guarantee.
* The resource pooling code provides a pattern for managing resources that goroutines may need to acquire and release.
* The search code provides a pattern for using multiple goroutines to perform concurrent work.

## Links

http://blog.golang.org/pipelines

https://talks.golang.org/2012/concurrency.slide#1

https://blog.golang.org/context

http://blog.golang.org/advanced-go-concurrency-patterns

http://talks.golang.org/2012/chat.slide

## Code Review

[Work](work)

[Pooling](pool)

[Chat](chat)
