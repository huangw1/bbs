/**
 * @Author: huangw1
 * @Date: 2019/7/27 15:31
 */

package interior

import (
	"github.com/huangw1/bbs/pkg/config"
	"os"
	"path/filepath"
	"strconv"
)

func BuildAbsUrl(path string) string {
	if len(path) == 0 {
		return config.Conf.BaseUrl
	}
	return config.Conf.BaseUrl + string(os.PathSeparator) + path
}

func BuildBaseUrl() string {
	return BuildAbsUrl("")
}

func BuildUserUrl(userId int64) string {
	return BuildAbsUrl(filepath.Join("user", strconv.FormatInt(userId, 10)))
}

func BuildArticleUrl(articleId int64) string {
	return BuildAbsUrl(filepath.Join("article", strconv.FormatInt(articleId, 10)))
}

func BuildArticlesUrl(page int) string {
	if page < 1 {
		page = 1
	}
	return BuildAbsUrl(filepath.Join("articles", strconv.Itoa(page)))
}

func BuildCategoryArticlesUrl(categoryId int64, page int) string {
	if page < 1 {
		page = 1
	}
	return BuildAbsUrl(filepath.Join("articles", "cat", strconv.FormatInt(categoryId, 10), strconv.Itoa(page)))
}

func BuildTagArticlesUrl(tagId int64, page int) string {
	if page < 1 {
		page = 1
	}
	return BuildAbsUrl(filepath.Join("articles", "tag", strconv.FormatInt(tagId, 10), strconv.Itoa(page)))
}

func BuildUserArticlesUrl(userId int64, page int) string {
	if page < 1 {
		page = 1
	}
	return BuildAbsUrl(filepath.Join("user", strconv.FormatInt(userId, 10), "articles", strconv.Itoa(page)))
}

func BuildUserFavoritesUrl(userId int64, page int) string {
	if page < 1 {
		page = 1
	}
	return BuildAbsUrl(filepath.Join("user", strconv.FormatInt(userId, 10), "favorites", strconv.Itoa(page)))
}

func BuildTagsUrl(page int) string {
	if page < 1 {
		page = 1
	}
	return BuildAbsUrl(filepath.Join("tags", strconv.Itoa(page)))
}

func BuildMessagesUrl(page int) string {
	if page < 1 {
		page = 1
	}
	return BuildAbsUrl(filepath.Join("messages", strconv.Itoa(page)))
}

func BuildTopicUrl(topicId int64) string {
	return BuildAbsUrl(filepath.Join("topic", strconv.FormatInt(topicId, 10)))
}

func BuildTopicsUrl(page int) string {
	if page < 1 {
		page = 1
	}
	return BuildAbsUrl(filepath.Join("topics", strconv.Itoa(page)))
}

func BuildTagTopicsUrl(tagId int64, page int) string {
	if page < 1 {
		page = 1
	}
	return BuildAbsUrl(filepath.Join("topics", "tag", strconv.FormatInt(tagId, 10), strconv.Itoa(page)))
}

func BuildUserTopicsUrl(userId int64, page int) string {
	if page < 1 {
		page = 1
	}
	return BuildAbsUrl(filepath.Join("user", strconv.FormatInt(userId, 10), "topics", strconv.Itoa(page)))
}
