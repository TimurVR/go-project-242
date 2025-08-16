package testdata_test

import (
    "testing"
    code1 "code"
    "github.com/stretchr/testify/require"
)
// with Txt
func TestGetPathSizeHumanTxt(t *testing.T) {
    ex1, err1 := code1.GetPathSize("../testdata/example.txt", true, true, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "45B", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanTxt(t *testing.T) {
    ex1, err1 := code1.GetPathSize("../testdata/example.txt", true, false, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "45B", ex1, "Expected format: size\\tfilename")
}
// with Doc
func TestGetPathSizeHumanDocx(t *testing.T) {
    ex1, err1 := code1.GetPathSize("../testdata/example.docx", true, true, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "16.5KB", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanDocx(t *testing.T) {
    ex1, err1 := code1.GetPathSize("../testdata/example.docx", true, false, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "16902B", ex1, "Expected format: size\\tfilename")
}
// with Dir All==false && Recursion==false
func TestGetPathSizeHumanDir(t *testing.T) {
    ex1, err1 := code1.GetPathSize("../testdata/example", false, true, false)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "18.9KB", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanDir(t *testing.T) {
    ex1, err1 := code1.GetPathSize("../testdata/example", false, false, false)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19386B", ex1, "Expected format: size\\tfilename")
}
// with Dir All==true && Recursion==false
func TestGetPathSizeHumanDirAll(t *testing.T) {
    ex1, err1 := code1.GetPathSize("../testdata/example", false, true, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19.0KB", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanDirAll(t *testing.T) {
    ex1, err1 := code1.GetPathSize("../testdata/example", false, false, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19472B", ex1, "Expected format: size\\tfilename")
}
// with Dir All==false && Recursion==true
func TestGetPathSizeHumanDirRecursion(t *testing.T) {
    ex1, err1 := code1.GetPathSize("../testdata/example", true, true, false)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19.5KB", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanRecursion(t *testing.T) {
    ex1, err1 := code1.GetPathSize("../testdata/example", true, false, false)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19944B", ex1, "Expected format: size\\tfilename")
}
// with Dir All==true && Recursion==true
func TestGetPathSizeHumanTrue(t *testing.T) {
    ex1, err1 := code1.GetPathSize("../testdata/example", true, true, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "19.8KB", ex1, "Expected format: size\\tfilename")
}

func TestGetPathSizeNotHumanTrue(t *testing.T) {
    ex1, err1 := code1.GetPathSize("../testdata/example", true, false, true)
    if err1 != nil {
        t.Errorf(`Error in testing of getting path size %v`, err1)
    }
    require.Equal(t, "20258B", ex1, "Expected format: size\\tfilename")
}