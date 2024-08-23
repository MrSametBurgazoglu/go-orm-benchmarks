package models

import (
	"context"
	"github.com/MrSametBurgazoglu/enterprise/client"
)

const UsersTableName = "users"

const (
	UsersIDField    string = "id"
	UsersNameField  string = "name"
	UsersEmailField string = "email"
)

func NewUsers(ctx context.Context, dc client.DatabaseClient) *Users {
	v := &Users{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.changedFields = make(map[string]any)
	v.result.Init()
	v.Default()
	return v
}

func NewRelationUsers(ctx context.Context, dc client.DatabaseClient) *Users {
	v := &Users{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.changedFields = make(map[string]any)
	v.result.Init()
	return v
}

type Users struct {
	id uint

	name string

	email string

	changedFields map[string]any
	serialFields  []*client.SelectedField

	ctx    context.Context
	client *client.Client
	UsersPredicate
	relations *client.RelationList

	PostsList *PostsList

	result UsersResult
}

func (t *Users) GetDBName() string {
	return UsersTableName
}

func (t *Users) GetSelector() *UsersResult {
	t.result.selectedFields = nil
	return &t.result
}

func (t *Users) GetRelationList() *client.RelationList {
	return t.relations
}

func (t *Users) IsExist() bool {
	var v uint
	return t.id != v
}

func (t *Users) GetPrimaryKey() uint {
	return t.id
}

func NewUsersList(ctx context.Context, dc client.DatabaseClient) *UsersList {
	v := &UsersList{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.result.Init()
	return v
}

func NewRelationUsersList(ctx context.Context, dc client.DatabaseClient) *UsersList {
	v := &UsersList{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.result.Init()
	return v
}

type UsersList struct {
	Items []*Users

	ctx    context.Context
	client *client.Client
	UsersPredicate
	order     []*client.Order
	paging    *client.Paging
	relations *client.RelationList
	result    UsersResult
}

func (t *UsersList) GetDBName() string {
	return UsersTableName
}

func (t *UsersList) GetRelationList() *client.RelationList {
	return t.relations
}

func (t *UsersList) IsExist() bool {
	return t.Items[len(t.Items)-1].IsExist()
}

func (t *Users) SetID(v uint) {
	t.id = v
	t.SetIDField()
}
func (t *Users) SetName(v string) {
	t.name = v
	t.SetNameField()
}
func (t *Users) SetEmail(v string) {
	t.email = v
	t.SetEmailField()
}

func (t *Users) SetIDNillable(v *uint) {
	if v == nil {
		return
	}
	t.SetID(*v)
}
func (t *Users) SetNameNillable(v *string) {
	if v == nil {
		return
	}
	t.SetName(*v)
}
func (t *Users) SetEmailNillable(v *string) {
	if v == nil {
		return
	}
	t.SetEmail(*v)
}

func (t *Users) IDIN(v ...uint) bool {
	for _, x := range v {
		if t.id == x {
			return true
		}
	}
	return false
}

func (t *Users) NameIN(v ...string) bool {
	for _, x := range v {
		if t.name == x {
			return true
		}
	}
	return false
}

func (t *Users) EmailIN(v ...string) bool {
	for _, x := range v {
		if t.email == x {
			return true
		}
	}
	return false
}

func (t *Users) IDNotIN(v ...uint) bool {
	for _, x := range v {
		if t.id == x {
			return false
		}
	}
	return true
}

func (t *Users) NameNotIN(v ...string) bool {
	for _, x := range v {
		if t.name == x {
			return false
		}
	}
	return true
}

func (t *Users) EmailNotIN(v ...string) bool {
	for _, x := range v {
		if t.email == x {
			return false
		}
	}
	return true
}

func (t *Users) GetID() uint {
	return t.id
}
func (t *Users) GetName() string {
	return t.name
}
func (t *Users) GetEmail() string {
	return t.email
}

func (t *Users) SetIDField() {
	t.changedFields[UsersIDField] = t.id
}
func (t *Users) SetNameField() {
	t.changedFields[UsersNameField] = t.name
}
func (t *Users) SetEmailField() {
	t.changedFields[UsersEmailField] = t.email
}

func (t *Users) WithPostsList(opts ...func(*PostsList)) {
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
				RelationValue: "user_id",
				TableValue:    "id",
			},
		},
	)
	t.relations.RelationMap["posts"] = t.relations.Relations[len(t.relations.Relations)-1]
}

func (t *UsersList) WithPostsList(opts ...func(*PostsList)) {
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
				RelationValue: "user_id",
				TableValue:    "id",
			},
		},
	)
	t.relations.RelationMap["posts"] = t.relations.Relations[len(t.relations.Relations)-1]
}

func (t *UsersList) cleanPostsList() {
	Relation := t.Items[len(t.Items)-1].relations
	p := 0
	for i, v := range Relation.Relations {
		if v.RelationTable == "posts" {
			p = i
		}
	}
	Relation.Relations = append(Relation.Relations[:p], Relation.Relations[p+1:]...)
}

func (t *Users) Default() {

	v := &client.SelectedField{Name: UsersIDField, Value: &t.id}
	t.serialFields = append(t.serialFields, v)

}

func (t *Users) SetResult(result client.Result) {
	if t == nil {
		v := NewRelationUsers(t.ctx, t.client.Database)
		*t = *v
	}
	t.result = *result.(*UsersResult)
}

func (t *UsersList) SetResult(result client.Result) {
	if t == nil {
		v := NewRelationUsersList(t.ctx, t.client.Database)
		*t = *v
	}
	t.result = *result.(*UsersResult)
}

func (t *Users) ScanResult() {
	t.id = t.result.id
	t.name = t.result.name
	t.email = t.result.email

	if _, ok := t.relations.RelationMap["posts"]; ok {
		if t.PostsList == nil {
			t.PostsList = NewRelationPostsList(t.ctx, t.client.Database)
		}
		t.PostsList.relations = t.relations.RelationMap["posts"].RelationModel.GetRelationList()
		t.PostsList.SetResult(t.result.relationsMap["posts"])
		t.PostsList.ScanResult()
	}
}

func (t *Users) CheckPrimaryKey(v uint) bool {
	return t.id == v
}

func (t *UsersList) ScanResult() {
	var v *Users
	if len(t.Items) == 0 {
		v = NewRelationUsers(t.ctx, t.client.Database)
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
		v = NewRelationUsers(t.ctx, t.client.Database)
		t.Items = append(t.Items, v)
	}
	v.result = t.result
	v.relations = t.relations
	v.ScanResult()
}

func (t *Users) Get() (error, bool) {
	return t.client.Get(t.ctx, t.where, t, &t.result)
}

func (t *Users) Refresh() error {
	return t.client.Refresh(t.ctx, t, &t.result, UsersIDField, t.id)
}

func (t *Users) Create() error {
	return t.client.Create(t.ctx, UsersTableName, t.changedFields, t.serialFields)
}

func (t *Users) Update() error {
	return t.client.Update(t.ctx, UsersTableName, t.changedFields, UsersIDField, t.id)
}

func (t *Users) Delete() error {
	return t.client.Delete(t.ctx, UsersTableName, UsersIDField, t.id)
}

func (t *UsersList) List() (error, bool) {
	return t.client.List(t.ctx, t.where, t, &t.result, t.order, t.paging)
}

func (t *UsersList) Aggregate(f func(aggregate *client.Aggregate)) (func() error, error) {
	a := new(client.Aggregate)
	f(a)
	return t.client.Aggregate(t.ctx, t.where, t, a)
}

func (t *UsersList) Order(field string) *UsersList {
	t.order = append(t.order, &client.Order{Field: field})
	return t
}

func (t *UsersList) OrderDesc(field string) *UsersList {
	t.order = append(t.order, &client.Order{Field: field, Desc: true})
	return t
}

func (t *UsersList) Paging(skip, limit int) *UsersList {
	t.paging = &client.Paging{Skip: skip, Limit: limit}
	return t
}

type UsersResult struct {
	id    uint
	name  string
	email string

	selectedFields []*client.SelectedField

	Posts *PostsResult

	relations    []client.Result
	relationsMap map[string]client.Result
}

func (t *UsersResult) Init() {
	t.selectedFields = []*client.SelectedField{}
	t.relationsMap = make(map[string]client.Result)
	t.prepare()
	t.SelectAll()
}

func (t *UsersResult) GetSelectedFields() []*client.SelectedField {
	return t.selectedFields
}

func (t *UsersResult) GetRelations() []client.Result {
	return t.relations
}

func (t *UsersResult) prepare() {

}

func (t *UsersResult) SelectID() {
	v := &client.SelectedField{Name: UsersIDField, Value: &t.id}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *UsersResult) SelectName() {
	v := &client.SelectedField{Name: UsersNameField, Value: &t.name}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *UsersResult) SelectEmail() {
	v := &client.SelectedField{Name: UsersEmailField, Value: &t.email}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *UsersResult) GetDBName() string {
	return UsersTableName
}

func (t *UsersResult) SelectAll() {
	t.SelectID()
	t.SelectName()
	t.SelectEmail()

}

func (t *UsersResult) IsExist() bool {
	if t == nil {
		return false
	}
	var v uint
	return t.id != v
}
