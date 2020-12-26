# Ming Kitchen Backend
This project is using Go, Gin, Graphql. 

[Frontend Design](https://www.figma.com/proto/1icMeG8q4iWH5aGSqjpm69/Ming-Kitchen-Dashboard?node-id=1%3A5&viewport=115%2C320%2C0.6200000047683716&scaling=scale-down)

## Environment Prepare
- [Go](https://golang.org/doc/install)  (version 1.12+ is required)
- [Gin](https://github.com/gin-gonic/gin#installation)   `go get -u github.com/gin-gonic/gin`
- [Gorm](https://gorm.io/docs/) `go get -u gorm.io/gorm`

## API Design
#### Sales Stats
- getAllSalesStats
    - today
    - difference between today and yesterday
    - difference between this year today and last year today
    - this year total
    - profit
- getSalesByDate (support date range)

#### Cost Stats
- getAllCostStats
    - this month
    - difference between this month and last month
    - total cost
- addCost
    - amount
    - type of cost
    - date
#### Uber Eats
- getUberSalesStats
    - today
    - difference between today and yesterday
    - difference between this year today and last year today
    - this year total
    - profit
- getAllErrorOrder (DESC order by date)
    - order id
    - amount


#### User (keep it until other finished)
- login
- register
- token management
    - renew token
    - expire token


