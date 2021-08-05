# go-blog

go-blog is a CMS with a WordPress-like template hierarchy and a Laravel-like controller.

## go-blog's Design

### Summary

This framework designed by MVC architecture and DDD architecture.

Forgive me, I'm a beginner.:wink:

### Domain

#### Post

|Name|Description|
|---|---|
|PostId|post id|
|PostTitle|post title|
|PostSlug|post slug (use in URL)|
|PostBody|post body (usable markdown)|
|PostStatus|post status|

#### Category

|Name|Description|
|---|---|
|CategoryId|category id|
|CategoryName|category name|
|CategorySlug|category slug (use in URL)|
|CategoryDescription|category description|
|CategoryParent|parent category id|
|CategoryStatus|category status|

#### Tag

|Name|Description|
|---|---|
|TagId|tag id|
|TagName|tag name|
|TagSlug|tag slug (use in URL)|
|TagDescription|tag description|
|TagStatus|tag status|

#### Author

|Name|Description|
|---|---|
|AuthorId|author id|
|AuthorName|author name|
|AuthorSlug|author slug (use in URL)|
|AuthorDescription|author description|
|AuthorStatus|Author status|

### View

You can build the view using only golang.

> using package is [html/template](https://pkg.go.dev/html/template).
