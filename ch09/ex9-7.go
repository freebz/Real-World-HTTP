// 예제 9-7 server-sent events로 재접속한다

for {
	select {
	case <- ctx.Done():
		fmt.Println("Connection close from server")
		events, ctx, err = EventSource("http://localhost:18888/prime")
		if err != nil {
			panic(err)
		}
		continue
	case event := <-events:
		fmt.Printf("Event(Id=%s, Event=%s): %s\n", event.Id, event.Name, event.Data)
	}
}
