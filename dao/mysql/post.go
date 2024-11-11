package mysql

import "web-app/models"

func InsertPost(date *models.Post) (err error) {
	sqlstr := "insert into post(post_id,title,content,author_id,community_id) values (?,?,?,?,?)"
	_, err = db.Exec(sqlstr, date.ID, date.Title, date.Content, date.AuthorID, date.CommunityID)
	if err != nil {
		return
	}
	return
}
