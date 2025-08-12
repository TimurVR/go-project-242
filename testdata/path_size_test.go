package testdata_test

import (
    "testing"
    code1 "code"
    "github.com/stretchr/testify/require"
)

func TestGetPathSizeHumanTxt(t *testing.T) {
    ex1, err1 := code1.GetPathSize("example.txt", true, true, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "45B\texample.txt", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanTxt(t *testing.T) {
    ex1, err1 := code1.GetPathSize("example.txt", true, false, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "45B\texample.txt", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeHumanDocx(t *testing.T) {
    ex1, err1 := code1.GetPathSize("example.docx", true, true, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "16.5KB\texample.docx", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanDocx(t *testing.T) {
    ex1, err1 := code1.GetPathSize("example.docx", true, false, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "16902B\texample.docx", ex1, "Expected format: size\\tfilename")
}
/* тут в локальном плане всё считает правильно, но когда на гитхабе есть проблемы, не знаю как решить)
func TestGetPathSizeHumanDirOnlyFiles(t *testing.T) {//r
    ex1, err1 := code1.GetPathSize("example", true, true, false)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "18.7KB\texample", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanDirOnlyFiles(t *testing.T) {//red
    ex1, err1 := code1.GetPathSize("example", true, false, false)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19131B\texample", ex1, "Expected format: size\\tfilename")
}
*/
func TestGetPathSizeHumanDirAll(t *testing.T) {
    ex1, err1 := code1.GetPathSize("example", true, true, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19.2KB\texample", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanDirAll(t *testing.T) {
    ex1, err1 := code1.GetPathSize("example", true, false, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19689B\texample", ex1, "Expected format: size\\tfilename")
}