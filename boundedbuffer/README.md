# Library
We implement a blocking queue, with maximum capacity, with multiple producers/consumers in a concurrent environment.

# Usage
You specify the capacity, and then you call Produce or Consume, from different threads.

The producer will block when all slots have been filled, until one becomes available.

The consumer will block when there isn't any data left, until new data is added by a producer.

# Implementation
Implementation wise, it uses a channel as a semaphore.

We implement a blocking queue, with maximum capacity, with multiple producers/consumers in a concurrent environment.

When the producer needs to write, it blocks on the number of empty slots to fill. When it finishes, it adds 1 to filled slots.

When the consumer needs to read, it blocks on the number of filled slots. When it finishes, it adds 1 to available empty slots.

This is so that producers correctly block other producers and consumers, and vice versa.

# Details about the problem
Have a look in the book "Modern Operating Systems" for "bounded buffer" concurrency problem.
