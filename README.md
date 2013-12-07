# STORM

STORM is Semantically Terse Object Relational Mapping for Go


## WARNING!

This software is NOT ready for production use. It is currently an experimental proof of concept. The
code is fragile, volatile and not reliable!

That said, please contribute your ideas and feedback!


## Getting Started

1. `go get github.com/brendensoares/storm`
2. `import "github.com/brendensoares/storm"`
3. Add `storm.Model` to a Go struct (e.g. `type ModelName struct {...}`)
4. Create a `func NewModelName() *ModelName` that returns `storm.Factory(&ModelName{}).(*ModelName)`
5. `import 	_ "github.com/brendensoares/storm/driver/mysql"`
6. `modelName := NewModelName()` and start using your new model!


## Roadmap Features By Example

### Define Model

```go
type Post struct {
	storm.Model           `container:"posts" driver:"mysql"`

	Id int64              `alias:"post_id"`
	Title string          `length:"100" nullable:"true"`
	TransientTitle string `ignore:"yes"`
	Content string        `type:"text"`

	CreatedAt time.Time
	UpdateAt time.Time
	DeletedAt time.Time

	Categories *Category  `relation:"hasMany" through:"categories_posts/post_id"`
	Comments *Comment     `relation:"hasMany" key:"post_id"`
	Author *Author        `relation:"belongsTo" key:"author_id"`
	ParentPost *Post      `relation:"hasOne" key:"parent_post_id"`
}
```


### Create

```go
post := NewPost()
post.Title = "Hello, World!"
saveError := post.Save()
```

### Update

```go
if post.IsLoaded() {
	post.Title = "Hello, Earth!"
	if saveError := post.Save(); saveError == nil {
    // Do stuff!
	}
}
```


### Read

#### Single Object By Identifier

```go
post1 := NewPost()
// For numeric `Id`
post1.Get(1)
post2 := NewPost()
// For string `Id`
post2.Get("2")
```

#### Single Matching Object By Criteria

```go
post := NewPost()
post.Where("title", "like", "Hello%").And("title", "not like", "%earth%").Or("title").Get()
fmt.Println(post.Title)
```

#### Many Matching Objects By Criteria

```go
post := NewPost()
matchedPosts := post.Where("title", "like", "hello%").Limit(5).Order("title", "asc").All()
for _, post := range matchedPosts {
	fmt.Println(post.Title)
}
```


### Delete

#### Loaded Object

```go
post.Delete()
```

### Many Objects By Criteria

```go
NewPost().Has("author", commentAuthor).Delete()
```


### Relationships

#### Read

##### Lazy Load "hasMany"

```go
if post.Load("Categories") {
	categoryTitle := post.Categories[0].Title
}
if post.Load("Comments") {
	commentAuthor := post.Comments[0].Author.DisplayName
}
```

##### "hasOne"

```go
if post.ParentPost != nil {
	parentTitle := post.ParentPost.Title
}
```

##### "belongsTo"

```go
if post.Author != nil {
	authorName := post.Author.DisplayName
}
```


#### Add

##### "hasMany"

```go
category1 := NewCategory()
category1.Title = "Test Category 1"
category2 := NewCategory()
category2.Title = "Test Category 2"
categoryCountBefore := len(post.Categories)
addError := post.Add("Categories", category1, category2)
categoryCountAfter := len(post.Categories)
if addError != nil || categoryCountAfter - categoryCountBefore != 2 {
	// Something smells fishy.
}
```

##### "hasOne"

```go
author := NewAuthor()
author.DisplayName = "John Smith"
post.Author = author
```

##### "belongsTo"

```go
parentPost := NewPost()
parentPost.Get(20)
post.ParentPost = parentPost
```

#### Remove

##### "hasMany"

```go
category := NewCategory()
category.Get(10)
if removeError := post.Remove("Categories", category); removeError != nil {
	// Error...The game is afoot!
}
```

##### "hasOne"

```go
post.Author = nil
```

##### "belongsTo"

```go
post.ParentPost = nil
```

