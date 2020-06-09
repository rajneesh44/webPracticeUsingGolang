## course outline 
> https://github.com/GoesToEleven/GolangTraining
> https://github.com/GoesToEleven/golang-web-dev

### Basics
```
1. Go is created at google in 2007.
2. Go is created to take advantage of multiple cores. only language to do so.
3. statically typed, fast build, fast execution, compiled langauge.
4. One of the best langauge to do web /server/.
5. It supports Polymorphism. Its not OOP.
```

### Prereuisites: 
```
- cmd
- github
- programming
- html/css/js
- go
```
### Fundamentals
```
-Slice
    x :=[]int{1,2,3,4,5}
    fmt.Println(x)

-Map
    m:=map[string]int{
        "raj":23,
    }
    fmt.Println(m)

-type person struct{
    var name string
    }
    main(){
        p1:=person{"rajneesh"}
        fmt.Println(p1)      
    }

-for loop
    1. for i := 1; i < 5; i++ {
            sum += i
        }
    2. for as a while loop 
        n := 1
        for n < 5 {
            n *= 2
        }
    3. for{
            // for as infiinite loop 
        } 
    4. for-each range loop 
            for i, s :=range strings{
                fmt.Println(i,s)
            }
- functions 
    1. reciever
        func (p person) speak(){
            fmt.Println(p.name)
            }
        p.speak(); // calling the function 

    2. method
        func speak string(){ // string be the return type

        }
        speak()

- Interfaces
    type human Interface{
        speak()
    }
    func saysomething(h human){
        h.speak()
    }
    saysomething(p1)


```
