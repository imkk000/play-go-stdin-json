package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Println("err:", err)
			continue
		}

		var m map[string]any
		if err := json.Unmarshal(l, &m); err != nil {
			fmt.Println("err:", err)
			continue
		}
		walk(m, 0)
	}
}

func walk(m map[string]any, indent int) {
	for k, v := range m {
		switch v.(type) {
		case map[string]any:
			walk(v.(map[string]any), indent+1)
		case []any:
			for _, v := range v.([]any) {
				vv, ok := v.(map[string]any)
				if !ok {
					continue
				}
				walk(vv, indent+1)
			}
		default:
			for i := 0; i < indent; i++ {
				fmt.Print("  ")
			}
			fmt.Printf("%s: %s\n",
				color.BlueString(k),
				color.RedString("%v", v),
			)
		}
	}
}
