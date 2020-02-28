package main

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/mapaiva/encoder/job"
	"github.com/mapaiva/encoder/server"
	"github.com/mapaiva/encoder/transcoder"
)

const directory = "/home/mapaiva/Downloads/As Aventuras de Jackie Chan/2ª Temporada/Sandbox"

func main() {
	jobs := job.NewDatastore()
	scheduler := job.NewScheduler(jobs)
	server := server.New(jobs)
	transcode(scheduler)
	log.Fatal(server.Start())
}

func transcode(scheduler *job.Scheduler) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for i, file := range files {
		if file.IsDir() {
			continue
		}
		options := transcoder.OptionsFrom(file, directory)
		transcoder := transcoder.NewFfmpeg(options)
		job := job.Job{
			ID:         strconv.Itoa(i),
			File:       file,
			Status:     job.StatusRunning,
			Transcoder: transcoder,
		}
		go scheduler.Schedule(job)
	}
}
