package options

import "fmt"

type Options struct {
	UID     string
	GID     string
	Flags   string
	Company string
	Gender  bool
}

type Option func(*Options)

func WithUID(userID string) Option {
	return func(args *Options) {
		args.UID = userID
	}
}

func WithGID(groupID string) Option {
	return func(args *Options) {
		args.GID = groupID
	}
}

func WithFlags(flag string) Option {
	return func(args *Options) {
		args.Flags = flag
	}
}

func WithCompany(company string) Option {
	return func(args *Options) {
		args.Company = company
	}
}

func WithGender(gender bool) Option {
	return func(args *Options) {
		args.Gender = gender
	}
}

func NewDefaultOptions(setters ...Option) {
	options := &Options{
		UID:     "default UID",
		GID:     "default GID",
		Flags:   "default Flags",
		Company: "default Company",
		Gender:  true,
	}

	for _, setter := range setters {
		setter(options)
	}
	gender := "famale"
	if options.Gender {
		gender = "male"
	}
	fmt.Println("----------------------")
	fmt.Println("im am: ", "\nfrom: ", options.Company, "\ngender: ", gender, "\nUID: ", options.UID)
}
