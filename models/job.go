package models

import (
	"log"
	"time"

	"github.com/ileyd/topaz/db"
	"gopkg.in/mgo.v2/bson"
)

/*
	Rationale:
		This information is stored/collected so we can balance fetching content
		across Sonarr servers without fear of work duplication or race conditions
*/

// Collection names
const (
	JobsCollection = "jobs"
)

// Job types
const (
	FetchJob  = 1
	EncodeJob = 2
)

// Job describes a task that to be or that has been completed
type Job struct {
	ID            string `json:"_id,omitempty" bson:"_id,omitempty"`
	UUID          string `json:"uuid" bson:"uuid"`                   // to identification of job entry by job runner
	EpisodeInfoID string `json:"episodeInfoID" bson:"episodeInfoID"` // to enable identification of matching sonarr job pairs
	Type          int    `json:"type" bson:"type"`

	TimeStarted  time.Time `json:"timeStarted" bson:"timeStarted"`
	TimeFinished time.Time `json:"timeFinished" bson:"timeFinished"`
	Completed    bool      `json:"completed" bson:"completed"`

	Series Series `json:"series" bson:"series"`

	Episodes map[int]map[int]bool `json:"episodes" bson:"episodes"` // Episodes[season]{epnum, epnum, ...}
}

// JobModel is used to group model functions relating to Job objects
type JobModel struct{}

// Create ...
func (m *JobModel) Create(job Job) (err error) {
	jobs := db.GetDB().C(JobsCollection)
	return jobs.Insert(job)
}

// Update ...
func (m *JobModel) Update(job Job) (err error) {
	jobs := db.GetDB().C(JobsCollection)
	return jobs.Update(bson.M{"_id": job.ID}, job)
}

// GetAll ...
func (m *JobModel) GetAll() []Job {
	var results []Job
	err := db.GetDB().C(JobsCollection).Find(nil).All(&results)
	log.Println(err)
	return results
}

// Get ...
func (m *JobModel) Get(id int64) Job {
	var result Job
	err := db.GetDB().C(JobsCollection).Find(bson.M{"_id": id}).One(&result)
	log.Println(err)
	return result
}

// GetByEpisodeInfoID ...
func (m *JobModel) GetByEpisodeInfoID(ID string) (j Job, err error) {
	err = db.GetDB().C(JobsCollection).Find(bson.M{"episodeInfoID": ID}).One(&j)
	return j, err
}
