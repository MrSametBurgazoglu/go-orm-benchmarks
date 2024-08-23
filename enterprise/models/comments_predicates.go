package models

import "github.com/MrSametBurgazoglu/enterprise/client"

type CommentsPredicate struct {
	where []*client.WhereList
}

func (t *CommentsPredicate) Where(w ...*client.Where) {
	t.where = nil
	wl := &client.WhereList{}
	wl.Items = append(wl.Items, w...)
	t.where = append(t.where, wl)
}

func (t *CommentsPredicate) ORWhere(w ...*client.Where) {
	wl := &client.WhereList{}
	wl.Items = append(wl.Items, w...)
	t.where = append(t.where, wl)
}

func (t *CommentsPredicate) IsIDEqual(v uint) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     CommentsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IsTextEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     CommentsTextField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IsPostIDEqual(v uint) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     CommentsPostIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IsIDNotEqual(v uint) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     CommentsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IsTextNotEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     CommentsTextField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IsPostIDNotEqual(v uint) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     CommentsPostIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IsIDIN(v ...uint) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     CommentsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IsTextIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     CommentsTextField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IsPostIDIN(v ...uint) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     CommentsPostIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IsIDNotIN(v ...uint) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     CommentsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IsTextNotIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     CommentsTextField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IsPostIDNotIN(v ...uint) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     CommentsPostIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IDGreaterThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.GT,
		Name:     CommentsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IDGreaterEqualThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.GTE,
		Name:     CommentsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IDLowerThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.LT,
		Name:     CommentsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) IDLowerEqualThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.LTE,
		Name:     CommentsIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) PostIDGreaterThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.GT,
		Name:     CommentsPostIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) PostIDGreaterEqualThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.GTE,
		Name:     CommentsPostIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) PostIDLowerThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.LT,
		Name:     CommentsPostIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *CommentsPredicate) PostIDLowerEqualThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.LTE,
		Name:     CommentsPostIDField,
		HasValue: true,
		Value:    v,
	}
}
