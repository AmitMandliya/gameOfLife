function life(rows, cols, some, generations) 
    now = trues(rows*cols) 
    for c=1:rows*cols
        now[c] = rand() < some 
    end
    live(now,rows,generations)
end

function live(a,r,gen)
    if gen < 1
        return nothing
    end
    sleep(0.1)
    homeScreen()
    println("Generation ",lpad(gen,3,"0"))
    for c=1:length(a)
        if(a[c])
          print("o")
        else 
          print(" ")
        end
        if mod(c,r)==0
            print("\n") 
        end
    end
    b = falses(1000) 
    for c=1:1000
        neighbours = 0
        if c-1 >= 1
            neighbours += a[c-1]
        end
        if c+1 <= 1000
            neighbours += a[c+1]
        end
        if c-r-1 >= 1
            neighbours += a[c-r-1]
        end
        if c-r >= 1 && c-r <=1000
            neighbours += a[c-r]
        end
        if c-r+1 >= 1  && c-r+1 <=1000
            neighbours += a[c-r+1]
        end
        if c+r-1 >=1 && c+r-1 <= 1000
            neighbours += a[c+r-1]
        end
        if c+r <= 1000
            neighbours += a[c+r]
        end
        if c+51 <= 1000
            neighbours += a[c+51]
        end
        b[c] = a[c]
        if a[c]==0
            b[c] = (neighbours==3)
        else
          b[c] = (neighbours==2 || neighbours== 3)
        end
    end
    gen = gen-1
    live(b,r,gen)
end
                        
function homeScreen()  
    print("\033[1;1H") 
end
                        
function clearScreen() 
    run(`clear`) 
end
                        
                        
sleep(0.1) 
clearScreen()
life(50,20,0.6,200) 