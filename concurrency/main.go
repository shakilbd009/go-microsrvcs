package concurrency

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/shakilbd009/go-microsrvcs/src/api/domain/repositories"
	"github.com/shakilbd009/go-microsrvcs/src/api/services"
	"github.com/shakilbd009/go-microsrvcs/src/api/utils/errors"
)

var (
	success map[string]string
	failed  map[string]errors.ApiError
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.ApiError
}

func main() {

	requests := getRequests()
	input := make(chan createRepoResult)
	buffer := make(chan bool, 10)
	var wg sync.WaitGroup
	fmt.Println(fmt.Sprintf("about to process %d requests", len(requests)))
	go handleResults(input, &wg)
	for _, req := range requests {
		buffer <- true
		wg.Add(1)
		go createRepo(buffer, req, input)
	}
	wg.Wait()
	close(input)
}

func handleResults(input chan createRepoResult, wg *sync.WaitGroup) {

	for result := range input {
		if result.Error != nil {
			failed[result.Result.Name] = result.Error
		}
		success[result.Result.Name] = result.Result.Name
		wg.Done()
	}
}

func createRepo(buff chan bool, request repositories.CreateRepoRequest, out chan createRepoResult) {
	result, err := services.RepositoryService.CreateRepo("header", request)
	out <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}
	<-buff
}

func getRequests() []repositories.CreateRepoRequest {

	results := make([]repositories.CreateRepoRequest, 0)
	file, err := os.Open("/Users/akram/Desktop/hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{
			Name: line,
		}
		results = append(results, request)
	}
	return results
}
