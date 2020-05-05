package store

import (
	"context"
	"database/sql"
	"github.com/c8112002/news-api/entities"
	"github.com/c8112002/news-api/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type ArticleStore struct {
	db *sql.DB
	ctx context.Context
}

func NewArticleStore(db *sql.DB, ctx context.Context) *ArticleStore {
	return &ArticleStore{
		db:  db,
		ctx: ctx,
	}
}

// TODO: LimitとAfterを渡せるようにする
func (as *ArticleStore) GetAllArticles() ([]*entities.Article, error) {
	articles, err := models.Articles(qm.Load(qm.Rels(models.ArticleRels.Site)), qm.Load(qm.Rels(models.ArticleRels.Tags))).All(as.ctx, as.db)
	if err != nil {
		return nil, err
	}
	var res []*entities.Article
	for _, a := range articles {
		s := a.R.Site
		site := newSite(s)
		tags := newTags(a.R.Tags)
		image := newArticleImage(a.Image.String)
		article := entities.NewArticle(a.ID, a.OriginalID, a.Title, a.URL, image, a.CrawledAt, site, tags)
		res = append(res, article)
	}
	return res, nil
}

func newArticleImage(image string) *entities.ArticleImage {
	return entities.NewArticleImage(image)
}

