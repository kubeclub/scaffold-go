package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"kubeclub/scaffold_go_web/dao"
)

func ApiGraphqlRegister(router *gin.RouterGroup) {
	//核心入口
	rootQuery := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			//要请求的元素
			"user": &graphql.Field{
				//该元素里面的元素类型声明
				Type: CreateType(),
				//query参数声明
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					//具体参数获取
					id := p.Args["id"]
					v, _ := id.(int)
					//数据返回
					return &dao.User{Id: v, Name: "aa", Addr: "bb"}, nil
				},
			},
		}}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	})
	if err != nil {
		panic(err)
	}

	router.POST("/graphql", GraphqlHandler(schema))
	router.GET("/graphql", GraphqlHandler(schema))

}
func CreateType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

// 为GraphQL handler 结合gin封装的func
func GraphqlHandler(schema graphql.Schema) gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema: &schema,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
