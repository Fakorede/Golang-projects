# MAPS

A map is a collection of indexable key-value pairs. It allows us to quickly access an element using a unique key.

## Properties

- Used to represent a collection of related properties.

- Keys are indexed - we can iterate over them.

- A map doesn't accept duplicate keys.

- Don't need to know all the keys at compile time.

- All keys must be of the same type.

- All values must be of the same type.

- Key and Value types don't need to match. They can be different.

- A map key should be a comparable type.

```
map[int]bool --> okay
map[[]float64]bool --> not okay
```

slice, map and function values are not comparable.

- Just like a slice, we can't compare a map to another map(even to itself). We can only compare it to a nil value.

- Don't use float types as a map key. Results may be inaccurate. integer and string keys are efficient.

- **Maps are fast-lookup tables. Looping over a map may indicate a design problem.**

- Reference Type
