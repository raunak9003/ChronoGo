package main

import "sync"

type Job struct {
	ID   string
	Data string
}

type Result struct {
	JobID     int
	Processed string
	Error     error
}

func WorkerPool(jobs []Job, numWorkers int) []Result {

	var wg sync.WaitGroup
	jobsChan := make(chan Job, len(jobs))
	resultsChan := make(chan Result, len(jobs))

	// starting workers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, jobsChan, resultsChan)
	}

	// Send jobs
	go func() {
		for _, job := range jobs {
			jobsChan <- job
		}
		close(jobsChan) // Close after all jobs are sent
	}()

	// Wait and collect results
	go func() {
		wg.Wait()
		close(resultsChan) // Close after all workers are done
	}()

	// Collect results
	var results []Result
	for result := range resultsChan {
		results = append(results, result)
	}
	return results
}

func worker(wg *sync.WaitGroup, jobs chan Job, results chan Result) {
	defer wg.Done()
	for job := range jobs {
		result := Result{
			JobID:     job.ID,
			Processed: processJob(job), // processing logic
		}
		results <- result
	}
}
