ch := make(chan int, 2) // 하지만 이렇게 큐에 사이즈를 줘서 빈공간을 만들어주면 대기하지 않고 종료됨. 즉, 빈공간이 없이 재고가 쌓이면 대기함.
		// go square2()
		// ch <- 9
		// fmt.Println("is it print?")