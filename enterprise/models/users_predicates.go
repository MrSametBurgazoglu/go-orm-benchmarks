package models

import "github.com/MrSametBurgazoglu/enterprise/client"

type UsersPredicate struct {
	where []*client.WhereList
}

func (t *UsersPredicate) Where(w ...*client.Where) {
	t.where = nil
	wl := &client.WhereList{}
	wl.Items = append(wl.Items, w...)
	t.where = append(t.where, wl)
}

func (t *UsersPredicate) ORWhere(w ...*client.Where) {
	wl := &client.WhereList{}
	wl.Items = append(wl.Items, w...)
	t.where = append(t.where, wl)
}

func (t *UsersPredicate) IsIDEqual(v uint) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     UsersIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IsNameEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     UsersNameField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IsEmailEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     UsersEmailField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IsIDNotEqual(v uint) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     UsersIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IsNameNotEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     UsersNameField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IsEmailNotEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     UsersEmailField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IsIDIN(v ...uint) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     UsersIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IsNameIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     UsersNameField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IsEmailIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     UsersEmailField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IsIDNotIN(v ...uint) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     UsersIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IsNameNotIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     UsersNameField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IsEmailNotIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     UsersEmailField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IDGreaterThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.GT,
		Name:     UsersIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IDGreaterEqualThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.GTE,
		Name:     UsersIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IDLowerThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.LT,
		Name:     UsersIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *UsersPredicate) IDLowerEqualThan(v uint) *client.Where {
	return &client.Where{
		Type:     client.LTE,
		Name:     UsersIDField,
		HasValue: true,
		Value:    v,
	}
}
