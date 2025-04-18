# Algorithms library

## Index

1. [Array Linear Search](arrs/01_linear_search.go) arrays basic generics avg O(n/2) best O(1) worst O(n)
2. [Array Append](arrs/02_append.go) arrays basic generics O(n) resizing
3. [Array Insert](arrs/03_insert.go) arrays basic generics O(n) resizing
4. [Array Delete](arrs/04_delete.go) basic generics resizing O(n)
5. [Linked List Append](linked_list/main.go) generics arrays append O(1)
6. [Linked List Merge](linked_list/main.go) generic ordered arrays merge O(n+m)
7. [Bubble Sort](sorting/bubble) generics swap array sorting
8. [Hoare Partitions](sorting/partitions/hoare.go) generics array hoare sorting O(n) space complexity O(1)
9. [Common Words](strings/01_common_words.go) strings O(n+m)
10. [Two Sum](arrs2/01_sum.go) arrays O(n) *
11. [K Closest Points](arrs2/02_k_distance.go) arrays O (N * Long(N)) * **
12. [Queue](ds/queue.go) generic queue using an slice
13. [Neuron](nn/simple.go) basic neuron


- \* study later
- \*\** try to make it simpler

## Array Vs Linked List
 
Arrays will be faster than linked lists for most applications.  This is because arrays are stored in contiguous 
memory locations, which makes it easy to access elements using an index.
Linked lists, on the other hand, require traversing the list to find an element, which can be slower.

Use arrays when the number of elements is known in advance and when you need to access elements by index.
When accessing the elements is a common operation, arrays are a better choice. Array access time is O(1) while 
linked list is O(n). Binary search, interpolation search, quick sort, and merge sort are all faster with arrays.

Use linked lists when the number of elements is not known in advance and when you need to insert or delete elements.
The most common use case of Linked List is the implementation of undo operation in Browsers and editing software.
If you require an array  of linked list for hashing, you can use a linked list to handle collisions.