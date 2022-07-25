package semvar

import (
	"fmt"
	"strconv"
	"strings"
)

type SemvarTag struct {
	Major int
	Minor int
	Patch int
}

func New(major, minor, patch int) *SemvarTag {
	return &SemvarTag{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func (sv *SemvarTag) ToString() string {
	return fmt.Sprintf("v%d.%d.%d", sv.Major, sv.Minor, sv.Patch)
}

func NewFromString(semvar string) (SemvarTag, error) {
	semvar = strings.TrimPrefix(semvar, "v")
	parts := strings.Split(semvar, ".")

	if len(parts) != 3 {
		return SemvarTag{}, fmt.Errorf("got invalid semvar tag: %s", semvar)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return SemvarTag{}, err
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return SemvarTag{}, err
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return SemvarTag{}, err
	}

	return SemvarTag{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}
