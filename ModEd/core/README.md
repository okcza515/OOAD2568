# BaseController Usage Example

## Overview
The `BaseController` is a generic controller that provides common CRUD (Create, Read, Update, Delete) operations for any GORM-based model. It simplifies controller implementation by offering reusable methods.

## Implementation Example
Using BaseController in ArticleController
The ArticleController extends BaseController to provide CRUD operations for Article models.

```go
package controller

import (
	"ModEd/core"
	"ModEd/project/model"

	"gorm.io/gorm"
)

type ArticleController struct {
	*core.BaseController[model.Article]
	db *gorm.DB
}

func NewArticleController(db *gorm.DB) *ArticleController {
	return &ArticleController{
		db:             db,
		BaseController: core.NewBaseController[model.Article](db),
	}
}

func (c *ArticleController) ListAllArticles() ([]model.Article, error) {
	return c.List(map[string]interface{}{})
}

func (c *ArticleController) ListArticlesWithPagination(page, pageSize int) ([]model.Article, int64, error) {
	return c.ListPagination(map[string]interface{}{}, page, pageSize)
}

func (c *ArticleController) RetrieveArticle(id uint) (*model.Article, error) {
	return c.RetrieveByID(id)
}

// RetrieveArticleWithPreloads demonstrates how to use preload to retrieve related fields
// For example, if Article has a relation to "Author" and "Comments", we can eager load them.
func (c *ArticleController) RetrieveArticleWithPreloads(id uint) (*model.Article, error) {
	return c.RetrieveByID(id, "Author", "Comments")
}

func (c *ArticleController) InsertArticle(article model.Article) error {
	return c.Insert(&article)
}

func (c *ArticleController) InsertArticles(articles []model.Article) error {
	return c.InsertMany(articles)
}

func (c *ArticleController) UpdateArticle(id uint, article *model.Article) error  {
	return c.UpdateByID(id, article)
}

func (c *ArticleController) DeleteArticle(id uint) error {
	return c.DeleteByID(id)
}
```

### Benefits
- Code Reusability: Reduces repetitive CRUD operations.
- Maintainability: Centralized logic for data operations.
- Flexibility: Can be reused for different models.