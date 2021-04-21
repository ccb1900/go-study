package gmp

type G struct {
	m *M
}

type M struct {
	g0   *G
	curg *G
	p    *P
	oldp *P
}

type P struct {
	m *M
}
