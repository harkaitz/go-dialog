package dialog

import (
	"bytes"
	"os/exec"
	"time"
	"strconv"
	"strings"
	"errors"
	"io/ioutil"
	"os"
)

var globalArgs1 []string = []string {}
var globalArgs2 []string = []string {}
var height string = "0"
var width  string = "0"

// AddG add global flags to dialog(1) to be used by all later
// calls. It can be used for example to set the backtitle
// with dialog.AddG("--backtitle", "MY Title").
func AddG(options ...string) {
	globalArgs1 = append(globalArgs1, options...)
}

// ConfigG overwrites or adds flags to dialog(1).
func ConfigG(options ...string) {
	for i, o := range globalArgs1 {
		if options[0] != o {
			continue
		}
		for j, a := range options {
			globalArgs1[i+j] = a
		}
		return
	}
	AddG(options...)
}

// Config add flags to dialog(1) that will only be used in the
// next call.
func Config(options ...string) {
	globalArgs2 = append(globalArgs2, options...)
}

// Size sets the default size for windows.
func Size(h, w int) {
	height = strconv.Itoa(h)
	width = strconv.Itoa(w)
}

// ExecDialog executes the dialog program and captures the return
// code and the standard error. The standard input and output use
// the terminal.
func ExecDialog(args ...string) (res string, ok bool, err error) {
	
	var stderr bytes.Buffer
	var nargs  []string = []string{ "-c", `dialog "$@" >/dev/tty </dev/tty`, "--"}
	
	nargs = append(nargs, globalArgs2...)
	nargs = append(nargs, globalArgs1...)
	nargs = append(nargs, args...)
	
	globalArgs2 = []string {}
	height = "0"
	width = "0"
	
	cmd := exec.Command("sh", nargs...)
	cmd.Stderr = &stderr
	
	err = cmd.Start()
	if err != nil {
		return
	}
	
	err = cmd.Wait()
	if err != nil {
		if _, isRetErr := err.(*exec.ExitError); isRetErr {
			return "", false, nil
		} else {
			return "", false, err
		}
	}
	
	return stderr.String(), true, nil
}

// ===================================================================
// ==== WIDGETS ======================================================
// ===================================================================

// Unimplemented:
// - buildlist
// - gauge
// - inputmenu
// - mixedform
// - mixedgauge
// - passwordform
// - prgbox
// - programbox
// - progressbox
// - tailbox
// - tailboxbg
// - timebox
// - treeview

// TagItemStatus describes an item box. When selected "Tag" is used
// to identify each item. "Item" is the actual name displayed. "Status"
// can be "on" or "off".
type TagItemStatus struct {
	Tag    string
	Item   string
	Status string
}

// FormField represents a form field. It is placed in the coordinates
// "YPos", "XPos". The question is "Label" and default answer "Value".
type FormField struct {
	Label string
	Value string
	YPos  int
	XPos  int
}

// MenuItem represents an item in a menu.
type MenuItem struct {
	Key  string
	Text string
}

// Calendar displays month, day and year and asks the user to
// choose a day.
func Calendar(text string, stime time.Time) (t time.Time, ok bool, err error) {
	var tS string
	year, month, day := stime.Date()
	tS, ok, err = ExecDialog(
		"--calendar", text, height, width,
		strconv.Itoa(day),
		strconv.Itoa(int(month)),
		strconv.Itoa(year),
	)
	if err != nil || ok == false {
		return
	}
	t, err = time.Parse("02/01/2006", tS)
	if err != nil {
		ok = false
		return
	}
	return
}

// CheckList a multiple entry menu.
func CheckList(msg string, items []TagItemStatus) (tags []string, ok bool, err error) {
	cmd := []string { "--checklist", msg, height, width, "0" }
	for _, i := range items {
		cmd = append(cmd, i.Tag, i.Item, i.Status)
	}
	str, ok, err := ExecDialog(cmd...)
	if !ok || err != nil {
		return []string{}, ok, err
	}
	return strings.Split(str, " "), ok, err
}

// DSelect asks the user for a directory.
func DSelect(dir string) (path string, ok bool, err error) {
	return ExecDialog("--dselect", dir, height, width)
}

// EditBox opens an small text editor.
func EditBox(file string) (content string, ok bool, err error) {
	return ExecDialog("--editbox", file, height, width)
}

// Form displays a form with multiple fields.
func Form(msg string, labelWidth, valueWidth int, fields []FormField) (data []string, ok bool, err error) {
	data = []string{}
	cmd := []string { "--form", msg, height, width, "0" }
	for _, f := range fields {
		cmd = append(
			cmd,
			f.Label,
			strconv.Itoa(f.YPos),
			strconv.Itoa(f.XPos),
			f.Value,
			strconv.Itoa(f.YPos),
			strconv.Itoa(labelWidth),
			strconv.Itoa(valueWidth),
			"0",
		)
	}
	str, ok, err := ExecDialog(cmd...)
	if !ok || err != nil {
		return
	}
	data = strings.Split(str, "\n")
	if len(data)-1 != len(fields) {
		err = errors.New("Invalid dialog output")
		return
	}
	return
}

// FSelect asks the user for a file.
func FSelect(fil string) (path string, ok bool, err error) {
	return ExecDialog("--fselect", fil, height, width)
}

// InfoBox shows a message to the user but doesn't clean the screen
// afterwards. This is usefull to display a message until some
// operation finishes.
func InfoBox(msg string) (res string, ok bool, err error) {
	return ExecDialog("--infobox", msg, height, width)
}

// InputBox queries for a string to the user.
func InputBox(msg string, initOpt string) (res string, ok bool, err error) {
	cmd := []string { "--inputbox", msg, height, width }
	if initOpt != "" {
		cmd = append(cmd, initOpt)
	}
	return ExecDialog(cmd...)
}

// Menu : As its name suggests, a menu box is a dialog box that can be
// used to present a list of choices in the form of a menu for the
// user to choose.
func Menu(msg string, menu []MenuItem) (key string, ok bool, err error) {
	cmd := []string { "--menu", msg, height, width, "0" }
	for _, i := range menu {
		cmd = append(cmd, i.Key, i.Text)
	}
	return ExecDialog(cmd...)
}

// MsgBox displays a message box.
func MsgBox(msg string) (ok bool, err error) {
	_, ok, err = ExecDialog("--msgbox", msg, height, width)
	return
}

// PasswordBox asks the user for a password.
func PasswordBox(msg string, initOpt string) (res string, ok bool, err error) {
	cmd := []string { "--passwordbox", msg, height, width }
	if initOpt != "" {
		cmd = append(cmd, initOpt)
	}
	return ExecDialog(cmd...)
}

// Pause asks the user to click enter.
func Pause(msg string, secs int) (ok bool, err error) {
	_, ok, err = ExecDialog("--pause", msg, height, width, strconv.Itoa(secs))
	return
}

// RadioList displays a radio list.
func RadioList(msg string, items []TagItemStatus) (sel string, ok bool, err error) {
	cmd := []string { "--radiolist", msg, height, width, "0" }
	for _, i := range items {
		cmd = append(cmd, i.Tag, i.Item, i.Status)
	}
	return ExecDialog(cmd...)
}

// RangeBox queries the user o select from a range of values using a slider.
func RangeBox(msg string, min, max, def int) (res int, ok bool, err error) {
	cmd := []string {
		"--rangebox", msg, height, width,
		strconv.Itoa(min),
		strconv.Itoa(max),
		strconv.Itoa(def),
	}
	resT, ok, err := ExecDialog(cmd...)
	if !ok || err != nil {
		return 0, ok, err
	}
	res, _ = strconv.Atoi(resT)
	return
}

// TextBox shows a text.
func TextBox(file string) (ok bool, err error) {
	_, ok, err = ExecDialog("--textbox", file, height, width)
	return
}

// YesNo asks the user a yes/no question.
func YesNo(msg string) (ok bool, err error) {
	_, ok, err = ExecDialog("--yesno", msg, height, width)
	return
}


// MenuList is like Menu but takes an array and returns the index
// of the selected item.
func MenuList(msg string, menu []string) (num int, ok bool, err error) {
	items := make([]MenuItem, len(menu))
	for i, s := range menu {
		items[i].Key = strconv.Itoa(i)
		items[i].Text = s
	}
	str, ok, err := Menu(msg, items)
	if !ok || err != nil {
		return
	}
	num, _ = strconv.Atoi(str)
	return
}

// TextBoxString shows msg in a text box.
func TextBoxString(msg string) (ok bool, err error) {
	var fp *os.File
	fp, err = ioutil.TempFile("", "dialog")
	if err != nil {
		return
	}
	defer fp.Close()
	defer os.Remove(fp.Name())
	fp.WriteString(msg)
	fp.Sync()
	_, ok, err = ExecDialog("--textbox", fp.Name(), height, width)
	return
}
