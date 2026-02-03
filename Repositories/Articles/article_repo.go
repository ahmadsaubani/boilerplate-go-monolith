package Articles

import (
	"boilerplate-go/DTO/Articles"
	"boilerplate-go/DTO/General"
	"boilerplate-go/Libraries/Helpers"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (a article) GetAllArticles(params General.RequestParamDTO) (resp []Articles.ArticleDTO, total int, err error) {
	conn := a.dbCon.PostgreMainCon()
	parts := strings.Split(params.OrderBy, ",")
	direction := "asc"
	field := "created_at"

	if len(parts) == 2 {
		field = strings.TrimSpace(strings.ToLower(parts[0]))
		direction = strings.TrimSpace(strings.ToLower(parts[1]))
	}

	query := fmt.Sprintf(`
	SELECT
		a.id,
		a.uuid,
		a.title,
		a.description,
		a.url,
		a.image_url,
		a.published_at,
		a.source,
		a.created_at,
		a.updated_at,
		COALESCE(
			json_agg(
				DISTINCT jsonb_build_object(
					'id', c.id,
					'uuid', c.uuid,
					'label', c.label,
					'value', c.value
				)
			) FILTER (WHERE c.id IS NOT NULL),
			'[]'
		) AS categories,
		COUNT(*) OVER() AS total
	FROM articles a
	LEFT JOIN article_categories ac ON ac.article_id = a.id
	LEFT JOIN categories c ON c.id = ac.category_id
	GROUP BY a.id
	ORDER BY %s %s
	LIMIT $1 OFFSET $2
`, field, direction)

	args := []interface{}{params.PerPage, params.Page}
	rows, err := conn.Query(query, args...)
	if err != nil {
		Helpers.LogInfo(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var res Articles.ArticleDTO
		var categories []byte
		err = rows.Scan(
			&res.ID,
			&res.UUID,
			&res.Title,
			&res.Description,
			&res.URL,
			&res.ImageURL,
			&res.PublishedAt,
			&res.Source,
			&res.CreatedAt,
			&res.UpdatedAt,
			&categories,
			&total,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return
			}
			return
		}

		err = json.Unmarshal(categories, &res.Categories)
		if err != nil {
			return
		}

		resp = append(resp, res)
	}
	return
}

func (a article) SaveArticles(items []Articles.ArticleDTO) (err error) {
	db := a.dbCon.PostgreMainCon()

	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	for _, item := range items {
		var articleID int
		err = tx.QueryRow(`
			SELECT id FROM articles WHERE url = $1
		`, item.URL).Scan(&articleID)

		if err != nil {
			if err != sql.ErrNoRows {
				return
			}

			articleUUID := uuid.New()

			err = tx.QueryRow(`
				INSERT INTO articles (
					uuid,
					title,
					description,
					url,
					image_url,
					published_at,
					source,
					created_at,
					updated_at
				) VALUES ($1,$2,$3,$4,$5,$6,$7,now(),now())
				RETURNING id
			`,
				articleUUID,
				item.Title,
				item.Description,
				item.URL,
				item.ImageURL,
				item.PublishedAt,
				item.Source,
			).Scan(&articleID)

			if err != nil {
				return
			}
		}
		for _, cat := range item.Categories {
			value := strings.ToLower(cat.Value)
			var categoryID int
			err = tx.QueryRow(`
				SELECT id FROM categories WHERE value = $1
			`, value).Scan(&categoryID)

			if err != nil {
				if err != sql.ErrNoRows {
					return
				}
				categoryUUID := uuid.New()

				err = tx.QueryRow(`
					INSERT INTO categories (
						uuid,
						label,
						value,
					    created_at,
					    updated_at
					)
					VALUES ($1,$2,$3,now(),now())
					RETURNING id
				`, categoryUUID, cat.Label, value).Scan(&categoryID)

				if err != nil {
					return
				}
			}
			_, err = tx.Exec(`
				INSERT INTO article_categories (article_id, category_id)
				SELECT $1, $2
				WHERE NOT EXISTS (
					SELECT 1 FROM article_categories
					WHERE article_id = $1 AND category_id = $2
				)
			`, articleID, categoryID)

			if err != nil {
				return
			}
		}
	}

	return tx.Commit()
}
