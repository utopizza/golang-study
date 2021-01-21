package main

import (
	"log"
	"os"
	"sync"
)

func createImageFile(filename string) (string, error) {
	// do something
	return filename + ".jpg", nil
}

// too slow
func makeThumbnails(filenames []string) {
	for _, filename := range filenames {
		if _, err := createImageFile(filename); err != nil {
			log.Println(err)
		}
	}
}

// incorrect
func makeThumbnails2(filenames []string) {
	for _, filename := range filenames {
		go createImageFile(filename) // the outer function will exit
	}
}

// use channel to wait for inner functions finished
func makeThumbnails3(filenames []string) {
	done := make(chan struct{})
	for _, filename := range filenames {
		go func(f string) {
			createImageFile(f)
			done <- struct{}{} //send one signal for each filename
		}(filename)
	}
	// wait for signals for each filename
	for range filenames {
		<-done
	}
}

// incorrect:goroutine leak
func makeThumbnails4(filenames []string) error {
	errors := make(chan error) //non-buffered
	for _, filename := range filenames {
		go func(f string) {
			_, err := createImageFile(f)
			errors <- err
		}(filename)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err
			//incorrect, the remaining goroutines will block forever
			//because no one will consume the next coming err
			//in the non-buffered channel
		}
	}
	return nil
}

// buffered channel
func makeThumbnails5(filenames []string) (thumbnails []string, err error) {
	type item struct {
		file string
		err  error
	}

	ch := make(chan item, len(filenames)) //buffered channel for all files
	for _, filename := range filenames {
		go func(f string) {
			var it item
			it.file, it.err = createImageFile(f)
			ch <- it
		}(filename)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
			//the chan is buffered, large enough to contain all goroutines errors
			//so that the remaining goroutine won't block, they can finish
			//their job then exit
		}
		thumbnails = append(thumbnails, it.file)
	}

	return thumbnails, nil
}

// finally, if we don't know the number of input jobs
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	//worker
	for filename := range filenames {
		wg.Add(1) // register a job
		go func(f string) {
			defer wg.Done() // finish a job
			thumb, err := createImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(filename)
	}

	//closer
	go func() {
		wg.Wait() //why here?
		close(sizes)
	}()

	//count
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
