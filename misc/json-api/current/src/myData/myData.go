package myData

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"myThings"
	"os"
	"path/filepath"
)

var (
	ErrNotFound      = errors.New("Not found.")
	InvalidType      = errors.New("Invalid struct type.")
	InvalidMode      = errors.New("Invalid access method selected.")
	resourceLocation = "../resources"
	putMap           = map[string]string{
		"*myThings.Character": "characters",
		"*myThings.Spell":     "spells",
	}
	meta      = "meta"
	maxID     = "id.json"
	accessGet = "get"
	accessPut = "put"
)

type object interface {
	GetID() int
	SetID(int)
}

func read(t interface{}, r io.Reader) error {
	b := &bytes.Buffer{}
	if _, err := b.ReadFrom(r); err != nil {
		return err
	}
	return json.Unmarshal(b.Bytes(), &t)
}
func write(t interface{}, w io.Writer) error {
	p, err := json.Marshal(t)
	if err != nil {
		return err
	}
	_, err = w.Write(p)
	return err
}

func access(t object, mode string) error {
	val, ok := putMap[fmt.Sprintf("%T", t)]
	if !ok {
		return InvalidType
	}
	file := fmt.Sprintf("%d.json", t.GetID())
	absFilePath := filepath.Join(resourceLocation, val, file)
	var f *os.File
	defer f.Close()
	switch mode {
	case accessPut:
		if t.GetID() == 0 {
			assignID(t)
		}
		f, err := os.Create(absFilePath)
		if err != nil {
			return err
		}
		return write(t, f)
	case accessGet:
		f, err := os.Open(absFilePath)
		if err != nil {
			return err
		}
		return read(t, f)
	default:
		return InvalidMode
	}
	return nil
}

func assignID(t object) error {
	objType := fmt.Sprintf("%T", t)
	mapper := map[string]int{}
	absFilePath := filepath.Join(resourceLocation, meta, maxID)
	f, err := os.Open(absFilePath)
	if err != nil {
		return err
	}
	if err = read(mapper, f); err != nil {
		return nil
	}
	max := mapper[objType]
	mapper[objType] += 1
	if err = write(mapper, f); err != nil {
		return err
	}
	defer f.Close()

	t.SetID(max)

	return nil
}
func PutCharacter(c *myThings.Character) error {
	return access(c, accessPut)
}

func GetCharacter(c *myThings.Character) error {
	return access(c, accessGet)
}
func PutSpell(s *myThings.Spell) error {
	return access(s, accessPut)
}

func GetSpell(s *myThings.Spell) error {
	return access(s, accessGet)
}
