package _defer

func f1() int {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}

func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f3() (r int) {
	defer func(r *int) {
		*r = *r + 5
	}(&r)
	return 1
}
