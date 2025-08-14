package testdata_test

import (
    "testing"
    code1 "code"
    "github.com/stretchr/testify/require"
)
// with Txt
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
// with Doc
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
// with Dir All==false && Recursion==false
func TestGetPathSizeHumanDir(t *testing.T) {//red
    ex1, err1 := code1.GetPathSize("example", false, true, false)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "18.9KB\texample", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanDir(t *testing.T) {//red
    ex1, err1 := code1.GetPathSize("example", false, false, false)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19386B\texample", ex1, "Expected format: size\\tfilename")
}
// with Dir All==true && Recursion==false
func TestGetPathSizeHumanDirAll(t *testing.T) {//red
    ex1, err1 := code1.GetPathSize("example", false, true, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19.0KB\texample", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanDirAll(t *testing.T) {//red
    ex1, err1 := code1.GetPathSize("example", false, false, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19472B\texample", ex1, "Expected format: size\\tfilename")
}
// with Dir All==false && Recursion==true
func TestGetPathSizeHumanDirRecursion(t *testing.T) {//red
    ex1, err1 := code1.GetPathSize("example", true, true, false)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19.5KB\texample", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanRecursion(t *testing.T) {//red
    ex1, err1 := code1.GetPathSize("example", true, false, false)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19944B\texample", ex1, "Expected format: size\\tfilename")
}
// with Dir All==true && Recursion==true
func TestGetPathSizeHumanTrue(t *testing.T) {//red
    ex1, err1 := code1.GetPathSize("example", true, true, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19.8KB\texample", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanTrue(t *testing.T) {//red
    ex1, err1 := code1.GetPathSize("example", true, false, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "20258B\texample", ex1, "Expected format: size\\tfilename")
}