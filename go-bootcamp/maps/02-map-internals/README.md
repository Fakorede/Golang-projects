# Map Internals and Cloning

When we create a Map, Go creates a **Map Header** behind the scenes then it returns the memory location of that header.

| Map Variable: dict |     |
| :----------------: | --- |
|      Pointer       | 40  |

| Map Variable: turkish |     |
| :-------------------: | --- |
|        Pointer        | 40  |

**Map Variabe(or Value)** is a pointer to the memory address of a map header which contains another pointer to the actual data.