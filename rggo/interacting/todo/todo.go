package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type item struct {
	Id   int
	Task string
	Done bool
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task: task,
		Done: false,
	}
	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exit", i)
	}
	ls[i-1].Done = true

	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *List) PrintList(verbose bool, clean bool) string {
	formatted := ""
	for k, t := range *l {
		prefix := " "
		if t.Done {
			prefix = "X "
		}
		if clean && t.Done {
			continue
		}
		if verbose {
			formatted += fmt.Sprintf("Adding Verbose: %s%d: %s\n", prefix, k+1, t.Task)
		} else {
			formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
		}
	}
	return formatted
}
