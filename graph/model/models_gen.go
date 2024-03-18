// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gmodel

import (
	"fmt"
	"io"
	"strconv"
)

type Flow struct {
	ID    uint    `json:"id"`
	Tasks []*Task `json:"tasks"`
}

type Mutation struct {
}

type Query struct {
}

type Subscription struct {
}

type Task struct {
	ID      uint       `json:"id"`
	Type    TaskType   `json:"type"`
	Status  TaskStatus `json:"status"`
	Args    any        `json:"args"`
	Results any        `json:"results"`
}

type TaskStatus string

const (
	TaskStatusInProgress TaskStatus = "inProgress"
	TaskStatusFinished   TaskStatus = "finished"
	TaskStatusStopped    TaskStatus = "stopped"
	TaskStatusFailed     TaskStatus = "failed"
)

var AllTaskStatus = []TaskStatus{
	TaskStatusInProgress,
	TaskStatusFinished,
	TaskStatusStopped,
	TaskStatusFailed,
}

func (e TaskStatus) IsValid() bool {
	switch e {
	case TaskStatusInProgress, TaskStatusFinished, TaskStatusStopped, TaskStatusFailed:
		return true
	}
	return false
}

func (e TaskStatus) String() string {
	return string(e)
}

func (e *TaskStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskStatus", str)
	}
	return nil
}

func (e TaskStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskType string

const (
	TaskTypeInput  TaskType = "input"
	TaskTypeAction TaskType = "action"
)

var AllTaskType = []TaskType{
	TaskTypeInput,
	TaskTypeAction,
}

func (e TaskType) IsValid() bool {
	switch e {
	case TaskTypeInput, TaskTypeAction:
		return true
	}
	return false
}

func (e TaskType) String() string {
	return string(e)
}

func (e *TaskType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskType", str)
	}
	return nil
}

func (e TaskType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
