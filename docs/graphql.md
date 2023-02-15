# Graphql
graphql 比 REST 更高效、强大和灵活的新一代 api 标准，由 facebook 开源。传统的 rest 接口根据不同功能需要提供不同接口，
而 graphql 始终只需要一个接口就够了。

# 终端语法
首先是介绍在前端查询时用的语法，分成Query 和Mutation 两部分，Subscription 的和Query 是类似的就不特别说明了。

## Query
该脚手架对应的 demo 接口

**URL**
`http://localhost:8880/graphql/graphql`

**查询**
```
query
{
    user (id:3){
        id
        name
    }
}

```

**响应**
```
{
    "data": {
        "user": {
            "id": 3,
            "name": "aa"
        }
    }
}
```
