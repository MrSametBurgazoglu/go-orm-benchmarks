package models

import (
	"context"
	"github.com/MrSametBurgazoglu/enterprise/client"
)

const PostsTableName = "posts"

const (
	PostsIDField      string = "id"
	PostsTitleField   string = "title"
	PostsContentField string = "content"
	PostsUserIDField  string = "user_id"
)

func NewPosts(ctx context.Context, dc client.DatabaseClient) *Posts {
	v := &Posts{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.changedFields = make(map[string]any)
	v.result.Init()
	v.Default()
	return v
}

func NewRelationPosts(ctx context.Context, dc client.DatabaseClient) *Posts {
	v := &Posts{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.changedFields = make(map[string]any)
	v.result.Init()
	return v
}

type Posts struct {
	id uint

	title string

	content string

	userid uint

	changedFields map[string]any
	serialFields  []*client.SelectedField

	ctx    context.Context
	client *client.Client
	PostsPredicate
	relations *client.RelationList

	Users        *Users
	CommentsList *CommentsList

	result PostsResult
}

func (t *Posts) GetDBName() string {
	return PostsTableName
}

func (t *Posts) GetSelector() *PostsResult {
	t.result.selectedFields = nil
	return &t.result
}

func (t *Posts) GetRelationList() *client.RelationList {
	return t.relations
}

func (t *Posts) IsExist() bool {
	var v uint
	return t.id != v
}

func (t *Posts) GetPrimaryKey() uint {
	return t.id
}

func NewPostsList(ctx context.Context, dc client.DatabaseClient) *PostsList {
	v := &PostsList{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.result.Init()
	return v
}

func NewRelationPostsList(ctx context.Context, dc client.DatabaseClient) *PostsList {
	v := &PostsList{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.result.Init()
	return v
}

type PostsList struct {
	Items []*Posts

	ctx    context.Context
	client *client.Client
	PostsPredicate
	order     []*client.Order
	paging    *client.Paging
	relations *client.RelationList
	result    PostsResult
}

func (t *PostsList) GetDBName() string {
	return PostsTableName
}

func (t *PostsList) GetRelationList() *client.RelationList {
	return t.relations
}

func (t *PostsList) IsExist() bool {
	return t.Items[len(t.Items)-1].IsExist()
}

func (t *Posts) SetID(v uint) {
	t.id = v
	t.SetIDField()
}
func (t *Posts) SetTitle(v string) {
	t.title = v
	t.SetTitleField()
}
func (t *Posts) SetContent(v string) {
	t.content = v
	t.SetContentField()
}
func (t *Posts) SetUserID(v uint) {
	t.userid = v
	t.SetUserIDField()
}

func (t *Posts) SetIDNillable(v *uint) {
	if v == nil {
		return
	}
	t.SetID(*v)
}
func (t *Posts) SetTitleNillable(v *string) {
	if v == nil {
		return
	}
	t.SetTitle(*v)
}
func (t *Posts) SetContentNillable(v *string) {
	if v == nil {
		return
	}
	t.SetContent(*v)
}
func (t *Posts) SetUserIDNillable(v *uint) {
	if v == nil {
		return
	}
	t.SetUserID(*v)
}

func (t *Posts) IDIN(v ...uint) bool {
	for _, x := range v {
		if t.id == x {
			return true
		}
	}
	return false
}

func (t *Posts) TitleIN(v ...string) bool {
	for _, x := range v {
		if t.title == x {
			return true
		}
	}
	return false
}

func (t *Posts) ContentIN(v ...string) bool {
	for _, x := range v {
		if t.content == x {
			return true
		}
	}
	return false
}

func (t *Posts) UserIDIN(v ...uint) bool {
	for _, x := range v {
		if t.userid == x {
			return true
		}
	}
	return false
}

func (t *Posts) IDNotIN(v ...uint) bool {
	for _, x := range v {
		if t.id == x {
			return false
		}
	}
	return true
}

func (t *Posts) TitleNotIN(v ...string) bool {
	for _, x := range v {
		if t.title == x {
			return false
		}
	}
	return true
}

func (t *Posts) ContentNotIN(v ...string) bool {
	for _, x := range v {
		if t.content == x {
			return false
		}
	}
	return true
}

func (t *Posts) UserIDNotIN(v ...uint) bool {
	for _, x := range v {
		if t.userid == x {
			return false
		}
	}
	return true
}

func (t *Posts) GetID() uint {
	return t.id
}
func (t *Posts) GetTitle() string {
	return t.title
}
func (t *Posts) GetContent() string {
	return t.content
}
func (t *Posts) GetUserID() uint {
	return t.userid
}

func (t *Posts) SetIDField() {
	t.changedFields[PostsIDField] = t.id
}
func (t *Posts) SetTitleField() {
	t.changedFields[PostsTitleField] = t.title
}
func (t *Posts) SetContentField() {
	t.changedFields[PostsContentField] = t.content
}
func (t *Posts) SetUserIDField() {
	t.changedFields[PostsUserIDField] = t.userid
}

func (t *Posts) WithUsers(opts ...func(*Users)) {
	t.Users = NewRelationUsers(t.ctx, t.client.Database)
	for _, opt := range opts {
		opt(t.Users)
	}
	t.result.Users = new(UsersResult)
	t.result.Users.Init()
	t.result.relations = append(t.result.relations, t.result.Users)
	t.result.relationsMap["users"] = t.result.Users
	for _, Relation := range t.Users.relations.Relations {
		t.result.Users.relations = append(t.result.Users.relations, Relation.RelationResult)
		t.result.Users.relationsMap[Relation.RelationTable] = Relation.RelationResult
	}
	t.relations.Relations = append(t.relations.Relations,
		&client.Relation{
			RelationModel:  t.Users,
			RelationTable:  "users",
			RelationResult: t.result.Users,
			Where:          t.Users.where,

			RelationWhere: &client.RelationCondition{
				RelationValue: "id",
				TableValue:    "user_id",
			},
		},
	)
	t.relations.RelationMap["users"] = t.relations.Relations[len(t.relations.Relations)-1]
}
func (t *Posts) WithCommentsList(opts ...func(*CommentsList)) {
	t.CommentsList = NewRelationCommentsList(t.ctx, t.client.Database)
	for _, opt := range opts {
		opt(t.CommentsList)
	}
	t.result.Comments = new(CommentsResult)
	t.result.Comments.Init()
	t.result.relations = append(t.result.relations, t.result.Comments)
	t.result.relationsMap["comments"] = t.result.Comments
	for _, Relation := range t.CommentsList.relations.Relations {
		t.result.Comments.relations = append(t.result.Comments.relations, Relation.RelationResult)
		t.result.Comments.relationsMap[Relation.RelationTable] = Relation.RelationResult
	}
	t.relations.Relations = append(t.relations.Relations,
		&client.Relation{
			RelationModel:  t.CommentsList,
			RelationTable:  "comments",
			RelationResult: t.result.Comments,
			Where:          t.CommentsList.where,

			RelationWhere: &client.RelationCondition{
				RelationValue: "post_id",
				TableValue:    "id",
			},
		},
	)
	t.relations.RelationMap["comments"] = t.relations.Relations[len(t.relations.Relations)-1]
}

func (t *PostsList) WithUsers(opts ...func(*Users)) {
	//t.Users = NewRelationUsers(t.ctx, t.client.Database)
	v := NewRelationUsers(t.ctx, t.client.Database)
	for _, opt := range opts {
		opt(v)
	}
	t.result.Users = new(UsersResult)
	t.result.Users.Init()
	t.result.relations = append(t.result.relations, t.result.Users)
	t.result.relationsMap["users"] = t.result.Users
	for _, Relation := range v.relations.Relations {
		t.result.Users.relations = append(t.result.Users.relations, Relation.RelationResult)
		t.result.Users.relationsMap[Relation.RelationTable] = Relation.RelationResult
	}
	t.relations.Relations = append(t.relations.Relations,
		&client.Relation{
			RelationModel:  v,
			RelationTable:  "users",
			RelationResult: t.result.Users,
			Where:          v.where,
			RelationWhere: &client.RelationCondition{
				RelationValue: "id",
				TableValue:    "user_id",
			},
		},
	)
	t.relations.RelationMap["users"] = t.relations.Relations[len(t.relations.Relations)-1]
}

func (t *PostsList) cleanUsers() {
	Relation := t.Items[len(t.Items)-1].relations
	p := 0
	for i, v := range Relation.Relations {
		if v.RelationTable == "users" {
			p = i
		}
	}
	Relation.Relations = append(Relation.Relations[:p], Relation.Relations[p+1:]...)
}
func (t *PostsList) WithCommentsList(opts ...func(*CommentsList)) {
	//t.CommentsList = NewRelationCommentsList(t.ctx, t.client.Database)
	v := NewRelationCommentsList(t.ctx, t.client.Database)
	for _, opt := range opts {
		opt(v)
	}
	t.result.Comments = new(CommentsResult)
	t.result.Comments.Init()
	t.result.relations = append(t.result.relations, t.result.Comments)
	t.result.relationsMap["comments"] = t.result.Comments
	for _, Relation := range v.relations.Relations {
		t.result.Comments.relations = append(t.result.Comments.relations, Relation.RelationResult)
		t.result.Comments.relationsMap[Relation.RelationTable] = Relation.RelationResult
	}
	t.relations.Relations = append(t.relations.Relations,
		&client.Relation{
			RelationModel:  v,
			RelationTable:  "comments",
			RelationResult: t.result.Comments,
			Where:          v.where,
			RelationWhere: &client.RelationCondition{
				RelationValue: "post_id",
				TableValue:    "id",
			},
		},
	)
	t.relations.RelationMap["comments"] = t.relations.Relations[len(t.relations.Relations)-1]
}

func (t *PostsList) cleanCommentsList() {
	Relation := t.Items[len(t.Items)-1].relations
	p := 0
	for i, v := range Relation.Relations {
		if v.RelationTable == "comments" {
			p = i
		}
	}
	Relation.Relations = append(Relation.Relations[:p], Relation.Relations[p+1:]...)
}

func (t *Posts) Default() {

	v := &client.SelectedField{Name: PostsIDField, Value: &t.id}
	t.serialFields = append(t.serialFields, v)

}

func (t *Posts) SetResult(result client.Result) {
	if t == nil {
		v := NewRelationPosts(t.ctx, t.client.Database)
		*t = *v
	}
	t.result = *result.(*PostsResult)
}

func (t *PostsList) SetResult(result client.Result) {
	if t == nil {
		v := NewRelationPostsList(t.ctx, t.client.Database)
		*t = *v
	}
	t.result = *result.(*PostsResult)
}

func (t *Posts) ScanResult() {
	t.id = t.result.id
	t.title = t.result.title
	t.content = t.result.content
	t.userid = t.result.userid

	if _, ok := t.relations.RelationMap["users"]; ok {
		if t.Users == nil {
			t.Users = NewRelationUsers(t.ctx, t.client.Database)
		}
		t.Users.relations = t.relations.RelationMap["users"].RelationModel.GetRelationList()
		t.Users.SetResult(t.result.relationsMap["users"])
		t.Users.ScanResult()
	}
	if _, ok := t.relations.RelationMap["comments"]; ok {
		if t.CommentsList == nil {
			t.CommentsList = NewRelationCommentsList(t.ctx, t.client.Database)
		}
		t.CommentsList.relations = t.relations.RelationMap["comments"].RelationModel.GetRelationList()
		t.CommentsList.SetResult(t.result.relationsMap["comments"])
		t.CommentsList.ScanResult()
	}
}

func (t *Posts) CheckPrimaryKey(v uint) bool {
	return t.id == v
}

func (t *PostsList) ScanResult() {
	var v *Posts
	if len(t.Items) == 0 {
		v = NewRelationPosts(t.ctx, t.client.Database)
		t.Items = append(t.Items, v)
	} else {
		for _, item := range t.Items {
			if item.CheckPrimaryKey(t.result.id) {
				v = item
				break
			}
		}
	}
	if v == nil {
		v = NewRelationPosts(t.ctx, t.client.Database)
		t.Items = append(t.Items, v)
	}
	v.result = t.result
	v.relations = t.relations
	v.ScanResult()
}

func (t *Posts) Get() (error, bool) {
	return t.client.Get(t.ctx, t.where, t, &t.result)
}

func (t *Posts) Refresh() error {
	return t.client.Refresh(t.ctx, t, &t.result, PostsIDField, t.id)
}

func (t *Posts) Create() error {
	return t.client.Create(t.ctx, PostsTableName, t.changedFields, t.serialFields)
}

func (t *Posts) Update() error {
	return t.client.Update(t.ctx, PostsTableName, t.changedFields, PostsIDField, t.id)
}

func (t *Posts) Delete() error {
	return t.client.Delete(t.ctx, PostsTableName, PostsIDField, t.id)
}

func (t *PostsList) List() (error, bool) {
	return t.client.List(t.ctx, t.where, t, &t.result, t.order, t.paging)
}

func (t *PostsList) Aggregate(f func(aggregate *client.Aggregate)) (func() error, error) {
	a := new(client.Aggregate)
	f(a)
	return t.client.Aggregate(t.ctx, t.where, t, a)
}

func (t *PostsList) Order(field string) *PostsList {
	t.order = append(t.order, &client.Order{Field: field})
	return t
}

func (t *PostsList) OrderDesc(field string) *PostsList {
	t.order = append(t.order, &client.Order{Field: field, Desc: true})
	return t
}

func (t *PostsList) Paging(skip, limit int) *PostsList {
	t.paging = &client.Paging{Skip: skip, Limit: limit}
	return t
}

type PostsResult struct {
	id      uint
	title   string
	content string
	userid  uint

	selectedFields []*client.SelectedField

	Users    *UsersResult
	Comments *CommentsResult

	relations    []client.Result
	relationsMap map[string]client.Result
}

func (t *PostsResult) Init() {
	t.selectedFields = []*client.SelectedField{}
	t.relationsMap = make(map[string]client.Result)
	t.prepare()
	t.SelectAll()
}

func (t *PostsResult) GetSelectedFields() []*client.SelectedField {
	return t.selectedFields
}

func (t *PostsResult) GetRelations() []client.Result {
	return t.relations
}

func (t *PostsResult) prepare() {

}

func (t *PostsResult) SelectID() {
	v := &client.SelectedField{Name: PostsIDField, Value: &t.id}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *PostsResult) SelectTitle() {
	v := &client.SelectedField{Name: PostsTitleField, Value: &t.title}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *PostsResult) SelectContent() {
	v := &client.SelectedField{Name: PostsContentField, Value: &t.content}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *PostsResult) SelectUserID() {
	v := &client.SelectedField{Name: PostsUserIDField, Value: &t.userid}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *PostsResult) GetDBName() string {
	return PostsTableName
}

func (t *PostsResult) SelectAll() {
	t.SelectID()
	t.SelectTitle()
	t.SelectContent()
	t.SelectUserID()

}

func (t *PostsResult) IsExist() bool {
	if t == nil {
		return false
	}
	var v uint
	return t.id != v
}
