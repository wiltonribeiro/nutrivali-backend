package models

type Article struct {
	Description string `bson:"description" json:"description"`
	Source string `bson:"source" json:"source"`
	Author interface{} `bson:"author" json:"author"`
	Title string `bson:"title" json:"title"`
	URL string `bson:"url" json:"url"`
	URLToImage string `bson:"urlToImage" json:"urlToImage"`
	PublishedAt string `bson:"publishedAt" json:"publishedAt"`
	Content interface{} `bson:"content" json:"content"`

}
