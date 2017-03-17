package init

import "testing"

func TestString2Lines(t *testing.T) {
	sampleStringAscii := "0123456789"
	sampleStringMultiByte := "あいうえお"

	if lines := string2Lines(sampleStringAscii, 0); len(lines) != 1 {
		t.Error(string2Lines(sampleStringAscii, 0))
	} else if lines[0].S != "0123456789" {
		t.Errorf("I wanted %v but it was %v.", "0123456789", lines[0].String())
	}

	if lines := string2Lines(sampleStringAscii, 5); len(lines) != 2 {
		t.Error(string2Lines(sampleStringAscii, 5))
	} else if lines[0].S != "01234" {
		t.Errorf("I wanted %v but it was %v.", "01234", lines[0].String())
	} else if lines[1].S != "56789" {
		t.Errorf("I wanted %v but it was %v.", "56789", lines[1].String())
	}

	if lines := string2Lines(sampleStringAscii, 4); len(lines) != 3 {
		t.Error(string2Lines(sampleStringAscii, 5))
	} else if lines[0].S != "0123" {
		t.Errorf("I wanted %v but it was %v.", "0123", lines[0].String())
	} else if lines[1].S != "4567" {
		t.Errorf("I wanted %v but it was %v.", "4567", lines[1].String())
	} else if lines[2].S != "89" {
		t.Errorf("I wanted %v but it was %v.", "89", lines[2].String())
	}

	if lines := string2Lines(sampleStringMultiByte, 0); len(lines) != 1 {
		t.Error(string2Lines(sampleStringMultiByte, 0))
	} else if lines[0].S != "あいうえお" {
		t.Errorf("I wanted %v but it was %v.", "あいうえお", lines[0].String())
	}

	if lines := string2Lines(sampleStringMultiByte, 5); len(lines) != 3 {
		t.Error(string2Lines(sampleStringMultiByte, 5))
	} else if lines[0].S != "あい" {
		t.Errorf("I wanted %v but it was %v.", "あい", lines[0].String())
	} else if lines[1].S != "うえ" {
		t.Errorf("I wanted %v but it was %v.", "うえ", lines[0].String())
	} else if lines[2].S != "お" {
		t.Errorf("I wanted %v but it was %v.", "お", lines[0].String())
	}
}

func TestLineString(t *testing.T) {
	line := NewLine("hoge")
	if line.String() != "  hoge" {
		t.Errorf(line.String())
	}

	line.SetCursor()
	if line.String() != "> hoge" {
		t.Errorf(line.String())
	}

	line.DeleteCursor()
	if line.String() != "  hoge" {
		t.Errorf(line.String())
	}

	line.ToggleCursor()
	if line.String() != "> hoge" {
		t.Errorf(line.String())
	}

	line.ToggleCursor()
	if line.String() != "  hoge" {
		t.Errorf(line.String())
	}
}

func TestItemString(t *testing.T) {
	if item := NewItem("test"); len(item.Lines) != 1 {
		t.Errorf("can not create item: %v", item)
	} else if item.SetCursor(); item.Lines[0].hasCursor == false {
		t.Errorf("can not set cursor: %v", item)
	} else if item.DeleteCursor(); item.Lines[0].hasCursor == true {
		t.Errorf("can not delete cursor: %v", item)
	} else if item.ToggleCursor(); item.Lines[0].hasCursor == false {
		t.Errorf("can not delete cursor: %v", item)
	} else if item.ToggleCursor(); item.Lines[0].hasCursor == true {
		t.Errorf("can not delete cursor: %v", item)
	}

	if item := NewItem("0123456789", 5); len(item.Lines) != 2 {
		t.Errorf("can not create item: %v", item)
	} else if item.SetCursor(); item.Lines[0].hasCursor == false {
		t.Errorf("can not set cursor: %v", item)
	} else if item.Lines[1].hasCursor == true {
		t.Errorf("can not set cursor: %v", item)
	} else if item.DeleteCursor(); item.Lines[0].hasCursor == true {
		t.Errorf("can not set cursor: %v", item)
	} else if item.Lines[1].hasCursor == true {
		t.Errorf("can not set cursor: %v", item)
	} else if item.ToggleCursor(); item.Lines[0].hasCursor == false {
		t.Errorf("can not set cursor: %v", item)
	} else if item.Lines[1].hasCursor == true {
		t.Errorf("can not set cursor: %v", item)
	} else if item.ToggleCursor(); item.Lines[0].hasCursor == true {
		t.Errorf("can not set cursor: %v", item)
	} else if item.Lines[1].hasCursor == true {
		t.Errorf("can not set cursor: %v", item)
	}
}

func TestNewSelectBox(t *testing.T) {
	elems := []string{}
	if box := NewSelectBox("", elems); box != nil {
		t.Errorf("box is not nil: %v", box)
	}

	if box := NewSelectBox("", elems, 1000); box != nil {
		t.Errorf("box is not nil: %v", box)
	}

	if box := NewSelectBox("", nil, 1000); box != nil {
		t.Errorf("box is not nil: %v", box)
	}

	elems = append(
		elems,
		[]string{
			"hoge",
			"huga",
			"foo",
			"bar",
		}...,
	)

	if box := NewSelectBox("", elems); box == nil {
		t.Errorf("box is nil: %v", box)
	} else if len(box.Items) != 4 {
		t.Errorf("The number of elements is abnormal. Although the desired number is %v, it is actually %v.: %v", 4, len(box.Items), box.Items)
	} else if box.Items[0].Lines[0].S != "hoge" {
		t.Error(box.Items[0].Lines[0].S)
	} else if box.Items[0].Lines[0].hasCursor != true {
		t.Error(box.Items[0].Lines[0].hasCursor)
	} else if box.Items[1].Lines[0].hasCursor != false {
		t.Error(box.Items[0].Lines[0].hasCursor)
	}

	elems = []string{
		"0123456789",
		"98765",
	}

	if box := NewSelectBox("", elems, 5); box == nil {
		t.Errorf("box is nil: %v", box)
	} else if len(box.Items) != 2 {
		t.Errorf("The number of elements is abnormal. Although the desired number is %v, it is actually %v.: %v", 2, len(box.Items), box.Items)
	} else if len(box.Items[0].Lines) != 2 {
		t.Errorf("The number of elements is abnormal. Although the desired number is %v, it is actually %v.: %v", 2, len(box.Items[0].Lines), box.Items[0].Lines)
	} else if len(box.Items[1].Lines) != 1 {
		t.Errorf("The number of elements is abnormal. Although the desired number is %v, it is actually %v.: %v", 2, len(box.Items[1].Lines), box.Items[1].Lines)
	}
}

func TestMoveCursor(t *testing.T) {
	elems := []string{
		"hoge",
		"huga",
		"foo",
		"bar",
	}
	if box := NewSelectBox("", elems); box == nil {
		t.Errorf("box is nil: %v", box)
	} else if box.Up(); box.cursorPlace != 0 ||
		!box.Items[box.cursorPlace].Lines[0].hasCursor ||
		box.Items[box.cursorPlace+1].Lines[0].hasCursor {
		t.Errorf("The predicted numerical value is %v but %v.", 0, box.cursorPlace)
	} else if box.Down(); box.cursorPlace != 1 ||
		!box.Items[box.cursorPlace].Lines[0].hasCursor ||
		box.Items[box.cursorPlace-1].Lines[0].hasCursor {
		t.Errorf("The predicted numerical value is %v but %v.", 1, box.cursorPlace)
	} else if box.Down(); box.cursorPlace != 2 ||
		!box.Items[box.cursorPlace].Lines[0].hasCursor ||
		box.Items[box.cursorPlace-1].Lines[0].hasCursor {
		t.Errorf("The predicted numerical value is %v but %v.", 2, box.cursorPlace)
	} else if box.Down(); box.cursorPlace != 3 ||
		!box.Items[box.cursorPlace].Lines[0].hasCursor ||
		box.Items[box.cursorPlace-1].Lines[0].hasCursor {
		t.Errorf("The predicted numerical value is %v but %v.", 3, box.cursorPlace)
	} else if box.Down(); box.cursorPlace != 3 ||
		!box.Items[box.cursorPlace].Lines[0].hasCursor ||
		box.Items[box.cursorPlace-1].Lines[0].hasCursor {
		t.Errorf("The predicted numerical value is %v but %v.", 3, box.cursorPlace)
	} else if box.Up(); box.cursorPlace != 2 ||
		!box.Items[box.cursorPlace].Lines[0].hasCursor ||
		box.Items[box.cursorPlace+1].Lines[0].hasCursor {
		t.Errorf("The predicted numerical value is %v but %v.", 2, box.cursorPlace)
	} else if box.Up(); box.cursorPlace != 1 ||
		!box.Items[box.cursorPlace].Lines[0].hasCursor ||
		box.Items[box.cursorPlace+1].Lines[0].hasCursor {
		t.Errorf("The predicted numerical value is %v but %v.", 1, box.cursorPlace)
	} else if box.Up(); box.cursorPlace != 0 ||
		!box.Items[box.cursorPlace].Lines[0].hasCursor ||
		box.Items[box.cursorPlace+1].Lines[0].hasCursor {
		t.Errorf("The predicted numerical value is %v but %v.", 0, box.cursorPlace)
	} else if box.Up(); box.cursorPlace != 0 ||
		!box.Items[box.cursorPlace].Lines[0].hasCursor ||
		box.Items[box.cursorPlace+1].Lines[0].hasCursor {
		t.Errorf("The predicted numerical value is %v but %v.", 0, box.cursorPlace)
	}
}

func TestTextBox(t *testing.T) {
	if box := NewTextBox("test"); box == nil {
		t.Error("box is nil")
	} else if box.Question != "test" {
		t.Errorf("The expected value is %v, but in fact %v.", "test", box.Question)
	} else if box.Subst("hoge"); box.input != "hoge" {
		t.Errorf("The expected value is %v, but in fact %v.", "hoge", box.input)
	} else if box.Add('0'); box.input != "hoge0" {
		t.Errorf("The expected value is %v, but in fact %v.", "hoge0", box.input)
	} else if box.BS(); box.input != "hoge" {
		t.Errorf("The expected value is %v, but in fact %v.", "hoge", box.input)
	} else if box.Answer() != "hoge" {
		t.Errorf("The expected value is %v, but in fact %v.", "hoge", box.input)
	} else {
		for i := 0; i < 0x20; i++ {
			box.Add(rune(uint8(i)))
			if box.input != "hoge" {
				t.Errorf("The expected value is %v, but in fact %v.", "hoge", box.input)
			}
		}
		for i := 0x80; i < 0xff; i++ {
			box.Add(rune(uint8(i)))
			if box.input != "hoge" {
				t.Errorf("The expected value is %v, but in fact %v.", "hoge", box.input)
			}
		}
		for i := 0; i < 2048; i++ {
			box.BS()
		}
		if box.input != "" {
			t.Errorf("The expected value is %v, but in fact %v.", "", box.input)
		}
		for i := 0; i < 4096; i++ {
			box.Add('a')
		}
		if len(box.input) != 4096 {
			t.Errorf("The expected value is %v, but in fact %v.", 4096, len(box.input))
		}
	}
}
