package derror

func (de *DError) Error() string {
	return de.Err.Error()
}
