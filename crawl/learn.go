
func main(){
	c :=  make(chan string)
	people := [2]string{"nico", "moai"}
	for _, person := range people{
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println("Received this message :", <- c)
	}

}

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 3)
	fmt.Println(person)
	channel <- person + "is sexy"
}
