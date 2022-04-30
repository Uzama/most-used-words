# most-used-words
A server program to process text to count most used words in given text or paragraph

### How to start 

- Clone the project and run the program by typing ```go run main.go```
- Server start accepting requests on ```localhost:8080```

## API Reference

```http
  POST /process_text
```

 - Request body


```json
    {
    "text": "In computer science, heapsort is a comparison-based sorting algorithm. Heapsort can be thought of as an improved selection sort: like selection sort, heapsort divides its input into a sorted and an unsorted region, and it iteratively shrinks the unsorted region by extracting the largest element from it and inserting it into the sorted region. Unlike selection sort, heapsort does not waste time with a linear-time scan of the unsorted region; rather, heap sort maintains the unsorted region in a heap data structure to more quickly find the largest element in each step.[1] Although somewhat slower in practice on most machines than a well-implemented quicksort, it has the advantage of a more favorable worst-case O(n log n) runtime. Heapsort is an in-place algorithm, but it is not a stable sort.Heapsort was invented by J. W. J. Williams in 1964.[2] This was also the birth of the heap, presented already by Williams as a useful data structure in its own right.[3] In the same year, R. W. Floyd published an improved version that could sort an array in-place, continuing his earlier research into the treesort algorithm"
    }
```

- Response 
```json
    [{"word":"the","count":11},{"word":"in","count":9},{"word":"a","count":8},{"word":"heapsort","count":6},{"word":"sort","count":6},{"word":"an","count":5},{"word":"region","count":5},{"word":"it","count":5},{"word":"unsorted","count":4},{"word":"of","count":4}]
```
