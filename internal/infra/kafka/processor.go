package kafka

//import (
//	"log"
//
//	"github.com/lovoo/goka"
//)
//
//func process(ctx goka.Context, msg any) {
//	var userLike *UserLike
//	var ok bool
//	var userPost *UserPost
//
//	if userLike, ok = msg.(*UserLike); !ok || userLike == nil {
//		return
//	}
//
//	if val := ctx.Value(); val != nil {
//		userPost = val.(*UserPost)
//	} else {
//		userPost = &UserPost{PostLike: make(map[int]bool)}
//	}
//
//	userPost.PostLike[userLike.PostId] = userLike.Like
//
//	ctx.SetValue(userPost)
//	log.Printf("[proc] key: %s,  msg: %v, data in group_table %v \n", ctx.Key(), userLike, userPost)
//}
//
//func runProcessor() {
//	g := goka.DefineGroup(group,
//		goka.Input(topic, new(userLikeCodec), process),
//		goka.Persist(new(userPostCodec)),
//	)
//	p, err := goka.NewProcessor(brokers,
//		g,
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = p.Run(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//}
