package values

type Value interface {
	Validate() error
}
