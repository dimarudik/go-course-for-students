# WaitGroup example

```
type WaitGroup struct {
	noCopy noCopy // запрещает копирование структуры

	state1 [3]uint32
}


wg:= &WaitGroup{}
wg.Add(delta) // добавляет в счетчик значение delta
wg.Done() // вычитает из счетчика 1
wg.Wait() // ожидает, пока счетчик не станет 0
```

runtime.GOMAXPROCS(1) разрешает только один логический процессор.  
В этом случае сначала будет выполняться горутина, вставшая в очередь первой.  
Управлять этим можно или с помощью задержек, либо с помощью runtime.Gosched()  