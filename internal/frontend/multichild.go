package frontend

// func MultiChild(s space.Space, handler msg.Callback) func() {
// 	return func() {
// 		queueName := "simpleQueue"
// 		s.AddQueue(queueName)
// 		s.Subscribe(queueName, handler)
// 		parent := msg.NewSimpleMsg(uuid.Nil, []byte("I am you father"))
// 		child1 := msg.NewSimpleMsg(parent.GetId(), []byte("I am you child"))
// 		child2 := msg.NewSimpleMsg(parent.GetId(), []byte("I am you child"))
// 		child3 := msg.NewSimpleMsg(parent.GetId(), []byte("I am you child"))
// 		s.Publish(queueName, parent)
// 		s.Publish(queueName, child1)
// 		s.Publish(queueName, child2)
// 		s.Publish(queueName, child3)

// 	}
// }
