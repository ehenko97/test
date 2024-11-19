package service

type producer interface {
	produce() ([]string, error)
}

type presenter interface {
	present([]string) error
}

type Service struct {
	prod producer
	pres presenter
}

func NewService(prod producer, pres presenter) *Service {
	return &Service{prod: prod, pres: pres}
}

func (s *Service) maskLink(lines []string) []string {
	maskedLines := make([]string, len(lines))
	for i, line := range lines {
		byteSlice := []byte(line)
		mask := []byte("http://")
		s.maskLinkBytes(byteSlice, mask)
		maskedLines[i] = string(byteSlice)
	}
	return maskedLines
}

func (s *Service) maskLinkBytes(byteSlice, mask []byte) {
	httpLength := len(mask)
	for i := 0; i <= len(byteSlice)-httpLength; i++ {
		match := true
		for j := 0; j < httpLength; j++ {
			if byteSlice[i+j] != mask[j] {
				match = false
				break
			}
		}
		if match {
			endIndex := i + httpLength
			for endIndex < len(byteSlice) && (byteSlice[endIndex] >= 'a' && byteSlice[endIndex] <= 'z' ||
				byteSlice[endIndex] >= 'A' && byteSlice[endIndex] <= 'Z' ||
				byteSlice[endIndex] >= '0' && byteSlice[endIndex] <= '9' ||
				byteSlice[endIndex] == '.' || byteSlice[endIndex] == '/' || byteSlice[endIndex] == ':' || byteSlice[endIndex] == '?') {
				endIndex++
			}
			for k := i + httpLength; k < endIndex; k++ {
				byteSlice[k] = '*'
			}
		}
	}
}

func (s *Service) Run() error {
	lines, err := s.prod.produce()
	if err != nil {
		return err
	}
	maskedLines := s.maskLink(lines)
	return s.pres.present(maskedLines)
}
