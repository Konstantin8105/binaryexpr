# binaryexpr
check *ast.BinaryExpr in golang code in according to number axis

Example of number axis(numbers from left to right):
```
-|---|---|---|---|---|--->
-2  -1   0   1   2   3

```

```
Invalid comparing:

if a >= 5 ....
if b > c ....
d := b > a
```

```
Valid comparing:

if 5 <= a ...
if c < b ...
d := a < b
```
