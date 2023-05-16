package main

import (
	"context"
	"errors"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
)

// errorが全て出力できるか確認
// oroutineでエラーを出しても他のgoroutineには影響ないか確認
func main() {
	ctx := context.Background()
	mg := new(multierror.Group)
	mg.Go(func() error {
		return worker1(ctx)
	})
	mg.Go(func() error {
		return worker2(ctx)
	})
	mg.Go(func() error {
		return worker3(ctx)
	})
	mgerr := mg.Wait()
	for _, err := range mgerr.Errors {
		fmt.Println(err.Error())
	}
}

func worker1(ctx context.Context) error {
	fmt.Println("##### worker1 #####")
	return errors.New("worker1 error")
}

func worker2(ctx context.Context) error {
	fmt.Println("##### worker2 #####")
	return errors.New("worker2 error")
}

func worker3(context.Context) error {
	fmt.Println("##### worker3 #####")
	return nil
}
