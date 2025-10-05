package metadata

type Builder struct {
	md *Metadata
}

func New() *Builder {
	return &Builder{md: &Metadata{}}
}

func (b *Builder) WithUserId(userId string) *Builder {
	b.md.UserId = userId
	return b
}

func (b *Builder) WithRole(role string) *Builder {
	b.md.Role = role
	return b
}

func (b *Builder) WithName(name string) *Builder {
	b.md.Name = name
	return b
}

func (b *Builder) WithEmail(email string) *Builder {
	b.md.Email = email
	return b
}

func (b *Builder) WithRegisterTimestamp(ts int64) *Builder {
	b.md.RegisterTimestamp = ts
	return b
}
