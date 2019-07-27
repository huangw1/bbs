/**
 * @Author: huangw1
 * @Date: 2019/7/30 13:06
 */

package model

type ArticleRequest struct {
	Tags       []string `json:"tags" form:"tags"`
	Title      string   `json:"title" form:"title"`
	Summary    string   `json:"summary" form:"summary"`
	Content    string   `json:"content" form:"content"`
	CategoryId int64    `json:"categoryId" form:"categoryId"`
}
