package models

import (
	"context"
	"github.com/MrSametBurgazoglu/enterprise/client"
)

const CommentsTableName = "comments"

const (
	CommentsIDField     string = "id"
	CommentsTextField   string = "text"
	CommentsPostIDField string = "post_id"
)

func NewComments(ctx context.Context, dc client.DatabaseClient) *Comments {
	v := &Comments{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.changedFields = make(map[string]any)
	v.result.Init()
	v.Default()
	return v
}

func NewRelationComments(ctx context.Context, dc client.DatabaseClient) *Comments {
	v := &Comments{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.changedFields = make(map[string]any)
	v.result.Init()
	return v
}

type Comments struct {
	id uint

	text string

	postid uint

	changedFields map[string]any
	serialFields  []*client.SelectedField

	ctx    context.Context
	client *client.Client
	CommentsPredicate
	relations *client.RelationList

	PostsList *PostsList

	result CommentsResult
}

func (t *Comments) GetDBName() string {
	return CommentsTableName
}

func (t *Comments) GetSelector() *CommentsResult {
	t.result.selectedFields = nil
	return &t.result
}

func (t *Comments) GetRelationList() *client.RelationList {
	return t.relations
}

func (t *Comments) IsExist() bool {
	var v uint
	return t.id != v
}

func (t *Comments) GetPrimaryKey() uint {
	return t.id
}

func NewCommentsList(ctx context.Context, dc client.DatabaseClient) *CommentsList {
	v := &CommentsList{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.result.Init()
	return v
}

func NewRelationCommentsList(ctx context.Context, dc client.DatabaseClient) *CommentsList {
	v := &CommentsList{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.result.Init()
	return v
}

type CommentsList struct {
	Items []*Comments

	ctx    context.Context
	client *client.Client
	CommentsPredicate
	order     []*client.Order
	paging    *client.Paging
	relations *client.RelationList
	result    CommentsResult
}

func (t *CommentsList) GetDBName() string {
	return CommentsTableName
}

func (t *CommentsList) GetRelationList() *client.RelationList {
	return t.relations
}

func (t *CommentsList) IsExist() bool {
	return t.Items[len(t.Items)-1].IsExist()
}

func (t *Comments) SetID(v uint) {
	t.id = v
	t.SetIDField()
}
func (t *Comments) SetText(v string) {
	t.text = v
	t.SetTextField()
}
func (t *Comments) SetPostID(v uint) {
	t.postid = v
	t.SetPostIDField()
}

func (t *Comments) SetIDNillable(v *uint) {
	if v == nil {
		return
	}
	t.SetID(*v)
}
func (t *Comments) SetTextNillable(v *string) {
	if v == nil {
		return
	}
	t.SetText(*v)
}
func (t *Comments) SetPostIDNillable(v *uint) {
	if v == nil {
		return
	}
	t.SetPostID(*v)
}

func (t *Comments) IDIN(v ...uint) bool {
	for _, x := range v {
		if t.id == x {
			return true
		}
	}
	return false
}

func (t *Comments) TextIN(v ...string) bool {
	for _, x := range v {
		if t.text == x {
			return true
		}
	}
	return false
}

func (t *Comments) PostIDIN(v ...uint) bool {
	for _, x := range v {
		if t.postid == x {
			return true
		}
	}
	return false
}

func (t *Comments) IDNotIN(v ...uint) bool {
	for _, x := range v {
		if t.id == x {
			return false
		}
	}
	return true
}

func (t *Comments) TextNotIN(v ...string) bool {
	for _, x := range v {
		if t.text == x {
			return false
		}
	}
	return true
}

func (t *Comments) PostIDNotIN(v ...uint) bool {
	for _, x := range v {
		if t.postid == x {
			return false
		}
	}
	return true
}

func (t *Comments) GetID() uint {
	return t.id
}
func (t *Comments) GetText() string {
	return t.text
}
func (t *Comments) GetPostID() uint {
	return t.postid
}

func (t *Comments) SetIDField() {
	t.changedFields[CommentsIDField] = t.id
}
func (t *Comments) SetTextField() {
	t.changedFields[CommentsTextField] = t.text
}
func (t *Comments) SetPostIDField() {
	t.changedFields[CommentsPostIDField] = t.postid
}

func (t *Comments) WithPostsList(opts ...func(*PostsList)) {
	t.PostsList = NewRelationPostsList(t.ctx, t.client.Database)
	for _, opt := range opts {
		opt(t.PostsList)
	}
	t.result.Posts = new(PostsResult)
	t.result.Posts.Init()
	t.result.relations = append(t.result.relations, t.result.Posts)
	t.result.relationsMap["posts"] = t.result.Posts
	for _, Relation := range t.PostsList.relations.Relations {
		t.result.Posts.relations = append(t.result.Posts.relations, Relation.RelationResult)
		t.result.Posts.relationsMap[Relation.RelationTable] = Relation.RelationResult
	}
	t.relations.Relations = append(t.relations.Relations,
		&client.Relation{
			RelationModel:  t.PostsList,
			RelationTable:  "posts",
			RelationResult: t.result.Posts,
			Where:          t.PostsList.where,

			RelationWhere: &client.RelationCondition{
				RelationValue: "post_id",
				TableValue:    "id",
			},
		},
	)
	t.relations.RelationMap["posts"] = t.relations.Relations[len(t.relations.Relations)-1]
}

func (t *CommentsList) WithPostsList(opts ...func(*PostsList)) {
	//t.PostsList = NewRelationPostsList(t.ctx, t.client.Database)
	v := NewRelationPostsList(t.ctx, t.client.Database)
	for _, opt := range opts {
		opt(v)
	}
	t.result.Posts = new(PostsResult)
	t.result.Posts.Init()
	t.result.relations = append(t.result.relations, t.result.Posts)
	t.result.relationsMap["posts"] = t.result.Posts
	for _, Relation := range v.relations.Relations {
		t.result.Posts.relations = append(t.result.Posts.relations, Relation.RelationResult)
		t.result.Posts.relationsMap[Relation.RelationTable] = Relation.RelationResult
	}
	t.relations.Relations = append(t.relations.Relations,
		&client.Relation{
			RelationModel:  v,
			RelationTable:  "posts",
			RelationResult: t.result.Posts,
			Where:          v.where,
			RelationWhere: &client.RelationCondition{
				RelationValue: "post_id",
				TableValue:    "id",
			},
		},
	)
	t.relations.RelationMap["posts"] = t.relations.Relations[len(t.relations.Relations)-1]
}

func (t *CommentsList) cleanPostsList() {
	Relation := t.Items[len(t.Items)-1].relations
	p := 0
	for i, v := range Relation.Relations {
		if v.RelationTable == "posts" {
			p = i
		}
	}
	Relation.Relations = append(Relation.Relations[:p], Relation.Relations[p+1:]...)
}

func (t *Comments) Default() {

	v := &client.SelectedField{Name: CommentsIDField, Value: &t.id}
	t.serialFields = append(t.serialFields, v)

}

func (t *Comments) SetResult(result client.Result) {
	if t == nil {
		v := NewRelationComments(t.ctx, t.client.Database)
		*t = *v
	}
	t.result = *result.(*CommentsResult)
}

func (t *CommentsList) SetResult(result client.Result) {
	if t == nil {
		v := NewRelationCommentsList(t.ctx, t.client.Database)
		*t = *v
	}
	t.result = *result.(*CommentsResult)
}

func (t *Comments) ScanResult() {
	t.id = t.result.id
	t.text = t.result.text
	t.postid = t.result.postid

	if _, ok := t.relations.RelationMap["posts"]; ok {
		if t.PostsList == nil {
			t.PostsList = NewRelationPostsList(t.ctx, t.client.Database)
		}
		t.PostsList.relations = t.relations.RelationMap["posts"].RelationModel.GetRelationList()
		t.PostsList.SetResult(t.result.relationsMap["posts"])
		t.PostsList.ScanResult()
	}
}

func (t *Comments) CheckPrimaryKey(v uint) bool {
	return t.id == v
}

func (t *CommentsList) ScanResult() {
	var v *Comments
	if len(t.Items) == 0 {
		v = NewRelationComments(t.ctx, t.client.Database)
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
		v = NewRelationComments(t.ctx, t.client.Database)
		t.Items = append(t.Items, v)
	}
	v.result = t.result
	v.relations = t.relations
	v.ScanResult()
}

func (t *Comments) Get() (error, bool) {
	return t.client.Get(t.ctx, t.where, t, &t.result)
}

func (t *Comments) Refresh() error {
	return t.client.Refresh(t.ctx, t, &t.result, CommentsIDField, t.id)
}

func (t *Comments) Create() error {
	return t.client.Create(t.ctx, CommentsTableName, t.changedFields, t.serialFields)
}

func (t *Comments) Update() error {
	return t.client.Update(t.ctx, CommentsTableName, t.changedFields, CommentsIDField, t.id)
}

func (t *Comments) Delete() error {
	return t.client.Delete(t.ctx, CommentsTableName, CommentsIDField, t.id)
}

func (t *CommentsList) List() (error, bool) {
	return t.client.List(t.ctx, t.where, t, &t.result, t.order, t.paging)
}

func (t *CommentsList) Aggregate(f func(aggregate *client.Aggregate)) (func() error, error) {
	a := new(client.Aggregate)
	f(a)
	return t.client.Aggregate(t.ctx, t.where, t, a)
}

func (t *CommentsList) Order(field string) *CommentsList {
	t.order = append(t.order, &client.Order{Field: field})
	return t
}

func (t *CommentsList) OrderDesc(field string) *CommentsList {
	t.order = append(t.order, &client.Order{Field: field, Desc: true})
	return t
}

func (t *CommentsList) Paging(skip, limit int) *CommentsList {
	t.paging = &client.Paging{Skip: skip, Limit: limit}
	return t
}

type CommentsResult struct {
	id     uint
	text   string
	postid uint

	selectedFields []*client.SelectedField

	Posts *PostsResult

	relations    []client.Result
	relationsMap map[string]client.Result
}

func (t *CommentsResult) Init() {
	t.selectedFields = []*client.SelectedField{}
	t.relationsMap = make(map[string]client.Result)
	t.prepare()
	t.SelectAll()
}

func (t *CommentsResult) GetSelectedFields() []*client.SelectedField {
	return t.selectedFields
}

func (t *CommentsResult) GetRelations() []client.Result {
	return t.relations
}

func (t *CommentsResult) prepare() {

}

func (t *CommentsResult) SelectID() {
	v := &client.SelectedField{Name: CommentsIDField, Value: &t.id}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *CommentsResult) SelectText() {
	v := &client.SelectedField{Name: CommentsTextField, Value: &t.text}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *CommentsResult) SelectPostID() {
	v := &client.SelectedField{Name: CommentsPostIDField, Value: &t.postid}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *CommentsResult) GetDBName() string {
	return CommentsTableName
}

func (t *CommentsResult) SelectAll() {
	t.SelectID()
	t.SelectText()
	t.SelectPostID()

}

func (t *CommentsResult) IsExist() bool {
	if t == nil {
		return false
	}
	var v uint
	return t.id != v
}
