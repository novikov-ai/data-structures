# Blockchain

## Разница в трудоёмкости вычисления нонса по мере увеличения количества нулей в конце хеша

### Результаты

~~~
goos: darwin
goarch: arm64
pkg: blockchain
cpu: Apple M1 Pro
BenchmarkHashWithNonce/nonceCount=1-10          1000000000               0.0000208 ns/op               0 B/op          0 allocs/op
BenchmarkHashWithNonce/nonceCount=2-10          1000000000               0.0000139 ns/op               0 B/op          0 allocs/op
BenchmarkHashWithNonce/nonceCount=3-10          1000000000               0.006951 ns/op        0 B/op          0 allocs/op
BenchmarkHashWithNonce/nonceCount=4-10          1000000000               0.09779 ns/op         0 B/op          0 allocs/op
BenchmarkHashWithNonce/nonceCount=5-10               374           4590645 ns/op         2735987 B/op      68375 allocs/op
BenchmarkHashWithNonce/nonceCount=6-10                 1        16855675542 ns/op       9853686176 B/op 246247922 allocs/op
BenchmarkHashWithNonce/nonceCount=7-10                 1        43690236292 ns/op       25723829656 B/op        642846564 allocs/op
~~~

### Выводы 

Проанализировав результаты в зависимости от количество нулей (nc) и выполняемых операций за единицу времени, то получили следующее:

- nc = 1 => 0.0000208 ns/op
- nc = 2 => 0.0000139 ns/op
- nc = 3 => 0.006951 ns/op
- nc = 4 => 0.09779 ns/op 
- nc = 5 => 4590645 ns/op => 47,000,000× ↑	~10⁷
- nc = 6 => 16855675542 ns/op (~16.8 seconds) => 3,670× ↑	~10³
- nc = 7 => 43690236292 ns/op (~43.7 seconds) => 2.59× ↑	~10⁰

Первые несколько значений, вероятно, не репрезентативны из-за очень быстрого выполнения операции. Нужно больше тестов.

Наблюдаем, что при значении в 5 нулей (nc=5) получаем супер-экспоненциальный рост количества операций, времени операции и количества аллокаций.  