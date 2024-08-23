package models

import "github.com/MrSametBurgazoglu/enterprise/client"

type PostsPredicate struct {
	where []*client.WhereList
}

func (t *PostsPredicate) Where(w ...*client.Where) {
	t.where = nil
	wl := &client.WhereList{}
	wl.Items = append(wl.Items, w...)
	t.where = append(t.where, wl)
}

func (t *PostsPredicate) ORWhere(w ...*client.Where) {
	wl := &client.WhereList{}
	wl.Items = append(wl.Items, w...)
	t.where = append(t.where, wl)
}

func (t *PostsPredicate) IsIDEqual(v uint) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     PostsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsTitleEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     PostsTitleField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsContentEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     PostsContentField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsUserIDEqual(v uint) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     PostsUserIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsIDNotEqual(v uint) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     PostsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsTitleNotEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     PostsTitleField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsContentNotEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     PostsContentField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsUserIDNotEqual(v uint) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     PostsUserIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsIDIN(v ...uint) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     PostsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsTitleIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     PostsTitleField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsContentIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     PostsContentField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsUserIDIN(v ...uint) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     PostsUserIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsIDNotIN(v ...uint) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     PostsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsTitleNotIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     PostsTitleField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsContentNotIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     PostsContentField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IsUserIDNotIN(v ...uint) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     PostsUserIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IDGreaterThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.GT,
		Name:     PostsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IDGreaterEqualThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.GTE,
		Name:     PostsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IDLowerThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.LT,
		Name:     PostsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) IDLowerEqualThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.LTE,
		Name:     PostsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) UserIDGreaterThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.GT,
		Name:     PostsUserIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) UserIDGreaterEqualThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.GTE,
		Name:     PostsUserIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) UserIDLowerThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.LT,
		Name:     PostsUserIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *PostsPredicate) UserIDLowerEqualThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.LTE,
		Name:     PostsUserIDField,
		HasValue: true,
		Value:    v,
	}
}
