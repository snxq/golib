package structs

type Option func(*searcher)

func WithDelitimter(delitimter string) Option {
	return func(s *searcher) {
		if delitimter == "" {
			return
		}
		s.delitimter = delitimter
	}
}
