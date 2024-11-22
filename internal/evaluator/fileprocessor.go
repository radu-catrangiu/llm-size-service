package evaluator

type gitLfsResponse struct {
	Files []*file
}

type file struct {
	Name string
	Size uint64
}

func (r gitLfsResponse) computeSizeBytes() uint64 {
	var totalBytes uint64 = 0

	for _, f := range r.Files {
		totalBytes += f.Size
	}

	return totalBytes
}
