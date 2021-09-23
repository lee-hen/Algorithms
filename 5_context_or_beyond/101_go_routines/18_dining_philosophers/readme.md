# Dining Philosophers

- 5 philosophers sitting at a round table
- 1 chopstick is placed between each adjacent pair
- Want to eat rice from their plate, but needs two chopsticks
- Only one philosopher can hold a chopstick at a time
- Not enough chopsticks for everyone to eat at once

Each chopstick is a mutex

Each philosopher is associated with a goroutine and two chopsticks
