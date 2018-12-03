package helloimpl

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"

	pb "github.com/golangstudy0/server/hello/hello"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

//create database user
/**
create table userinfo (
	`uid` bigint(20) NOT NULL AUTO_INCREMENT comment '用户id' ,
	`name` varchar(120) NOT NULL comment '用户名' ,
	`age` int(11) NOT NULL  comment '年龄',
	`create_time` timestamp NOT NULL  DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment '创建时间',

	PRIMARY KEY (`uid`),
	KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1000 comment '用户表';
 */

// server is used to implement hello.GreeterServer.
type Hello struct{}

var (
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	db, _ = sql.Open("mysql", "root:root@/user")
)



// SayHello implements hello.GreeterServer
func (s *Hello) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET name=?,age=?")
	res, err := stmt.Exec("testname", 1)
	if nil != res {
		nr ,_:= res.RowsAffected()
		fmt.Println("rows affect : " ,nr )
	}

	rows, err := db.Query("SELECT * FROM userinfo")
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(age)
	}

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	set, err := client.SetNX("key", "value", 5*time.Second).Result()
	fmt.Println(set, err)

	if !set {
		return &pb.HelloReply{Message:"wait..." + in.Name},nil
	}

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
